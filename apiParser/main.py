import requests
from bs4 import BeautifulSoup
from datetime import datetime

headers={'X-Fsign': 'SW9D1eZo'}
leagues = {}
r = requests.get("https://d.flashscore.com/x/feed/f_1_0_3_en_1", headers=headers)
lis = r.text.split("ZA÷")[1:]
for item in lis:
	l = item.split("AA÷")
	print(l[0].split("¬")[0])
	for m in l[1:]:
		ID = m.split('¬')[0]
		time = datetime.fromtimestamp(int(m.split('¬')[1][3:-1]))
		t1 = m.split('CX÷')[1].split("¬")[0]
		t2 = m.split('AF÷')[1].split("¬")[0]
		print('\t', ID, time, " : ", t1, " vs ", t2)