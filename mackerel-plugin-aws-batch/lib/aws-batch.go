package mpawsbatch

import (
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
	mp "github.com/mackerelio/go-mackerel-plugin"
)


type BatchPlugin struct {
	Prefix string
}


func (p *) MetricKeyPrefix() string {
	if p.Prefix == "" {
		p.Prefix = "sample"
	}
	return p.Prefix
}


func (p *SamplePlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.Prefix)
	return map[string]mp.Graphs{
		"dice": {
			Label: labelPrefix + " Dice Value",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "d6", Label: "Dice(d6)"},
				{Name: "d20", Label: "Dice(d20)"},
			},
		},
	}
}


func (p *SamplePlugin) FetchMetrics() (map[string]float64, error) {
	rand.Seed(time.Now().UnixNano())
	metrics := map[string]float64{
		"d6":  float64(rand.Intn(6) + 1),
		"d20": float64(rand.Intn(20) + 1),
	}
	return metrics, nil
}


func Do() {
	optPrefix := flag.String("metric-key-prefix", "", "Metric key prefix")
	flag.Parse()

	plugin := mp.NewMackerelPlugin(&SamplePlugin{
		Prefix: *optPrefix,
	})
	plugin.Run()
}
