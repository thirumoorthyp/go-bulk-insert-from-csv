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

This go-lang code took ***2 sec*** to import 5k records from CSV file into the SQLite database. 

It look long time compared to Python(https://github.com/thirumoorthyp/import_csv_into_sqlite_using_python. It took ***250ms***) and PHP(https://github.com/thirumoorthyp/php_import_csv_into_sqlite. It took ***481ms***) for importing 5k record. So tried with 1 record insert with same go-sqlite3 cgo package. It took ***1.5 sec***(https://github.com/thirumoorthyp/go-lang-csv-import-sqlite).

So the huge performance difference is because of go-sqlite3 cgo package. There is an existing issue raised(https://github.com/mattn/go-sqlite3/issues/1145)

***Conclusion:***
SQLite is a file-based database engine that provides a lightweight and embedded solution for managing databases. It is designed to be simple and self-contained, making it suitable for smaller-scale applications or scenarios where a full client-server database is not required.

When it comes to connecting to a SQLite database, Python offers the sqlite3 module in its standard library, which provides a convenient way to interact with SQLite databases. The sqlite3 module is implemented in C and is highly optimized for performance.

On the other hand, Go also provides a package called database/sql in its standard library that allows you to work with various databases, including SQLite. Go's database/sql package leverages the go-sqlite3 driver, which is a popular and efficient SQLite driver for Go.

In terms of performance, both Python and Go can achieve satisfactory results when connecting to a SQLite database. However, there may be cases where one language performs better than the other due to differences in implementation details, such as how the database driver is designed or how the code is written.

In terms of raw performance, Go tends to be faster than Python because it is compiled to machine code, whereas Python is an interpreted language. Additionally, Go was designed with a focus on efficiency and concurrent programming, making it well-suited for handling concurrent tasks and parallelism.



