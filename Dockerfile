FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /github.com/ssoyyoung.p/BlackHeart-Golang

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN go build -o ./github.com/ssoyyoung.p/BlackHeart-Golang .


EXPOSE 8080

#CMD ["./github.com/ssoyyoung.p/BlackHeart-Golang"]
CMD go run main.go
