FROM golang:1.17-alpine

## buat folder yang berisi file kita
RUN mkdir /app

## set direktori kerja yang akan digunakan
WORKDIR /app

## copy seluruh file dari root ke /app
ADD . /app

## jalankan build untuk memebuat executeable
RUN go build -o main .

## run executeable yang sudah dibuat
CMD ["/app/main"]