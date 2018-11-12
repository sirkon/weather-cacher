# weather-cacher
gRPC сервис для получения и кеширования погодных данных от разных источников

для запуска необходимы следующие пункты:

1. go 1.11+
2. Установленный компилятор protobuf-а (protoc)
3. Установленный плагин go для protoc (`go get github.com/golang/protobuf/protoc-gen-go`)
4. Установленный плагин с поддержкой gRPC для go (`go get -u google.golang.org/grpc`)
5. Сгенерировать код из прилагаемой схемы: `go generate github.com/sirkon/weather-cacher`
6. Иметь доступ к установленным postgres с postgis и redis

