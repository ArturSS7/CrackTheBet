from datetime import datetime

class Match:
	def __init__(self, ID, time, t1, t2):
		self.ID = ID
		self.time = time
		self.t1 = t1
		self.t2 = t2
	def __str__(self):
		return ("\t" + self.ID + str(self.time) + " : " + self.t1 + " vs " + self.t2)

class League:
	def __init__(self, data):
		#print(data)
		self.matches = []
		self.name = data.split("¬")[0]
		for m in data.split("AA÷")[1:]:
			#print(m)
			ID = m.split('¬')[0]
			#print(ID)
			time = datetime.fromtimestamp(int(m.split('¬')[1][3:-1]))
			t1 = m.split('CX÷')[1].split("¬")[0]
			t2 = m.split('AF÷')[1].split("¬")[0]
			self.matches.append(Match(ID,time,t1,t2))

