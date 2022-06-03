// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tailerfactory

import (
	"fmt"

	coreConfig "github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/logs/internal/util/containersorpods"
	"github.com/DataDog/datadog-agent/pkg/logs/sources"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// useFile determines whether the user would like to log this source with
// file or socket.  It does not handle fallback in the presence of permissions
// errors.
func (tf *TailerFactory) useFile(source *sources.LogSource) bool {
	// The user configuration consulted is different depending on what we are
	// logging.  Note that we assume that by the time we have gotten a source,
	// the necessary utilities are started, thus calling Get and not Wait.
	logWhat := tf.cop.Get()

	switch logWhat {
	case containersorpods.LogContainers:
		// docker_container_use_file is a suggestion
		if !coreConfig.Datadog.GetBool("logs_config.docker_container_use_file") {
			return false
		}

		// docker_container_force_use_file is a requirement
		if coreConfig.Datadog.GetBool("logs_config.docker_container_force_use_file") {
			return true
		}

		// if file was suggested, but this source has a registry entry with a
		// docker socket path, use socket.
		if source.Config.Identifier != "" {
			// TODO: this should be a utility function
			registryID := fmt.Sprintf("%s:%s", source.Config.Type, source.Config.Identifier)
			if tf.registry.GetOffset(registryID) != "" {
				return false
			}
		}

		return true

	case containersorpods.LogPods:
		return coreConfig.Datadog.GetBool("logs_config.k8s_container_use_file")

	default:
		// if this occurs, then sources have been arriving before the
		// container interfaces to them are ready.  Something is wrong.
		log.Warnf("LogWhat = %s; not ready to log containers", logWhat.String())
		return false
	}
}
