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

var ErrorAdd = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "error_add",
	Help:        "error_add",
	ConstLabels: prometheus.Labels{},
})

var ErrorGet = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "error_get",
	Help:        "error_get",
	ConstLabels: prometheus.Labels{},
})

var ErrorPut = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "error_put",
	Help:        "error_put",
	ConstLabels: prometheus.Labels{},
})

var ErrorDelete = prometheus.NewCounter(prometheus.CounterOpts{
	Name:        "error_delete",
	Help:        "error_delete",
	ConstLabels: prometheus.Labels{},
})

var LatencyAdd = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "latency_add",
	Help: "latency_add",
})

var LatencyGet = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "latency_get",
	Help: "latency_get",
})

var LatencyPut = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "latency_put",
	Help: "latency_put",
})

var LatencyDelete = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "latency_delete",
	Help: "latency_delete",
})
