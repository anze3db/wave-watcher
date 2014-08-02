from lxml import html
import requests

SITE = 'http://prognoza.hr/tri_karta.php?id=tri&param=Istarska&code=Medulin'

page = requests.get(SITE)
tree = html.fromstring(page.text)

row = tree.xpath('//div[@class="sadrzajContents"]/table/tr')[0]
all_rows = row.xpath('//tr[@align="center"]/td')

weather_wind = ["weather", "wind"]
j = 0

for i, row in enumerate(all_rows):
  imgs = [c for c in row.iterchildren()]
  if len(imgs) == 0:
    continue

  src = imgs[0].get('src')
  if not src:
    continue

  j += 1
  if j % 2 == 1:
    continue
  src = src.replace('alasimboli/', '')
  src = src.replace('.gif', '')
  if "SE" in src and "1" not in src:
    print j, "STRONG JUGO, MAYBE GO SURFING"
