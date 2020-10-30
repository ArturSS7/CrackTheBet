import psycopg2
conn = psycopg2.connect(dbname='BetDB', user='keker', password='everybodykissmybody', host='localhost')
cursor = conn.cursor()

cursor.execute('''CREATE TABLE IF NOT EXISTS events (
	id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    event_type varchar(255),
    player1 varchar(255),
    player2 varchar(255),
    odds1 float,
    odds2 float,
    status varchar(255),
    time integer,
    flashscore_id varchar(255)
);''')

cursor.execute('''CREATE TABLE IF NOT EXISTS bets (
	id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	user_id integer,
	match_id integer,
	bet varchar(255),
    odds float,
    amount integer,
    prize integer
);''')

cursor.close()
conn.close()