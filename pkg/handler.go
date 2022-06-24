package pkg

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	webFS embed.FS
}

func NewHandler(webFS embed.FS) *Handler {
	return &Handler{
		webFS: webFS,
	}
}

func (h *Handler) Handle() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"ping": "pong"}`))
			return
		})
		r.Post("/decode-eth-raw", h.decodeEthRawTx())
	})
	router.Handle("/*", h.handleStatic("web/dist"))
	return router
}

func (h *Handler) decodeEthRawTx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := Dict{}
		i := DictFromReader(r.Body)
		raw := i.GetString("raw_tx")
		if raw == "" {
			res.Set("error", "raw_tx is empty")
			w.Write(ToByte(res))
		} else {
			tx, err := DecodeEthRawTx(raw)
			if err != nil {
				res.Set("error", err.Error())
				w.Write(ToByte(res))
			} else {
				d, _ := tx.MarshalJSON()
				transaction := ToDict(d)
				res.Set("tx", transaction)
				transaction["valueHex"] = transaction["value"]
				transaction["value"] = tx.Value()
				transaction["cost"] = tx.Cost()
				transaction["gasHex"] = transaction["gas"]
				transaction["gas"] = tx.Gas()
				transaction["gasPriceHex"] = fmt.Sprintf("0x%x", tx.GasPrice())
				transaction["gasPrice"] = tx.GasPrice()
				transaction["nonce"] = tx.Nonce()
				transaction["size"] = tx.Size()
				transaction["cost"] = tx.Cost()
				transaction["chainId"] = tx.ChainId()
				transaction["chainIdHex"] = fmt.Sprintf("0x%x", tx.ChainId())
				transaction["from"] = strings.ToLower(GetFromAddress(tx))
				w.Write(ToByte(res))
			}
		}
	}
}

func (h *Handler) handleStatic(sub string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		normalPath := r.URL.Path
		if !strings.HasPrefix(normalPath, "/") {
			normalPath = "/" + normalPath
		}
		normalPath = path.Clean(normalPath)
		fSys := fs.FS(h.webFS)
		contentStatic, _ := fs.Sub(fSys, sub)
		if _, err := contentStatic.Open(strings.TrimLeft(normalPath, "/")); err != nil {
			r.URL.Path = "/"
		}
		http.FileServer(http.FS(contentStatic)).ServeHTTP(w, r)
	}
}
