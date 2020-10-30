FROM python:3

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

RUN python apiParser/init_db.py
EXPOSE 5555

CMD ["python", "./main.py"]