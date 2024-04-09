## Cara menjalankan proyek ini

### Instalasi
- Clone repository ini
- Buka terminal dan arahkan ke direktori proyek
- Jalankan perintah berikut untuk menginstal dependensi
```bash
go mod tidy
```
- Masukkan environment variable yang diperlukan pada file `.env`
- Salin .env.example menjadi .env
```bash
cp .env.example .env
```
- Jalankan perintah berikut untuk menjalankan aplikasi
```bash
go run main.go
```
