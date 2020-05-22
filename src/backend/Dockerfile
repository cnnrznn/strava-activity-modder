FROM golang:1.14

WORKDIR /src/backend

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /app/backend

ENTRYPOINT ["/app/backend"]