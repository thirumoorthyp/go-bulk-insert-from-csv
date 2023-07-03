# Go-lang CSV import performance test
This go-lang code is used to bulk import data from CSV file into the SQLite database.

# Manual Steps: 
``` shell
apt install golang
```

``` shell
go version
```
# go version go1.20.3

# config
``` shell
 set GOOS=linux
 set GOARCH=amd64
 set CGO_ENABLED=0
```

# install sqlite
``` shell
go get github.com/modernc.org/sqlite
```

# executing go program
```shell
go run main.go
```

This go-lang code is used to bulk import data from CSV file into the SQLite database.

It took ***2 sec*** to import 5k records. 

It look long time compared to Python and PHP. 
Slow performance on bulk insert - https://github.com/mattn/go-sqlite3/issues/1145
So tried with single row insert with same go-sqlite3 cgo package. It took ***1.5 sec***(https://github.com/thirumoorthyp/go-lang-csv-import-sqlite.)

Also Tried performance test in Python and PHP 

Python Import: https://github.com/thirumoorthyp/import_csv_into_sqlite_using_python. It took ***250ms*** for 5k record

PHP: https://github.com/thirumoorthyp/php_import_csv_into_sqlite. It took ***481ms*** for 5k records

***Conclusion:***
In terms of raw performance, Go tends to be faster than Python because it is compiled to machine code, whereas Python is an interpreted language. Additionally, Go was designed with a focus on efficiency and concurrent programming, making it well-suited for handling concurrent tasks and parallelism.

However, when it comes to database connections and performance, the key factor is often not the language itself but rather the database driver or library being used to connect and interact with the database. Both Go and Python have mature and efficient database libraries that can connect to various database systems.
