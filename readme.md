# Golang ready-to-go template 

> based on [`echo`](https://echo.labstack.com/guide/) router


### Quick launch

1. to rename your project replace "go-template" in go.mod and in all imports (optional)
2. duplicate ".env.example" and rename it to ".env"

```sh
# build and run app (auto download deps)
> ./scripts/run.bat

# bundle prod app into folder (it can take some time at first time)
> ./scripts/build.bat
```
<br />

### [Flags](https://pkg.go.dev/flag)

```sh
# default value = 8000
> ./app --port 8000 
```
<br />


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