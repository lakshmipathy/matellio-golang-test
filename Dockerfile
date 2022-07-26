FROM golang:1.18-alpine3.15
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
CMD ["./matellio-golang-test"]