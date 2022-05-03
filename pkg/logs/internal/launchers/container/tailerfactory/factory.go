// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package tailerfactory implements the complex logic required to determine which
// kind of tailer to use for a container-related LogSource.
package tailerfactory

import (
	"github.com/DataDog/datadog-agent/pkg/logs/auditor"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/tailers"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/util/containersorpods"
	"github.com/DataDog/datadog-agent/pkg/logs/pipeline"
	"github.com/DataDog/datadog-agent/pkg/logs/sources"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// TailerFactory encapsulates the information required to determine which kind
// of tailer to use for a container-related LogSource.
type TailerFactory struct {
	// pipelineProvider provides pipelines for the instantiated tailers
	pipelineProvider pipeline.Provider

	// registry is the auditor/registry, used both to look for existing offsets
	// and to create new tailers.
	registry auditor.Registry

	// cop allows the factory to determine whether the agent is logging
	// containers or pods.
	cop containersorpods.Chooser
}

// New creates a new TailerFactory.
func New(pipelineProvider pipeline.Provider, registry auditor.Registry) *TailerFactory {
	return &TailerFactory{
		pipelineProvider: pipelineProvider,
		registry:         registry,
		cop:              containersorpods.NewChooser(),
	}
}

// MakeTailer creates a new tailer for the given LogSource.
func (tf *TailerFactory) MakeTailer(source *sources.LogSource) (tailers.Tailer, error) {
	if tf.useFile(source) {
		t, err := tf.makeFileTailer(source)
		if err == nil {
			return t, nil
		}
		log.Warnf("Could not make file tailer for source %s (falling back to socket): %v", source.Name, err)
		return tf.makeSocketTailer(source)
	}

	t, err := tf.makeSocketTailer(source)
	if err == nil {
		return t, nil
	}
	log.Warnf("Could not make socket tailer for source %s (falling back to file): %v", source.Name, err)
	return tf.makeFileTailer(source)
}
