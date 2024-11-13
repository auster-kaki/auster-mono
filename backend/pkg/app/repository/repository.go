package repository

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type Repository interface {
	User() UserRepository
	UserItineraryDiary() UserItineraryDiaryRepository
	Itinerary() ItineraryRepository
	ItineraryResult() ItineraryResultRepository
	TravelSpot() TravelSpotRepository
	TravelSpotPhoto() TravelSpotPhotoRepository
	TravelSpotItem() TravelSpotItemRepository
	Vendor() VendorRepository
}

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
}

type UserItineraryDiaryRepository interface {
	Create(ctx context.Context, userItineraryDiaries ...entity.UserItineraryDiary) error
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserItineraryDiaries, error)
}

type ItineraryRepository interface {
	Create(ctx context.Context, itineraries ...entity.Itinerary) error
	GetByUserItineraryDiaryID(ctx context.Context, userItineraryDiaryID entity.UserItineraryDiaryID) (entity.Itineraries, error)
}

type ItineraryResultRepository interface {
	Create(ctx context.Context, itineraryResults ...entity.ItineraryResult) error
	GetByItineraryID(ctx context.Context, itineraryID entity.ItineraryID) (entity.ItineraryResults, error)
}

type TravelSpotRepository interface {
	Create(ctx context.Context, travelSpots ...entity.TravelSpot) error
	GetByVendorID(ctx context.Context, vendorID entity.VendorID) (entity.TravelSpots, error)
}

type TravelSpotPhotoRepository interface {
	Create(ctx context.Context, travelSpotPhotos ...entity.TravelSpotPhoto) error
	GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotPhotos, error)
}

type TravelSpotItemRepository interface {
	Create(ctx context.Context, travelSpotItems ...entity.TravelSpotItem) error
	GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotItems, error)
}

type VendorRepository interface {
	Create(ctx context.Context, vendors ...entity.Vendor) error
	GetAll(ctx context.Context) (entity.Vendors, error)
}
