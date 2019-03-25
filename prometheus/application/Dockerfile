from golang:1.12.0

RUN mkdir /application
ADD . /application
WORKDIR /application
RUN go build -o application .

CMD ["/application/application"]
