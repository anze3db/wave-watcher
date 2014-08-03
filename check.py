import requests
from src.parser import Parser

SITE = 'http://www.yr.no/place/Croatia/Istria/Medulin/forecast.xml'

page = requests.get(SITE)
result = Parser.parse(page.text)
print result
