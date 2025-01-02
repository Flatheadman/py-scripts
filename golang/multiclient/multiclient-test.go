package multiclient

import (
	"fmt"
	"time"
)

func Test() {
	// Create multiple LocalClient instances
	client1 := NewLocalClient(100)
	client2 := NewLocalClient(101)
	client3 := NewLocalClient(102)

	// Create a MultiClient with the LocalClients
	m := New([]Client{client1, client2, client3})

	// Run for a while to observe the behavior
	for i := 0; i < 10; i++ {
		//因为*LocalClient实现了Client函数的GetLatestValue方法
		//所以断言成功，并返回了断言类型的变量。参照以下函数实现:
		//func (c *LocalClient) GetLatestValue() (int64, error) 
		bestClient := m.BestClient().(*LocalClient) //断言加类型转换
		bestClient.Mutex.Lock()
		fmt.Printf("Best Client Index: %d, Value: %d\n", m.BestIndex.Load(), bestClient.CurrentValue)
		bestClient.Mutex.Unlock()
		time.Sleep(2 * time.Second)
	}
}
