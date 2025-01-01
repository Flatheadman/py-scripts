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
		bestClient := m.BestClient().(*LocalClient) // Type assertion needed
		bestClient.Mutex.Lock()
		fmt.Printf("Best Client Index: %d, Value: %d\n", m.BestIndex.Load(), bestClient.CurrentValue)
		bestClient.Mutex.Unlock()
		time.Sleep(2 * time.Second)
	}
}
