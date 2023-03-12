from golang:1.20

RUN mkdir /application
ADD . /application
WORKDIR /application
RUN go build -o application .

CMD ["/application/application"]
