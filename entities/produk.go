package entities

import "time"

type Produk struct {
	Id         uint
	Name       string
	Category   Category
	Stok       int64
	Keterangan string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
