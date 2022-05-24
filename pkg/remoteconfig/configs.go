// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022-present Datadog, Inc.

package remoteconfig

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/remoteconfig/products/apmsampling"
	"github.com/theupdateframework/go-tuf/data"
)

var allProducts = []string{ProductAPMSampling, ProductCWSDD, ProductFeatures, ProductLiveDebugging}

const (
	// ProductAPMSampling is the apm sampling product
	ProductAPMSampling = "APM_SAMPLING"
	// ProductCWSDD is the cloud workload security product managed by datadog employees
	ProductCWSDD = "CWS_DD"
	// ProductFeatures is a pseudo-product that lists whether or not a product should be enabled in a tracer
	ProductFeatures = "FEATURES"
	// ProductLiveDebugging is the live debugging product
	ProductLiveDebugging = "LIVE_DEBUGGING"
)

var ErrNoConfigVersion = errors.New("version missing in custom file meta")

// APMSamplingConfig is a deserialized APM Sampling configuration file
// along with its associated remote config metadata.
type APMSamplingConfig struct {
	Config   apmsampling.APMSampling
	Metadata Metadata
}

func parseConfigAPMSampling(data []byte) (APMSamplingConfig, error) {
	var apmConfig apmsampling.APMSampling
	_, err := apmConfig.UnmarshalMsg(data)
	if err != nil {
		return APMSamplingConfig{}, fmt.Errorf("could not parse apm sampling config: %v", err)
	}
	return APMSamplingConfig{
		Config: apmConfig,
	}, nil
}

// ConfigCWSDD is a deserialized CWS DD configuration file along with its
// associated remote config metadata
type ConfigCWSDD struct {
	Config   []byte
	Metadata Metadata
}

func parseConfigCWSDD(data []byte) (ConfigCWSDD, error) {
	return ConfigCWSDD{
		Config: data,
	}, nil
}

type LDConfig struct {
	Config   []byte
	Metadata Metadata
}

func parseLDConfig(data []byte) (LDConfig, error) {
	return LDConfig{
		Config: data,
	}, nil
}

type FeaturesConfig struct {
	Config   []byte
	Metadata Metadata
}

func parseFeaturesConfing(data []byte) (FeaturesConfig, error) {
	return FeaturesConfig{
		Config: data,
	}, nil
}

type Metadata struct {
	Product   string
	ID        string
	Name      string
	Version   uint64
	RawLength uint64
	Hashes    map[string][]byte
}

func newConfigMetadata(parsedPath configPath, tfm data.TargetFileMeta) (Metadata, error) {
	var m Metadata
	m.ID = parsedPath.ConfigID
	m.Product = parsedPath.Product
	m.Name = parsedPath.Name
	m.RawLength = uint64(tfm.Length)
	m.Hashes = make(map[string][]byte)
	for k, v := range tfm.Hashes {
		m.Hashes[k] = []byte(v)
	}
	v, err := fileMetaVersion(tfm)
	if err != nil {
		return Metadata{}, err
	}
	m.Version = v

	return m, nil
}

type fileMetaCustom struct {
	Version *uint64 `json:"v"`
}

func fileMetaVersion(fm data.TargetFileMeta) (uint64, error) {
	fmc, err := parseFileMetaCustom(*fm.Custom)
	if err != nil {
		return 0, err
	}

	return *fmc.Version, nil
}

func parseFileMetaCustom(rawCustom []byte) (fileMetaCustom, error) {
	var custom fileMetaCustom
	err := json.Unmarshal(rawCustom, &custom)
	if err != nil {
		return fileMetaCustom{}, err
	}
	if custom.Version == nil {
		return fileMetaCustom{}, ErrNoConfigVersion
	}
	return custom, nil
}
