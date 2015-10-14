FROM orchardup/python:2.7
RUN pip install Flask uwsgi requests
ADD . /code
WORKDIR /code
CMD uwsgi --http 0.0.0.0:8080 --wsgi-file whereisbot/whereisbot.py --callable app --master --no-default-app