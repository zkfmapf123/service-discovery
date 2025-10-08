package dispatcher

import (
	"context"
	"log"
	"time"
)

type QueueEvents struct {
	Type    string         `json:"type"`
	Data    map[string]any `json:"data"`
	Error   string         `json:"error"`
	Created int64          `json:"created"`
}

type Queue struct {
	ch          chan QueueEvents
	ChannelSize int
}

func NewQueue(channelSize int) *Queue {
	return &Queue{
		ch:          make(chan QueueEvents, channelSize),
		ChannelSize: channelSize,
	}
}

func (q *Queue) Enqueue(data QueueEvents) {
	q.ch <- data
}

func (q *Queue) Size() int {
	return len(q.ch)
}

func (q *Queue) Dispatcher(ctx context.Context) {

	for {
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)

		select {
		case res := <-q.ch:

			switch res.Type {

			case "log":

			case "func":

			case "debug":

			default:
				log.Printf("Current Queue Size : %d , MaxSize : %v", q.Size(), q.ChannelSize)
			}

		case <-timeoutCtx.Done():
			cancel()

			if ctx.Err() != nil {
				log.Println("Context Is Timeout ...")
			}
		}
	}
}
