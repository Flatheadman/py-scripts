package multiclient

import (
	"math/rand"
	"sync"
	"time"

	"go.uber.org/atomic"
)


type Client interface {
	GetLatestValue() (int64, error)
}


type MultiClient struct {
	clients   []Client
	//将变量声明为原子类型的整型
	BestIndex atomic.Int32 
	//如果这里变量名首字母小写，在其他的package中将不能直接访问。
}


func New(clients []Client) *MultiClient {
	m := &MultiClient{
		clients: clients,
	}
	if len(clients) > 1 {
		go m.sniffLoop()
	}
	return m  //这里不会导致上面的子协程被迫终止，因为本函数不是主协程，只有main函数是主协程。
}


func (m *MultiClient) BestClient() Client {
	//Load函数:原子地读取变量值。
	return m.clients[m.BestIndex.Load()]
}

func (m *MultiClient) sniffLoop() {
	t := time.NewTimer(0)
	// Select模块的任意分支命中后，他的执行就结束了，所以需要加for循环。
	for {
		select {
		case <-t.C:
			m.sniff()
			t.Reset(time.Second)//将心跳间隔重置为一秒。
		}
	}
}

// sniff determines the best performing client based on value and latency
func (m *MultiClient) sniff() {
	var (
		values = make([]int64, len(m.clients))
		times  = make([]int64, len(m.clients))
		l      sync.Mutex
		wg     sync.WaitGroup
	)
	wg.Add(len(m.clients))
	for i, client := range m.clients {
		i, client := i, client
		go func() {
			defer wg.Done()
			start := time.Now().UnixNano()
			value, _ := client.GetLatestValue()
			l.Lock()
			values[i] = value
			times[i] = time.Now().UnixNano() - start
			l.Unlock()
		}()
	}
	wg.Wait()

	var (
		maxValue   = values[0]
		minTime    = times[0]
		bestClient = 0
	)
	for i := 1; i < len(m.clients); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
			minTime = times[i]
			bestClient = i
		} else if values[i] == maxValue {
			if times[i] < minTime {
				minTime = times[i]
				bestClient = i
			}
		}
	}
	m.BestIndex.Store(int32(bestClient))
}

// LocalClient simulates an external client with random values and latency
type LocalClient struct {
	CurrentValue int64
	Mutex        sync.Mutex
}

// NewLocalClient creates a new LocalClient
func NewLocalClient(initialValue int64) *LocalClient {
	return &LocalClient{
		CurrentValue: initialValue,
	}
}

// GetLatestValue simulates getting the latest value from an external source
func (c *LocalClient) GetLatestValue() (int64, error) {
	// Simulate network latency
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// Simulate value updates
	c.CurrentValue += int64(rand.Intn(10) - 3) // Increase or decrease randomly
	if c.CurrentValue < 0 {
		c.CurrentValue = 0
	}

	return c.CurrentValue, nil
}
