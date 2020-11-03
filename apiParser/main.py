import requests
import psycopg2
import time
import multiprocessing
from datetime import datetime
from updater.updater import get_status, get_odds, update_db
import logging
logging.basicConfig(filename='loggg.log', level=logging.INFO)
logging.info("connected")

conn = None
cur = None
conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
cur = conn.cursor()

class Worker(multiprocessing.Process):

	def __init__(self, job_queue):
		super().__init__()
		self._job_queue = job_queue

	def run(self):
		while True:
			ID = self._job_queue.get()
			if ID is None:
				break
			status, winner = get_status(ID)
			try:
				odds1, odds2 = get_odds(ID)
				cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s)", (odds1, odds2, status, winner, ID))
				#conn.commit()
			except AttributeError:
				cur.execute("delete from events where flashscore_id = '{}'".format(ID))
				#conn.commit()
			conn.commit()
			
matches = update_db(cur, conn)

jobs = []
job_queue = multiprocessing.Queue()

for i in range(5):
	p = Worker(job_queue)
	jobs.append(p)
	p.start()

while(True):
	start_time = time.time()

	for ID in matches:
		# status, winner = get_status(ID)
		# try:
		# 	odds1, odds2 = get_odds(ID)
		# except AttributeError:
		# 	cur.execute("delete from events where flashscore_id = '{}'".format(ID))
		# 	print("no bets for ", ID)
		# 	conn.commit()
		# cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s)", (odds1, odds2, status, winner, ID))
		# conn.commit()
		job_queue.put(ID)
		print()
	while(not job_queue.empty()):
		print("not empty")
		pass
	conn.commit()
	logging.info("Update took : {}".format(time.time() - start_time))
	time.sleep(30)

cur.close()
conn.close()