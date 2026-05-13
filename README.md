# Guess the Number - Technical Test

## Service Overview

A simple service for a **number guessing** game. The user sends 2 numbers, then the service adds both numbers together and compares the result with the answer stored in the database.

### Flow

1. User sends `bilangan1` and `bilangan2` via POST request
2. Service adds both numbers: `total = bilangan1 + bilangan2`
3. Service retrieves the correct answer from the database
4. Service compares the total with the answer and returns the result:
   - `"lebih besar"` — if the total is greater than the answer
   - `"lebih kecil"` — if the total is less than the answer
   - `"tepat sekali"` — if the total equals the answer
5. Each attempt is logged to the `history` table in the database

### Tech Stack

| Version | File | Framework | Database |
|---------|------|-----------|----------|
| Python | `main.py` | Flask | SQLite |
| Go | `main.go` | net/http | SQLite |

### Database Schema

**Table `answer`** — stores the correct answer

| Column | Type | Description |
|--------|------|-------------|
| id | INTEGER (PK) | ID |
| number | INTEGER | Answer number |

**Table `history`** — stores attempt history

| Column | Type | Description |
|--------|------|-------------|
| id | INTEGER (PK, AUTO) | ID |
| bil1 | TEXT | First number |
| bil2 | TEXT | Second number |
| total | INTEGER | Sum result |
| result | TEXT | Comparison result |
| ts | TEXT | Timestamp |

---

## How to Run

### Python

```bash
pip install flask
python main.py
```

Server runs at `http://localhost:8081`

### Go

```bash
go mod init tebak-angka
go mod tidy
go build -o tebak-bin main.go
./tebak-bin
```

Server runs at `http://localhost:5000`

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

### 1. Guess is less than the answer

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

### 2. Guess is greater than the answer

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

### 3. Guess is exact

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

### 4. Validation error — missing field

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

### 5. Validation error — empty body

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
