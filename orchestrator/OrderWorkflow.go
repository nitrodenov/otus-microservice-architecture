package main

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
	"time"
)

type Item struct {
	itemId      int64
	title       string
	description string
	quantity    int
	price       int64
}

type DeliveryDetail struct {
	deliveryType string
	city         string
	date         time.Time
}

type CreateOrder struct {
	orderId        uuid.UUID
	items          []Item
	deliveryDetail DeliveryDetail
	price          int64
	consumerId     uuid.UUID
}

func Workflow(ctx workflow.Context, order CreateOrder) (string, error) {
	ctx = New(ctx, SagaOptions{
		ParallelCompensation: false,
		ContinueWithError:    false,
	})

	// make order
	var orderId string
	err := workflow.ExecuteActivity(ctx, PlaceOrder).Get(ctx, &orderId)
	if err != nil {
		// If the placing order failed, fail the Workflow.
		return "", err
	}
	AddCompensation(ctx, func(ctx workflow.Context) error {
		cancelErr := workflow.ExecuteActivity(ctx, DeclineOrder, orderId)
		if cancelErr != nil {

		}
		return nil
	})

	// reserve
	var reservationId string
	err = workflow.ExecuteActivity(ctx, ReserveItems, order.consumerId, orderId, order.items).Get(ctx, &reservationId)
	if err != nil {
		// If the placing order failed, fail the Workflow.
		return "", err
	}
	AddCompensation(ctx, func(ctx workflow.Context) error {
		cancelErr := workflow.ExecuteActivity(ctx, CancelReservation, reservationId)
		if cancelErr != nil {

		}
		return nil
	})

	// delivery
	var deliveryId string
	err = workflow.ExecuteActivity(ctx, ReserveDelivery, orderId, order.deliveryDetail).Get(ctx, &deliveryId)
	if err != nil {
		// If the placing order failed, fail the Workflow.
		return "", err
	}
	AddCompensation(ctx, func(ctx workflow.Context) error {
		cancelErr := workflow.ExecuteActivity(ctx, CancelDelivery, deliveryId)
		if cancelErr != nil {

		}
		return nil
	})

	// payment
	var paymentId string
	err = workflow.ExecuteActivity(ctx, MakePayment, order.consumerId, orderId, order.items).Get(ctx, &paymentId)
	if err != nil {
		err = Compensate(ctx)
	}
	// confirm
	err = workflow.ExecuteActivity(ctx, ConfirmOrder, orderId, order.deliveryDetail, paymentId, order.items, reservationId).Get(ctx, &paymentId)

	if err != nil {
		// If the payment failed, fail the Workflow.
		return "", err
	}

	return "", nil
}
