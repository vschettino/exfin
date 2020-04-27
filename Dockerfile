FROM golang:alpine

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
WORKDIR /exfin
RUN chmod +x /wait
COPY . .
RUN apk add git
RUN go get -d -v ./...
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-pg/pg/v9
RUN go get -u github.com/go-pg/migrations/v7
RUN go get -u github.com/appleboy/gin-jwt/v2
RUN go get -u github.com/stretchr/testify
CMD /wait && /exfin/scripts/entrypoint.sh