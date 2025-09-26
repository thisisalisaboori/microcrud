package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/thisisalisaboori/microcrud/api/proto/microcrudproto"
	"google.golang.org/grpc"
	 "google.golang.org/protobuf/types/known/structpb"
)

type mydata struct{
	Name string `json:"fname"`
	LastName string `json:"lname"`
 }

func main(){
   conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()




    client := pb.NewCrudServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
    defer cancel()
	
	
	client.Init(ctx, &pb.InitRequst{Bucket: "b" , Collection: "person" , CreateIndex: true})
	
	//return 
	res, err := client.CreateItem(ctx, &pb.CreateItemRequest{
        Entity: "person",
		Bucket: "b",
		Data : &structpb.Struct{ Fields: map[string]*structpb.Value {"name" :structpb.NewStringValue("elahe"),
		 "lastname" :structpb.NewStringValue("saboori"), "age": structpb.NewNumberValue(38) } },
    })

	
	fmt.Println(res.Ok)

	//x,_:=client.GetItemById(ctx, &pb.GetItemRequest{Id: "90f57f37-0354-4fae-a3d2-3e5efc494573" , Entity: "newcol"})
	
	x,_:=client.GetItems(ctx, &pb.GetItemsRequest{Entity: "person" , PageSize: 10 ,PageIndex: 1 ,Bucket: "b"  })
	fmt.Println(x.Data)

	for _,data:= range(x.Data){
		js:= data.Data.AsMap()
		id:=fmt.Sprintf("%s",js["id"])
		fmt.Printf(id)
		//id:= string(js["id"], )  
		//client.DeleteItem(ctx, &pb.DeleteItemRequest{ Id: id  , Entity: "mycollection" })
		_, _ = client.UpdateItem(ctx, &pb.UpdateItemRequest{
        Entity: "person",
		Bucket: "b",
		Id: id,
		Data : &structpb.Struct{ Fields: map[string]*structpb.Value {"name" :structpb.NewStringValue("ali"),
		 "lastname" :structpb.NewStringValue("saboori"), "age": structpb.NewNumberValue(39),"city":structpb.NewStringValue("tehran") } },
    })
	}

}