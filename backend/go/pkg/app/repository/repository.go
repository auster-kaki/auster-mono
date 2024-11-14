package repository

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type Repository interface {
	User() UserRepository
	UserItineraryHistory() UserItineraryHistoryRepository
	Itinerary() ItineraryRepository
	ItineraryResult() ItineraryResultRepository
	TravelSpot() TravelSpotRepository
	TravelSpotPhoto() TravelSpotPhotoRepository
	TravelSpotItem() TravelSpotItemRepository
	Vendor() VendorRepository
	Diary() DiaryRepository
	DiaryTag() DiaryTagRepository
	DiaryUser() DiaryUserRepository
}

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
}

type UserItineraryHistoryRepository interface {
	Create(ctx context.Context, userItineraryDiaries ...entity.UserItineraryHistory) error
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserItineraryHistories, error)
}

type ItineraryRepository interface {
	Create(ctx context.Context, itineraries ...entity.Itinerary) error
	GetByUserItineraryHistoryID(ctx context.Context, userItineraryDiaryID entity.UserItineraryHistoryID) (entity.Itineraries, error)
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
	FindByID(ctx context.Context, id entity.VendorID) (*entity.Vendor, error)
}

type DiaryRepository interface {
	Create(ctx context.Context, diaries ...entity.Diary) error
	FindByID(ctx context.Context, id entity.DiaryID) (*entity.Diary, error)
	GetByID(ctx context.Context, id entity.DiaryID) (entity.Diaries, error)
}

type DiaryTagRepository interface {
	Create(ctx context.Context, diaryTags ...entity.DiaryTag) error
	GetByDiaryID(ctx context.Context, diaryID entity.DiaryID) (entity.DiaryTags, error)
}

type DiaryUserRepository interface {
	Create(ctx context.Context, diaryUsers ...entity.DiaryUser) error
	GetByUserID(ctx context.Context, userID entity.UserID) (entity.DiaryUsers, error)
}
