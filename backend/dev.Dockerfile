FROM golang:1.15.6

WORKDIR /code

RUN go get -u github.com/cosmtrek/air

COPY go.mod go.sum /code/
RUN go mod download
COPY . /code

CMD air
