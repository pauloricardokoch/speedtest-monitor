package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/pauloricardokoch/speedtest-monitor/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type Metrics struct {
	Isp        string  `json:"isp"`
	Host       string  `json:"host"`
	Ip         string  `json:"ip"`
	Location   string  `json:"location"`
	Country    string  `json:"country"`
	IntIp      string  `json:"int-ip"`
	IntName    string  `json:"int-name"`
	IntMacAddr string  `json:"int-mac-addr"`
	IntIsVpn   bool    `json:"int-is-vpn"`
	DBytes     float64 `json:"d-bytes"`
	DElapsed   float64 `json:"d-elapsed"`
	UBytes     float64 `json:"u-bytes"`
	UElapsed   float64 `json:"u-elapsed"`
	Latency    float64 `json:"latency"`
}

var (
	labels = []string{
		"isp",
		"host",
		"ip",
		"location",
		"country",
		"intIp",
		"intName",
		"intMacAddr",
		"intIsVpn",
	}
	SpeedBuckets  = []float64{0, 25, 50, 75, 100, 125, 150, 175, 200, 225, 250, 275, 300, 325, 350, 375, 400, 425, 450, 475, 500}
	downloadSpeed = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "speedtest_download_speed",
			Buckets: SpeedBuckets,
			Help:    "Current download speed",
		},
		labels)
	uploadSpeed = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "speedtest_upload_speed",
			Buckets: SpeedBuckets,
			Help:    "Current upload speed",
		},
		labels)
	latency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "speedtest_latency_speed",
			Buckets: []float64{0, 2, 4, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50},
			Help:    "Current upload speed",
		},
		labels)
)

func register(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.Method == http.MethodPost {
		var m Metrics
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		downloadSpeed.WithLabelValues(
			m.Isp,
			m.Host,
			m.Ip,
			m.Location,
			m.Country,
			m.IntIp,
			m.IntName,
			m.IntMacAddr,
			strconv.FormatBool(m.IntIsVpn),
		).Observe(m.DBytes / m.DElapsed * 8 / 1000)

		uploadSpeed.WithLabelValues(
			m.Isp,
			m.Host,
			m.Ip,
			m.Location,
			m.Country,
			m.IntIp,
			m.IntName,
			m.IntMacAddr,
			strconv.FormatBool(m.IntIsVpn),
		).Observe(m.UBytes / m.UElapsed * 8 / 1000)

		latency.WithLabelValues(
			m.Isp,
			m.Host,
			m.Ip,
			m.Location,
			m.Country,
			m.IntIp,
			m.IntName,
			m.IntMacAddr,
			strconv.FormatBool(m.IntIsVpn),
		).Observe(m.Latency)
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	port := viper.GetString("Port")

	fmt.Printf("Server is running on port %s...", port)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/register", register)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
