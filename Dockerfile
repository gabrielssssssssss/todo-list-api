FROM golang:1.24.10

WORKDIR /app

ENV GODEBUG=netdns=go+1
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o to-do-list-api ./cmd/to-do-list-api /main.go

CMD sh -c "docker run --name to-do-list-api -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres && ./to-do-list-api"

# CREATE TABLE tasks(
#     id SERIAL PRIMARY KEY,
#     title VARCHAR(1000) NOT NULL,
#     description VARCHAR(100) NOT NULL,
#     status VARCHAR(10) DEFAULT "todo",
#     createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
#     updateAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
# );