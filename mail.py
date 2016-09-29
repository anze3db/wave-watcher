import smtplib
import os
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

recipients = os.environ.get('WW_EMAIL_LIST', "").split(',')
me = "wave-watcher@psywerx.net"


def _get_msg_alert(txt):
  msg = MIMEMultipart('alternative')
  first_date = txt.split("\n<br>")[0]
  msg['Subject'] = "Strong wind on {}!".format(first_date)
  msg['From'] = me
  msg['To'] = ", ".join(recipients)

  with open('template.html', 'r') as template:
    html = template.read().replace("{times}", txt)
  part2 = MIMEText(html, 'html')

  msg.attach(part2)

  return msg


def send_email_alert(msg):
  s = smtplib.SMTP('smtp.postmarkapp.com', 587)
  s.login(os.environ.get('POSTMARK_API_KEY'),
          os.environ.get('POSTMARK_API_KEY'))

  msg = _get_msg_alert(msg)
  s.sendmail(me, recipients, msg.as_string())
  s.quit()
