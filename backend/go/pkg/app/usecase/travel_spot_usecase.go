package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/rpc"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type TravelSpotUseCase struct {
	repository repository.Repository
	rpc        rpc.RPC
}

func NewTravelSpotUseCase(r repository.Repository, rpc rpc.RPC) *TravelSpotUseCase {
	return &TravelSpotUseCase{repository: r, rpc: rpc}
}

func (u *TravelSpotUseCase) GetTravelSpots(ctx context.Context, userID entity.UserID, hobbyID entity.HobbyID) (entity.TravelSpots, error) {
	travelSpotHobbies, err := u.repository.TravelSpotHobby().GetByHobbyID(ctx, hobbyID)
	if err != nil {
		return nil, err
	}
	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, travelSpotHobbies.TravelSpotIDs())
	if err != nil {
		return nil, err
	}
	return travelSpots, nil
}

type CreateDiaryOutput struct {
	ID          entity.TravelSpotDiaryID
	Title       string
	PhotoPath   string
	Description string
}

// CreateDiary TODO: 画像・本文・タイトルの生成周りが未完成
func (u *TravelSpotUseCase) CreateDiary(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID, date time.Time) (*CreateDiaryOutput, error) {
	user, err := u.repository.User().FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	travelSpot, err := u.repository.TravelSpot().FindByID(ctx, travelSpotID)
	if err != nil {
		return nil, fmt.Errorf("failed to find travel spot: %w", err)
	}
	/*
		日記を生成する条件
		1. 初めて体験を選択した場合
		2. 予約後、もう一度体験を選択した場合

		日記を生成しない条件
		1. 初めて体験を選択した後、一度予約をせずに戻り、もう一度同じ体験を選択した場合
	*/
	rs, err := u.repository.Reservation().FindByUserIDAndTravelSpotID(ctx, userID, travelSpotID)
	if err != nil {
		return nil, fmt.Errorf("failed to find reservation: %w", err)
	}

	travelSpotDiaries, dErr := u.repository.TravelSpotDiary().FindByUserIDAndTravelSpotID(ctx, userID, travelSpotID)
	if dErr != nil && dErr != repository.ErrNotFound {
		return nil, fmt.Errorf("failed to find travel spot diary: %w", dErr)
	}
	for _, travelSpotDiary := range travelSpotDiaries {
		// 日記を生成しない条件
		// 1. 初めて体験を選択した後、一度予約をせずに戻り、もう一度同じ体験を選択した場合
		if !slices.Contains(rs.TravelSpotDiaryIDs(), travelSpotDiary.ID) {
			return &CreateDiaryOutput{
				ID:          travelSpotDiary.ID,
				Title:       travelSpotDiary.Title,
				PhotoPath:   travelSpotDiary.PhotoPath,
				Description: travelSpotDiary.Description,
			}, nil
		}
	}

	gOut, err := u.generateDiary(ctx, user, travelSpot)
	if err != nil {
		fmt.Println("failed to generate diary: %w", err)
		// 画像生成に失敗した場合は元の体験画像をそのまま返す
		photo, err := austerstorage.Get(travelSpot.PhotoPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get photo: %w", err)
		}
		paths := strings.Split(travelSpot.PhotoPath, "/")
		gOut = &rpc.GetImagePathOutput{
			GeneratedImage: photo,
			Filename:       paths[len(paths)-1],
		}
	}
	id := austerid.Generate[entity.TravelSpotDiaryID]()
	// goサーバにも画像を保存
	path, err := austerstorage.Save(
		austerstorage.ContentType(austerstorage.PNG),
		filepath.Join("travel_spot_diaries", string(userID), string(id), gOut.Filename),
		gOut.GeneratedImage,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	var (
		title       = "ダミータイトル１"
		description = "ダミー本文１"
	)
	// TODO: タイトルと本文を生成AIから取得したい
	if d, ok := dummyDiaries[travelSpotID]; ok {
		title = d.Title
		description = d.Description
	}
	if err := u.repository.TravelSpotDiary().Create(ctx, entity.TravelSpotDiary{
		ID:           id,
		UserID:       userID,
		TravelSpotID: travelSpotID,
		Title:        title,
		Date:         date,
		PhotoPath:    path,
		Description:  description,
	}); err != nil {
		return nil, fmt.Errorf("failed to create travel spot diary: %w", err)
	}

	return &CreateDiaryOutput{
		ID:          id,
		Title:       title,
		PhotoPath:   path,
		Description: description,
	}, nil
}

func (u *TravelSpotUseCase) generateDiary(ctx context.Context, user *entity.User, travelSpot *entity.TravelSpot) (*rpc.GetImagePathOutput, error) {
	cOut, err := u.rpc.Diary().CreateImage(ctx, rpc.CreateImageInput{
		SourcePath: travelSpot.PhotoPath,
		TargetPath: user.ProfilePath,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create image: %w", err)
	}

	// ポーリング処理の設定
	var (
		maxRetries   = 90              // 最大リトライ回数
		pollInterval = time.Second * 2 // ポーリング間隔
	)
	// ジョブの完了を待つポーリング処理
	for i := 0; i < maxRetries; i++ {
		gOut, err := u.rpc.Diary().GetStatus(ctx, rpc.GetStatusInput{
			JobID: cOut.JobID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get job status: %w", err)
		}
		if gOut.Status == "ok" {
			// 処理完了
			fmt.Println("image generation completed. job id: ", cOut.JobID)
			break
		} else if gOut.Status == "error" {
			return nil, fmt.Errorf("image generation failed: %s", gOut.Status)
		}

		// 最大リトライ回数に達した場合
		if i == maxRetries-1 {
			return nil, fmt.Errorf("timeout: image generation did not complete within %d attempts", maxRetries)
		}

		// 次のポーリングまで待機
		time.Sleep(pollInterval)
	}

	gOut, err := u.rpc.Diary().GetImagePath(ctx, rpc.GetImagePathInput{
		JobID: cOut.JobID,
	})
	if err != nil {
		return nil, fmt.Errorf("u.rpc.Diary(). failed to get image path: %w", err)
	}
	return &gOut, nil
}

// TODO: CreateDiary で何かを一時的に返す場合は↓を参考を元に CreateDiaryOutput に変換して返す
// 趣味: 釣り の 1 ~ 6 だけ用意
var dummyDiaries = map[entity.TravelSpotID]entity.TravelSpotDiary{
	"1": {
		Title:       "大物ヒラマサとの出会い",
		PhotoPath:   "/assets/images/dairy_1.jpg",
		Description: "今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！",
	},
	"2": {
		Title:       "大物カジキとの戦い",
		PhotoPath:   "/assets/images/dairy_2.jpg",
		Description: "今日は銚子港から漁船に乗り、念願のカジキ釣りに挑戦。朝もやの中、期待と不安を胸に出港。3時間の格闘の末、体長2メートルを超える大物を釣り上げることができた！船長のアドバイスのおかげで、生涯最高の釣果。あの引きの強さは一生忘れられない。記念写真を撮って、達成感に浸った最高の1日だった。",
	},
	"3": {
		Title:       "特大ヒラメ、大興奮！",
		PhotoPath:   "/assets/images/dairy_3.jpg",
		Description: "今日は銚子港から漁船に乗り、念願のカジキ釣りに挑戦。朝もやの中、期待と不安を胸に出港。3時間の格闘の末、体長2メートルを超える大物を釣り上げることができた！船長のアドバイスのおかげで、生涯最高の釣果。あの引きの強さは一生忘れられない。記念写真を撮って、達成感に浸った最高の1日だった。",
	},
	"4": {
		Title:       "初めての競り体験",
		PhotoPath:   "/assets/images/dairy_4.jpg",
		Description: "早朝4時、銚子の市場で初めて競りを見学。大量のマグロが整然と並べられ、その迫力に圧倒された。仲買人たちが真剣な眼差しで品定めをする中、独特の手の合図と威勢の良い掛け声が響き渡る。漁師さんに教わりながら、マグロの品質を見分けるポイントも学べた。市場の熱気に触れ、新鮮な魚の流通を肌で感じた貴重な体験だった。",
	},
	"5": {
		Title:       "漁師の仕事を体験",
		PhotoPath:   "/assets/images/dairy_5.jpg",
		Description: "銚子の漁船に乗せてもらい、本物の漁を体験。早朝3時の出港で眠かったけど、大量の魚を網で引き上げる瞬間の興奮で目が覚めた。漁師さんの手際の良さに感動。魚を選別する作業も手伝わせてもらい、プロの技を間近で見られた。海の上での仕事の大変さと、新鮮な魚が食卓に届くまでの過程を知る貴重な経験になった。",
	},
	"6": {
		Title:       "美しい海を守る休日",
		PhotoPath:   "/assets/images/dairy_6.jpg",
		Description: "銚子の海岸でゴミ拾いボランティアに参加。いつも釣りで楽しませてもらっている海への恩返しだ。2時間で大きなゴミ袋2つが一杯になった。プラスチックごみが特に目立ち、海の環境問題を実感。地元の漁師さんも「きれいな海を守りたい」と一緒に活動してくれて心強かった。次回は釣り仲間も誘って参加したい。",
	},
}

type GetItinerariesOutput struct {
	TravelSpotItineraries entity.TravelSpotItineraries
	Items                 entity.TravelSpotItineraryItems
}

func (u *TravelSpotUseCase) GetItineraries(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (*GetItinerariesOutput, error) {
	travelSpotItineraries, err := u.repository.TravelSpotItinerary().GetByTravelSpotID(ctx, travelSpotID)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(travelSpotItineraries, func(a, b *entity.TravelSpotItinerary) int {
		return a.Order - b.Order
	})

	travelSpotItineraryItems, err := u.repository.TravelSpotItineraryItem().GetByTravelSpotItineraryIDs(ctx, travelSpotItineraries.IDs())
	if err != nil {
		return nil, err
	}

	return &GetItinerariesOutput{
		TravelSpotItineraries: travelSpotItineraries,
		Items:                 travelSpotItineraryItems,
	}, nil
}
