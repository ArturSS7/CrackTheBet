import threading
import requests
import psycopg2
import time
import multiprocessing
from datetime import datetime
from updater.updater import get_status, get_odds, update_db
import logging
import socket

manager = multiprocessing.Manager()
matches = manager.list()

conn_str = "user='keker' host='db' dbname='betdb' password='everybodykissmybody'"

conn = psycopg2.connect(conn_str)
cur = conn.cursor()

logging.basicConfig(filename='loggg.log', level=logging.INFO)
logging.info("connected")

for i in update_db(cur, conn):
	matches.append(i)
cur.close()
conn.close()

def status_resp():
	sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	server_address = ('0.0.0.0', 5555)
	sock.bind(server_address)
	sock.listen(1)
	while True:
		connection, client_addres = sock.accept()
		print("Accept")
		try:
			while True:
				flashscore_id = connection.recv(1024).decode('utf-8').strip('\n')
				print(flashscore_id)
				status, num = get_status(flashscore_id)
				print(status)
				connection.sendall(str.encode(status+'\n'))
				break
		finally:
			connection.close()


class Worker(multiprocessing.Process):

	def __init__(self, job_queue):
		super().__init__()
		self._job_queue = job_queue

	def run(self):
		conn = psycopg2.connect(conn_str)
		cur = conn.cursor()
		while True:
			ID = self._job_queue.get()
			if ID is None:
				break
			status, winner = get_status(ID)
			# try:
			# 	odds1, odds2 = get_odds(ID)
			# 	cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s)", (odds1, odds2, status, winner, ID))
			# 	#conn.commit()
			# except AttributeError:
			# 	cur.execute("delete from events where flashscore_id = '{}'".format(ID))
			# 	matches.remove(ID)
			# 	#conn.commit()
			odds1, odds2 = get_odds(ID, status)
			if ((odds1 == '-') and (odds2 == '-')) or (status == 'canceled'):
				cur.execute("delete from events where flashscore_id = '{}'".format(ID))
				matches.remove(ID)
			elif (status == 'finished'):
				cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s)", (odds1, odds2, status, winner, ID))
				matches.remove(ID)
			else:
				cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s)", (odds1, odds2, status, winner, ID))
			conn.commit()

jobs = []
job_queue = multiprocessing.Queue()

for i in range(10):
	p = Worker(job_queue)
	jobs.append(p)
	p.start()
#Start socket
sock = threading.Thread(target=status_resp, args=())
sock.start()
###

while(True):
	start_time = time.time()
	for ID in matches:
		job_queue.put(ID)
	while(not job_queue.empty()):
		pass
	logging.info("Update took : {}".format(time.time() - start_time))
	print("cycle done")
	time.sleep(10)