import requests
import psycopg2
from bs4 import BeautifulSoup
from models.models import Match, League

#connect to db
conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
cur = conn.cursor()

#erase events db and restart ID
cur.execute("TRUNCATE events RESTART IDENTITY")
conn.commit()

#gets events from FS and insert into db
headers={'X-Fsign': 'SW9D1eZo'}
r = requests.get("https://d.flashscore.com/x/feed/f_1_0_3_en_1", headers=headers)
lis = r.text.split("ZA÷")[1:]
for item in lis:
	league = League(item)
	for match in league.matches:
		cur.execute("insert into events (event_type, league, player1, player2, odds1, odds2, status, time, flashscore_id) values \
					(%s, %s, %s, %s, %s, %s, %s, %s, %s)", ('soccer', league.name, match.t1, match.t2, 1.5, 1.5, 'finished', match.time, match.ID))
		conn.commit()

cur.close()
conn.close()