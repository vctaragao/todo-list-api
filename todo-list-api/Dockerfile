FROM golang:1.20

WORKDIR /user/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]