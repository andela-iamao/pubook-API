FROM google/golang

WORKDIR /gopath/src/github.com/topcoderinc/cribs
ADD . /gopath/src/github.com/topcoderinc/cribs/

# go get all of the dependencies
RUN go get github.com/emicklei/forest
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin
RUN go get github.com/astaxie/beego/orm

RUN go get github.com/andela-iamao/pubook-api

# set env variables to mongo
ENV MONGO_DB YOUR-MONGO-DB
ENV MONGO_URL YOUR-MONGO-URL

EXPOSE 8080
CMD []
ENTRYPOINT ["/gopath/bin/cribs"]