package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisAuthStore struct {
	clientAuth *redis.Client
}

func NewRedisAuthStore(redisAddr string) *RedisAuthStore {
	var client *redis.Client

	for retries := 0; retries < 5; retries++ {
		client = redis.NewClient(&redis.Options{
			Addr: redisAddr,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		status := client.Ping(ctx)
		if status.Err() == nil {
			fmt.Println("Connected to Redis successfully")
			store := &RedisAuthStore{clientAuth: client}
			go store.reconnect(redisAddr)
			return store
		}

		fmt.Printf("Failed to connect to Redis: %v. Retrying in 5 seconds...", status.Err())
		time.Sleep(5 * time.Second)
	}

	fmt.Printf("Failed to connect to Redis after 5 attempts")
	return nil
}

func (store *RedisAuthStore) reconnect(redisAddr string) {
	for {
		time.Sleep(30 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		status := store.clientAuth.Ping(ctx)
		cancel()
		if status.Err() != nil {
			fmt.Printf("Lost connection to Redis: %v. Attempting to reconnect...", status.Err())
			for retries := 0; retries < 5; retries++ {
				client := redis.NewClient(&redis.Options{
					Addr: redisAddr,
				})

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				status := client.Ping(ctx)
				cancel()
				if status.Err() == nil {
					fmt.Println("Reconnected to Redis successfully")
					store.clientAuth = client
					break
				}

				fmt.Printf("Failed to reconnect to Redis: %v. Retrying in 5 seconds...", status.Err())
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (s *RedisAuthStore) Get(key string) (string, error) {
	localCtx := context.Background()
	value, err := s.clientAuth.Get(localCtx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *RedisAuthStore) Set(key string, value interface{}) error {
	localCtx := context.Background()
	err := s.clientAuth.Set(localCtx, key, value, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisAuthStore) PushQueue(key string, value interface{}) error {
	ctx := context.Background()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.clientAuth.LPush(ctx, key, data).Err()
}

// func (s *RedisAuthStore) PopQueue(key string) (string, error) {
// 	ctx := context.Background()
// 	result, err := s.clientAuth.RPop(ctx, key).Result()
// 	if err == redis.Nil {
// 		return "", nil // คิวว่าง
// 	}
// 	return result, err
// }

func (s *RedisAuthStore) PopQueue(key string) (string, error) {
	ctx := context.Background()
	result, err := s.clientAuth.BLPop(ctx, 0*time.Second, key).Result()
	if err == redis.Nil {
		return "", nil // คิวว่าง
	}
	if err != nil {
		return "", err
	}
	// BLPop คืน slice [key, value] → เราเอา value มาใช้
	if len(result) > 1 {
		return result[1], nil
	}
	return "", err
}

func (s *RedisAuthStore) SubscriberEvent(ch string) *redis.PubSub {
	ctx := context.Background()
	return s.clientAuth.Subscribe(ctx, ch)
}

func (s *RedisAuthStore) StartSubscriber(ch string) {
	ctx := context.Background()
	pubsub := s.clientAuth.Subscribe(ctx, ch)
	defer pubsub.Close()

	log.Println("👂 Listening for messages on", ch)

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("❌ Error receiving message:", err)
			continue
		}
		// var result map[string][]interface{}
		// if err := json.Unmarshal([]byte(msg.Payload), &result); err != nil {
		// 	log.Println("unmarshal error:", err)
		// 	continue
		// }
		fmt.Printf("📨 Received on %s: %s\n", msg.Channel, msg.Payload)
	}

}

func (s *RedisAuthStore) StartPublisher(ch string, msg string) error {
	ctx := context.Background()
	err := s.clientAuth.Publish(ctx, ch, msg).Err()
	if err != nil {
		log.Println("❌ Publish error:", err)
		return err
	} else {
		fmt.Println("📤 Published:", msg)
		return nil
	}
	// for {
	// 	message := fmt.Sprintf("Hello at %v", time.Now().Format(time.RFC3339))
	// 	err := s.clientAuth.Publish(ctx, ch, message).Err()
	// 	if err != nil {
	// 		log.Println("❌ Publish error:", err)
	// 	} else {
	// 		fmt.Println("📤 Published:", message)
	// 	}
	// 	time.Sleep(3 * time.Second)
	// }
}

// //example pub
// func (r *ReceivedDataController) SendDataRealtime(num int, pubName string, data map[string]interface{}) {
// 	jsonRealtime := make(map[int]interface{})
// 	jsonRealtime[num] = data
// 	jsonBytes, err := json.Marshal(jsonRealtime)
// 	if err != nil {
// 		log.Println("marshal error:", err)
// 		return
// 	}
// 	jsonPayloadStr := string(jsonBytes)
// 	if err := r.RedisStore.StartPublisher(fmt.Sprintf("%s_%v", pubName, num), jsonPayloadStr); err != nil {
// 		log.Println("publish error:", err)
// 		return
// 	}
// 	log.Println("publish success:")
// }

// //example sub
// func (s *SubScriberDeviceController) PumpControl() {
// 	ctx := context.Background()
// 	pubsub := s.RedisStore.SubscriberEvent("data-send")
// 	defer pubsub.Close()

// 	log.Println("👂 Listening for messages on data-send")

// 	for {
// 		msg, err := pubsub.ReceiveMessage(ctx)
// 		if err != nil {
// 			log.Println("❌ Error receiving message:", err)
// 			continue
// 		}
// 		var result map[string]interface{}
// 		if errJson := json.Unmarshal([]byte(msg.Payload), &result); errJson != nil {
// 			log.Println("unmarshal error:", errJson)
// 			continue
// 		}

// 	}
// }
