package metaldb

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

type MetalRecommendation struct {
	URL        string `db:"url"`
	VideoID    string `db:"video_id"`
	VideoTitle string `db:"video_title"`
	Timestamp  string `db:"timestamp"`
}

type MetalDatabase struct {
	DB *sqlx.DB
}

func NewMetalDatabase(
	host string,
	port int,
	dbname string,
	dbuser string,
	dbpass string) (*MetalDatabase, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port,
		dbuser, dbpass, dbname)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		return nil, errors.Errorf("error creating database with info: %s", dbinfo)
	}

	return &MetalDatabase{
		DB: db,
	}, nil
}

func (db *MetalDatabase) StoreRecommendation(mr *MetalRecommendation) error {
	tsNow := time.Now().Format(time.RFC3339)

	query := (`
	INSERT INTO metal_links (
		video_id, 
		video_title, 
		url, 
		timestamp
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) ON CONFLICT DO NOTHING`)

	tx := db.DB.MustBegin()
	db.DB.MustExec(query, mr.VideoID, mr.VideoTitle, mr.URL, tsNow)
	err := tx.Commit()
	if err != nil {
		return errors.Wrapf(err, "could not insert values via: %v", query)
	}

	return nil
}

func (db *MetalDatabase) GetRecommendations() ([]MetalRecommendation, error) {
	query := (`SELECT * FROM metal_links ORDER BY id`)

	rows, err := db.DB.Queryx(query)
	if err != nil {
		log.WithField("query_error", err.Error()).Info("error selecting query")
		return nil, errors.Errorf("error executing select query")
	}
	defer rows.Close()

	var recommendations []MetalRecommendation
	var recommendation MetalRecommendation
	for rows.Next() {
		err = rows.StructScan(&recommendation)
		if err != nil {
			//log.Debug("error scanning row")
		}
		recommendations = append(recommendations, recommendation)
	}

	for _, r := range recommendations {
		//fmt.Printf("%v\n", p)
		fmt.Println(r.URL, r.VideoID, r.VideoTitle, r.Timestamp)
	}

	return recommendations, nil
}
