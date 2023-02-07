package handlers

import (
	dto "be/dto/result"
	transactiondto "be/dto/transaction"
	"be/models"
	"be/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go/coreapi"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey: os.Getenv("CLIENT_KEY"),
}

func (h *handlerTransaction) FindOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	//get user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	transaction, err := h.TransactionRepository.FindOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var responseTransaction []transactiondto.TransactionResponse
	for _, t := range transaction {
		responseTransaction = append(responseTransaction, convertResponseTransaction(t))
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: responseTransaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	transaction, err := h.TransactionRepository.GetOrderByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var data []models.Transaction
	// for _, order := range transaction {
	// 	dataGet := models.Transaction{
	// 		ID:       order.ID,
	// 		CheckIn:  order.CheckIn,
	// 		CheckOut: order.CheckOut,
	// 		HousyID:  order.HousyID,
	// 		UserID:   order.UserID,
	// 		Total:    order.Total,
	// 		Status:   order.Status,
	// 	}
	// 	data = append(data, dataGet)
	// }
	// fmt.Println(data)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var request transactiondto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//create id unix transaction
	// var TransIdIsMatch = false
	// var TransactionId int
	// for !TransIdIsMatch {
	// 	TransactionId = int(time.Now().Unix())
	// 	transactionData, _ := h.TransactionRepository.GetOrder(TransactionId)
	// 	if transactionData.ID == 0 {
	// 		TransIdIsMatch = true
	// 	}
	// }

	transaction := models.Transaction{
		// ID:       TransactionId,
		CheckIn:  request.CheckIn,
		CheckOut: request.CheckOut,
		HousyID:  request.HousyID,
		Total:    request.Total,
		UserID:   userId,
		Status:   "pending",
	}

	newTransaction, err := h.TransactionRepository.AddTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataTransactions, err := h.TransactionRepository.GetOrder(newTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// s := snap.Client{}
	// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// req := &snap.Request{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(dataTransactions.ID), //112233
	// 		GrossAmt: int64(dataTransactions.Total),
	// 	},
	// 	CreditCard: &snap.CreditCardDetails{
	// 		Secure: true,
	// 	},
	// 	CustomerDetail: &midtrans.CustomerDetails{
	// 		FName: dataTransactions.User.Username, //frm buyer
	// 		Email: dataTransactions.User.Email,    //buyer
	// 	},
	// }

	// snapResp, err := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: dataTransactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	delete, err := h.TransactionRepository.DeleteOrder(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: delete}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerTransaction) CreatePayment(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// }

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", orderId)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			h.TransactionRepository.UpdateTransaction("success", orderId)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		h.TransactionRepository.UpdateTransaction("success", orderId)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", orderId)
	}

	w.WriteHeader(http.StatusOK)
}

func convertResponseTransaction(t models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID:       t.ID,
		CheckIn:  t.CheckIn,
		CheckOut: t.CheckOut,
		Housy:    t.Housy,
		UserID:   t.UserID,
		Total:    t.Total,
		Status:   t.Status,
	}
}
