FROM golang:1.23

WORKDIR /app

RUN go install golang.org/x/vuln/cmd/govulncheck@latest

COPY . .

RUN go build -o main main.go

RUN ls -l /app

ENTRYPOINT ["/app/main"]    