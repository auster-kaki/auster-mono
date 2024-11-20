package rdb

import (
	"cmp"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb/table"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunslog"
)

type rdb struct {
	user                 *table.User
	hobby                *table.Hobby
	userHobby            *table.UserHobby
	vendor               *table.Vendor
	travelSpotItem       *table.TravelSpotItem
	travelSpotPhoto      *table.TravelSpotPhoto
	itinerary            *table.Itinerary
	itineraryResult      *table.ItineraryResult
	travelSpot           *table.TravelSpot
	userItineraryHistory *table.UserItineraryHistory
	diary                *table.Diary
	diaryTag             *table.DiaryTag
	diaryUser            *table.DiaryUser
	encounter            *table.Encounter
	travelSpotHobby      *table.TravelSpotHobby
}

func NewDB() (*rdb, error) {
	var (
		host           = cmp.Or(os.Getenv("MYSQL_HOST"), "127.0.0.1")
		dataSourceName = fmt.Sprintf("root:@tcp(%s:3306)/auster?charset=utf8mb4&parseTime=true", host)
	)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqlDB, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	db.AddQueryHook(bunslog.NewQueryHook(bunslog.WithLogger(slog.Default())))

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &rdb{
		user:                 table.NewUser(db),
		hobby:                table.NewHobby(db),
		userHobby:            table.NewUserHobby(db),
		vendor:               table.NewVendor(db),
		travelSpotItem:       table.NewTravelSpotItem(db),
		travelSpotPhoto:      table.NewTravelSpotPhoto(db),
		itinerary:            table.NewItinerary(db),
		itineraryResult:      table.NewItineraryResult(db),
		travelSpot:           table.NewTravelSpot(db),
		userItineraryHistory: table.NewUserItineraryHistory(db),
		diary:                table.NewDairy(db),
		diaryTag:             table.NewDiaryTag(db),
		diaryUser:            table.NewDairyUser(db),
		encounter:            table.NewEncounter(db),
		travelSpotHobby:      table.NewTravelSpotHobby(db),
	}, nil
}

func (r *rdb) User() repository.UserRepository {
	return r.user
}

func (r *rdb) Hobby() repository.HobbyRepository { return r.hobby }

func (r *rdb) UserHobby() repository.UserHobbyRepository { return r.userHobby }

func (r *rdb) UserItineraryHistory() repository.UserItineraryHistoryRepository {
	return r.userItineraryHistory
}

func (r *rdb) Itinerary() repository.ItineraryRepository {
	return r.itinerary
}

func (r *rdb) ItineraryResult() repository.ItineraryResultRepository {
	return r.itineraryResult
}

func (r *rdb) TravelSpot() repository.TravelSpotRepository {
	return r.travelSpot
}

func (r *rdb) TravelSpotPhoto() repository.TravelSpotPhotoRepository {
	return r.travelSpotPhoto
}

func (r *rdb) TravelSpotItem() repository.TravelSpotItemRepository {
	return r.travelSpotItem
}

func (r *rdb) Vendor() repository.VendorRepository {
	return r.vendor
}

func (r *rdb) Diary() repository.DiaryRepository {
	return r.diary
}

func (r *rdb) DiaryTag() repository.DiaryTagRepository {
	return r.diaryTag
}

func (r *rdb) DiaryUser() repository.DiaryUserRepository {
	return r.diaryUser
}

func (r *rdb) Encounter() repository.EncounterRepository {
	return r.encounter
}

func (r *rdb) TravelSpotHobby() repository.TravelSpotHobbyRepository {
	return r.travelSpotHobby
}
