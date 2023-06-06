FROM golang:alpine


WORKDIR /go/src/app
COPY . .

RUN go build -o server main.go


# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
#ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
#RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh
#ENTRYPOINT [ "sh", "/entrypoint.sh" ]

CMD ["./server"]