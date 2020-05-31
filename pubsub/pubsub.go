package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
)

//PublishService 发布服务
type PublishService struct {
	pub *pubsub.Publisher
}

type NewPublisherService() *PublishService {
	return &PublishService {
		pub: pubsub.NewPublisher(time.Microsecond*100, 10)
	}
}

func (p *PublishService) Publish(ctx context.Conetxt, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

func (p *PublishService) Subscribe(arg *String, stream PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}

		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	p := pubsub.NewPublisher(100*time.Microsecond, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}

		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}

		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: hello golang")
	go p.Publish("docker: hellp docker")

	time.Sleep(2)

	go func() {
		fmt.Println("golang topic: ", <-golang)
	}()

	go func() {
		fmt.Println("docker topic: ", <-docker)
	}()

	time.Sleep(time.Second * 5)

}
