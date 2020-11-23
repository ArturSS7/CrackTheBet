import requests
import psycopg2
from models.models import Match, League
from bs4 import BeautifulSoup


def get_status(ID):
	r = requests.get("https://www.flashscore.com/match/{}".format(ID))
	#print(ID)
	soup = BeautifulSoup(r.text, features="lxml")
	raw_status = soup.find("div", {"class": "info-status mstat"}).text
	if ("Finished" in raw_status) or ("Awarded" in raw_status) or ("After Penalties" in raw_status):
		t1 = int(soup.find("div", {"id": "event_detail_current_result"}).findAll("span", {"class":"scoreboard"})[0].text)
		t2 = int(soup.find("div", {"id": "event_detail_current_result"}).findAll("span", {"class":"scoreboard"})[1].text)
		if t1 > t2:
			return "finished", 1
		elif t1 < t2:
			return "finished", 2
		else:
			return "finished", 0
	elif ("Cancelled" in raw_status) or ("Postponed" in raw_status):
		return "cancelled", -1
	elif '0a09090909090909c2a00a090909090909' == raw_status.encode().hex():
		return "hasn't started", -1
	else:
		return "active", -1

def get_odds(ID, status):
	headers={'X-Fsign': 'SW9D1eZo'}
	if (status == "hasn't started") or (status == "finished"):
		r = requests.get("https://d.flashscore.com/x/feed/df_dos_1_{}_".format(ID), headers=headers)
		data = r.text.split('[')[1].split(']')[0].split('"')
		odds1 = data[1]
		odds2 = data[5]
		return odds1, odds2
	else:
		try:
			r = requests.get("https://d.flashscore.com/x/feed/df_lod2_453_{}".format(ID), headers=headers)
			data = r.text.split('÷')
			odds1 = data[3].split('¬')[0][1:]
			odds2 = data[7].split('¬')[0][1:]
			return odds1, odds2
		except IndexError:
			r = requests.get("https://d.flashscore.com/x/feed/df_dos_1_{}_".format(ID), headers=headers)
			data = r.text.split('[')[1].split(']')[0].split('"')
			odds1 = data[1]
			odds2 = data[5]
			return odds1, odds2
		

def update_db(cur, conn):
	matches = []

	#erase events db and restart ID
	cur.execute("TRUNCATE events RESTART IDENTITY")
	conn.commit()
	print("db is cleared")
	#gets events from FS and insert into db
	headers={'X-Fsign': 'SW9D1eZo'}
	r = requests.get("https://d.flashscore.com/x/feed/f_1_0_3_en_1", headers=headers)
	print("got data from FS")
	lis = r.text.split("ZA÷")[1:]
	for item in lis:
		league = League(item)
		for match in league.matches:
			matches.append(match.ID)
			cur.execute("insert into events (event_type, league, player1, player2, odds1, odds2, status, time, flashscore_id) values \
						(%s, %s, %s, %s, %s, %s, %s, %s, %s)", ('soccer', league.name, match.t1, match.t2, 228, 228, 'finished', match.time, match.ID))
			conn.commit()
	print("done")
	return matches