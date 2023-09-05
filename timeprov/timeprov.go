package main

import (
	"fmt"
	"time"
)

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
	return m.MockTime
}

type Subscription struct {
	BillingDate time.Time
}

type Customer struct {
	Subscription Subscription
}

type BillingSystem struct {
	TimeProvider TimeProvider
}

func (bs BillingSystem) CalculateBill(c Customer) string {
	currentTime := bs.TimeProvider.TimeNow()
	billingDate := c.Subscription.BillingDate
	fs := fmt.Sprintf("You have been billed on %s. Your next bill will be on %s.", currentTime, billingDate)
	return fs
}

func main() {

	mockCustomer := Customer{
		Subscription: Subscription{
			BillingDate: time.Now().AddDate(0, 1, 0),
		},
	}

	mockTime := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	mockTimeProvider := MockTimeProvider{MockTime: mockTime}

	mockBillingSystem := BillingSystem{TimeProvider: mockTimeProvider}
	mockBill := mockBillingSystem.CalculateBill(mockCustomer)
	fmt.Println(mockBill)
}
