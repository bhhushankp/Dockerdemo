FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8181

CMD [ "./main" ]

