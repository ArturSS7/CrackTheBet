import psycopg2
conn = psycopg2.connect(dbname='betdb', user='keker', password='everybodykissmybody', host='localhost')
cursor = conn.cursor()


cursor.execute('select * from bets;')
print(cursor.fetchall())
cursor.close()
conn.close()
