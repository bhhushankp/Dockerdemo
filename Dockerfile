FROM golang:1.21.6-alpine3.19

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8181

CMD [ "./main" ]

