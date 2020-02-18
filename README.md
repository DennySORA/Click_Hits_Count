# Simple Click Hits Count 

<p>
<img src="https://img.shields.io/badge/Golang-1.13-brightgreen">
<img src="https://img.shields.io/badge/Database-SQLite-brightgreen">
<img src="https://img.shields.io/badge/license-Apache--2.0-blue">
<img src="https://hits.dennysora.me/hits?name=ClickHitsCount&amp;chapter=0">
</p>

## Demo

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
http://127.0.0.1:8223/create/chapter?name=Testid=Test&ep=0&chapter=0
```
SP. **id** 
1. name is any name.
2. id name need exist in Novel.
3. ep is any number(int).
4. chapter is any number(int).

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
