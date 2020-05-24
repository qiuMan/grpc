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
