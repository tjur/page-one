FROM golang:1.27.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY api/ api/

RUN CGO_ENABLED=0 GOOS=linux go build -o page-one

CMD ["./page-one"]
