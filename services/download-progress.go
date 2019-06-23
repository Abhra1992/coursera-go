package services

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type DownloadProgress struct {
	total    float64
	current  float64
	start    time.Time
	now      time.Time
	finished bool
}

func (dp DownloadProgress) Start() {
	dp.now = time.Now()
	dp.start = dp.now
}

func (dp DownloadProgress) Stop() {
	dp.now = time.Now()
	dp.finished = true
	dp.total = dp.current
	dp.ReportProgress()
}

func (dp DownloadProgress) Read(bytes float64) {
	dp.now = time.Now()
	dp.current += bytes
	dp.ReportProgress()
}

func (dp DownloadProgress) Report(bytes float64) {
	dp.now = time.Now()
	dp.current += bytes
	dp.ReportProgress()
}

func (dp DownloadProgress) CalcPercent() string {
	if dp.total < 0 {
		return "--%"
	}
	if dp.total == 0 {
		return "100% done"
	}
	perc := int(float64(dp.current) / float64(dp.total) * 100.0)
	marks := int(perc / 2)
	display := fmt.Sprintf("[%-50s] %d%%", strings.Repeat("#", marks), perc)
	return display
}

func (dp DownloadProgress) CalcSpeed() string {
	diff := float64(dp.now.Sub(dp.start) / time.Second)
	if dp.current == 0 || diff < 0.001 {
		return "---b/s"
	}
	fbytes := formatBytes(float64(dp.current) / diff)
	return fmt.Sprintf("%s/s", fbytes)
}

func (dp DownloadProgress) ReportProgress() {
	percent := dp.CalcPercent()
	total := formatBytes(dp.total)
	speed := dp.CalcSpeed()
	tsr := fmt.Sprintf("%s at %s", total, speed)
	report := fmt.Sprintf("\r%-56s %30s", percent, tsr)
	if dp.finished {
		fmt.Println(report)
	} else {
		fmt.Print(report)
	}
}

func formatBytes(bytes float64) string {
	if bytes < 0 {
		return "N/A"
	}
	var exponent float64
	if bytes == 0 {
		exponent = 0
	} else {
		exponent = math.Floor(math.Log2(bytes) / 10)
	}
	suffixes := [...]string{"B", "KB", "MB", "GB", "TB"}
	suffix := suffixes[int(exponent)]
	converted := bytes / float64(math.Pow(1024, exponent))
	return fmt.Sprintf("%.2f%s", converted, suffix)
}
