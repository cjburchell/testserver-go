FROM golang:1.12-alpine as serverbuilder
WORKDIR /testserver-go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=serverbuilder /testserver-go/main  /server/main

WORKDIR  /server

CMD ["./main"]
