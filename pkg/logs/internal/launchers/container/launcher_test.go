// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package container

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/logs/auditor"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/launchers"
	"github.com/DataDog/datadog-agent/pkg/logs/pipeline"
	"github.com/stretchr/testify/require"
)

func TestStartStop(t *testing.T) {
	l := NewLauncher()

	sp := launchers.NewMockSourceProvider()
	pl := pipeline.NewMockProvider()
	reg := auditor.New("/run", "agent", 0, nil)
	l.Start(sp, pl, reg)

	require.NotNil(t, l.cancel)
	require.NotNil(t, l.stopped)

	l.Stop()

	require.Nil(t, l.cancel)
	require.Nil(t, l.stopped)
}
