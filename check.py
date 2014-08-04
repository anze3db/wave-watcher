import requests
from src.parser import Parser
from datetime import datetime

SITE = 'http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml'

page = requests.get(SITE)
parser = Parser.get_yr_parser()
parsed = parser.parse(page.text)
output_str = "Strong JUGO {time} (ws: {windSpeed}, wd: {windDirection})"
result = [output_str.format(**p.__dict__) for p in parsed
          if p.is_strong_jugo()]
if len(result) > 0:
  print datetime.now(), "> Surf! ", ", ".join(result)
else:
  print datetime.now(), "> No surf :/"
