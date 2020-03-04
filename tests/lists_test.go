package tests

import (
	pb "amqp-subscriber/router"
	"context"
	"testing"
)

func TestLists(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Lists(
		context.Background(),
		&pb.ListsParameter{
			Identity: []string{"a1", "a2"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	println(response.Error)
}
