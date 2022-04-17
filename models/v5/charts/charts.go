package charts

import (
	"time"

	v5 "github.com/grokify/elastirad-go/models/v5"
	"github.com/grokify/gocharts/v2/charts/c3"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/gocharts/v2/data/timeseries/interval"
	"github.com/grokify/mogo/time/timeutil"
)

func C3ChartForEsAggregationSimple(agg v5.AggregationResRad) c3.C3Chart {
	c3Chart := c3.C3Chart{
		Data: c3.C3ChartData{
			Columns: [][]interface{}{},
		},
	}
	for _, bucket := range agg.AggregationData.Buckets {
		c3Column := []interface{}{bucket.Key, bucket.DocCount}
		c3Chart.Data.Columns = append(c3Chart.Data.Columns, c3Column)
	}
	return c3Chart
}

func EsAggsToTimeSeriesSet(aggs []v5.AggregationResRad, timeInterval timeutil.Interval, weekStart time.Weekday) (interval.TimeSeriesSet, error) {
	set := interval.NewTimeSeriesSet(timeInterval, weekStart)

	for _, agg := range aggs {
		seriesName := agg.AggregationName
		for _, bucket := range agg.AggregationData.Buckets {
			set.AddItem(timeseries.TimeItem{
				SeriesName: seriesName,
				Time:       timeutil.UnixMillis(int64(bucket.Key.(float64))),
				Value:      int64(bucket.DocCount)})
		}
	}
	err := set.Inflate()
	return set, err
}
