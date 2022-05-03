// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package container

import (
	"context"
	"time"

	"github.com/DataDog/datadog-agent/pkg/logs/auditor"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/launchers"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/tailers"
	"github.com/DataDog/datadog-agent/pkg/logs/pipeline"
	"github.com/DataDog/datadog-agent/pkg/logs/sources"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

const (
	backoffInitialDuration = 1 * time.Second
	backoffMaxDuration     = 60 * time.Second
)

// A Launcher starts and stops new tailers for every new containers discovered by autodiscovery.
type Launcher struct {
	// cancel will cause the launcher loop to stop
	cancel context.CancelFunc

	// once the loop stops, this channel will be closed
	stopped chan struct{}

	// pipelineProvider provides pipelines to which log messages are sent
	pipelineProvider pipeline.Provider

	// registry is used to record positions in logs over process restarts
	registry auditor.Registry

	// tailers contains the tailer for each source
	tailers map[*sources.LogSource]tailers.Tailer
}

// NewLauncher returns a new launcher
func NewLauncher() *Launcher {
	launcher := &Launcher{}
	return launcher
}

// Start starts the Launcher
func (l *Launcher) Start(sourceProvider launchers.SourceProvider, pipelineProvider pipeline.Provider, registry auditor.Registry) {
	// only start this launcher once it's determined that we should be logging containers, and not pods.
	ctx, cancel := context.WithCancel(context.Background())
	l.cancel = cancel
	l.stopped = make(chan struct{})
	l.pipelineProvider = pipelineProvider
	l.registry = registry
	go l.run(ctx, sourceProvider)
}

// Stop stops the Launcher. This call returns when the launcher has stopped.
func (l *Launcher) Stop() {
	if l.cancel != nil {
		l.cancel()
		l.cancel = nil
		<-l.stopped
		l.stopped = nil
	}
}

// run is the main loop for this launcher.  It monitors for sources added or
// removed to the agent and starts or stops tailers appropriately.
func (l *Launcher) run(ctx context.Context, sourceProvider launchers.SourceProvider) {
	log.Info("Starting Container launcher")

	addedSources, removedSources := sourceProvider.SubscribeForType(sources.DockerSourceType)
	// TODO: register for all runtimes?

	for {
		select {
		case source := <-addedSources:
			l.startSource(source)

		case source := <-removedSources:
			l.stopSource(source)

		case <-ctx.Done():
			log.Info("Stopping Container launcher")
			// TODO: shut down sources (stop them? or does logs-agent do that??)
			close(l.stopped)
			return
		}
	}
}

// startSource starts tailing from a source.
func (l *Launcher) startSource(source *sources.LogSource) {
	if _, exists := l.tailers[source]; exists {
		return
	}

	panic("TODO")
	/*
		tailer := tailerfactory.MakeTailerForSource(source)
		if tailer == nil {
			source.Status.Error(errors.New("No tailer implementation defined for this source"))
			return
		}
	*/

	l.tailers[source] = tailer
	tailer.Start(l.pipelineProvider.NextPipelineChan(), l.registry)
}

// stopSource stops tailing from a source.
func (l *Launcher) stopSource(source *sources.LogSource) {
	if tailer, exists := l.tailers[source]; exists {
		tailer.Stop()
		l.tailers[source] = nil
	}
}
