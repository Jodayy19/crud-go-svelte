# Simple CRUD App (Go Fiber + SvelteKit)

## Deskripsi
Project ini merupakan implementasi CRUD sederhana dengan **backend Go Fiber** dan **Frontend SvelteKit**. Fungsinya untuk manajemen data user (create, read, update, delete). 

## Tech Stack
- Go Fiber (Backend)
- SvelteKit (Frontend)

## Cara Menjalankan

### Backend
cd backend  
go run main.go  

### Frontend
cd frontend  
npm install  
npm run dev  

## API
- GET /users - menampilkan semua users
- POST /users - menambahkan user baru
- PUT /users/:id - update user
- DELETE /users/:id - hapus user

## Catatan
Pastikan backend jalan sebelum menjalankan frontend
Semua data user tersimpan di memory sementara, belum menggunakan database permanen.

