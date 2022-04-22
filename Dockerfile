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

#RUN go env
RUN go get github.com/gorilla/mux
RUN go get github.com/gin-gonic/gin
RUN go build -o ${BINNAME} main.go


FROM scratch
ENV BINNAME="app"

COPY --from=gobuild /go/${BINNAME} /bin/

CMD ["/bin/app"]
