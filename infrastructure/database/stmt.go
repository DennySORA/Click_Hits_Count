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

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT 	
		novel_name as "Novel Name",
		chapter_name as "Chapter Name", 
		episode as "Episode", 
		chapter as "Chapter", 
		createtime as "Update Time", 
		ipaddress as "IP Address",  
		Count(*) AS "Hits Counts"
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
		OR chapter_id =? ) 
	GROUP  BY ipaddress 
	ORDER  BY "Hits Counts" DESC;  
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetChapterHitsCounts"] = stmt
	}

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
		HitsDatbase.Stmt["GetAllData"] = stmt
	}

	if stmt, err := HitsDatbase.DB.Prepare(`
	SELECT novel_name   AS "Novel Name", 
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
	GROUP  BY chapter_id 
	ORDER  BY "hits counts" DESC; 
	`); err != nil {
		return err
	} else {
		HitsDatbase.Stmt["GetAllChapterHitsCount"] = stmt
	}

	return nil
}
