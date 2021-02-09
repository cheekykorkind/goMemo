FROM golang:1.15-alpine3.12

RUN apk add sudo

ARG UID
ARG GID
ARG UNAME

ENV UID ${UID}
ENV GID ${GID}
ENV UNAME ${UNAME}

ENV GO111MODULE=on

RUN adduser -D -u ${UID} -s /bin/sh -G wheel ${UNAME}
RUN echo '%wheel ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers
