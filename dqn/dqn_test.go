package dqn

import (
	"testing"
	"time"

	"github.com/sinmetal/slog"
)

// DQNDummyClient is UnitTestのためのDQN Dummy実装
type DQNDummyClient struct {
	PredictionCount int
	Body            *Payload
	DummyAnswer     *Answer
}

func (client *DQNDummyClient) Prediction(log *slog.Log, body *Payload) (*Answer, error) {
	client.PredictionCount++
	client.Body = body
	return client.DummyAnswer, nil
}

func TestSetDummyClient(t *testing.T) {
	da := &Answer{
		X:      -1,
		Y:      0,
		Paused: false,
		Angle:  AngleRight,
		Speed:  speed,
	}
	SetDummyClient(&DQNDummyClient{
		DummyAnswer: da,
	})

	client := NewClient()

	payload := &Payload{
		Instances: []Instance{
			Instance{},
		},
	}

	// Playerが常に右隣にいる状態
	payload.Instances[0].State[(SenseRangeRow / 2)][(SenseRangeCol / 2)][0] = 1

	// DQNが中心ぐらいにいる状態
	payload.Instances[0].State[(SenseRangeRow / 2)][(SenseRangeCol / 2)][1] = 1

	slog := slog.Start(time.Now())
	a, err := client.Prediction(&slog, payload)
	if err != nil {
		t.Fatalf("failed Prediction. err = %+v", err)
	}
	if e, g := da.X, a.X; e != g {
		t.Fatalf("expected X is %f; got %f", e, g)
	}
	if e, g := da.Y, a.Y; e != g {
		t.Fatalf("expected Y is %f; got %f", e, g)
	}
	if e, g := da.Paused, a.Paused; e != g {
		t.Fatalf("expected Paused is %t; got %t", e, g)
	}
	if e, g := da.Angle, a.Angle; e != g {
		t.Fatalf("expected Angle is %f; got %f", e, g)
	}
	if e, g := da.Speed, a.Speed; e != g {
		t.Fatalf("expected Speed is %f; got %f", e, g)
	}
}