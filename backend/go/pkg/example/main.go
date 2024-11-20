package main

import (
	"context"
	"fmt"
	"time"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/http"
)

func main() {
	client := http.NewClient()

	output, err := client.CreateImage(
		context.Background(),
		appRPC.CreateImageInput{
			SourcePath: "camp-gia.png",
			TargetPath: "ryo.jpg",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
	for {
		time.Sleep(1 * time.Second)

		output2, err := client.GetStatus(
			context.Background(),
			appRPC.GetStatusInput{
				JobID: output.JobID,
			},
		)
		if err != nil {
			panic(err)
		}

		if output2.Status == "ok" {
			break
		}
		fmt.Println(output2.Status)

	}
	// ここで1秒待つ
	output3, err := client.GetImagePath(
		context.Background(),
		appRPC.GetImagePathInput{
			JobID: output.JobID,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(output3.ImageURL)

	fmt.Println("ok")
}
