import psycopg2
conn = psycopg2.connect(dbname='betdb', user='keker', password='everybodykissmybody', host='localhost')
cursor = conn.cursor()

cursor.execute("insert into bets (user_id, match_id, bet, odds, amount, prize) values (5,1,'qwe', 2.2, 123,123)")
cursor.execute('select * from bets;')
print(cursor.fetchall())
cursor.close()
conn.close()
