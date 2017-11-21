package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
	"./protobuf/maxwell_smarts"
	"github.com/golang/protobuf/proto"
)

func main() {

	broker := "kafka.docker:9092"
	group := "foo"
	topics := []string{"maxwellsmarts.ticket_events"}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           group,
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
		"default.topic.config": kafka.ConfigMap{
			"auto.offset.reset":  "earliest",
			"enable.auto.commit": false}})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(topics, nil)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				//fmt.Printf("%% Message on %s:\n%s\n",
				//	e.TopicPartition, string(e.Value))

				ticketEvents := &com_zendesk_maxwellsmarts_zendesk_ticketevents.TicketEvents{}
				err = proto.Unmarshal(e.Value, ticketEvents)
				//if err != nil {
				//	log.Fatal("unmarshaling error: ", err)
				//}
				fmt.Println("account ID: ", ticketEvents.AccountId)
				fmt.Println("ticket ID", ticketEvents.TicketId)

			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}
