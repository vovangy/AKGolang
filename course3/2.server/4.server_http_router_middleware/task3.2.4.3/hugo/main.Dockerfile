FROM golang:1.19.0-alpine AS builder

COPY go.mod go.sum /github.com/AKGolang/
WORKDIR /github.com/AKGolang/

RUN go mod download

COPY . .

#RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./main.go

FROM scratch AS runner

WORKDIR /docker-avtomakon/

COPY --from=builder /github.com/AKGolang/.bin .

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip

EXPOSE 80 443

ENTRYPOINT ["./.bin"]