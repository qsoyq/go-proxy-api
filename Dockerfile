FROM golang:alpine AS builder

WORKDIR /home/app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/

RUN go build -o bin/app src/main.go

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    apt-get clean

ENV TZ=Asia/Shanghai

WORKDIR /home/app

EXPOSE 8000

COPY --from=builder /home/app/bin/app ./
COPY --from=builder /home/app/src/docs ./docs

CMD ["./app", "--swagger", "./docs/swagger.json"]
