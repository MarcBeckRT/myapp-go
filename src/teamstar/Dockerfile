FROM golang:1.16-buster
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go install
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["teamstar"]
EXPOSE 8080