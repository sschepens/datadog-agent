// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tailerfactory

import (
	"github.com/DataDog/datadog-agent/pkg/logs/internal/tailers"
	"github.com/DataDog/datadog-agent/pkg/logs/sources"
)

// makeFileTailer makes a file-based tailer for the given source, or returns
// an error if it cannot do so (e.g., due to permission errors)
func (tf *TailerFactory) makeFileTailer(source *sources.LogSource) (tailers.Tailer, error) {
	panic("TODO")
}
