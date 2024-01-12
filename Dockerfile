#FROM ubuntu:latest
#LABEL authors="Анна"

#ENTRYPOINT ["top", "-b"]

#RUN go version
#ENV GOPATH=/

#COPY ../House ./

#RUN go mod download
#RUN go build -0 Structs ./config/main.go

#CMD ["./Structs"]

# текстовый файл, содержащий набор операций, которые могут быть использованы для создания докер образа
    # имя используемого образа
FROM golang:1.21-alpine
# инструкция фром определяет базовый образ операционной системы
RUN go version # иструкция ран определяет команды, выполняемые в командной оболочке внутри образа
ENV GOPATH=/

#COPY docker ./
COPY ./ ./

RUN go mod download
#RUN go build -o awesome_project ./main.go
RUN go build -o Structs ./main.go
#CMD ["./awesome_project"]
CMD ["./Structs"]