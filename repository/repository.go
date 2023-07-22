package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/mikeleo03/Course-Scheduler_Backend/models"
	// Uncomment for local development
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Repo struct {
	db *sql.DB
}

// CreateConnection creates a connection with the PostgreSQL database
func CreateConnection() (*Repo, error) {
	// Load .env file
	// Uncomment for local development
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")

	// Return the connection
	return &Repo{db: db}, nil
}

// CloseConnection closes the connection with the PostgreSQL database
func (r *Repo) CloseConnection() error {
	return r.db.Close()
}

// InsertMataKuliah inserts one mata kuliah into the DB
func (r *Repo) InsertMataKuliah(matkul models.MataKuliahAdd, fakultasData []models.Fakultas) int64 {
	// create the insert sql query
	// returning userid will return the id of the inserted matkul
	sqlStatement := `
		INSERT INTO mata_kuliah (nama, sks, jurusan, fakultas, semmin, prediksi) 
		SELECT $1, $2, $3::VARCHAR, f.Nama, $4, $5
		FROM jurusan j JOIN fakultas f ON j.Fakultas = f.Nama
		WHERE j.Nama = $3::VARCHAR
		ON CONFLICT (nama) DO NOTHING
		RETURNING ID`

	// execute the sql statement
	// Scan function will save the insert id in the id
	commandTag, err := r.db.Exec(sqlStatement, matkul.Nama, matkul.SKS, matkul.Jurusan, matkul.SemMin, matkul.Prediksi)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	row, err := commandTag.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to retrieve the affected row count. %v", err)
	}

	// return the inserted row
	return row
}

// InsertFakultas inserts one fakultas into the DB
func (r *Repo) InsertFakultas(fakul models.Fakultas) int64 {
	// create the insert sql query
	// returning userid will return the id of the inserted fakul
	sqlStatement := `
		INSERT INTO fakultas (nama)
		VALUES ($1)
		ON CONFLICT (nama) DO NOTHING
		RETURNING ID`

	// execute the sql statement
	// Scan function will save the insert id in the id
	commandTag, err := r.db.Exec(sqlStatement, fakul.Nama)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	row, err := commandTag.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to retrieve the affected row count. %v", err)
	}

	// return the inserted row
	return row
}

// InsertJurusan inserts one jurusan into the DB
func (r *Repo) InsertJurusan(jurusan models.Jurusan) int64 {
	// create the insert sql query
	// returning userid will return the id of the inserted jurusan
	sqlStatement := `
		INSERT INTO jurusan (nama, fakultas)
		VALUES ($1, $2)
		ON CONFLICT (nama) DO NOTHING
		RETURNING ID`

	// execute the sql statement
	// Scan function will save the insert id in the id
	commandTag, err := r.db.Exec(sqlStatement, jurusan.Nama, jurusan.Fakultas)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	row, err := commandTag.RowsAffected()
	if err != nil {
		log.Fatalf("Unable to retrieve the affected row count. %v", err)
	}

	// return the inserted row
	return row
}

// GetAllMatkul returns all matkuls from the DB
func (r *Repo) GetAllMatkul() ([]models.MataKuliah, error) {
	// create a var base on models
	var matkuls []models.MataKuliah

	// create the select sql query
	sqlStatement := `SELECT * FROM mata_kuliah ORDER BY id`

	// execute the sql statement
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var matkul models.MataKuliah

		// unmarshal the row object to matkul
		err = rows.Scan(&matkul.ID, &matkul.Nama, &matkul.SKS, &matkul.Jurusan, &matkul.Fakultas, &matkul.SemMin, &matkul.Prediksi)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the matkul in the matkuls slice
		matkuls = append(matkuls, matkul)
	}

	// check for any errors during iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
	}

	// return empty matkul on error
	return matkuls, nil
}

// GetAllFakul returns all fakuls from the DB
func (r *Repo) GetAllFakul() ([]models.Fakultas, error) {
	// create a var base on models
	var fakuls []models.Fakultas

	// create the select sql query
	sqlStatement := `SELECT * FROM fakultas ORDER BY id`

	// execute the sql statement
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var fakul models.Fakultas

		// unmarshal the row object to fakul
		err = rows.Scan(&fakul.ID, &fakul.Nama)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the fakul in the fakuls slice
		fakuls = append(fakuls, fakul)
	}

	// check for any errors during iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
	}

	// return empty matkul on error
	return fakuls, nil
}

// GetAllJurus returns all jurus from the DB
func (r *Repo) GetAllJurus() ([]models.Jurusan, error) {
	// create a var base on models
	var jurus []models.Jurusan

	// create the select sql query
	sqlStatement := `SELECT * FROM jurusan ORDER BY id`

	// execute the sql statement
	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var jurusan models.Jurusan

		// unmarshal the row object to jurusan
		err = rows.Scan(&jurusan.ID, &jurusan.Nama, &jurusan.Fakultas)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the jurusan in the jurus slice
		jurus = append(jurus, jurusan)
	}

	// check for any errors during iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
	}

	// return empty matkul on error
	return jurus, nil
}

// GetMatkulData returns filtered matkuls based on jurusan and semester
func (r *Repo) GetMatkulData(jurusan string, semester int) ([]models.MataKuliah, error) {
	// create a var based on models
	var matkuls []models.MataKuliah

	// create the select SQL query with filtering
	sqlStatement := 
		`SELECT * 
		FROM mata_kuliah 
		WHERE Fakultas IN 
			(SELECT f.Nama 
			FROM jurusan j JOIN fakultas f ON j.Fakultas = f.Nama
			WHERE j.Nama = $1) 
		AND SemMin <= $2`

	// execute the SQL statement with the provided parameters
	rows, err := r.db.Query(sqlStatement, jurusan, semester)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var matkul models.MataKuliah

		// unmarshal the row object to matkul
		err = rows.Scan(&matkul.ID, &matkul.Nama, &matkul.SKS, &matkul.Jurusan, &matkul.Fakultas, &matkul.SemMin, &matkul.Prediksi)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the matkul in the matkuls slice
		matkuls = append(matkuls, matkul)
	}

	// check for any errors during iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
	}

	// return empty matkul on error
	return matkuls, nil
}
