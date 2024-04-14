FROM golang:latest

WORKDIR /app

COPY ./main .
COPY conf .

CMD ["./main"]