FROM golang:1.14

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /app/backend

ENTRYPOINT ["/app/backend"]