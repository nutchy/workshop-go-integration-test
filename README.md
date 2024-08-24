## Start docker compose for local developement
```
$ docker compose up -f docker/docker-compose.yml up
```

```
$ docker build -t go-test-junit -f docker/Dockerfile .
$ docker run -v $(pwd):/app go-test-junit
```

## Run integration test using docker compose
```
$ docker compose -f docker/docker-compose.yml up --build --abort-on-container-exit
```