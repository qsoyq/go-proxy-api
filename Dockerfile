FROM golang:latest as builder

WORKDIR  /home/app

ADD . .

RUN go build -o bin/main.app src/main.go

RUN chmod +x bin/main.app

FROM debian:latest

WORKDIR  /home/app

EXPOSE 8000

COPY --from=builder /home/app/bin/ ./

CMD ./main.app
