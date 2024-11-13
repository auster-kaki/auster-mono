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

type UserItineraryDiaryRepository interface{}

type ItineraryRepository interface{}

type ItineraryResultRepository interface{}

type TravelSpotRepository interface{}

type TravelSpotPhotoRepository interface{}

type TravelSpotItemRepository interface{}

type VendorRepository interface{}
