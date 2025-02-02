FROM python:3-alpine

RUN apk add --no-cache --update \
    curl \
    bash

WORKDIR /app

COPY main.py .

RUN pip install requests

CMD ["python", "main.py"]
