package main

import "context"

func PlaceOrder(ctx context.Context) (string, error) {

}

func DeclineOrder(ctx context.Context, orderId string) (string, error) {

}
func ReserveItems(ctx context.Context, consumerId string, orderId string, items []Item) (string, error) {

}

func CancelReservation(ctx context.Context, reservationId int64) (string, error) {

}

func ReserveDelivery(ctx context.Context, orderId string, delivery DeliveryDetail) (string, error) {

}

func CancelDelivery(ctx context.Context, deliveryId string) (string, error) {

}

func MakePayment(ctx context.Context, consumerId string, orderId string, items []Item) (string, error) {

}

func ConfirmOrder(ctx context.Context, orderId string, delivery int64, paymentId int64, reservationId int64) (string, error) {

}
