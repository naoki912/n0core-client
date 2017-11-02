package n0core

import (
	"bufio"
	"fmt"
	"os"

	"github.com/takebayashi/go-pulsar"
	"github.com/takebayashi/go-pulsar/websocket"
)

type Client struct {
	Url          string
	Topic        string
	Subscription string
}

func NewClient(url string, topic string, subscription string) *Client {
	client := &Client{
		Url:          url,
		Topic:        topic,
		Subscription: subscription,
	}
	client.Initialize()

	return client
}

func (client *Client) Initialize() {}

func (client *Client) GetConsumerUrl() string {
	// "ws://localhost:8080/ws/consumer/persistent/property/cluster/namespace/topic/subscription",
	return client.Url + "/ws/consumer/persistent" + client.Topic + "/" + client.Subscription
}

func (client *Client) GetProducerUrl() string {
	// "ws://localhost:8080/ws/producer/persistent/property/cluster/namespace/topic",
	return client.Url + "/ws/producer/persistent" + client.Topic
}

func (client *Client) subscribe() {
	c, err := websocket.NewAsyncConsumer(client.GetConsumerUrl())
	if err != nil {
		panic(err)
	}

	for m := range c.Messages() {
		fmt.Printf("%+v: %s\n", m.Timestamp, string(m.Data))
		c.Ack(m)
	}
}

func (client *Client) publish() {
	c, err := websocket.NewAsyncProducer(client.GetProducerUrl())
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		c.Input() <- &pulsar.Message{Data: sc.Bytes()}
	}
}
