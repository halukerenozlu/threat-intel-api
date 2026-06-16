# threat-intel-api

Çoklu kaynaktan tehdit istihbaratı (threat intelligence) verisi toplayan,
birleştiren ve kimlik doğrulamalı bir REST API üzerinden sunan bir Go projesi.

## Ne işe yarar?

Bir IP adresi, URL veya dosya hash'i sorgulandığında sistem şu kaynaklara
aynı anda istek atar ve sonuçları birleştirip tek bir tehdit puanı döner:

- AbuseIPDB
- VirusTotal
- URLhaus
- AlienVault OTX

## Teknoloji yığını

- **Dil**: Go 1.23
- **Veritabanı**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)
- **Dahili iletişim** (planlanan): gRPC
- **AI özelliği** (planlanan): LLM tabanlı rapor özetleme

## Proje durumu

🚧 Geliştirme aşamasında — henüz çalışır bir sürüm yok.

## Kurulum

```bash
git clone https://github.com/halukerenozlu/threat-intel-api.git
cd threat-intel-api
cp .env.example .env
# .env dosyasını kendi API anahtarlarınla doldur
go mod download
go run cmd/server/main.go
```

## Mimari

Detaylı mimari ve veritabanı şeması için `docs/` klasörüne bakabilirsin.

## Lisans

TBD
