package main

import (
	"fmt"

	score_server "github.com/davidddw2017/panzer/gostudy/demo06/01_protobuf/src_gen"
	"github.com/golang/protobuf/proto"
)

func main() {
	score_info := &score_server.BaseScoreInfoT{}
	//score_info.WinCount = new(int32)
	score_info.WinCount = 1
	//score_info.LoseCount = new(int32)
	score_info.LoseCount = 2
	//score_info.ExceptionCount = new(int32)
	score_info.ExceptionCount = 3
	//score_info.KillCount = new(int32)
	score_info.KillCount = 4
	//score_info.DeathCount = new(int32)
	score_info.DeathCount = 5
	//score_info.AssistCount = new(int32)
	score_info.AssistCount = 6
	//score_info.Rating = new(int64)
	score_info.Rating = 1800

	fmt.Printf("original data{%s}\n", score_info.String())

	// encode
	data, err := proto.Marshal(score_info)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	// decode
	score_info_1 := &score_server.BaseScoreInfoT{}
	err = proto.Unmarshal(data, score_info_1)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", err.Error())
		return
	}

	score_info_1.Rating = 2000
	fmt.Printf("after decode:{%s}\n", score_info_1.String())
}
