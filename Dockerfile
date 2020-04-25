FROM golang:alpine

WORKDIR /exfin
COPY . .
RUN apk add git
RUN go get -d -v ./...
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-pg/pg/v9
RUN go get -u github.com/go-pg/migrations/v7
RUN go get -u github.com/appleboy/gin-jwt/v2
CMD ["./scripts/entrypoint.sh"]