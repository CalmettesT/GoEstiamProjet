package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type LogData struct {
	ID       int64
	DH       time.Time
	MF       string
	Argument string
	Statut   string
}

var db *sql.DB

func ConnectDataBase() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "go",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connexion réussie")

}

func AddLog(log LogData) (int64, error) {
	result, err := db.Exec("INSERT INTO journal (dh, mf, argument, statut) VALUES (?,?,?,?)", log.DH, log.MF, log.Argument, log.Statut)
	if err != nil {
		return 0, fmt.Errorf("addLog %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addLog %v", err)
	}

	return id, nil
}

// func albumByArtist(name string) ([]Album, error) {
// 	var albums []Album

// 	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
// 	if err != nil {
// 		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var album Album
// 		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
// 			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
// 		}

// 		albums = append(albums, album)
// 	}

// 	return albums, nil
// }
