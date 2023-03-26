package services

import (
	"database/sql"
	"rest-api-postgres/config"
	"rest-api-postgres/models"
)

func GetAllBook() []models.Book {
	var db = config.CreateConnection()
	defer db.Close()

	var result []models.Book

	sqlStatement := `SELECT * FROM "Book"`

	rows, err := db.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			panic(err)
		}

		result = append(result, book)
	}

	return result
}

func GetBookById(Id int) (*models.Book, error) {
	db := config.CreateConnection()
	defer db.Close()

	book := &models.Book{}

	sqlStatement := `
	SELECT * FROM "Book" 
		WHERE "ID" = $1
	`
	err := db.QueryRow(sqlStatement, Id).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		panic(err)
	}

	return book, nil
}

func CreatedBook(collection models.Book) *models.Book {
	var db = config.CreateConnection()
	defer db.Close()

	book := &models.Book{}

	sql := `
	INSERT INTO "Book" ("Title", "Author", "Desc")
		VALUES($1, $2, $3)
	RETURNING *
	`

	err := db.QueryRow(sql, collection.Title, collection.Author, collection.Desc).Scan(
		&book.ID, &book.Title, &book.Author, &book.Desc,
	)

	if err != nil {
		panic(err)
	}

	return book
}

func UpdatedBook(id int, collection models.Book) int64 {
	var db = config.CreateConnection()
	defer db.Close()

	sqlStatement := `
	UPDATE "Book"
		SET "Title" = $2, "Author" = $3, "Desc" = $4
	WHERE "ID" = $1
	`

	res, err := db.Exec(sqlStatement, id, collection.Title, collection.Author, collection.Desc)
	if err != nil {
		panic(err)
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return RowsAffected
}

func DeletedBook(Id int) int64 {
	var db = config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM "Book" WHERE "ID" = $1`
	res, err := db.Exec(sqlStatement, Id)
	if err != nil {
		panic(err)
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return RowsAffected
}
