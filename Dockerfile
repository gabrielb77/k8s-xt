#FROM golang:1.18.1-alpine3.15 as gobuild
FROM golang:bullseye as gobuild

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto
ENV GOBIN=$GOPATH/bin
#ENV GOPATH=/go
ENV BINNAME="app"

#WORKDIR /build

COPY . .

RUN go env
RUN go get
RUN go build -o ${BINNAME} .


# ---
FROM scratch AS run
ENV BINNAME="app"

COPY --from=gobuild /go/${BINNAME} /bin/

CMD ["${BINNAME}"]
