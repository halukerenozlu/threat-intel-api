// Bu dosya uygulamanın giriş noktasıdır.
// Sorumluluğu sadece "wiring" yapmaktır: config yükle, bağlantıları kur,
// router'ı oluştur, sunucuyu başlat. İş mantığı burada YAZILMAZ.
package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: config.Load() ile ortam değişkenlerini oku
	// TODO: database.Connect() ile PostgreSQL bağlantısını kur
	// TODO: router'ı oluştur ve handler'ları bağla

	port := ":8080"
	log.Printf("threat-intel-api sunucusu %s portunda başlatılıyor...\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("sunucu başlatılamadı: %v", err)
	}
}

// healthCheckHandler basit bir sağlık kontrolü endpoint'i.
// Deploy sonrası sunucunun ayakta olduğunu doğrulamak için kullanılır.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
