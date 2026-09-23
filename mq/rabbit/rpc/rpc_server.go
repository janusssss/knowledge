package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "faild to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("rpc_queue", false, false, false, false, nil)
	failOnError(err, "fialed to declare a queue")

	err = ch.Qos(1, 0, false)
	failOnError(err, "failed to set QoS")

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")
			log.Printf("[.] fib(%d)", n)
			response := fib(n)

			ch.PublishWithContext(ctx, "", d.ReplyTo, false, false,
				amqp.Publishing{ContentType: "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response))})
			failOnError(err, "Failed to publish a message")
			d.Ack(false)
		}
	}()

	log.Println(" [*] Awaiting RPC requests")
	<- forever

}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
