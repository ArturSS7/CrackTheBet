import requests
import psycopg2
from datetime import datetime
from updater.updater import get_status, get_odds, update_db

print("ready")
conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
cur = conn.cursor()

print("connected")

matches = update_db(cur, conn)

print(matches)

while(True):
	for ID in matches:
		status = get_status(ID)
		try:
			odds1, odds2 = get_odds(ID)
		except AttributeError:
			cur.execute("delete from events where flashscore_id = '{}'".format(ID))
			conn.commit()
		print(status, odds1, odds2)

cur.close()
conn.close()