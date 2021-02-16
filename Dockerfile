From golang:latest

RUN apt-get update && apt-get install -y
WORKDIR /temp/work

ENV TZ=Asia/Tokyo
ENV USER user1

RUN useradd -m ${USER} --uid 1000
RUN gpasswd -a ${USER} sudo
RUN echo "${USER}:password" | chpasswd

RUN go get -v golang.org/x/tools/gopls
RUN go get -v github.com/go-delve/delve/cmd/dlv
RUN go get github.com/gin-gonic/gin
USER ${USER}