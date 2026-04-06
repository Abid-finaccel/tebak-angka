# Tebak Angka - Technical Test

## Garis Besar Service

Service sederhana untuk permainan **tebak angka**. User mengirimkan 2 bilangan, lalu service menjumlahkan kedua bilangan tersebut dan membandingkan hasilnya dengan jawaban yang tersimpan di database.

### Flow

1. User mengirim `bilangan1` dan `bilangan2` via POST request
2. Service menjumlahkan kedua bilangan: `total = bilangan1 + bilangan2`
3. Service mengambil jawaban yang benar dari database
4. Service membandingkan total dengan jawaban dan mengembalikan hasil:
   - `"lebih besar"` — jika total lebih besar dari jawaban
   - `"lebih kecil"` — jika total lebih kecil dari jawaban
   - `"tepat sekali"` — jika total sama dengan jawaban
5. Setiap percobaan dicatat ke tabel `history` di database

### Tech Stack

| Versi | File | Framework | Database |
|-------|------|-----------|----------|
| Python | `main.py` | Flask | SQLite |
| Go | `main.go` | net/http | SQLite |

### Database Schema

**Tabel `answer`** — menyimpan jawaban yang benar

| Column | Type | Description |
|--------|------|-------------|
| id | INTEGER (PK) | ID |
| number | INTEGER | Angka jawaban |

**Tabel `history`** — menyimpan riwayat percobaan

| Column | Type | Description |
|--------|------|-------------|
| id | INTEGER (PK, AUTO) | ID |
| bil1 | TEXT | Bilangan pertama |
| bil2 | TEXT | Bilangan kedua |
| total | INTEGER | Hasil penjumlahan |
| result | TEXT | Hasil perbandingan |
| ts | TEXT | Timestamp |

---

## Cara Menjalankan

### Python

```bash
pip install flask
python main.py
```

Server berjalan di `http://localhost:5000`

### Go

```bash
go run main.go
```

Server berjalan di `http://localhost:5000`

---

## API

### POST `/tebak`

#### Request

```
Content-Type: application/json
```

```json
{
  "bilangan1": <number>,
  "bilangan2": <number>
}
```

---

## Example Curl & Output

### 1. Tebakan lebih kecil dari jawaban

```bash
curl -X POST http://localhost:5000/tebak \
  -H "Content-Type: application/json" \
  -d '{"bilangan1": 10, "bilangan2": 20}'
```

**Output:**

```json
{
  "status": "ok",
  "result": "lebih kecil"
}
```

### 2. Tebakan lebih besar dari jawaban

```bash
curl -X POST http://localhost:5000/tebak \
  -H "Content-Type: application/json" \
  -d '{"bilangan1": 50, "bilangan2": 40}'
```

**Output:**

```json
{
  "status": "ok",
  "result": "lebih besar"
}
```

### 3. Tebakan tepat

```bash
curl -X POST http://localhost:5000/tebak \
  -H "Content-Type: application/json" \
  -d '{"bilangan1": 25, "bilangan2": 50}'
```

**Output:**

```json
{
  "status": "ok",
  "result": "tepat sekali"
}
```

### 4. Validasi error — field kosong

```bash
curl -X POST http://localhost:5000/tebak \
  -H "Content-Type: application/json" \
  -d '{"bilangan1": 10}'
```

**Output:**

```json
{
  "status": "error",
  "message": "bilangan1 dan bilangan2 harus diisi dan tidak boleh kosong"
}
```

### 5. Validasi error — body kosong

```bash
curl -X POST http://localhost:5000/tebak \
  -H "Content-Type: application/json" \
  -d '{}'
```

**Output:**

```json
{
  "status": "error",
  "message": "bilangan1 dan bilangan2 harus diisi dan tidak boleh kosong"
}
```
