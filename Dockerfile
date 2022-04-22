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
RUN ls -l
RUN go build
# -o goapp main.go


FROM scratch
ENV BINNAME="app"

COPY --from=gobuildNO /go/${BINNAME} /bin/

CMD ["/go/${BINNAME}"]
