import requests
import psycopg2
import time
from datetime import datetime
from updater.updater import get_status, get_odds, update_db
import logging
logging.basicConfig(filename='loggg.log', encoding='utf-8', level=logging.INFO)


print("ready")
conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
cur = conn.cursor()

logging.info("connected")

matches = update_db(cur, conn)

while(True):
	start_time = time.time()
	for ID in matches:
		status = get_status(ID)
		try:
			odds1, odds2 = get_odds(ID)
		except AttributeError:
			cur.execute("delete from events where flashscore_id = '{}'".format(ID))
			print("no bets for ", ID)
			conn.commit()
		cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s) where flashscore_id = (%s)", (odds1, odds2, status, ID))
		print(status, odds1, odds2)
	logging.info("Update took : {}".format(time.time() - start_time))
	time.sleep(30)

cur.close()
conn.close()