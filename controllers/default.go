package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"os"
	"path/filepath"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	if IsLinux() {
		path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configProvider := config.FromFile(path + "/conf/config.yaml")
		sdk, err := fabsdk.New(configProvider)

		fmt.Println("------------------------")
		if err != nil {
			log.Fatalf("create sdk fail:%s\n",err.Error())
		}

		mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("org1.example.com"))

		if err != nil {
		     log.Fatalf("create msp client fail:%s\n",err.Error())
		}

		adminIdentity, err := mspClient.GetSigningIdentity("Admin")
		if err != nil {
		     log.Fatalf("get admin identity fail:%s\n",err.Error())
		}else {
		     fmt.Println("AdminIdentity is found:")
		     fmt.Println(adminIdentity)
		}

		channelProvider := sdk.ChannelContext("mychannel",
			fabsdk.WithUser("Admin"),
			fabsdk.WithOrg("org1.example.com"))

		channelClient, err := channel.New(channelProvider)
		if err != nil {
			log.Fatalf("create channel client fail:%s\n",err.Error())
		}

		var args [][]byte
		args = append(args,[]byte("key1"))

		request := channel.Request{
			ChaincodeID: "demo",
			Fcn:         "query",
			Args:        args,
		}

		response, err := channelClient.Query(request)

		if err != nil {
			log.Fatalf("query fail:",err.Error())
		}else {
			fmt.Printf("response is %s\n",response.Payload)
		}
	}



	c.TplName = "manager/index.html"
}
