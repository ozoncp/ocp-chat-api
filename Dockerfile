FROM golang:1.16.5-alpine3.13 as build-stage

RUN set -e

ARG SRC_PATH

RUN apk add make git --no-cache

ADD . ${SRC_PATH}
#todo don't forget dockerignore
WORKDIR ${SRC_PATH}
RUN make build

FROM alpine:3.13

ARG RUNNING_BIN=ocp-chat-api
ARG SRC_PATH

COPY --from=build-stage ${SRC_PATH}/bin/${RUNNING_BIN} /usr/local/bin/${RUNNING_BIN}
COPY --from=build-stage ${SRC_PATH}/migrations ${SRC_PATH}/migrations
WORKDIR ${SRC_PATH}

EXPOSE 1337

CMD ["ocp-chat-api"]