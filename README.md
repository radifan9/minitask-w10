# Week 10 Beginner Backend

## Minitask Day 1
- Buat rute untuk **LOGIN** dan **REGISTER**
- Masing masing rute bisa menerima body
- Berikan validasi untuk body baik di login ataupun di register
- Berikan respon yang sesuai
- Respon sukses cukup berisikan data body dan status keberhasilan
- Lengkapi error handling

## Minitask Day 2
- Buatlah Update berbentuk **PATCH**, hanya berubah sebagian
- Selesaikan Code Splitting

# 🛍️ Simple Store Database (PostgreSQL)

---

## 📦 Tables

### Users
- `id` (uuid, PK)
- `email` (text, unique, not null)
- `password` (text, not null)
- `created_at` (timestamptz, default now)
- `updated_at` (timestamptz, default now)

### Products
- `id` (serial, PK)
- `name` (text)
- `price` (int)

### Transactions
- `id` (uuid, PK)
- `user_id` (uuid, FK → users.id)
- `product_id` (int, FK → products.id)
- `created_at` (timestamptz, default now)
- `updated_at` (timestamptz, default now)

---

## 🔗 Relationships
- One **user** → many **transactions**
- One **product** → many **transactions**

---
