FROM golang:latest as builder

WORKDIR  /home/app

ADD . .

RUN go mod tidy

RUN go build -o bin/app src/main.go

RUN chmod +x bin/app

FROM debian:latest

ENV TZ=Asia/Shanghai

WORKDIR  /home/app

EXPOSE 8000

COPY --from=builder /home/app/bin/ ./

COPY --from=builder /home/app/src/docs ./docs

CMD ./app --swagger ./docs/swagger.json
