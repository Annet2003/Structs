FROM ubuntu:latest
LABEL authors="Анна"

ENTRYPOINT ["top", "-b"]

RUN go version
ENV GOPATH=/

COPY ../House ./

RUN go mod download
RUN go build -0 Structs ./config/main.go

CMD ["./Structs"]