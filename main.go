package main

import (
	"fmt"
	"ghulammuzz/go-grpc/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	user := &pb.Users{
		Data: []*pb.User{
			{
				Id:   "uuid-1",
				Name: "hendrasubila",
				Age:  20,
				Category: &pb.Category{
					Id:   "uuid-1",
					Name: "tua",
				},
			},
			{
				Id:   "uuid-2",
				Name: "purtrame",
				Age:  10,
				Category: &pb.Category{
					Id:   "uuid-2",
					Name: "muda",
				},
			},
		},
	}

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal("Marshal Error")
	}
	// compact binary wire format

	fmt.Println(data)

	// decode

	testUser := &pb.Users{}
	if err = proto.Unmarshal(data, testUser); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println(testUser)

	// loop

	for _, user := range testUser.GetData() {
		fmt.Println(user.GetName())
		fmt.Println(user.GetCategory().GetName())
	}
}
