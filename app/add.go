package app

import (
	"ClickHitsCount/infrastructure/database"
)

func AddAccessIP(uid int, ip string) error {
	_, err := database.HitsDatbase.Stmt["CreateCounts"].Exec(ip, uid)
	if err != nil {
		return err
	}
	return nil
}
