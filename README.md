file main อยู่ที่ workshop.go

แก้ไขรายละเอียด database ก่อนรันโค้ด ได้ที่ workshop.go
```go
dsn := "user:password@tcp(localhost:3306)/databaseName?charset=utf8mb4&parseTime=True&loc=Local"
```

คำสั่งในการรันโปรเจค
```bash
go mod tidy
go run workshop.go
```