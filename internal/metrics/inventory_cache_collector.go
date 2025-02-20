package metrics

import (
	"github.com/openshift/assisted-service/internal/common"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	gaugeInventoryCacheEntriesTotal = "assisted_installer_inventory_cache_entries_total"
	counterInventoryCacheHits       = "assisted_installer_inventory_cache_hit_count"
	counterInventoryCacheMisses     = "assisted_installer_inventory_cache_miss_count"

	descGaugeInventoryCacheEntriesTotal = "Total number of entries in the inventory cache"
	descCounterInventoryCacheHits       = "Number of hits in inventory cache"
	descCounterInventoryCacheMisses     = "Number of misses in inventory cache"
)

type inventoryCacheCollector struct {
	descGaugeInventoryCacheEntriesTotal *prometheus.Desc
	descCounterInventoryCacheHits       *prometheus.Desc
	descCounterInventoryCacheMisses     *prometheus.Desc
}

func newInventoryCacheCollector() *inventoryCacheCollector {
	return &inventoryCacheCollector{
		descGaugeInventoryCacheEntriesTotal: prometheus.NewDesc(gaugeInventoryCacheEntriesTotal, descGaugeInventoryCacheEntriesTotal, []string{}, nil),
		descCounterInventoryCacheHits:       prometheus.NewDesc(counterInventoryCacheHits, descCounterInventoryCacheHits, []string{}, nil),
		descCounterInventoryCacheMisses:     prometheus.NewDesc(counterInventoryCacheMisses, descCounterInventoryCacheMisses, []string{}, nil),
	}
}

func (c *inventoryCacheCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.descGaugeInventoryCacheEntriesTotal
	ch <- c.descCounterInventoryCacheHits
	ch <- c.descCounterInventoryCacheMisses
}

func (c *inventoryCacheCollector) Collect(ch chan<- prometheus.Metric) {
	if common.InventoryCache == nil {
		return
	}
	ch <- prometheus.MustNewConstMetric(c.descGaugeInventoryCacheEntriesTotal, prometheus.GaugeValue, float64(common.InventoryCache.Len()))
	ch <- prometheus.MustNewConstMetric(c.descCounterInventoryCacheHits, prometheus.CounterValue, float64(common.InventoryCache.Metrics().Hits))
	ch <- prometheus.MustNewConstMetric(c.descCounterInventoryCacheMisses, prometheus.CounterValue, float64(common.InventoryCache.Metrics().Misses))
}
