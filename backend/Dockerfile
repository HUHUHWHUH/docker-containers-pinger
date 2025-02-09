FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o backend .
EXPOSE 8000
CMD ["./backend"]