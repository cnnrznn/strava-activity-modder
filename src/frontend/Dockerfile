FROM golang:1.14

WORKDIR /src/frontend

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /app/frontend

ENTRYPOINT ["/app/frontend"]