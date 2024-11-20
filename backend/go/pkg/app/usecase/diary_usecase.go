package usecase

import (
	"context"
	"errors"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
)

type DiaryUseCase struct {
	repository repository.Repository
}

func NewDiaryUseCase(r repository.Repository) *DiaryUseCase {
	return &DiaryUseCase{repository: r}
}

func (u *DiaryUseCase) GetDiaries(ctx context.Context, userID entity.UserID) (entity.Diaries, error) {
	diaryUsers, err := u.repository.DiaryUser().GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	diaries, err := u.repository.Diary().GetByIDs(ctx, diaryUsers.DiaryIDs())
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

func (u *DiaryUseCase) CreateDiary(ctx context.Context, req *request.Diary) (*entity.Diary, error) {
	// TODO: 生成AIと接続されるまでこちらのデータを利用
	// 生成AIと接続処理を追加し、title, photo_path, body を生成する
	dairy, ok := diaries[entity.TravelSpotID(req.TravelSpotID)]
	if !ok {
		return nil, errors.New("diaries map dummy data not found. travel_spot_id: " + req.TravelSpotID)
	}

	res := entity.Diary{
		ID:        austerid.Generate[entity.DiaryID](),
		Title:     dairy.Title,
		Date:      req.Date,
		PhotoPath: dairy.PhotoPath,
		Body:      dairy.Body,
	}

	if err := u.repository.Diary().Create(ctx, []entity.Diary{res}...); err != nil {
		return nil, err
	}
	return &res, nil
}

// TODO: 生成AIと接続されるまでこちらのデータを利用
// 趣味: 釣り の 1 ~ 6 だけ用意
var diaries = map[entity.TravelSpotID]entity.Diary{
	"1": {
		Title:     "大物ヒラマサとの出会い",
		PhotoPath: "/assets/images/dairy_1.jpg",
		Body:      "今日は早朝から漁船に乗り、期待に胸を膨らませて出航しました。風は少し冷たかったけれど、海の静けさが心地よかったです。そして、ついに大物のヒラマサがヒット！かなりの引きで、腕がパンパンになりましたが、無事に釣り上げることができました。この魚の力強さと美しさには感動しました。次回もこのサイズを狙いたいと思います！",
	},
	"2": {
		Title:     "大物カジキとの戦い",
		PhotoPath: "/assets/images/dairy_2.jpg",
		Body:      "今日は銚子港から漁船に乗り、念願のカジキ釣りに挑戦。朝もやの中、期待と不安を胸に出港。3時間の格闘の末、体長2メートルを超える大物を釣り上げることができた！船長のアドバイスのおかげで、生涯最高の釣果。あの引きの強さは一生忘れられない。記念写真を撮って、達成感に浸った最高の1日だった。",
	},
	"3": {
		Title:     "特大ヒラメ、大興奮！",
		PhotoPath: "/assets/images/dairy_3.jpg",
		Body:      "今日は銚子港から漁船に乗り、念願のカジキ釣りに挑戦。朝もやの中、期待と不安を胸に出港。3時間の格闘の末、体長2メートルを超える大物を釣り上げることができた！船長のアドバイスのおかげで、生涯最高の釣果。あの引きの強さは一生忘れられない。記念写真を撮って、達成感に浸った最高の1日だった。",
	},
	"4": {
		Title:     "初めての競り体験",
		PhotoPath: "/assets/images/dairy_4.jpg",
		Body:      "早朝4時、銚子の市場で初めて競りを見学。大量のマグロが整然と並べられ、その迫力に圧倒された。仲買人たちが真剣な眼差しで品定めをする中、独特の手の合図と威勢の良い掛け声が響き渡る。漁師さんに教わりながら、マグロの品質を見分けるポイントも学べた。市場の熱気に触れ、新鮮な魚の流通を肌で感じた貴重な体験だった。",
	},
	"5": {
		Title:     "漁師の仕事を体験",
		PhotoPath: "/assets/images/dairy_5.jpg",
		Body:      "銚子の漁船に乗せてもらい、本物の漁を体験。早朝3時の出港で眠かったけど、大量の魚を網で引き上げる瞬間の興奮で目が覚めた。漁師さんの手際の良さに感動。魚を選別する作業も手伝わせてもらい、プロの技を間近で見られた。海の上での仕事の大変さと、新鮮な魚が食卓に届くまでの過程を知る貴重な経験になった。",
	},
	"6": {
		Title:     "美しい海を守る休日",
		PhotoPath: "/assets/images/dairy_6.jpg",
		Body:      "銚子の海岸でゴミ拾いボランティアに参加。いつも釣りで楽しませてもらっている海への恩返しだ。2時間で大きなゴミ袋2つが一杯になった。プラスチックごみが特に目立ち、海の環境問題を実感。地元の漁師さんも「きれいな海を守りたい」と一緒に活動してくれて心強かった。次回は釣り仲間も誘って参加したい。",
	},
}
