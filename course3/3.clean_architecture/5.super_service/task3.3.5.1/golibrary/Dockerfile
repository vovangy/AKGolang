FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./

RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o golibrary .

EXPOSE 8080

CMD ["./golibrary"]
