package ipbanner

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

// BanIPs bans your banip.txt's ips
func BanIPs(next http.Handler) http.Handler {
	file := readtxtfile()
	r := chi.NewRouter()
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		ips := strings.Split(file, "\n")
		for ip := 0; ip < len(ips); ip++ {
			if ips[ip] == r.RemoteAddr {
				w.WriteHeader(400)
				w.Write([]byte(r.RemoteAddr + " has been unathorized to visit this page."))
			}
		}
		next.ServeHTTP(w, r)
	})
	return next
}

func readtxtfile() string {
	wd, _ := os.Getwd()
	text, _ := ioutil.ReadFile(wd + "/banip.txt")
	return string(text)
}
