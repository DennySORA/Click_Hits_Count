package app

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOneValue(row *sql.Row) (int, error) {
	var data int
	if err := row.Scan(&data); err != nil {
		return 0, err
	}
	return data, nil
}

func CheckString(str string) interface{} {
	if str == "" {
		return nil
	}
	return str
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
