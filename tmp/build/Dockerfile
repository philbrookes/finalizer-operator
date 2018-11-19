FROM alpine:3.6

RUN adduser -D finalizer-operator
USER finalizer-operator

ADD tmp/_output/bin/finalizer-operator /usr/local/bin/finalizer-operator
