FROM golang:1.21.5-alpine3.17

WORKDIR /app 

COPY .  .

RUN go get -d -v ./...
# build Image
RUN go build -o api .


EXPOSE 8085
CMD ["./api"]