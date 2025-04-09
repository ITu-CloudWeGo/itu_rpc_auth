package kafka

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/service/jwt"
	"github.com/segmentio/kafka-go"
)

func ReaderUidAndGenerateToken() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "uid_topic",
		CommitInterval: 1 * time.Second,
		GroupID:        "uid_group",
		StartOffset:    kafka.FirstOffset,
	})

	writer := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "token_topic",
		Balancer:               &kafka.Hash{},
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}

	defer reader.Close()
	defer writer.Close()

	for {
		// 读取消息
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		var uidMessage map[string]string
		err = json.Unmarshal(msg.Value, &uidMessage)
		if err != nil {
			log.Printf("Failed to unmarshal UID message: %v", err)
			continue
		}

		uidStr := uidMessage["uid"]
		uid, err := strconv.ParseUint(uidStr, 10, 64)
		if err != nil {
			log.Printf("Failed to parse UID: %v", err)
			continue
		}

		// 生成访问令牌和刷新令牌
		accessToken, accessExpire, err := jwt.GenerateAccessToken(uid)
		if err != nil {
			log.Printf("Failed to generate access token: %v", err)
			continue
		}

		refreshToken, refreshExpire, err := jwt.GenerateRefreshToken(uid)
		if err != nil {
			log.Printf("Failed to generate refresh token: %v", err)
			continue
		}

		// 构造要发送到 Kafka 的消息
		tokenMessage := map[string]interface{}{
			"uid":                  uid,
			"access_token":         accessToken,
			"refresh_token":        refreshToken,
			"access_token_expire":  accessExpire,
			"refresh_token_expire": refreshExpire,
		}

		value, err := json.Marshal(tokenMessage)
		if err != nil {
			log.Printf("Failed to marshal token message: %v", err)
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: value,
		})
		if err != nil {
			log.Printf("Failed to send token to Kafka: %v", err)
			continue
		}

		log.Printf("Tokens sent to Kafka for UID: %d", uid)
	}
}
