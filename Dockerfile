FROM golang:alpine 

WORKDIR /app

COPY . .
COPY  .env /app

RUN go mod tidy 
RUN go build -o  gin

COPY gin /app

ENTRYPOINT [ "/app/gin" ]