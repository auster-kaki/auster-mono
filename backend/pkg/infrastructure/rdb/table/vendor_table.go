package table

import (
	"github.com/uptrace/bun"
)

type Vendor struct {
	db *bun.DB
}

func NewVendor(db *bun.DB) *Vendor {
	return &Vendor{db: db}
}
