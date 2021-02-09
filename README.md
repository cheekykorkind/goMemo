[한국어](README.kr.md)
[日本語](README.jp.md)
# go-memo

## command sequence

### start docker compose
- `DOCKER_UID=$(id -u $USER) DOCKER_GID=$(id -g $USER) docker-compose up -d --build`

### after docker-compose up -d --build
- write terraform template at host machine `./workspace`
- execute terraform template 
  - in docker container
    - `docker exec -it go-memo /bin/sh`
  - outside of docker container
    - `docker exec go-memo go run main.go`

