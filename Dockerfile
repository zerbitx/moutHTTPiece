FROM golang:1.12-alpine

RUN mkdir -p /usr/src/app
COPY bin/mouthttpiece /usr/src/app
WORKDIR /usr/src/app

ENV PORT=8000
ENV ECHO=false
ENV STATUS_CODE_DEFAULT=200
EXPOSE ${PORT}

CMD ["./mouthttpiece"]
