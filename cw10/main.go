package main

import (
	"course/cw10/wallet"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// /wallets GET, POST
// /wallets/{id} GET, PUT, DELETE

// /exchange GET

func main() {
	r := mux.NewRouter()

	walletRes := &walletResource{
		storage: *wallet.NewStorage(),
	}

	exchangeRes := &exchangeResource{}

	r.HandleFunc("/wallets", walletRes.getWallets).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}", walletRes.getWalletById).Methods(http.MethodGet)
	r.HandleFunc("/wallets", walletRes.createWallet).Methods(http.MethodPost)
	r.HandleFunc("/exchange", exchangeRes.getExchange).Methods(http.MethodGet)
	r.HandleFunc("/wallets/{id}/amount", walletRes.putWalletAmount).Methods(http.MethodPut)
	r.HandleFunc("/wallets/{id}", walletRes.deleteWallet).Methods(http.MethodDelete)

	log.Default().Println("Server starting..")
	log.Fatal(http.ListenAndServe(":8082", r))
}

type walletResource struct {
	storage wallet.Storage
}

func (wr *walletResource) createWallet(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("handling request wallets")

	var wallet wallet.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		log.Default().Println("Something wrong while read wallet")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if exists, _ := wr.storage.GetWalletById(wallet.ID); exists != nil {
		log.Default().Println("Wallet already exists")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.storage.AddWallet(&wallet)

	respondWithJson(w, wallet)
}

func (wr *walletResource) getWallets(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("handling request wallets")

	respondWithJson(w, wr.storage.GetAll())
}

type exchangeResource struct {
}

const frankUrl = "https://api.frankfurter.app/latest?from=USD&to=EUR"

func (er *exchangeResource) getExchange(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(frankUrl)
	if err != nil {
		log.Default().Println("Failed to get course")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Default().Println("Non OK status from exchange")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type respBody struct {
		Rates map[string]float64 `json:"rates"`
	}

	var rates respBody

	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		log.Default().Println("cant decode rates %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respondWithJson(w, rates)

}

func (wr *walletResource) getWalletById(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("handling request wallets")
	walletId := mux.Vars(r)["id"]

	walletById, err := wr.storage.GetWalletById(walletId)
	if err != nil {
		log.Default().Println("Something wrong while read wallet")
		w.WriteHeader(http.StatusInternalServerError)
	}

	respondWithJson(w, walletById)
}

func (wr *walletResource) putWalletAmount(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("handling request wallets")
	walletId := mux.Vars(r)["id"]

	type amountBody struct {
		Amount int `json:"amount"`
	}

	var amount amountBody

	err := json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		log.Default().Println("Something wrong while read wallet %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	walletById, err := wr.storage.GetWalletById(walletId)
	if err != nil {
		log.Default().Println("Something wrong while read wallet %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.storage.UpdateAmount(walletById, amount.Amount)

	respondWithJson(w, walletById)
}

func (wr *walletResource) deleteWallet(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("handling request wallets")
	walletId := mux.Vars(r)["id"]

	wr.storage.DeleteWallet(walletId)

	type result struct {
		success bool `json:"success"`
	}

	res := result{success: true}

	respondWithJson(w, res)
}

func respondWithJson(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Default().Println("Something wrong while read wallet")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
