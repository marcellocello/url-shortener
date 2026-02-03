# URL Shortener - Sistem URL Shortener

URL Shortener adalah backend layanan pemendek URL yang dirancang dengan fokus pada kecepatan dan skalabilitas. Proyek ini menggabungkan ekosistem **Go** yang efisien dengan strategi **Multi-Layer Storage** (PostgreSQL & Redis) untuk menangani jutaan pengalihan tanpa membebani database utama.

## 🏆 Fitur Unggulan
- **Custom Short Links**: Pengguna dapat menentukan alias sendiri untuk tautan mereka.
- **Microsecond Redirects**: Menggunakan **Redis Caching** untuk menyimpan tautan populer, mengurangi beban PostgreSQL hingga 90%.
- **Real-time Analytics**: Pelacakan statistik jumlah klik secara atomik untuk mencegah *race conditions*.
- **Rate Limiting**: Melindungi server dari serangan *brute-force* atau bot menggunakan middleware `tollbooth`.
- **Server-Side Rendering (SSR)**: Menggunakan Go Templates untuk performa *frontend* yang ringan dan SEO-friendly.

## 🛠 Tech Stack
- **Language**: [Go](https://go.dev/) (Golang) - Fokus pada efisiensi memori dan konkurensi.
- **API Framework**: [Gin Gonic](https://gin-gonic.com/) - Untuk routing yang cepat dan middleware yang fleksibel.
- **Database**: [PostgreSQL](https://www.postgresql.org/) - Sebagai penyimpanan data permanen.
- **Cache Layer**: [Redis](https://redis.io/) - Untuk akses data secepat kilat (In-memory storage).
- **CSS Framework**: [Tailwind CSS](https://tailwindcss.com/) - Untuk UI minimalis dan responsif.

---

## 🏗 Arsitektur Sistem
Sistem ini menggunakan alur **Cache-Aside Pattern**:
1. Request masuk -> Cek **Redis**.
2. Jika ada (Cache Hit): Langsung redirect.
3. Jika tidak ada (Cache Miss): Cek **PostgreSQL**, simpan ke Redis untuk akses berikutnya, lalu redirect.
# url-shortener
