package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@172.16.10.14:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"hello", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// failOnError(err, "Failed to declare a queue")
	queue := "hello"
	_ = ch.Confirm(false)
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	ctx, cancel := context.WithCancel(context.Background())

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			fmt.Println("goroutine loop...")
			select {
			case confirmed := <-confirms:
				if confirmed.Ack {
					log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
				} else {
					log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
				}
			case <-ctx.Done():
				fmt.Println("confirm cancel done...")
				return ctx.Err()
			}
		}
	})

	g.Go(func() error {
		for {
			select {
			case <-time.After(time.Second * 3):
				body := "Hello World!"
				id, _ := uuid.NewUUID()
				err = ch.Publish(
					"",    // exchange
					queue, // routing key
					false, // mandatory
					false, // immediate
					amqp.Publishing{
						ContentType:   "text/plain",
						CorrelationId: id.String(),
						Body:          []byte(body),
					})
				failOnError(err, "Failed to publish a message")
				log.Printf(" [x] Sent %s", body)
			case <-ctx.Done():
				fmt.Println("publish cancel done")
				return ctx.Err()
			}
		}
	})

	// if err := g.Wait(); err != nil {
	// 	fmt.Println("wait error:", err)
	// }
	fmt.Println("started")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
	<-c
	cancel()
	fmt.Println("quite")
	time.Sleep(time.Second * 3)
}
