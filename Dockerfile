FROM golang:latest
EXPOSE 8080

COPY ./ ./
RUN go build -o main
CMD ["./main"]