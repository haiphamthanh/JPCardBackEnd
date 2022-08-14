# JPCardBackEnd

`This is a project to handle japanse learning logic`

## -----> GO-LANG <-----

```Go
// To int project with `github.com/boilerplate` package name
go mod init github.com/boilerplate
```

```Go
// To get new module
go get -u github.com/julienschmidt/httprouter
```

```Go
// To download all modules in `mod` file (mod file like pod.lock in cocoapod)
go mod download
```

### [Migrate CLI link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
```Go
// To install `migrate CLI`
brew install golang-migrate
``` 

## -----> DOCKER <-----

```bash
# Remove all containers
docker container rm -f $(docker container ps -a -q)
```

```bash
# Clear volumes
docker volume prune -f
```


```bash
# Start up docker-compose
docker-compose up --build
```

## -----> MARKER <-----

```bash
# TODO: Make this bash as utils to create multilple project
# To create new project (Check template sh files)
sh make
```

## -----> PERMISSION <-----
```bash
# To give a grant access to db folder for postgres tool
chmod -R 777 db/migrations
chmod -R 777 db/scripts
```


## -----> Reference links <-----
1. [Create rest api layout golang (part-1)](https://hackernoon.com/how-to-create-golang-rest-api-project-layout-configuration-part-1-am733yi7?ref=hackernoon.com)
2. [Create rest api layout golang (part-2)](https://hackernoon.com/how-to-create-golang-rest-api-project-layout-configuration-part-2-wh2z3y5z)
3. [Create rest api layout golang (part-3)](https://hackernoon.com/how-to-create-golang-rest-api-project-layout-configuration-part-3-pr453ylt)
4. []()
5. []()
6. []()
7. []()
8. []()
9. []()
10. []()


## -----> Resolve problems links <-----
1. [Scripts permission problems](https://stackoverflow.com/questions/66229898/docker-entrypoint-initdb-permission-denied)
2. [Chmod là gì??](https://blog.hostvn.net/chia-se/chmod-la-gi-cach-chmod-nhu-the-nao.html)
3. []()
4. []()
5. []()
6. []()
7. []()
8. []()
9. []()
10. []()

psql -U postgres -W