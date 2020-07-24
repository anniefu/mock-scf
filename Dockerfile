FROM golang:1.14 as proxy

WORKDIR /app/proxy
COPY ./proxy ./

RUN go build -o server

FROM python:3.7
COPY --from=proxy /app/proxy/server /app/proxy/server

RUN pip install functions-framework gunicorn

COPY start_both.sh /app/
RUN chmod +x /app/start_both.sh
COPY ./python /app/python

WORKDIR /app
CMD ["/app/start_both.sh"]
