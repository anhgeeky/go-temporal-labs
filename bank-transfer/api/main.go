package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	app "github.com/anhgeeky/go-temporal-labs/bank-transfer"
	"github.com/anhgeeky/go-temporal-labs/bank-transfer/config"
	"github.com/anhgeeky/go-temporal-labs/bank-transfer/domain"

	"github.com/bojanz/httpx"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.temporal.io/sdk/client"
)

type (
	ErrorResponse struct {
		Message string
	}

	UpdateEmailRequest struct {
		Email string
	}

	CheckoutRequest struct {
		Email string
	}
)

var (
	temporal client.Client
	PORT     string
)

func main() {
	PORT := os.Getenv("PORT")
	var err error
	temporal, err = client.NewLazyClient(client.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	log.Println("Temporal client connected")

	r := mux.NewRouter()
	r.Handle("/accounts", http.HandlerFunc(GetAccountsHandler)).Methods("GET")
	r.Handle("/bank-transfer", http.HandlerFunc(CreateCartHandler)).Methods("POST")
	r.Handle("/bank-transfer/{workflowID}", http.HandlerFunc(GetCartHandler)).Methods("GET")
	r.Handle("/bank-transfer/{workflowID}/add", http.HandlerFunc(AddToCartHandler)).Methods("PUT")
	r.Handle("/bank-transfer/{workflowID}/remove", http.HandlerFunc(RemoveFromCartHandler)).Methods("PUT")
	r.Handle("/bank-transfer/{workflowID}/checkout", http.HandlerFunc(CheckoutHandler)).Methods("PUT")
	r.Handle("/bank-transfer/{workflowID}/email", http.HandlerFunc(UpdateEmailHandler)).Methods("PUT")

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	var cors = handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))

	http.Handle("/", cors(r))
	server := httpx.NewServer(":"+PORT, http.DefaultServeMux)
	server.WriteTimeout = time.Second * 240

	log.Println("Starting server on port: " + PORT)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	res := domain.AccountList{}
	res.Accounts = domain.Accounts

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateCartHandler(w http.ResponseWriter, r *http.Request) {
	workflowID := "CART-" + fmt.Sprintf("%d", time.Now().Unix())

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.Workflows.BANK_TRANSFER,
	}

	cart := app.TransferState{Items: make([]app.CartItem, 0)}
	we, err := temporal.ExecuteWorkflow(context.Background(), options, app.TransferWorkflow, cart)
	if err != nil {
		WriteError(w, err)
		return
	}

	res := make(map[string]interface{})
	res["cart"] = cart
	res["workflowID"] = we.GetID()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := temporal.QueryWorkflow(context.Background(), vars["workflowID"], "", "getCart")
	if err != nil {
		WriteError(w, err)
		return
	}
	var res interface{}
	if err := response.Get(&res); err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item app.CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		WriteError(w, err)
		return
	}

	update := app.AddToCartSignal{Route: app.RouteTypes.ADD_TO_CART, Item: item}

	err = temporal.SignalWorkflow(context.Background(), vars["workflowID"], "", app.SignalChannels.ADD_TO_CART_CHANNEL, update)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["ok"] = 1
	json.NewEncoder(w).Encode(res)
}

func RemoveFromCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item app.CartItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		WriteError(w, err)
		return
	}

	update := app.RemoveFromCartSignal{Route: app.RouteTypes.REMOVE_FROM_CART, Item: item}

	err = temporal.SignalWorkflow(context.Background(), vars["workflowID"], "", app.SignalChannels.REMOVE_FROM_CART_CHANNEL, update)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["ok"] = 1
	json.NewEncoder(w).Encode(res)
}

func UpdateEmailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body UpdateEmailRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteError(w, err)
		return
	}

	updateEmail := app.UpdateEmailSignal{Route: app.RouteTypes.UPDATE_EMAIL, Email: body.Email}

	err = temporal.SignalWorkflow(context.Background(), vars["workflowID"], "", app.SignalChannels.UPDATE_EMAIL_CHANNEL, updateEmail)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["ok"] = 1
	json.NewEncoder(w).Encode(res)
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteError(w, err)
		return
	}

	checkout := app.CheckoutSignal{Route: app.RouteTypes.CHECKOUT, Email: body.Email}

	err = temporal.SignalWorkflow(context.Background(), vars["workflowID"], "", app.SignalChannels.CHECKOUT_CHANNEL, checkout)
	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["sent"] = true
	json.NewEncoder(w).Encode(res)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	res := ErrorResponse{Message: "Endpoint not found"}
	json.NewEncoder(w).Encode(res)
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	res := ErrorResponse{Message: err.Error()}
	json.NewEncoder(w).Encode(res)
}
