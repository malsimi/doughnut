package main

import "testing"

func TestGetOrders(t *testing.T) {

	want := []order{
		{ID: 1, Quantity: 3, Total: 30},
		{ID: 2, Quantity: 4, Total: 40},
		{ID: 3, Quantity: 5, Total: 50},
	}

	assertion := func(got []order, want []order, t *testing.T) {
		t.Helper()
		isFailed := false
		for i, value := range got {
			if len(want) > i {
				wantedOrder := want[i]
				if value.ID != wantedOrder.ID || value.Quantity != wantedOrder.Quantity || value.Total != wantedOrder.Total {
					isFailed = true
					break
				}
			} else {
				isFailed = true
				break
			}
		}
		if isFailed {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("Get the list of available orders", func(t *testing.T) {
		got := getOrders()
		assertion(got, want, t)
	})

	t.Run("Get order by id", func(t *testing.T) {
		sampleOrder := want[2]
		got, err := getOrderById(sampleOrder.ID)
		if err != nil {
			t.Errorf("got %q want %q", got, want)
		}

		if got.ID != sampleOrder.ID || got.Quantity != sampleOrder.Quantity || got.Total != sampleOrder.Total {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func TestAddOrder(t *testing.T) {

	newOrder := order{
		ID:       7,
		Quantity: 7,
		Total:    70,
	}

	t.Run("add a valid order", func(t *testing.T) {
		isSuccess, err := addOrder(newOrder)

		if isSuccess == false || err != nil {
			t.Errorf("got %v want %v", isSuccess, true)
		}
	})

	t.Run("check if the newOrder is available", func(t *testing.T) {
		got := getOrders()

		retrievedNewOrder := got[3]

		if retrievedNewOrder.ID != newOrder.ID || retrievedNewOrder.Quantity != newOrder.Quantity || retrievedNewOrder.Total != newOrder.Total {
			t.Errorf("got %q want %q", got, newOrder)
		}
	})

}
