FROM golang:1.16

LABEL maintainer="carlostrejo2308"

WORKDIR /usr/local/go/app

COPY . .

RUN go mod init github.com/carlostrejo2308/gomongo && go get github.com/CarlosTrejo2308/peopleApiResource@v1.0.0 && go mod tidy && go mod vendor && go mod download

RUN go build -o ./out/go-app .

CMD ["./out/go-app"]