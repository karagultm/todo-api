# 📝 Todo API (Go + Echo + GORM + SQLite)

Basit bir **RESTful Todo API** uygulaması.  
Go dili kullanılarak [Echo](https://echo.labstack.com/) frameworkü üzerinde geliştirildi ve [GORM](https://gorm.io/) ile SQLite veritabanı entegre edildi.  

---

## 🚀 Özellikler
- ✅ Tüm todoları listeleme (`GET /todos`)  
- ➕ Yeni todo ekleme (`POST /todos`)  
- 🔄 Todo durumunu güncelleme / toggle etme (`PUT /todos/:id`)  
- ❌ Todo silme (`DELETE /todos/:id`)  
- 🗄️ SQLite üzerinde kalıcı veri saklama  

---

## 🛠️ Kurulum

### Gereksinimler
- [Go](https://go.dev/dl/) (1.21+ önerilir)  
- [Git](https://git-scm.com/)  

### Adımlar
```bash
# Depoyu klonla
git clone https://github.com/kullanici-adin/todoapi.git

# Dizine gir
cd todoapi

# Gerekli bağımlılıkları indir
go mod tidy

# Uygulamayı çalıştır
go run server.go

```
---
## 📡 API Endpointleri

### 1️⃣ Tüm todoları listele

``` http
GET /todos
```

### 2️⃣ Yeni todo ekle

``` http
POST /todos
Content-Type: application/json

{
  "content": "Github'a projelerini yükle!",
  "done": false
}
```

### 3️⃣ Todo durumunu değiştir

``` http
PUT /todos/1
```

### 4️⃣ Todo sil

``` http
DELETE /todos/1
```

---

## 🛠 Kullanılan Teknolojiler

-   [Go](https://go.dev/)
-   [Echo Framework](https://echo.labstack.com/)
-   [GORM](https://gorm.io/)
-   [SQLite](https://www.sqlite.org/)
