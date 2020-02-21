package app

import (
	"ClickHitsCount/infrastructure/database"
)

func GetNovelUIDFromDatabase(id string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetNovelUIDFromDatabase"].QueryRow(id))
}

func GetChapterUIDFromDatabase(name string, chapter string) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetChapterUIDFromDatabase"].QueryRow(name, chapter))
}

func GetHitsFromDatabase(uid int) (int, error) {
	return GetOneValue(database.HitsDatbase.Stmt["GetHitsFromDatabase"].QueryRow(uid))
}

func GetChapterIPFromDatabase(novel Novel) ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetChapterIPFromDatabase"].Query(
		novel.chapter_id, novel.chapter_id,
		novel.novel_name, novel.novel_name,
		novel.chapter_name, novel.chapter_name,
		novel.ipaddress, novel.ipaddress,
		novel.episode, novel.episode,
		novel.chapter, novel.chapter,
	))
}

func GetChapterHitsFromDatabase(novel Novel) ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetChapterHitsFromDatabase"].Query(
		novel.chapter_id, novel.chapter_id,
		novel.novel_name, novel.novel_name,
		novel.chapter_name, novel.chapter_name,
		novel.ipaddress, novel.ipaddress,
		novel.episode, novel.episode,
		novel.chapter, novel.chapter,
	))
}

func GetNovelHitsFromDatabase(novelName interface{}) ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetNovelHitsHitsFromDatabase"].Query(novelName, novelName))
}

func GetEpisodeHitsFromDatabase(novelName interface{})([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetEpisodeHitsFromDatabase"].Query(novelName, novelName))
}

func GetAllHitsIPFromDatabase() ([]map[string]interface{}, error) {
	return GetAllValue(database.HitsDatbase.Stmt["GetAllHitsIPFromDatabase"].Query())
}
