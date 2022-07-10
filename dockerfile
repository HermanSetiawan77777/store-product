FROM golang:1.16 as base

FROM base as dev

RUN mkdir /app/src

WORKDIR /app/src
CMD ["air"]