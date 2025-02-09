FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o pinger .
EXPOSE 7000
CMD ["./pinger"]
