from check import check


def test_check():
  assert '> No surf :/' == check()[1]
