// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build orchestrator
// +build orchestrator

package processors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Item struct {
	UID string
}

func TestChunkOrchestratorMetadataBySizeAndWeight(t *testing.T) {
	orchestratorMetadata := []interface{}{
		Item{UID: "1"},
		Item{UID: "2"},
		Item{UID: "3"},
		Item{UID: "4"},
		Item{UID: "5"},
	}
	tests := []struct {
		name                     string
		maxChunkSize             int
		maxChunkWeight           int
		orchestratorMetadata     []interface{}
		orchestratorMetadataYaml [][]byte
		expectedChunks           [][]interface{}
	}{
		{
			name:                 "chunk by size and weight, one high weight",
			maxChunkSize:         3,
			maxChunkWeight:       1000,
			orchestratorMetadata: orchestratorMetadata,
			orchestratorMetadataYaml: [][]byte{
				make([]byte, 1001),
				make([]byte, 100),
				make([]byte, 100),
				make([]byte, 100),
				make([]byte, 100),
			},
			expectedChunks: [][]interface{}{
				{Item{UID: "1"}},
				{Item{UID: "2"}, Item{UID: "3"}, Item{UID: "4"}},
				{Item{UID: "5"}},
			},
		},
		{
			name:                 "chunk by size and weight, weight exceeded",
			maxChunkSize:         3,
			maxChunkWeight:       1000,
			orchestratorMetadata: orchestratorMetadata,
			orchestratorMetadataYaml: [][]byte{
				make([]byte, 2000),
				make([]byte, 2000),
				make([]byte, 2000),
				make([]byte, 2000),
				make([]byte, 2000),
			},
			expectedChunks: [][]interface{}{
				{Item{UID: "1"}},
				{Item{UID: "2"}},
				{Item{UID: "3"}},
				{Item{UID: "4"}},
				{Item{UID: "5"}},
			},
		},
		{
			name:                 "chunk by size and weight, low weight",
			maxChunkSize:         3,
			maxChunkWeight:       1000,
			orchestratorMetadata: orchestratorMetadata,
			orchestratorMetadataYaml: [][]byte{
				make([]byte, 100),
				make([]byte, 100),
				make([]byte, 100),
				make([]byte, 100),
				make([]byte, 100),
			},
			expectedChunks: [][]interface{}{
				{Item{UID: "1"}, Item{UID: "2"}, Item{UID: "3"}},
				{Item{UID: "4"}, Item{UID: "5"}},
			},
		},
		{
			name:                 "chunk by size and weight, mixed",
			maxChunkSize:         3,
			maxChunkWeight:       1000,
			orchestratorMetadata: orchestratorMetadata,
			orchestratorMetadataYaml: [][]byte{
				make([]byte, 200),
				make([]byte, 400),
				make([]byte, 800),
				make([]byte, 300),
				make([]byte, 2000),
			},
			expectedChunks: [][]interface{}{
				{Item{UID: "1"}, Item{UID: "2"}},
				{Item{UID: "3"}},
				{Item{UID: "4"}},
				{Item{UID: "5"}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			chunker := &collectorOrchestratorMetadataChunker{}
			chunkOrchestratorMetadataBySizeAndWeight(tc.orchestratorMetadata, tc.orchestratorMetadataYaml, tc.maxChunkSize, tc.maxChunkWeight, chunker)
			assert.Equal(t, tc.expectedChunks, chunker.collectorOrchestratorMetadataList)
		})
	}
}
