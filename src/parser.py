from lxml import etree
from time import strptime


class Parser(object):

  def __init__(self, **kw):
    self.__dict__.update(kw)

  @staticmethod
  def parse(xml):
    tree = etree.fromstring(str(xml))
    time = tree.xpath('//time')
    wind_direction = tree.xpath('//windDirection')
    wind_speed = tree.xpath('//windSpeed')

    def parse_values(args):
      (t, wd, ws) = args
      speed = float(ws.get('mps'))
      direction = float(wd.get('deg'))
      time = strptime(t.get('from'), '%Y-%m-%dT%H:%M:%S')
      from weather import Weather
      return Weather(
          wind_speed=speed,
          wind_direction=direction,
          time=time,
      )

    return map(parse_values, zip(time, wind_direction, wind_speed))
