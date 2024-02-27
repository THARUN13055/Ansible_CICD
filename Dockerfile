FROM golang:latest AS Builder

WORKDIR /golang

COPY /Golang/* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

#main image

FROM alpine:latest

COPY --from=Builder /golang/main ./

EXPOSE 8080

CMD [ "./main" ]

