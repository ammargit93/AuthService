import requests
import sqlite3


data = {
    'Username': 'yourUsername',
    'Password': 'yourPassword',
    'DBTag': 'firstdatabase'
}

response = requests.post('http://localhost:8080/register', json=data)

print(response.json())

conn = sqlite3.connect("models/UserDB.db")
cur = conn.cursor()
cur.execute("select * from Users")
row = cur.fetchall()
for r in row:
    print(r)