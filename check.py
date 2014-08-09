import requests
from src.parser import Parser
from datetime import datetime
from mail import send_email_alert

SITE = 'http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml'


def check():
  page = requests.get(SITE)
  parser = Parser.get_yr_parser()
  parsed = parser.parse(page.text)
  output_str = "On: {time} (ws: {windSpeed}, wd: {windDirection})"
  result = [output_str.format(**p.__dict__) for p in parsed
            if p.is_strong_jugo()]
  if len(result) > 0:
    return datetime.now(), "> Surf! Strong JUGO:", ", ".join(result)
    send_email_alert("\n".join(result))
  else:
    return datetime.now(), "> No surf :/"

if __name__ == '__main__':
  print check()
