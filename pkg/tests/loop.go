package tests

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
	"vs/pkg/logger"
)

// Status 상태
type Status struct {
	Success int
	Failed  int
	mu      sync.Mutex
}

// IncrementSuccess 성공 카운트 증가
func (s *Status) IncrementSuccess() {
	s.mu.Lock()
	s.Success++
	s.mu.Unlock()
}

// IncrementFailed 실패 카운트 증가
func (s *Status) IncrementFailed() {
	s.mu.Lock()
	s.Failed++
	s.mu.Unlock()
}

// LoopFunction 반복
func LoopFunction(arg interface{}, loopCnt, rps int, f func(arg interface{}) error) *Status {
	rpsCnt := int64(0)
	maxRpsCnt := int64(0)
	status := &Status{}
	wg := sync.WaitGroup{}
	// 초당 호출 수 적용
	var ticker *time.Ticker
	if rps != 0 {
		ticker = time.NewTicker(time.Second / time.Duration(rps))
	}
	go func() {
		for {
			<-time.Tick(1 * time.Second)
			maxRpsCnt = atomic.LoadInt64(&rpsCnt)
			atomic.StoreInt64(&rpsCnt, 0)
		}
	}()
	for i := 0; i < loopCnt; i++ {
		if ticker != nil {
			<-ticker.C
		}
		wg.Add(1)
		atomic.AddInt64(&rpsCnt, 1)
		logger.ProgressBar(i, loopCnt, &maxRpsCnt)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Fatalln(fmt.Sprintf("고루틴 동작간 오류 발생\n 오류 내용 : %v", err))
				}
			}()

			if err := f(arg); err != nil {
				status.IncrementFailed()
			} else {
				status.IncrementSuccess()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return status
}
