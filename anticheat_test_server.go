package main

import (
	"context"
	"fmt"
	"github.com/blackjack200/xyron/xyron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	c := xyron.NewAnticheatClient(conn)
	p, err := c.AddPlayer(context.TODO(), &xyron.AddPlayerRequest{
		Player: &xyron.Player{
			Os:   0,
			Name: "IPlayfordev",
		},
		Data: map[int32]*xyron.TimestampedReportData{
			0: {
				Timestamp: 0,
				Data: []*xyron.WildcardReportData{
					{Data: &xyron.WildcardReportData_GameModeData{GameModeData: &xyron.PlayerGameModeData{GameMode: xyron.GameMode_Survival}}},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	jdmt, err := c.Report(context.TODO(), &xyron.PlayerReport{
		Player: p,
		Data: map[int32]*xyron.TimestampedReportData{
			2: {
				Timestamp: 2,
				Data: []*xyron.WildcardReportData{
					{Data: &xyron.WildcardReportData_ActionData{ActionData: &xyron.PlayerActionData{
						Position: &xyron.EntityPositionData{
							Location: &xyron.Loc3F{
								Position:  &xyron.Vec3F{},
								Direction: &xyron.Vec3F{},
							},
							BoundingBox: &xyron.AxisAlignedBoundingBox{
								Min: &xyron.Vec3F{},
								Max: &xyron.Vec3F{X: 1, Y: 2, Z: 1},
							},
							Below:             nil,
							IsImmobile:        false,
							IsOnGround:        false,
							AllowFlying:       false,
							IsFlying:          false,
							CollidedBlocks:    nil,
							IntersectedBlocks: nil,
						},
						Action: xyron.PlayerAction_Jump,
					}}},
				},
			},
		},
	})
	for _, j := range jdmt.Judgements {
		fmt.Printf("judgement: %v: %v message:%v\n", j.Type, j.Judgement.String(), j.Message)
	}
	_, err = c.RemovePlayer(context.TODO(), p)
	if err != nil {
		panic(err)
	}
	conn.Close()
}
