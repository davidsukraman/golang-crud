package produkmodel

import (
	"golang-crud/config"
	"golang-crud/entities"
)

func GetAll() []entities.Produk {

	rows, err := config.DB.Query(`
	select 
	produk.id,
	produk.name, 
	categories.name as category_name,
	produk.stok,
	produk.keterangan,
	produk.createat,
	produk.updateat
	from produk
	join categories on produk.idcat = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var produk []entities.Produk

	for rows.Next() {
		var p entities.Produk
		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Category.Name,
			&p.Stok,
			&p.Keterangan,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		produk = append(produk, p)
	}
	return produk
}

func Create(produk entities.Produk) bool {

	result, err := config.DB.Exec(`
	insert into produk (name, idcat, stok, keterangan, createat, updateat) values (?, ?, ?, ?, ?, ?)
	`, produk.Name, produk.Category.Id, produk.Stok, produk.Keterangan, produk.CreatedAt, produk.UpdatedAt)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}
