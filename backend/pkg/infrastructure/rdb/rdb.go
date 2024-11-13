package rdb

import (
	"cmp"
	"database/sql"
	"fmt"
	"os"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb/table"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type rdb struct {
	user               *table.User
	vendor             *table.Vendor
	travelSpotItem     *table.TravelSpotItem
	travelSpotPhoto    *table.TravelSpotPhoto
	itinerary          *table.Itinerary
	itineraryResult    *table.ItineraryResult
	travelSpot         *table.TravelSpot
	userItineraryDiary *table.UserItineraryDiary
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
	return &rdb{
		user:               table.NewUser(db),
		vendor:             table.NewVendor(db),
		travelSpotItem:     table.NewTravelSpotItem(db),
		travelSpotPhoto:    table.NewTravelSpotPhoto(db),
		itinerary:          table.NewItinerary(db),
		itineraryResult:    table.NewItineraryResult(db),
		travelSpot:         table.NewTravelSpot(db),
		userItineraryDiary: table.NewUserItineraryDiary(db),
	}, nil
}

func (r *rdb) User() repository.UserRepository {
	return r.user
}

func (r *rdb) UserItineraryDiary() repository.UserItineraryDiaryRepository {
	return r.userItineraryDiary
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
