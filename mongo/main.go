package main

import (
	"context"
	"fmt"
	"log"
	"mongo/domain"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Mongo Connection Boiler Plate
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	ipspec := domain.NewIpAddrSpec([]string{"192.168.1.1/24"}, []string{"192.168.2.1/24"})
	res, err := client.Database("asxtest").Collection("IpAddrSpec").InsertOne(ctx, ipspec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", *res)

}

// func mkLineItem() *LineItem {

// 	li := &domain.LineItem{
// 		Base:         NewBase(),
// 		CampaignId:   randx.ULID(),
// 		AdvertiserId: randx.ULID(),
// 		Name:         "LineItem-1",
// 		FCap:         domain.FreqencyCap{MaxImpressions: 10, NumTimeUnits: 10, TimeUnit: domain.TIMEUNIT_DAY},
// 		Targeting:    []domain.TargetingSpec{},
// 	}

// 	return nil
// }
