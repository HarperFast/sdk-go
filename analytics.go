package harper

import (
	"time"
)

type GetAnalyticsRequest struct {
	Metric        string           `json:"metric"`
	GetAttributes AttributeList    `json:"get_attributes"`
	StartTime     int64            `json:"start_time"`
	EndTime       int64            `json:"end_time"`
	Conditions    SearchConditions `json:"conditions"`
}

type GetAnalyticsResult map[string]interface{}

func (c *Client) GetAnalytics(req GetAnalyticsRequest) ([]GetAnalyticsResult, error) {
	op := operation{
		Operation:     OP_GET_ANALYTICS,
		Metric:        req.Metric,
		GetAttributes: req.GetAttributes,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Conditions:    req.Conditions,
	}

	results := make([]GetAnalyticsResult, 0)

	err := c.opRequest(op, &results)
	if err != nil {
		return nil, err
	}

	for i, r := range results {
		if r["id"] != nil {
			timeMillis := r["id"].(float64)
			idTime := time.UnixMilli(int64(timeMillis))
			results[i]["id"] = idTime
		}
	}

	return results, nil
}

type ListMetricsResult string

type MetricType int

const (
	MetricTypeCustom MetricType = iota
	MetricTypeBuiltin
)

var metricTypeName = map[MetricType]string{
	MetricTypeCustom:  "custom",
	MetricTypeBuiltin: "builtin",
}

type ListMetricsRequest struct {
	MetricTypes         []MetricType `json:"metric_types"`
	CustomMetricsWindow int64        `json:"custom_metrics_window"`
}

func (c *Client) ListMetrics(req ListMetricsRequest) ([]ListMetricsResult, error) {
	typeNames := make([]string, 0, len(req.MetricTypes))
	for _, typ := range req.MetricTypes {
		typeNames = append(typeNames, metricTypeName[typ])
	}

	op := operation{
		Operation:   OP_LIST_METRICS,
		MetricTypes: typeNames,
	}

	if req.CustomMetricsWindow != 0 {
		op.CustomMetricsWindow = req.CustomMetricsWindow
	}

	results := make([]ListMetricsResult, 0)

	err := c.opRequest(op, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

type AttributeDesc struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DescribeMetricResult struct {
	Attributes []AttributeDesc `json:"attributes"`
}

func (c *Client) DescribeMetric(metric string) (*DescribeMetricResult, error) {
	op := operation{
		Operation: OP_DESCRIBE_METRIC,
		Metric:    metric,
	}

	var result DescribeMetricResult
	err := c.opRequest(op, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
