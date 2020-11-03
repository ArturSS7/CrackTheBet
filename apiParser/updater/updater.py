import requests
import psycopg2
from models.models import Match, League
from bs4 import BeautifulSoup


def get_status(ID):
	r = requests.get("https://www.flashscore.com/match/{}".format(ID))
	soup = BeautifulSoup(r.text, features="lxml")
	raw_status = soup.find("div", {"class": "info-status mstat"}).text
	if "Finished" in raw_status:
		t1 = int(soup.find("div", {"id": "event_detail_current_result"}).findAll("span", {"class":"scoreboard"})[0].text)
		t2 = int(soup.find("div", {"id": "event_detail_current_result"}).findAll("span", {"class":"scoreboard"})[1].text)
		if t1 > t2:
			return "finished", 1
		elif t1 < t2:
			return "finished", 2
		else:
			return "finished", 0
	elif '0a09090909090909c2a00a090909090909' == raw_status.encode().hex():
		return "hasn't started", -1
	else:
		return "active", -1

def get_odds(ID):
	headers={'X-Fsign': 'SW9D1eZo'}
	r = requests.get("https://d.flashscore.com/x/feed/d_od_{}_en_1_eu".format(ID), headers=headers)
	soup = BeautifulSoup(r.text, features="lxml")
	raw_odd = soup.find("tr", {"class": "odd"})
	odds1 = raw_odd.find("td", {"onclick": "e_t.track_click('bookmaker-button-click', 'block-1x2_ft_1');"}).find(\
		"span", {"class": "odds-wrap"}).text
	odds2 = raw_odd.find("td", {"onclick": "e_t.track_click('bookmaker-button-click', 'block-1x2_ft_2');"}).find(\
		"span", {"class": "odds-wrap"}).text
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
						(%s, %s, %s, %s, %s, %s, %s, %s, %s)", ('soccer', league.name, match.t1, match.t2, 1.5, 1.5, 'finished', match.time, match.ID))
			conn.commit()
	print("done")
	return matches