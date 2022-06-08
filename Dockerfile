FROM REPOSITORY_BASE_URI/golang:1.18.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum main.go /app/
COPY controller  /app/controller
COPY core /app/core
COPY model /app/model

RUN apk update \
    && apk add --no-cache git \
    && go get -d -v \
    && go build -o go-rest-api

#----

FROM REPOSITORY_BASE_URI/alpine:latest
LABEL "author"="Himanshu Shekhar <himanshu.shekhar.in@gmail.com>"

WORKDIR /app

COPY --from=builder /app/go-rest-api ./

CMD [ "/app/go-rest-api" ]
