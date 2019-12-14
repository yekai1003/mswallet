FROM golang:1.12-stretch

WORKDIR $GOPATH/src/hdwallet

ADD . $GOPATH/src/hdwallet 

ENV GO111MODULE on 
ENV GOPROXY https://goproxy.io

RUN go build . 

RUN cp ./hdwallet /go/bin 

ENTRYPOINT ["/bin/bash"]