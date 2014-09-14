from flask import Flask
from flask import render_template
import os
import requests
from src.parser import Parser

app = Flask(__name__)


@app.route("/")
def index():

  SITE = 'http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml'

  page = requests.get(SITE)
  parser = Parser.get_yr_parser()
  parsed = parser.parse(page.text)
  has_waves = len([p for p in parsed if p.is_strong_jugo()]) > 0
  return render_template("index.html", results=parsed, has_waves=has_waves)

if __name__ == "__main__":
  app.debug = os.environ.get('WV_DEBUG', False)
  app.run()
