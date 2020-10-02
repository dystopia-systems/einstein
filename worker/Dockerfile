FROM golang:1.14

WORKDIR einstein-worker/src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["worker"]
