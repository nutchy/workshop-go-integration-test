```
$ docker build -t go-test-junit -f docker/Dockerfile .
$ docker run -v $(pwd):/app go-test-junit
```

```
$ docker compose -f docker/docker-compose.yml up --build --abort-on-container-exit
```