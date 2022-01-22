FROM golang:1.17

RUN apt-get update && apt-get install -y
ENV APP /web
WORKDIR ${APP}

ENV TZ=Asia/Tokyo
ENV USER user1
ENV PATH ${PATH}:${APP}


RUN go get -v golang.org/x/tools/gopls
RUN go get -v github.com/go-delve/delve/cmd/dlv
RUN go get github.com/gin-gonic/gin