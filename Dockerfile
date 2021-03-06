ARG BINNAME="app"

FROM golang:bullseye as gobuild

ARG BINNAME

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto
#ENV GOBIN=$GOPATH/bin
#ENV GOPATH=/go
#ENV BINNAME="app"

#WORKDIR /build

COPY . .

#RUN go env
#RUN go get github.com/gorilla/mux
RUN go get github.com/gin-gonic/gin
RUN go get k8s.io/api/admission/v1
RUN go build -o ${BINNAME} main.go

RUN du -hs *

FROM scratch
ARG BINNAME

COPY --from=gobuild /go/${BINNAME} /bin/

CMD ["/bin/app"]
