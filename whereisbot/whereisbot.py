from flask import Flask, request
import datetime
import requests
import os

source_json = os.environ["SOURCE_JSON"]

def _get_by_date(name, date):
    locations_map = _load_locations_map(name)
    if date.date() in locations_map:
        return "{0} is {1}".format(name, locations_map[date.date()]
    else:
        return "undefined"

def _load_locations_map(name):
    r = requests.get(source_json)
    locations_file = r.json()
    print ("Fetched locations file")
    locations_map = {}
    for person in locations_file["people"]:
        if person["name"].lower() == name.lower():
            for location in person["locations"]:
                unparsed_date = location["date"]
                parsed_date = datetime.datetime.strptime(unparsed_date, "%Y-%m-%d")
                locations_map[parsed_date.date()] = location["location"]
            break
    return locations_map

app = Flask(__name__)

@app.route('/')
def index():
    return 'Hello! See my source at <a href="https://github.com/ssk2/whereisbot">https://github.com/ssk2/whereisbot</a>.'

@app.route('/whereis', methods=['GET','POST'])
def whereis():
    today = datetime.datetime.today()
    if request.method == 'POST':
        user = request.form.get('user_name')
        name = request.form.get('text').lower()
    else:
        user = request.args.get('user_name')
        name = request.args.get('text')
    print ("Fetching location for name {0} date {1} from user {2}".format(name, today, user))
    return _get_by_date(name, today)

if __name__ == '__main__':
    app.run()
