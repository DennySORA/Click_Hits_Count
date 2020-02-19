CREATE TABLE IF NOT EXISTS IP(
  uid INTEGER PRIMARY KEY AUTOINCREMENT,
  IPAddress TEXT NOT NULL UNIQUE,
  Novel_ID INTEGER NULL UNIQUE,
  Counts INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS Novel(
  uid INTEGER PRIMARY KEY AUTOINCREMENT,
  Novel TEXT NOT NULL UNIQUE
);


SELECT uid FROM Novel_Chapter WHERE 
  Novel_ID = (SELECT uid as id FROM Novel WHERE Name = ?)
AND
  Chapter = ?

SELECT uid FROM hits WHERE `Novel` = ?

SELECT Counts FROM IP WHERE IPAddress = ? AND Novel_ID = ? 

INSERT INTO IP(IPAddress,Novel_ID,Counts) VALUES ("127.0.0.1","557","1");

UPDATE IP SET Counts=(SELECT Counts AS cou FROM IP WHERE IPAddress = "127.0.0.1")+1 WHERE IPAddress = "127.0.0.1";

SELECT count(*) as hist FROM IP WHERE Novel_ID = ?;


SELECT 
       ip.createtime, 
       ip.ipaddress, 
       ip.chapter_id, 
       novel_chapter.name, 
       novel_chapter.episode, 
       novel_chapter.novel_id, 
       novel_chapter.chapter, 
       novel.name 
FROM   ip 
       LEFT JOIN novel_chapter 
              ON ip.chapter_id = novel_chapter.uid 
       LEFT JOIN novel 
              ON novel_chapter.novel_id = novel.uid 
ORDER  BY ip.createtime DESC;


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
	WHERE  ( Ifnull(10, 'empty') = 'empty' 
		OR chapter_id =10 ) 
	GROUP  BY ipaddress 
	ORDER  BY "Hits Counts" DESC; 


  
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
        WHERE  ( Ifnull(NULL, 'empty') = 'empty' 
                  OR chapter_id = 13 ) 
        GROUP  BY ipaddress, 
                  chapter_id) 
GROUP  BY chapter_id 
ORDER  BY "hits counts" DESC; 