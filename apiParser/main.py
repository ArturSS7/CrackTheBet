import requests
import psycopg2
from bs4 import BeautifulSoup
from models.models import Match, League

headers={'X-Fsign': 'SW9D1eZo'}
leagues = []
r = requests.get("https://d.flashscore.com/x/feed/f_1_0_3_en_1", headers=headers)
lis = r.text.split("ZA÷")[1:]
for item in lis:
	leagues.append(League(item))
# for league in leagues:
# 	print(league.name)
# 	for match in league.matches:
# 		print(match)
