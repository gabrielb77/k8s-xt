FROM golang:bullseye as gobuild

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto
#ENV GOBIN=$GOPATH/bin
#ENV GOPATH=/go
ENV BINNAME="app"

#WORKDIR /build

COPY . .

RUN type go
RUN go version
RUN pwd
RUN ls -lR
#RUN go env
#RUN go get
RUN ls -l bin

RUN go build -o ${BINNAME} main.go

RUN ls -lR

FROM scratch
ENV BINNAME="app"

COPY --from=gobuild ${BINNAME} /bin/

CMD ["${BINNAME}"]
