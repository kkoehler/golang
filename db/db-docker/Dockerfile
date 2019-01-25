FROM ibmcom/db2express-c:latest

RUN su - db2inst1 -c "db2start \
    && db2 'create db GOTEST' \
    && db2 'connect to GOTEST' \
    && db2 +c 'create table customer (id integer not null, firstname varchar(256), lastname varchar(256), birthday date, data blob)' \
    && db2 LIST TABLES"
