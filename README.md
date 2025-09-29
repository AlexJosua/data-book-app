📚 Task Quiz 3 Sanbercode - API Documentation

🚀 Cara Penggunaan
Sebelum melakukan pengetesan API di Postman, harus login terlebih dahulu untuk mendapatkan token.

🔑 Login
POST https://data-book-app-production.up.railway.app/api/users/login
Body (JSON):
{
  "username": "admin",
  "password": "secret"
}

Jika berhasil login, akan menerima JWT Token.
Salin token tersebut, lalu di Postman:
- Buka tab Authorization
- Pilih Bearer Token
- Masukkan token ke dalam kolom


========================
📂 Categories API
========================

GET ALL
GET https://data-book-app-production.up.railway.app/api/categories

GET BY ID
GET https://data-book-app-production.up.railway.app/api/categories/{id}

CREATE
POST https://data-book-app-production.up.railway.app/api/categories
Body (JSON):
{
  "name": "Komik"
}

UPDATE
PUT https://data-book-app-production.up.railway.app/api/categories/{id}
Body (JSON):
{
  "name": "Novel"
}

DELETE
DELETE https://data-book-app-production.up.railway.app/api/categories/{id}


========================
📖 Books API
========================

GET ALL
GET https://data-book-app-production.up.railway.app/api/books

GET BY ID
GET https://data-book-app-production.up.railway.app/api/books/{id}

CREATE
POST https://data-book-app-production.up.railway.app/api/books
Body (JSON):
{
  "title": "My First Book",
  "description": "Belajar Golang dengan mudah",
  "image_url": "https://example.com/book.png",
  "release_year": 2024,
  "price": 50000,
  "total_page": 120,
  "category_id": 1
}

UPDATE
PUT https://data-book-app-production.up.railway.app/api/books/{id}
Body (JSON):
{
  "title": "My Updated Book",
  "description": "Update deskripsi buku",
  "image_url": "https://example.com/newbook.png",
  "release_year": 2023,
  "price": 75000,
  "total_page": 150,
  "category_id": 1
}

DELETE
DELETE https://data-book-app-production.up.railway.app/api/books/{id}
