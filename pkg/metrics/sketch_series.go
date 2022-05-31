// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package metrics

import (
	"bytes"
	"encoding/json"

	"github.com/DataDog/datadog-agent/pkg/aggregator/ckey"
	"github.com/DataDog/datadog-agent/pkg/quantile"
	"github.com/DataDog/datadog-agent/pkg/tagset"
)

// A SketchSeries is a timeseries of quantile sketches.
type SketchSeries struct {
	Name       string               `json:"metric"`
	Tags       tagset.CompositeTags `json:"tags"`
	Host       string               `json:"host"`
	Interval   int64                `json:"interval"`
	Points     []SketchPoint        `json:"points"`
	ContextKey ckey.ContextKey      `json:"-"`
}

// A SketchPoint represents a quantile sketch at a specific time
type SketchPoint struct {
	Sketch *quantile.Sketch `json:"sketch"`
	Ts     int64            `json:"ts"`
}

// SketchSeriesList is a collection of SketchSeries
type SketchSeriesList []*SketchSeries

// MarshalJSON serializes sketch series to JSON.
func (sl SketchSeriesList) MarshalJSON() ([]byte, error) {
	type SketchSeriesAlias SketchSeriesList
	data := map[string]SketchSeriesAlias{
		"sketches": SketchSeriesAlias(sl),
	}
	reqBody := &bytes.Buffer{}
	err := json.NewEncoder(reqBody).Encode(data)
	return reqBody.Bytes(), err
}

// String returns the JSON representation of a SketchSeriesList as a string
// or an empty string in case of an error
func (sl SketchSeriesList) String() string {
	json, err := sl.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(json)
}

// SketchesSink is a sink for sketches.
// It provides a way to append a sketches into `SketchSeriesList`
type SketchesSink interface {
	Append(*SketchSeries)
}

var _ SketchesSink = (*SketchSeriesList)(nil)

// Append appends a sketches.
func (sl *SketchSeriesList) Append(sketches *SketchSeries) {
	*sl = append(*sl, sketches)
}

// SketchesSource is a source of sketches used by the serializer.
type SketchesSource interface {
	MoveNext() bool
	Current() *SketchSeries
	Count() uint64
}

// These functions are removed in a later commit
func (sl SketchSeriesList) MoveNext() bool {
	panic("NOT IMPLEMENTED")
}

func (sl SketchSeriesList) Current() *SketchSeries {
	panic("NOT IMPLEMENTED")

}

func (sl SketchSeriesList) Count() uint64 {
	panic("NOT IMPLEMENTED")
}
