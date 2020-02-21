package app

import (
	"ClickHitsCount/infrastructure/database"
)

func AddIPAccess(uid int, ip string) error {
	_, err := database.HitsDatbase.Stmt["AddIPAccessFromDatabase"].Exec(ip, uid)
	if err != nil {
		return err
	}
	return nil
}
