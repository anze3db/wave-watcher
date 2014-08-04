from lxml import etree
from time import strptime, mktime
from datetime import datetime


class Parser(object):

  def __init__(self, **kw):
    self.values = kw

  def _xml_to_values(self, tree, valfun):
    val, fun = valfun
    return [fun(node) for node in tree.xpath('//' + val)]

  def _values_to_cls(self, keys, values, cls):
    return map(lambda vals: cls(**dict(zip(keys, vals))), values)

  def parse(self, xml):
    tree = etree.fromstring(str(xml))
    all_values = zip(*map(
        lambda x: self._xml_to_values(tree, x),
        self.values.iteritems()
    ))

    from weather import Weather
    return self._values_to_cls(self.values.keys(), all_values, Weather)

  @staticmethod
  def get_yr_parser():
    return Parser(
        time=lambda x: datetime.fromtimestamp(
            mktime(strptime(x.get('from'), '%Y-%m-%dT%H:%M:%S'))),
        windDirection=lambda x: float(x.get('deg')),
        windSpeed=lambda x: float(x.get('mps')),
    )
