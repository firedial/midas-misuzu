FROM golang:stretch

WORKDIR /go/src/github.com/firedial/midas-misuzu

# 後で調整する
# RUN useradd midas
# RUN chown midas:midas -R /go/
# RUN mkdir /home/midas
# RUN chown midas:midas -R /home/midas/

# USER midas

COPY go.mod ./
COPY go.sum ./

RUN go mod download

