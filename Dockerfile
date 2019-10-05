FROM golang:1.13.1-alpine3.10
RUN mkdir /home/app
COPY . /home/app
WORKDIR /home/app
RUN go build main.go
ENV environment=development
CMD ["./main"]