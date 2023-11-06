FROM ubuntu:latest
LABEL authors="feddos"

ENTRYPOINT ["top", "-b"]