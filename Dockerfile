FROM golang:1.17.9-alpine as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make gcc \
    libc-dev tzdata && \
    mkdir /app

WORKDIR /app

ENV TZ=Asia/Jakarta

COPY . .

RUN go mod download

RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "./bin/apiactivity" .

EXPOSE 3030

CMD ["./bin/apiactivity"]
