FROM golang:1.21.1

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GO mod tidy

CMD ["go", "run", "cmd/app/app.go"]