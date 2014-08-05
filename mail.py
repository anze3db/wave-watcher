import smtplib
import os
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

me = "wave-watcher@psywerx.net"
you = "ensmotko@gmail.com"


def _get_msg_no_alert():
  msg = MIMEMultipart('alternative')
  msg['Subject'] = "No Jugo :("
  msg['From'] = me
  msg['To'] = you

  text = "No strong Jugo detected"
  html = """\
  <html>
    <head></head>
    <body>
      <h1>Sorry!</h1>
      <p>
        No strong Jugo detected :(
      </p>
    </body>
  </html>
  """

  part1 = MIMEText(text, 'plain')
  part2 = MIMEText(html, 'html')

  msg.attach(part1)
  msg.attach(part2)

  return msg


def _get_msg_alert(txt):
  msg = MIMEMultipart('alternative')
  msg['Subject'] = "Jugo Alert! "
  msg['From'] = me
  msg['To'] = you

  text = """\
    Jugo Alert!

    We have detected strong a jugo at the following times:
    {times}
    Please visit the following links to confirm the forecast:
    http://prognoza.hr/karte.php?id=prizemne&param=vjtl&it=anim
    http://prognoza.hr/karte.php?id=ecmwf&param=valovi&it=anim
  """.format(times=txt)
  html = """\
  <html>
    <head></head>
    <body>
      <h1>Jugo Alert!</h1>
      <p>
        We have detected strong a jugo at the following times:
      </p>
      <p>
         {times}
      </p>
      <p>
        Please visit the following links to confirm the forecast:<br>
        <a href="http://prognoza.hr/karte.php?id=prizemne&param=vjtl&it=anim">
        Prognoza.hr Wind
        </a><br>
        <a href="http://prognoza.hr/karte.php?id=ecmwf&param=valovi&it=anim">
        Prognoza.hr Waves
        </a>
      </p>
    </body>
  </html>
  """.format(times=txt.replace("\n", '<br>'))

  part1 = MIMEText(text, 'plain')
  part2 = MIMEText(html, 'html')

  msg.attach(part1)
  msg.attach(part2)

  return msg


def send_email_alert(msg):
  s = smtplib.SMTP('smtp.mandrillapp.com', 587)
  s.login(os.environ.get('MANDRILL_USERNAME'),
          os.environ.get('MANDRILL_APIKEY'))

  msg = _get_msg_alert(msg) if len(msg) > 0 else _get_msg_no_alert()
  s.sendmail(me, you, msg.as_string())
  s.quit()
