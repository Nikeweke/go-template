# Sturm - golang ready-to-go template 

> based on [`echo`](https://echo.labstack.com/guide/) router



### Quick launch

```sh
# build and run app
> ./scripts/run.bat

# bundle prod app into folder 
> ./scripts/release.bat
```


### Golang migrate  

> Download exe from releases choose `amd64` from - https://github.com/golang-migrate/migrate/releases and choose

1. Create migration 

```sh
migrate create -ext .sql -dir migrations [name-migration]
```

2. Migrate or use `scripts/migrate-up.bat` or `scripts/migrate-down.bat`

> be sure place migrate.exe into `./bin/migrate.exe`

```sh
migrate.exe -path ./migrations -database "mysql://root:12345@/test" -verbose up
```

<br />