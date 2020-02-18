package database

func createStmt() error {
	if stmt, err := HitsDatbase.DB.Prepare(`
		SELECT uid FROM Novel WHERE Name = ?
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetNovel"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT uid FROM Novel_Chapter WHERE 
		Novel_ID = (SELECT uid as id FROM Novel WHERE Name = ?)
  	AND
		Chapter = ?
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetChapter"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO Novel(Name) VALUES (?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["CreateNovel"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO Novel_Chapter(Novel_ID,Episode,Chapter,Name) VALUES (?,?,?,?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["CreateNovelChapter"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO IP(IPAddress,Chapter_ID) VALUES (?,?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["CreateCounts"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT count(*) as hist FROM (SELECT DISTINCT IPAddress FROM IP WHERE Chapter_ID = ?);
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetHits"] = stmt
	}

	return nil
}
