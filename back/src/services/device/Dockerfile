FROM golang:alpine
RUN mkdir /service
COPY . /service
WORKDIR /service
RUN go build -o main . 
CMD ["/service/main"]