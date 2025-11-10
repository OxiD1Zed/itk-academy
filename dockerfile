FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -o itk-app ./cmd/main.go

CMD ["./itk-app"]