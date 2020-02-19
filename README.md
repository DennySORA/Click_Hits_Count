# Simple Click Hits Count 

<p>
<img src="https://img.shields.io/badge/Golang-1.13-brightgreen">
<img src="https://img.shields.io/badge/Database-SQLite-brightgreen">
<img src="https://img.shields.io/badge/license-Apache--2.0-blue">
</p>

[![HitCount](https://hits.dennysora.me/hits?name=ClickHitsCount&chapter=0)](https://hits.dennysora.me/hits?name=ClickHitsCount&chapter=0)

## Demo

https://hits.dennysora.me/hits?name=ClickHitsCount&chapter=0

**Github has cache so...this is not bug, Is feature.**

https://hackmd.io/@dennySORA/ClickHitsCount

## Simple install
### Git

```git clone git@github.com:dennySORA/Click_Hits_Count ```

### TLS/SSL

Plesae put TLS/SSL in ssl folder in root folder.

Need to rename server in crt and key.

Example:
```
    root folder
         |d------- app
         |d------- command
         |d------- infrastructure
         |d------- server
         |d------- ssl (**Need Create**)
              |f------- sercer.key (**Rename**)
              |f------- server.crt (**Rename**)
         |f------- go.mod
         |f------- go.sum
         |f------- main.go
```

---

## Simple to use

### From Local

#### Create Novel

```
http://127.0.0.1:8223/create/novel?name=Test
```
1. name is any name.

#### Create chapter

```
http://127.0.0.1:8223/create/chapter?name=Test&id=Test&ep=0&chapter=0
```
SP. **id is create novel name** 
1. name is any name.
2. id name need exist in Novel.
3. ep is any number(int).
4. chapter is any number(int).

#### Get hits data 

##### Get all data
```
http://127.0.0.1:8223/get/all
```
1. This is get all data.


##### Get all chapter hits data

```
http://127.0.0.1:8223/get/all/chapter/hits
```
1. This is get all chapter hits counts data.

##### Get chapter detailed data 

```
http://127.0.0.1:8223/get/chapter/hits?name=Test&chapter=0
```

**If name or chapter not input parameters then return all chapter detailed data.**

1. name is "**Novel Name**".
2. chapter is "**Create chapter number(int)**".
3. This is get ip hits and hits last time.

### From Open

```
https://127.0.0.1:8123/hits?name=Test&chapter=0
```

1. name is "**Novel Name**".
2. chapter is "**Create chapter number(int)**".

---

## Contributors 
- **DennySORA**
    - [GitHub](https://github.com/dennySORA)
    - Email : dennysora.main@gmail.com
