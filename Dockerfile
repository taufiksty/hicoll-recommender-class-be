FROM golang:1.23.4 as builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

CMD [ "./server" ]