// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tailerfactory

import (
	"errors"
	"fmt"
	"time"

	coreConfig "github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/tailers"
	dockerTailer "github.com/DataDog/datadog-agent/pkg/logs/internal/tailers/docker"
	"github.com/DataDog/datadog-agent/pkg/logs/sources"
	dockerutilpkg "github.com/DataDog/datadog-agent/pkg/util/docker"
)

// makeSocketTailer makes a socket-based tailer for the given source, or returns
// an error if it cannot do so (e.g., due to permission errors)
func (tf *TailerFactory) makeSocketTailer(source *sources.LogSource) (tailers.Tailer, error) {
	identifier := source.Config.Identifier

	// sanity check; this should never be true
	if identifier == "" {
		return nil, errors.New("source.Config.Identifier is not set")
	}

	dockerutil, err := dockerutilpkg.GetDockerUtil()
	// sanity check: having gotten here logWhat = LogContainers, so dockerutil should
	// already be available
	if err != nil {
		return nil, fmt.Errorf("Could not use docker client, logs for container %s won’t be collected: %w",
			identifier, err)
	}

	pipeline := tf.pipelineProvider.NextPipelineChan()
	readTimeout := time.Duration(coreConfig.Datadog.GetInt("logs_config.docker_client_read_timeout")) * time.Second

	tailer := dockerTailer.NewTailer(
		dockerutil, identifier, source,
		pipeline,
		nil, // TODO: errored container ID channel
		readTimeout)
	return tailer, nil
}
