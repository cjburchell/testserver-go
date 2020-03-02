FROM golang:1.13 as serverbuilder
WORKDIR /test
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=serverbuilder /test/main  /server/main

WORKDIR  /server

CMD ["./main"]
