FROM --platform=linux/amd64 golang:1.20-alpine3.18
WORKDIR /app
COPY . .
RUN go build -o krs-agg
EXPOSE 8080
CMD ./krs-agg