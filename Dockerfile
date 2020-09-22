FROM golang:1.14

WORKDIR einstein/src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["einsten"]
