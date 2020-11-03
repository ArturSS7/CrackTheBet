import threading
import requests
import psycopg2
import time
import multiprocessing
from datetime import datetime
from updater.updater import get_status, get_odds, update_db
import logging
import socket


logging.basicConfig(filename='loggg.log', level=logging.INFO)
logging.info("connected")

conn = None
cur = None
conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
cur = conn.cursor()


def status_answer():
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
		conn = psycopg2.connect("user='keker' host='db' dbname='betdb' password='kek'")
		cur = conn.cursor()
		while True:
			ID = self._job_queue.get()
			if ID is None:
				break
			status, winner = get_status(ID)
			try:
				odds1, odds2 = get_odds(ID)
				cur.execute("update events set odds1 = (%s), odds2 = (%s), status = (%s), winner = (%s) where flashscore_id = (%s);", (odds1, odds2, status, winner, ID))
				#conn.commit()
			except AttributeError:
				cur.execute("delete from events where flashscore_id = '{}';".format(ID))
				#conn.commit()

			conn.commit()

matches = update_db(cur, conn)

jobs = []
job_queue = multiprocessing.Queue()

for i in range(5):
	p = Worker(job_queue)
	jobs.append(p)
	p.start()

#Start socket
sock = threading.Thread(target=status_answer, args=())
sock.start()
###

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
		#print("not empty")
		pass
	conn.commit()
	logging.info("Update took : {}".format(time.time() - start_time))
	time.sleep(30)

cur.close()
conn.close()