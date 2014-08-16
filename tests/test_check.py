from check import check
import requests
import os
import smtplib
from mock import Mock

BASE_DIR = os.path.abspath(os.path.dirname(__file__))
FIXTURES_DIR = os.path.join(BASE_DIR, 'fixtures')


class RequestsGetMock():
  def __init__(self, file):
    self.text = open(os.path.join(FIXTURES_DIR, file), 'r').read()


def requests_get_false(path):
  return RequestsGetMock('forecast.xml')


def requests_get_true(path):
  return RequestsGetMock('forecast-positive.xml')


def mock_SMTP():
  return Mock()


def test_check_false(monkeypatch):
  monkeypatch.setattr(requests, 'get', requests_get_false)
  m = Mock()
  monkeypatch.setattr(smtplib, 'SMTP', lambda *_: m)
  assert '> No surf :/' == check()[1]
  assert m.sendmail.call_count == 0


def test_check_true(monkeypatch):
  monkeypatch.setattr(requests, 'get', requests_get_true)
  m = Mock()
  monkeypatch.setattr(smtplib, 'SMTP', lambda *_: m)
  assert '> Surf! Strong JUGO:' == check()[1]
  assert m.sendmail.call_count == 1
