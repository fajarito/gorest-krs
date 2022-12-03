FROM golang:1.17-alpine3.15
WORKDIR /app
COPY . .
RUN go build -o gorest-pk
EXPOSE 8080
CMD ./gorest-pk