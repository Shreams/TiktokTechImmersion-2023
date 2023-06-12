package main

import (
	"context"
	"log"

	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	// sqldb = &SqlClient{} // make the sql client global in the 'main' package
	rdb = &RedisClient{}
)

func main() {

	ctx := context.Background()

	err := rdb.InitClient(ctx, "redis:6379", "") // "" for no password

	if err != nil {
		errMsg := "failed to init Redis client, err: " + err.Error()
		log.Fatal(errMsg)
	}

	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

// main(){
// 	// err := sqldb.InitClient(ctx, "localhost:3306", "root", "password")

// 	// if err != nil {
// 	// 	errMsg := fmt.Sprintf("failed to init SQL client, err: %v", err)
// 	// 	log.Fatal(errMsg)
// 	// }
// }

// type SqlClient struct {
// 	db *sql.DB
// }

// func (s *SqlClient) InitClient(ctx context.Context, address, user, password string) error {

// 	sqlType := "mysql"
// 	sqlUser := user
// 	sqlPassword := password
// 	sqlAddress := address

// 	db, err := sql.Open(sqlType, sqlUser+":"+sqlPassword+"@tcp("+sqlAddress+")/")

// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	fmt.Println("Successfully connected to MySQL database")

// 	s.db = db

// 	return nil
// }
