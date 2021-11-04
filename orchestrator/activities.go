package main

import (
	"context"
)

func PlaceOrder(ctx context.Context) {

}

func DeclineOrder(ctx context.Context, orderId string) {

}
func ReserveItems(ctx context.Context, consumerId string, orderId string, items []Item) {

}

func CancelReservation(ctx context.Context, reservationId int64) {

}

func ReserveDelivery(ctx context.Context, orderId string, delivery DeliveryDetail) {

}

func CancelDelivery(ctx context.Context, deliveryId string) {

}

func MakePayment(ctx context.Context, consumerId string, orderId string, items []Item) {

}

func ConfirmOrder(ctx context.Context, orderId string, delivery int64, paymentId int64, reservationId int64) {

}
