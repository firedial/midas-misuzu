FROM golang:stretch

WORKDIR /go/src/github.com/firedial/midas-misuzu
RUN useradd midas
RUN chown midas:midas -R /go/
RUN mkdir /home/midas
RUN chown midas:midas -R /home/midas/

USER midas

RUN go get github.com/gin-contrib/cors
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql

COPY . .
# COPY env.go midas-misuzu/config

RUN go build ./main.go

CMD ["./main"]

