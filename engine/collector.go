package engine

import (
	"fmt"
	client2 "github.com/influxdata/influxdb1-client/v2"
)

type Collector struct {
	client client2.Client
	query client2.Query
}

func NewCollector() (Collector, error) {
	c, err := client2.NewHTTPClient(client2.HTTPConfig{
		Addr:               "http://192.168.107.160:8086",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	q := client2.NewQuery("select last(*) from aggregate_stats", "sdn","")
	collector := Collector{
		client:c,
		query:q,
	}


	return collector, err
}

func (collector Collector) Collect() {
	if response, err := collector.client.Query(collector.query); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}
