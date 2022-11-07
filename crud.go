package main

import (
	"database/sql"
	"fmt"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) DBManager {
	return DBManager{db}
}

type Car struct {
	id       int64
	model    string
	color    string
	price    int64
	year     int32
	imageUrl string
	images   []*CarImage
}

type CarImage struct {
	id             int64
	imageUrl       string
	sequenceNumber int32
}

type GetCarParam struct {
	model string
	limit int32
	page  int32
}

type GetCarResponse struct {
	cars  []*Car
	count int32
}

func (m *DBManager) CreateCar(car *Car) (int64, error) {
	var carID int64
	queryCar := `
		INSERT INTO car(
			model,
			color,
			price,
			year,
			image_url
		)VALUES($1, $2, $3, $4, $5)
		RETURNING id
	`
	row := m.db.QueryRow(
		queryCar,
		car.model,
		car.color,
		car.price,
		car.year,
		car.imageUrl,
	)

	err := row.Scan(&carID)
	if err != nil {
		return 0, err
	}

	queryImage := `
		INSERT INTO carimage (
			id,
			imageurl,
			sequencenumber
		)VALUES ($1, $2, $3)
	`
	for _, image := range car.images {
		_, err := m.db.Exec(
			queryImage,
			carID,
			image.imageUrl,
			image.sequenceNumber,
		)
		if err != nil {
			return 0, err
		}
	}
	return carID, nil
}

func (m *DBManager) GetCar(id int64) (*Car, error) {
	var car Car
	car.images = make([]*CarImage, 0)

	query := `
		SELECT
			car.id,
			car.model,
			car.color,
			car.price,
			car.year,
			car.image_url
		FROM car
		INNER JOIN carimage im on car.id=im.id
		WHERE car.id=$1
	`
	row := m.db.QueryRow(query, id)
	err := row.Scan(
		&car.id,
		&car.model,
		&car.color,
		&car.price,
		&car.year,
		&car.imageUrl,
	)
	if err != nil {
		return nil, err
	}

	queryImage := `
		SELECT
			id,
			imageurl,
			sequencenumber
		FROM carimage
		WHERE id=$1
	`
	rows, err := m.db.Query(queryImage, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var image CarImage
		err := rows.Scan(
			&image.id,
			&image.imageUrl,
			&image.sequenceNumber,
		)
		if err != nil {
			return nil, err
		}
		car.images = append(car.images, &image)
	}
	return &car, nil
}

func (m *DBManager) GetAllCars(param *GetCarParam) (*GetCarResponse, error) {
	var result GetCarResponse

	result.cars = make([]*Car, 0)

	limit := fmt.Sprintf(" LIMIT %d", param.limit)

	filter := " WHERE true"
	if param.model != "" {
		filter += " AND model ilike '%" + param.model + "%' "
	}

	query := `
		SELECT
			c.id,
			c.model,
			c.color,
			c.price,
			c.year,
			c.image_url
		FROM car c
		INNER JOIN carimage im on im.id=c.id
		` + filter + `
		ORDER BY c.id
		` + limit
	
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var cr Car

		err := rows.Scan(
			&cr.id,
			&cr.model,
			&cr.color,
			&cr.price,
			&cr.year,
			&cr.imageUrl,
		)
		if err != nil {
			return nil, err
		}
		result.cars = append(result.cars, &cr)
	}

	return &result, nil
}

func (m *DBManager) UpdateCar(car *Car) error {
	query := `
		UPDATE car set
			model=$1,
			color=$2,
			price=$3,
			year=$4,
			image_url=$5
		WHERE id=$6
	`
	result, err := m.db.Exec(
		query,
		car.id,
		car.color,
		car.price,
		car.year,
		car.imageUrl,
	)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	queryDeleteImages := `DELETE FROM carimage WHERE id=$1`
	_, err = m.db.Exec(queryDeleteImages, car.id)
	if err != nil {
		return err
	}
	queryInsertImage := `
		INSERT INTO carimage (
			id,
			imageurl,
			sequencenumber
		)VALUES($1, $2, $3)
	`

	for _, image := range car.images {
		_, err := m.db.Exec(
			queryInsertImage,
			car.id,
			image.imageUrl,
			image.sequenceNumber,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *DBManager) DeleteCar(id int64) error {
	queryDelete := `DELETE FROM car WHERE id=$1`
	_, err := m.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}
	queryDel := `DELETE FROM car WHERE id=$1`
	result, err := m.db.Exec(queryDel, id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}
	return nil

}