package g

import (
	"sync"
	"time"
)

type SafeStrategyMap struct {
	sync.RWMutex
	// endpoint:metric => [strategy1, strategy2 ...]
	M map[string][]Strategy
}

type SafeEventMap struct {
	sync.RWMutex
	M map[string]*Event
}

var (
	HbsClient     *SingleConnRpcClient
	StrategyMap   = &SafeStrategyMap{M: make(map[string][]Strategy)}
	LastEvents    = &SafeEventMap{M: make(map[string]*Event)}
)

func InitHbsClient() {
	HbsClient = &SingleConnRpcClient{
		RpcServers: Config().Hbs.Servers,
		Timeout:    time.Duration(Config().Hbs.Timeout) * time.Millisecond,
	}
}

func (this *SafeStrategyMap) ReInit(m map[string][]Strategy) {
	this.Lock()
	defer this.Unlock()
	this.M = m
}

func (this *SafeStrategyMap) Get() map[string][]Strategy {
	this.RLock()
	defer this.RUnlock()
	return this.M
}

func (this *SafeEventMap) Get(key string) (*Event, bool) {
	this.RLock()
	defer this.RUnlock()
	event, exists := this.M[key]
	return event, exists
}

func (this *SafeEventMap) Set(key string, event *Event) {
	this.Lock()
	defer this.Unlock()
	this.M[key] = event
}
