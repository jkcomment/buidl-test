package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	resp "github.com/nicklaw5/go-respond"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RequestData struct {
	Message string `json:message`
}

var CIPHER_KEY = []byte("0123456789012345")

func main() {
	router := chi.NewRouter()
	router.Post("/encrypt", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		req := RequestData{}
		if err := json.Unmarshal(body, &req); err != nil {
			resp.NewResponse(w).UnprocessableEntity(err.Error())
			return
		}

		encmess, err := Encrypt(CIPHER_KEY, req.Message)
		if err != nil {
			resp.NewResponse(w).UnprocessableEntity(err.Error())
			return
		}

		resp.NewResponse(w).Created(encmess)
	})

	router.Post("/decrypt", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		req := RequestData{}
		if err := json.Unmarshal(body, &req); err != nil {
			resp.NewResponse(w).UnprocessableEntity(err.Error())
			return
		}

		decodedmess, err := Decrypt(CIPHER_KEY, req.Message)
		if err != nil {
			resp.NewResponse(w).UnprocessableEntity(err.Error())
			return
		}

		resp.NewResponse(w).Created(decodedmess)
	})

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
