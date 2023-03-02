# Multi-stage gw-mtls-server build
# Многоэтапная сборка gw-mtls-service

FROM golang AS build

ENV location /go/src/github.com/blablatov/mtls-grpc-gateway

WORKDIR ${location}/gw-mtls-service

ADD ./main.go ${location}/gw-mtls-service
ADD ./gw-mtls-proto  ${location}/gw-mtls-proto

#ADD ./go.mod  ${location}/gw-mtls-service
#ADD ./go.sum  ${location}/gw-mtls-service
RUN go mod init github.com/blablatov/mtls-grpc-gateway/gw-mtls-service

# Download all the dependencies. Загрузка зависимостей
RUN go get -d ./...
# Install the package. Установка пакетов
RUN go install ./...

# Build server app. Сборка серверного приложения.
RUN CGO_ENABLED=0 go build -o /bin/gw-mtls-service 

# Go binaries are self-contained executables. Используя директиву FROM scratch - 
# Go образы  не должны содержать ничего, кроме одного двоичного исполняемого файла.
FROM scratch
COPY --from=build /bin/gw-mtls-service /bin/gw-mtls-service

ENTRYPOINT ["/bin/gw-mtls-service"]
EXPOSE 50051