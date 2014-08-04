import os
import unittest
from parser import Parser
from weather import Weather
from datetime import datetime

BASE_DIR = os.path.abspath(os.path.dirname(__file__))


class TestBase(unittest.TestCase):
  def setUp(self):
    self.fixtures_dir = os.path.join(BASE_DIR, 'fixtures')

  def read_fixture(self, name):
    f = open(os.path.join(self.fixtures_dir, name), 'r')
    return f.read()


class TestParser(TestBase):

  def test_yr_parser(self):
    parser = Parser.get_yr_parser()
    parsed = parser.parse(self.read_fixture('forecast.xml'))
    self.assertTrue(len(parsed) == 38)
    first = parsed[0]
    self.assertIsInstance(first, Weather)
    self.assertIsInstance(first.windDirection, float)
    self.assertIsInstance(first.windSpeed, float)
    self.assertIsInstance(first.time, datetime)


class TestWeather(TestBase):

  def setUp(self):
    super(TestWeather, self).setUp()
    self.parser = Parser.get_yr_parser()

    self.negative = self.read_fixture('forecast.xml')
    self.positive = self.read_fixture('forecast-positive.xml')

  def test_strong_jugo_negative(self):
    parsed = self.parser.parse(self.negative)
    self.assertNotIn(True, map(lambda x: x.is_strong_jugo(), parsed))

  def test_strong_jugo_positive(self):
    parsed = self.parser.parse(self.positive)
    self.assertIn(True, map(lambda x: x.is_strong_jugo(), parsed))


if __name__ == '__main__':
  unittest.main()
