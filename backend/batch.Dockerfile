FROM python:3

WORKDIR /app

COPY batch/requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY batch .