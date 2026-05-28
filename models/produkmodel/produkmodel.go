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
	produk.createdat,
	produk.updatedat
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
	insert into produk (name, idcat, stok, keterangan, createdat, updatedat) values (?, ?, ?, ?, ?, ?)
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

func Update(id int, produk entities.Produk) bool {

	query, err := config.DB.Exec(`
	UPDATE produk SET name=?, idcat=?, stok=?, keterangan=?,updatedat=? where id=?
	`, produk.Name, produk.Category.Id, produk.Stok, produk.Keterangan, produk.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Detail(id int) entities.Produk {
	row := config.DB.QueryRow(`
	select 
	produk.id,
	produk.name, 
	categories.name as category_name,
	produk.stok,
	produk.keterangan,
	produk.createdat,
	produk.updatedat
	from produk
	join categories on produk.idcat = categories.id
	where  produk.id = ?
	`, id)

	var produk entities.Produk
	err := row.Scan(
		&produk.Id,
		&produk.Name,
		&produk.Category.Name,
		&produk.Stok,
		&produk.Keterangan,
		&produk.CreatedAt,
		&produk.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return produk
}

func Edit(id int) entities.Produk {
	row := config.DB.QueryRow(`
	select 
	produk.id,
	produk.name, 
	categories.name as category_name,
	produk.stok,
	produk.keterangan,
	produk.createdat,
	produk.updatedat
	from produk
	join categories on produk.idcat = categories.id
	where  produk.id = ?
	`, id)

	var produk entities.Produk
	err := row.Scan(
		&produk.Id,
		&produk.Name,
		&produk.Category.Name,
		&produk.Stok,
		&produk.Keterangan,
		&produk.CreatedAt,
		&produk.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return produk
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM produk WHERE id = ?`, id)

	return err
}