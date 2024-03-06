package main

import (
	"encoding/json"
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/utils"
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
	// log.Printf("Temporal: RequestTopic: %v, Msg: %v\n", requestTopic, fMsg)
	// ======================== SEND REQUEST ========================

	// ======================== GET RESPONSE ========================
	for i := 0; i < 10; i++ {
		msg, _ := bk.PublishAndReceive(
			requestTopic,
			&fMsg,
			broker.WithPublishReplyToTopic(replyTopic),
			broker.WithReplyConsumerGroup(utils.GetConsumerGroup(workflowID, action)),
		)

		log.Printf("PublishAndReceive Lan %d: RequestTopic: %v, ReplyTopic: %v, Msg: %v\n", i, requestTopic, replyTopic, msg)
	}
	select {}
	// ======================== GET RESPONSE ========================
}
