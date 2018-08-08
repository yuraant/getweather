FROM golang:1.9.3-alpine
WORKDIR /go/src/getweather
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...


CMD ["getweather"]