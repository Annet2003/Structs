#FROM ubuntu:latest
#LABEL authors="Анна"

#ENTRYPOINT ["top", "-b"]

#RUN go version
#ENV GOPATH=/

#COPY ../House ./

#RUN go mod download
#RUN go build -0 Structs ./config/main.go

#CMD ["./Structs"]


FROM golang:1.21-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o awesome_project ./main.go

CMD ["./awesome_project"]