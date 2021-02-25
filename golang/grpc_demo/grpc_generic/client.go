package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes"

	pb "grpc_demo/grpc_generic/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := ":50000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	req := new(pb.LoginRequest)

	// value
	req.JsCode = "value"
	r, err := c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("str generic failed: %v", err)
	}
	fmt.Print("\n\n***** value *****")
	fmt.Printf("\nret: %v", r)
	fmt.Printf("\nret.Msg: %v", r.Msg)
	fmt.Printf("\nret.Err: %v", r.Err)
	fmt.Printf("\nret.DataValue.Fields[open_id]: %v", r.DataValue.GetStructValue().Fields["open_id"].GetStringValue())
	fmt.Printf("\nret.DataValue.Fields[sessionkey]: %v", r.DataValue.GetStructValue().Fields["session_key"].GetStringValue())

	// any
	fmt.Print("\n\n***** any *****")
	req.JsCode = "any"
	r, err = c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("str generic failed: %v", err)
	}
	fmt.Printf("\nret: %v", r)
	fmt.Printf("\nret.Msg: %v", r.Msg)
	fmt.Printf("\nret.Err: %v", r.Err)
	if r.DataAny != nil {
		fmt.Printf("\nret.DataAny: %+v", r.DataAny)
		fmt.Printf("\nret.DataAny type_url: %v", r.DataAny.TypeUrl)
		fmt.Printf("\nret.DataAny value: %v", r.DataAny.Value)

		var result pb.LoginResult
		_ = ptypes.UnmarshalAny(r.DataAny, &result)
		fmt.Printf("\nret.DataAny.OpenId: [any]:%+v", result.OpenId)
		fmt.Printf("\nret.DataAny.SessionKey: [any]:%+v", result.SessionKey)
	}

	//output:
	//***** value *****
	//ret: err:1 msg:"test" data_value:{struct_value:{fields:{key:"open_id" value:{string_value:"openid..."}} fields:{key:"session_key" value:{string_value:"sessionkey..."}} fields:{key:"sizeCache" value:{number_value:0}} fields:{key:"state" value:{struct_value:{fields:{key:"DoNotCompare" value:{}} fields:{key:"DoNotCopy" value:{}} fields:{key:"NoUnkeyedLiterals" value:{}} fields:{key:"atomicMessageInfo" value:{}}}}} fields:{key:"unknownFields" value:{}}}}
	//ret.Msg: test
	//ret.Err: 1
	//ret.DataValue.Fields[open_id]: openid...
	//ret.DataValue.Fields[sessionkey]: sessionkey...

	//***** any *****
	//ret: err:1 msg:"test" data_value:{string_value:"test"} data_any:{[type.googleapis.com/user.LoginResult]:{open_id:"openid--any" session_key:"sessionkey-any"}}
	//ret.Msg: test
	//ret.Err: 1
	//ret.DataAny: [type.googleapis.com/user.LoginResult]:{open_id:"openid--any" session_key:"sessionkey-any"}
	//ret.DataAny type_url: type.googleapis.com/user.LoginResult
	//ret.DataAny value: [10 11 111 112 101 110 105 100 45 45 97 110 121 18 14 115 101 115 115 105 111 110 107 101 121 45 97 110 121]
	//ret.DataAny.OpenId: [any]:openid--any
	//ret.DataAny.SessionKey: [any]:sessionkey-any
}
