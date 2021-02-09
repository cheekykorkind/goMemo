#!/bin/bash

docker-compose down
# docker container stop $(docker container ls -aq)
# docker container rm $(docker container ls -aq)

# yes | docker system prune --volumes

# docker rmi $(docker images -q) --force

# sudo rm -rf ./workspace

DOCKER_UID=$(id -u $USER) DOCKER_GID=$(id -g $USER) docker-compose up --build