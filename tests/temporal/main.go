package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/core/broker/kafka"
)

// Temporal: Gửi request từ Temporal -> Nhận request từ Micro
func main() {
	// ======================= BROKER =======================
	bk := kafka.ConnectBrokerKafka("127.0.0.1:9092")
	// ======================= BROKER =======================

	workflowID := "BANK_TRANSFER-1709525114"
	requestTopic := config.Messages.CHECK_BALANCE_REQUEST_TOPIC
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC
	action := config.Messages.CHECK_BALANCE_ACTION

	// ======================== SEND REQUEST ========================
	req := account.CheckBalanceReq{}
	body, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	fMsg := broker.Message{
		Body: body,
		Headers: map[string]string{
			"workflow_id": workflowID,
			"activity-id": action,
		},
	}

	bk.Publish(requestTopic, &fMsg)
	log.Printf("Temporal: requestTopic: %v, msg %v\n", requestTopic, fMsg)
	// ======================== SEND REQUEST ========================

	// ======================== GET RESPONSE ========================

	isReceived := false
	var res account.CheckBalanceRes // TODO: check lại với Sơn
	// csGroupOpt := broker.WithSubscribeGroup(config.Messages.GROUP)

	// Loop -> khi nào có message phù hợp -> Nhận + parse message -> Done activity
	// TODO: Trường hợp không tìm thấy được message phù hợp -> Timeout
	for {
		bk.Subscribe(replyTopic, func(e broker.Event) error {
			headers := e.Message().Headers
			fmt.Printf("Received message from topic %v: Header: %v\n", replyTopic, headers)
			// TODO: Nhận response từ API Microservice push vào topic Reply

			// Kiểm tra theo điều kiện phù hợp
			if headers["workflow_id"] == workflowID && headers["activity-id"] == action { // TODO: check lại với Sơn
				body := string(e.Message().Body)
				if body != "" {
					err := json.Unmarshal(e.Message().Body, &res)
					if err != nil {
						return err // Đúng message + Payload res bị sai struct -> Fail Activity
					} else {
						isReceived = true
						log.Println("Temporal: CheckBalance success", res.Balance == 8888) // Check ok
					}
				}
			}

			return nil
		}) //, csGroupOpt)
		// }, csGroupOpt)

		if isReceived {
			break
		}
	}

	if isReceived {
		log.Println("Temporal: CheckBalance done", res)
	}
}
