package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"

	//"github.com/couchbase/go-couchbase"
	pb "github.com/thisisalisaboori/microcrud/api/proto/microcrudproto"
)

type Server struct {
	pb.UnimplementedCrudServiceServer
}

func Connect(collection string) (*gocb.Cluster, *gocb.Collection) {
	c, err := ConnectCluster()
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	bucket := c.Bucket("foo")
	if err := bucket.WaitUntilReady(15*time.Second, nil); err != nil {
		log.Fatal("bucket not ready:", err)
	}
	CreateCollection(bucket, collection)
	col := bucket.Collection(collection)

	return c, col

}
func ConnectCluster() (*gocb.Cluster, error) {

	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: "administrator",
			Password: "11235813",
		},
	}
	c, err := gocb.Connect("couchbase://127.0.0.1", opts)

	return c, err

}
func CreateCollection(bucket *gocb.Bucket, name string) {
	cm := bucket.Collections()
	// ساخت Collection در Scope
	_ = cm.CreateCollection(gocb.CollectionSpec{
		ScopeName: "_default",
		Name:      name,
	}, nil)

}

func (s *Server) CreateItem(ctx context.Context, request *pb.CreateItemRequest) (*pb.BaseResponse, error) {
	fmt.Printf("create .... %s %v , ", request.Entity, request.Data)
	cluster, c := Connect(request.Entity)
	defer cluster.Close(nil)

	id := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := request.Data.AsMap()
	fmt.Println(data)
	_, err := c.Insert(id.String(), data, &gocb.InsertOptions{Context: ctx})
	if err != nil {
		log.Fatalf("failed to upsert document: %v", err)
		return &pb.BaseResponse{Ok: false}, err
	}
	return &pb.BaseResponse{Ok: true}, nil
}

func (s *Server) UpdateItem(ctx context.Context, request *pb.UpdateItemRequest) (*pb.BaseResponse, error) {
	fmt.Println("update ....")
	cluster, c := Connect(request.Entity)
	defer cluster.Close(nil)

	id := request.Id
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := request.Data.AsMap()
	_, err := c.Upsert(id, data, &gocb.UpsertOptions{Context: ctx})
	if err != nil {
		log.Fatalf("failed to upsert document: %v", err)
		return &pb.BaseResponse{Ok: false}, err
	}
	return &pb.BaseResponse{Ok: true}, nil
}

func (s *Server) DeleteItem(ctx context.Context ,request *pb.DeleteItemRequest) (*pb.BaseResponse, error) {
 	cluster,col:=Connect(request.Entity)
	defer cluster.Close(nil)
	_,err:= col.Remove(request.Id , &gocb.RemoveOptions{}  )
	if err !=nil{
		return  &pb.BaseResponse{Ok: false},err
	}
	return  &pb.BaseResponse{Ok: true},nil
}
func (s *Server) GetItemById(ctx context.Context, request *pb.GetItemRequest) (*pb.GetByIdResponse, error) {
	fmt.Println("get item by id ....")
	cluster, c := Connect(request.Entity)
	defer cluster.Close(nil)
	q, err := c.Get(request.Id, nil)
	if err != nil{
		return &pb.GetByIdResponse{Ok: false, Data: nil}, nil
	
	}
	var data map[string]interface{}
	q.Content(&data)
	js, _ := structpb.NewStruct(data)
	return &pb.GetByIdResponse{Ok: true, Data: js}, nil
}





func (s *Server) GetItems(ctx context.Context, request *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	fmt.Println("get items....%s  %d %d", request.Entity ,request.PageIndex ,request.PageSize )
	cluster, col := Connect(request.Entity)
	defer cluster.Close(nil)
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dataResult := make([]*pb.GetByIdResponse, 0)
	scope := col.ScopeName()
	colnmae := col.Name()
	if request.PageIndex == 0{
		request.PageIndex=1
	}
	skip := (request.PageIndex - 1) * request.PageSize
	q := fmt.Sprintf("SELECT  Meta().id,c.* FROM `foo`.%s.%s  c offset %d  LIMIT %d;", scope, colnmae, skip, request.PageSize)
	fmt.Println(q)
	result, q_err := cluster.Query(q, &gocb.QueryOptions{Adhoc: true})
	if q_err != nil {
		log.Fatalf("query failed: %v", q_err)
	}
	var row map[string]interface{}
	for result.Next() {
		var record pb.GetByIdResponse
		read_error := result.Row(&row)
		if read_error == nil {
			js, _ := structpb.NewStruct(row)
			record.Data = js
			dataResult = append(dataResult, &record)
		}
		//record= row["Data"]

	}
	return &pb.GetItemsResponse{Ok: true, Data: dataResult}, nil
}
