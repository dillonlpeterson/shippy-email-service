package main

import (
	"context"
	"log"

	pb "github.com/dillonlpeterson/shippy-user-service/proto/user"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(crx context.Context, user *pb.User) error {
	log.Println("Picked up a message")
	log.Println("Sending email to: ", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}
