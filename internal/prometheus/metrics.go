package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    UploadsCounter = promauto.NewCounter(prometheus.CounterOpts{
        Name: "upload_handler_total",
        Help: "The total number of processed uploads",
    })

    FailedUploadsCounter = promauto.NewCounter(prometheus.CounterOpts{
        Name: "upload_handler_failed_total",
        Help: "The total number of failed uploads",
    })

    UploadDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name:    "upload_handler_duration_seconds",
        Help:    "Upload duration in seconds",
        Buckets: prometheus.DefBuckets,
    })
)