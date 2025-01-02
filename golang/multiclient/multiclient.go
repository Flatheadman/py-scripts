package multiclient

import (
	"math/rand"
	"sync"
	"time"

	"go.uber.org/atomic"//主要为整型数据提供原子操作。
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

func (m *MultiClient) sniff() {
	var (
		values = make([]int64, len(m.clients))//创建当前长度和容量相同的切片数组。
		times  = make([]int64, len(m.clients))
		l      sync.Mutex //互斥锁。
		wg     sync.WaitGroup //等待组。
	)
	wg.Add(len(m.clients))
	for i, client := range m.clients {
		i, client := i, client
		go func() {
			defer wg.Done()
			start := time.Now().UnixNano()
			value, _ := client.GetLatestValue()
			l.Lock()
			//切片数组的本质是一个结构体，虽然访问不同的切片数据索引，但是都是在读取这个结构体。
			//这个结构体的内部运作是复杂的，不是1+1=2，所以需要互斥锁来保护。
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

// 实际上的客户端，会继承client接口和其中定义的函数。
type LocalClient struct {
	CurrentValue int64 //本变量天生被多线程修改，标配一把锁。
	Mutex        sync.Mutex
}


func NewLocalClient(initialValue int64) *LocalClient {
	return &LocalClient{
		CurrentValue: initialValue,
	}
}

// 本方法继承自client接口。
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
