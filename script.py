import requests
import sqlite3


data = {
    'Username':"u2",
    'Password':'p2',
    'DBPassword': 'dbp2',
    'DBTag': 'dbt2'
}

response = requests.post('http://localhost:8080/register', json=data)

print(response.json())

conn = sqlite3.connect("models/UserDB.db")
cur = conn.cursor()
cur.execute("select * from Users")
conn.commit()
row = cur.fetchall()
for r in row:
    print(r)