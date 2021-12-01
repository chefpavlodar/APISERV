package main

import (
	APISERV "APISERV/proto"
	"context"
	"flag"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":50001", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	c := APISERV.NewConsoleClient(conn)

	res, err := c.Add(context.Background(), &APISERV.AddRequest{X: int32(x), Y: int32(y)})

	if err != nil {
		log.Fatal(err)

	}

	log.Println(res.GetResult())
}
