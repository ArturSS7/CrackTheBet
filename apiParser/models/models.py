from datetime import datetime

class Match:
	def __init__(self, ID, time=666, t1='papich', t2='aloha'):
		self.ID = ID
		self.time = time
		self.t1 = t1
		self.t2 = t2
	def __str__(self):
		return ("\t" + self.ID + str(self.time) + " : " + self.t1 + " vs " + self.t2)

class League:
	def __init__(self, data):
		self.matches = []
		self.name = data.split("¬")[0]
		for m in data.split("AA÷")[1:]:
			ID = m.split('¬')[0]
			#time = datetime.fromtimestamp(int(m.split('¬')[1][3:-1]))
			time = int(m.split('¬')[1][3:])
			t1 = m.split('CX÷')[1].split("¬")[0]
			t2 = m.split('AF÷')[1].split("¬")[0]
			self.matches.append(Match(ID,time,t1,t2))

