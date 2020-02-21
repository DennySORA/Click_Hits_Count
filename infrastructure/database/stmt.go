package database

func createStmt() error {

	// Get Novel UID From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
		SELECT uid FROM Novel WHERE Name = ?
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetNovelUIDFromDatabase"] = stmt
	}

	// Get Chapter UID From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT uid FROM Novel_Chapter WHERE 
		Novel_ID = (SELECT uid as id FROM Novel WHERE Name = ?)
  	AND
		Chapter = ?
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetChapterUIDFromDatabase"] = stmt
	}

	// Get Chapter Hits From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT count(*) as hist FROM (SELECT DISTINCT IPAddress FROM IP WHERE Chapter_ID = ?);
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetHitsFromDatabase"] = stmt
	}

	// Get Chapter IP From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT chapter_id   AS "Chapter ID", 
		novel_name   AS "Novel Name", 
		chapter_name AS "Chapter Name", 
		episode      AS "Episode", 
		chapter      AS "Chapter", 
		createtime   AS "Update Time", 
		ipaddress    AS "IP Address", 
		Count(*)     AS "Hits Counts" 
	FROM   (SELECT ip.createtime, 
				ip.ipaddress, 
				ip.chapter_id, 
				novel_chapter.name AS "chapter_name", 
				novel_chapter.episode, 
				novel_chapter.novel_id, 
				novel_chapter.chapter, 
				novel.name         AS "novel_name" 
		FROM   ip 
				LEFT JOIN novel_chapter 
					ON ip.chapter_id = novel_chapter.uid 
				LEFT JOIN novel 
					ON novel_chapter.novel_id = novel.uid 
		ORDER  BY ip.createtime DESC) 
	WHERE  ( Ifnull(?, 'empty') = 'empty' 
		OR chapter_id = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR novel_name = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR chapter_name = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR ipaddress = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR episode = ? )
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR chapter = ? ) 
	GROUP  BY novel_id, 
		chapter_id, 
		ipaddress 
	ORDER  BY "hits counts" DESC; 
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetChapterIPFromDatabase"] = stmt
	}

	// Get Chapter Hits From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT 
	    chapter_id	 AS "Chapter ID",
		novel_name   AS "Novel Name", 
		chapter_name AS "Chapter Name", 
		episode      AS "Episode", 
		chapter      AS "Chapter", 
		createtime   AS "Last Access Time", 
		ipaddress    AS "Last Access IP Address", 
		Count(*)     AS "Hits Counts" 
	FROM   (SELECT * 
			FROM   (SELECT ip.createtime, 
						ip.ipaddress, 
						ip.chapter_id, 
						novel_chapter.name AS "chapter_name", 
						novel_chapter.episode, 
						novel_chapter.novel_id, 
						novel_chapter.chapter, 
						novel.name         AS "novel_name" 
					FROM   ip 
						LEFT JOIN novel_chapter 
								ON ip.chapter_id = novel_chapter.uid 
						LEFT JOIN novel 
								ON novel_chapter.novel_id = novel.uid 
					ORDER  BY ip.createtime DESC) 
			GROUP  BY ipaddress, 
					chapter_id) 
	WHERE 
		(( Ifnull(?, 'empty') = 'empty' 
		OR "Chapter ID" = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR novel_name = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR chapter_name = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR ipaddress = ? ) 
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR episode = ? )
		AND ( Ifnull(?, 'empty') = 'empty' 
			OR chapter = ? ))
	GROUP  BY "Chapter ID" 
	ORDER  BY "hits counts" DESC; 
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetChapterHitsFromDatabase"] = stmt
	}

	// Get Novel Hits From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT 
	    novel_id	 AS "Novel ID",
		novel_name   AS "Novel Name", 
		createtime   AS "Last Access Time", 
		ipaddress    AS "Last Access IP Address", 
		Count(*)     AS "Hits Counts" 
	FROM   (SELECT * 
			FROM   (SELECT 
						ip.createtime, 
						ip.ipaddress,
						ip.chapter_id,
						novel_chapter.novel_id,
						novel_chapter.Episode,
						novel.name         AS "novel_name"
					FROM   ip 
						LEFT JOIN novel_chapter 
								ON ip.chapter_id = novel_chapter.uid 
						LEFT JOIN novel 
								ON novel_chapter.novel_id = novel.uid 
					ORDER  BY ip.createtime DESC) 
			GROUP  BY ipaddress, 
					chapter_id) 
	WHERE 
	 	( Ifnull(?, 'empty') = 'empty' 
			OR "Novel Name" = ? )
		AND NOT Episode = 0
	GROUP  BY "Novel Name" 
	ORDER  BY "hits counts" DESC; 
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetNovelHitsHitsFromDatabase"] = stmt
	}

	// Get Episode Hits From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT 
		novel_id	 AS "Novel ID",
		novel_name   AS "Novel Name", 
		Episode,
		createtime   AS "Last Access Time", 
		ipaddress    AS "Last Access IP Address", 
		Count(*)     AS "Hits Counts" 
	FROM   (SELECT * 
			FROM   (SELECT 
						ip.createtime, 
						ip.ipaddress,
						ip.chapter_id,
						novel_chapter.novel_id,
						Novel_Chapter.Episode,
						novel.name         AS "novel_name"
					FROM   ip 
						LEFT JOIN novel_chapter 
								ON ip.chapter_id = novel_chapter.uid 
						LEFT JOIN novel 
								ON novel_chapter.novel_id = novel.uid 
					ORDER  BY ip.createtime DESC) 
			GROUP  BY ipaddress, 
					chapter_id) 
	WHERE 
	( Ifnull(?, 'empty') = 'empty' 
	OR "Novel Name" = ? )
	GROUP  BY "Novel Name","Episode" 
	ORDER  BY "hits counts" DESC; 
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetEpisodeHitsFromDatabase"] = stmt
	}

	// Get All Hits From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT
		novel.name as novel_name, 
		novel_chapter.name as chapter_name, 
		novel_chapter.episode, 
		novel_chapter.chapter, 
		ip.createtime, 
		ip.ipaddress, 
		ip.chapter_id
	FROM      ip 
	LEFT JOIN novel_chapter 
	ON        ip.chapter_id = novel_chapter.uid 
	LEFT JOIN novel 
	ON        novel_chapter.novel_id = novel.uid 
	ORDER BY  ip.createtime DESC
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetAllHitsIPFromDatabase"] = stmt
	}

	// Create Novel From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO Novel(Name) VALUES (?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["CreateNovelFromDatabase"] = stmt
	}

	// Create Chapter From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO Novel_Chapter(Novel_ID,Episode,Chapter,Name) VALUES (?,?,?,?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["CreateChapterFromDatabase"] = stmt
	}

	// Add IP Access From Database

	if stmt, err := HitsDatbase.DB.Prepare(`
	INSERT INTO IP(IPAddress,Chapter_ID) VALUES (?,?)
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["AddIPAccessFromDatabase"] = stmt
	}

	return nil
}
