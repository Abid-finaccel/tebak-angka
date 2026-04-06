from flask import Flask, request, jsonify
import sqlite3
import datetime

app = Flask(__name__)

db_path = "game.db"

# setup table
conn = sqlite3.connect(db_path)
c = conn.cursor()
c.execute("CREATE TABLE IF NOT EXISTS answer (id INTEGER PRIMARY KEY, number INTEGER)")
c.execute("CREATE TABLE IF NOT EXISTS history (id INTEGER PRIMARY KEY AUTOINCREMENT, bil1 TEXT, bil2 TEXT, total INTEGER, result TEXT, ts TEXT)")
# seed answer if empty
c.execute("SELECT count(*) FROM answer")
if c.fetchone()[0] == 0:
    c.execute("INSERT INTO answer (id, number) VALUES (1, 75)")
conn.commit()
conn.close()

@app.route("/tebak", methods=["POST"])
def tebak():
    body = request.get_json()
    b1 = body.get("bilangan1")
    b2 = body.get("bilangan2")

    # validasi
    if b1 == None or b1 == "" or b2 == None or b2 == "":
        return jsonify({"status": "error", "message": "bilangan1 dan bilangan2 harus diisi dan tidak boleh kosong"})

    total = b1 + b2

    # get answer from db
    conn2 = sqlite3.connect(db_path)
    cur = conn2.cursor()
    cur.execute("SELECT number FROM answer WHERE id = 1")
    row = cur.fetchone()
    jawaban = row[0]
    conn2.close()

    if total > jawaban:
        r = "lebih besar"
    elif total < jawaban:
        r = "lebih kecil"
    else:
        r = "tepat sekali"

    # save history
    conn3 = sqlite3.connect(db_path)
    cur3 = conn3.cursor()
    now = str(datetime.datetime.now())
    cur3.execute("INSERT INTO history (bil1, bil2, total, result, ts) VALUES ('" + str(b1) + "', '" + str(b2) + "', " + str(total) + ", '" + r + "', '" + now + "')")
    conn3.commit()
    conn3.close()

    return jsonify({"status": "ok", "result": r})

if __name__ == "__main__":
    app.run(debug=True)
