FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN go build -o /mrf

ENTRYPOINT [ "/mrf" ]
