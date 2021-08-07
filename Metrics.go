package main

import "github.com/prometheus/client_golang/prometheus"

var RequestCountAdd = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "request_count_add",
	Help:        "request_count_add",
	ConstLabels: prometheus.Labels{},
})

var RequestCountGet = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "request_count_get",
	Help:        "request_count_get",
	ConstLabels: prometheus.Labels{},
})

var RequestCountPut = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "request_count_put",
	Help:        "request_count_put",
	ConstLabels: prometheus.Labels{},
})

var RequestCountDelete = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "request_count_delete",
	Help:        "request_count_delete",
	ConstLabels: prometheus.Labels{},
})
