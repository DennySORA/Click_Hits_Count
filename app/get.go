package app

import (
	"ClickHitsCount/infrastructure/database"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetNovelUID(id string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetNovel"].QueryRow(id))
}

func GetChapterUID(name string, chapter string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetChapter"].QueryRow(name, chapter))
}

func GetNovelHits(uid int) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetHits"].QueryRow(uid))
}

func GetChapterHitsCounts(chapterID int) ([]map[string]interface{}, error) {
	if chapterID == -1 {
		return GetAllValue(database.HitsDatbase.Stmt["GetChapterHitsCounts"].Query(nil, nil))
	}
	return GetAllValue(database.HitsDatbase.Stmt["GetChapterHitsCounts"].Query(chapterID, chapterID))
}

func GetAllChapterHitsCounts() ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetAllChapterHitsCount"].Query())
}

func GetAllHitsData() ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetAllData"].Query())
}

func GetOneValue(row *sql.Row) (int, error) {
	var data int
	if err := row.Scan(&data); err != nil {
		return 0, err
	}
	return data, nil
}

func GetAllValue(rows *sql.Rows, err error) ([]map[string]interface{}, error) {
	// Cheakc error.
	if err != nil {
		return nil, err
	}
	// Create return box.
	returnBox := []map[string]interface{}{}
	// Create value box.
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtr := make([]interface{}, count)
	// Get values.
	for rows.Next() {
		// -------------------[Link point make page]
		for i, _ := range columns {
			valuePtr[i] = &values[i]
		}
		// -------------------[For all page get data]
		err := rows.Scan(valuePtr...)
		if err != nil {
			return nil, err
		}
		box := map[string]interface{}{}
		// -------------------[Classification key and value]
		for i, col := range columns {
			b, ok := values[i].([]byte)
			// ---------------[Regex relpace key word]
			if ok {
				// -----------[Byte convter to string]
				box[col] = string(b)
			} else {
				box[col] = values[i]
			}
		}
		returnBox = append(returnBox, box)
	}
	return returnBox, nil
}

func GetImage(count int) ([]byte, error) {
	url := fmt.Sprintf("https://img.shields.io/badge/hits-%d-green", count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
