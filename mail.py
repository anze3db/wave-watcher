import smtplib
import os
import re
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

recipients = ['ensmotko@gmail.com', 'zidarsk8@gmail.com']
me = "wave-watcher@psywerx.net"


def _get_msg_alert(txt):
  msg = MIMEMultipart('alternative')
  msg['Subject'] = "Wind Alert!"
  msg['From'] = me
  msg['To'] = ", ".join(recipients)

  html = """\
  <html>
    <head></head>
    <body>
      <h1>Wind Alert!</h1>
      <p>
        We have detected a strong wind at the following times:
      </p>
      <p>
         {times}
      </p>
      <p>
        Please visit the following links to confirm the forecast:<br>
        <a
          href="http://prognoza.hr/karte.php?id=prizemne&param=vjtl&it=anim">
        Prognoza.hr Wind
        </a><br>
        <a href="http://prognoza.hr/karte.php?id=ecmwf&param=valovi&it=anim">
        Prognoza.hr Waves
        </a>
      </p>
    </body>
  </html>
  """

  text = re.sub('<[^<]+?>', '', html).format(times=txt).strip()
  html = html.format(times=txt.replace("\n", '<br>'))
  part1 = MIMEText(text, 'plain')
  part2 = MIMEText(html, 'html')

  msg.attach(part1)
  msg.attach(part2)

  return msg


def send_email_alert(msg):
  s = smtplib.SMTP('smtp.mandrillapp.com', 587)
  s.login(os.environ.get('MANDRILL_USERNAME'),
          os.environ.get('MANDRILL_APIKEY'))

  msg = _get_msg_alert(msg)
  s.sendmail(me, recipients, msg.as_string())
  s.quit()
