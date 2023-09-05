package main

import (
	"fmt"
	"time"
)

// Both RealTimeProvider and MockTimeProvider have a TimeNow method
// If a struct has a method with the same name as an interface, it implements that interface
type TimeProvider interface {
	TimeNow() time.Time
}

type RealTimeProvider struct {
	RealTime time.Time
}

type MockTimeProvider struct {
	MockTime time.Time
}

func (r RealTimeProvider) TimeNow() time.Time {
	r.RealTime = time.Now()
	return r.RealTime
}

func (m MockTimeProvider) TimeNow() time.Time {
	// m.MockTime = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	return m.MockTime
}

type Subscription struct { // Subscription has a billing date
	BillingDate time.Time
}

type Customer struct { // Customer has a subscription
	Subscription Subscription
}

type BillingSystem struct {
	TimeProvider TimeProvider
}

func (bs BillingSystem) CalculateBill() string {
	currentTime := bs.TimeProvider.TimeNow()

	return "You have been billed on " + currentTime.Format(time.Kitchen)
}

func main() {
	realTimeProvider := RealTimeProvider{}
	formatTime := realTimeProvider.TimeNow().Format(time.Kitchen)
	fmt.Println("whatTimeIsIt", formatTime)

	mockTime := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)  // Create a mock time
	mockTimeProvider := MockTimeProvider{MockTime: mockTime} // Pass the mock time into the mock time provider
	formatTime = mockTimeProvider.TimeNow().Format(time.Kitchen)
	fmt.Println("whatMockTimeIsIt", formatTime)

	times := []TimeProvider{RealTimeProvider{}, MockTimeProvider{}}
	for index, timeProvider := range times {
		formatTime = timeProvider.TimeNow().Format(time.Kitchen)
		fmt.Println("The time at", index, "is", formatTime)
	}

	fmt.Println(BillingSystem{TimeProvider: RealTimeProvider{}}.CalculateBill())

	// Create an mock customer with a subscription
	customer := Customer{Subscription: Subscription{BillingDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)}}
	fmt.Println("Customer billing date", customer.Subscription.BillingDate.Format(time.Kitchen))
}
