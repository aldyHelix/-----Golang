package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "Nike Air Jordan",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("marshal error", err)
	}

	fmt.Println(data)
}
