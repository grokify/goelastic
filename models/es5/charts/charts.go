package charts

import (
	"time"

	"github.com/grokify/gocharts/v2/charts/c3"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/gocharts/v2/data/timeseries/interval"
	"github.com/grokify/mogo/time/timeutil"

	"github.com/grokify/goelastic/models/es5"
)

func C3ChartForEsAggregationSimple(agg es5.AggregationResRad) c3.C3Chart {
	c3Chart := c3.C3Chart{
		Data: c3.C3ChartData{
			Columns: [][]interface{}{},
		},
	}
	for _, bucket := range agg.AggregationData.Buckets {
		c3Chart.Data.Columns = append(
			c3Chart.Data.Columns,
			[]interface{}{bucket.Key, bucket.DocCount}, // c3Column := []interface{}{bucket.Key, bucket.DocCount}
		)
	}
	return c3Chart
}

func EsAggsToTimeSeriesSet(aggs []es5.AggregationResRad, timeInterval timeutil.Interval, weekStart time.Weekday) (interval.TimeSeriesSet, error) {
	set := interval.NewTimeSeriesSet(timeInterval, weekStart)

	for _, agg := range aggs {
		seriesName := agg.AggregationName
		for _, bucket := range agg.AggregationData.Buckets {
			set.AddItem(timeseries.TimeItem{
				SeriesName: seriesName,
				Time:       time.UnixMilli(int64(bucket.Key.(float64))),
				Value:      int64(bucket.DocCount)})
		}
	}
	err := set.Inflate()
	return set, err
}
