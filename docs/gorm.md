### GORM Official Website 
```
https://gorm.io/
```

### Initialize App
```
go mod init
```
- envpackage
```
go get github.com/joho/godotenv
```
- install latest gorm package
```
go get -u gorm.io/gorm
```
- postgresql driver
```
go get -u gorm.io/driver/postgres
```

### migration
```
go run main.go migrate
```