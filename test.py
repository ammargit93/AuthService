import sqlite3

# Path to your database
db_path = "api/models/UserDB.db"

# Connect to the database
conn = sqlite3.connect(db_path)
cursor = conn.cursor()

# Get the list of tables
cursor.execute("SELECT * FROM Users")
row = cursor.fetchall()

for r in row:
    print(r)
