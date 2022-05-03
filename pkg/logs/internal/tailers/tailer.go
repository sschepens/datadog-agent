// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tailers

import (
	"github.com/DataDog/datadog-agent/pkg/logs/auditor"
	"github.com/DataDog/datadog-agent/pkg/logs/message"
)

// Tailer abstracts types which can tail logs.
type Tailer interface {

	// Start starts the tailer, sending messages to the given pipeline and
	// using the given registry.
	Start(pipeline chan *message.Message, registry auditor.Registry)

	// Stop stops the tailer, waiting for the operations to finish.
	Stop()
}
