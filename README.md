# Kasir API

API sederhana untuk sistem kasir dengan fitur manajemen produk dan kategori.

## Tech Stack
- Go
- PostgreSQL
- Viper (config management)

## Cara Menjalankan

1. Setup file `.env`:
```
PORT=8080
DB_CONN=postgresql://user:password@host:port/database?sslmode=require
```

2. Jalankan aplikasi:
```bash
go run main.go
```

### Development dengan Live-Reload

Install air untuk live-reload:
```bash
go install github.com/air-verse/air@latest
```

Jalankan dengan air:
```bash
air -c .air.toml
```

Air akan otomatis restart aplikasi setiap kali ada perubahan file.

## Endpoints

### Products

**GET** `/api/produk`  
Ambil semua produk (dengan data kategori jika ada)

**POST** `/api/produk`  
Buat produk baru
```json
{
  "name": "Indomie",
  "price": 3500,
  "stock": 100,
  "category_name": "makanan"
}
```

**GET** `/api/produk/{id}`  
Ambil produk berdasarkan ID (dengan data kategori jika ada)

**PUT** `/api/produk/{id}`  
Update produk
```json
{
  "name": "Indomie Goreng",
  "price": 3500,
  "stock": 50,
  "category_name": "makanan"
}
```

**DELETE** `/api/produk/{id}`  
Hapus produk

### Categories

**GET** `/api/categories`  
Ambil semua kategori

**POST** `/api/categories`  
Buat kategori baru
```json
{
  "name": "makanan",
  "description": "Makanan dan minuman"
}
```

**GET** `/api/categories/{id}`  
Ambil kategori berdasarkan ID

**PUT** `/api/categories/{id}`  
Update kategori

**DELETE** `/api/categories/{id}`  
Hapus kategori

### Transactions

**POST** `/api/checkout`  
Checkout transaksi
```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 3,
      "quantity": 1
    }
  ]
}
```

Response:
```json
{
  "id": 1,
  "total_amount": 45000,
  "details": [
    {
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Indomie Goreng",
      "quantity": 2,
      "subtotal": 7000
    }
  ]
}
```

### Reports

**GET** `/api/report/hari-ini`  
Laporan transaksi hari ini

Response:
```json
{
  "total_revenue": 45000,
  "total_transaksi": 5,
  "produk_terlaris": {
    "nama": "Indomie Goreng",
    "qty_terjual": 12
  }
}
```

### Health Check

**GET** `/health`  
Cek status API

**GET** `/`  
Root endpoint
