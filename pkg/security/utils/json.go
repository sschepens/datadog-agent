// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package utils

import (
	"errors"
	"time"

	"github.com/mailru/easyjson/jwriter"
)

type EasyjsonTime struct {
	inner time.Time
}

func NewEasyjsonTime(t time.Time) EasyjsonTime {
	return EasyjsonTime{inner: t}
}

func (t EasyjsonTime) MarshalEasyJSON(w *jwriter.Writer) {
	if y := t.inner.Year(); y < 0 || y >= 10000 {
		if w.Error == nil {
			w.Error = errors.New("Time.MarshalJSON: year outside of range [0,9999]")
		}
		return
	}

	w.Buffer.EnsureSpace(len(time.RFC3339Nano) + 2)
	w.Buffer.AppendByte('"')
	w.Buffer.Buf = t.inner.AppendFormat(w.Buffer.Buf, time.RFC3339Nano)
	w.Buffer.AppendByte('"')
}

func (t *EasyjsonTime) UnmarshalJSON(b []byte) error {
	return t.inner.UnmarshalJSON(b)
}
