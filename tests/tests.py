import os
import unittest
from parser import Parser

BASE_DIR = os.path.abspath(os.path.dirname(__file__))


class TestBase(unittest.TestCase):
  def setUp(self):
    self.fixtures_dir = os.path.join(BASE_DIR, 'fixtures')


class TestFetcher(TestBase):

  def test_parse(self):
    f = open(os.path.join(self.fixtures_dir, 'forecast.xml'), 'r')
    parsed = Parser.parse(f.read())
    self.assertTrue(len(parsed) == 38)

if __name__ == '__main__':
  unittest.main()
