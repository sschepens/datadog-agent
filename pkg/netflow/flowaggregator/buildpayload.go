package flowaggregator

import (
	"context"
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/netflow/enrichment"
	"github.com/DataDog/datadog-agent/pkg/util/log"

	coreutil "github.com/DataDog/datadog-agent/pkg/util"

	"github.com/DataDog/datadog-agent/pkg/netflow/common"
	"github.com/DataDog/datadog-agent/pkg/netflow/payload"
)

func buildPayload(aggFlow *common.Flow) payload.FlowPayload {
	hostname, err := coreutil.GetHostname(context.TODO())
	if err != nil {
		log.Warnf("Error getting the hostname: %v", err)
		hostname = ""
	}

	// TODO: format ipProtocol
	// TODO: format etherType
	ipProtocol := fmt.Sprintf("%d", aggFlow.IPProtocol)
	etherType := fmt.Sprintf("%d", aggFlow.EtherType)

	return payload.FlowPayload{
		// TODO: Implement Tos
		FlowType:     string(aggFlow.FlowType),
		SamplingRate: aggFlow.SamplingRate,
		Direction:    enrichment.RemapDirection(aggFlow.Direction),
		Exporter: payload.Exporter{
			IP: aggFlow.ExporterAddr,
		},
		Start:      aggFlow.StartTimestamp,
		End:        aggFlow.EndTimestamp,
		Bytes:      aggFlow.Bytes,
		Packets:    aggFlow.Packets,
		EtherType:  etherType,
		IPProtocol: ipProtocol,
		Source: payload.Endpoint{
			IP:   aggFlow.SrcAddr,
			Port: aggFlow.SrcPort,
			// TODO: implement Mask
			Mac: enrichment.FormatMacAddress(aggFlow.SrcMac),
			//Mask: enrichment.FormatMask(aggFlow.SrcMask),
			Mask: fmt.Sprintf("%d", aggFlow.SrcMask),
		},
		Destination: payload.Endpoint{
			IP:   aggFlow.DstAddr,
			Port: aggFlow.DstPort,
			Mac:  enrichment.FormatMacAddress(aggFlow.DstMac),
			Mask: fmt.Sprintf("%d", aggFlow.DstMask),
		},
		Ingress: payload.ObservationPoint{
			Interface: payload.Interface{
				Index: aggFlow.InputInterface,
			},
		},
		Egress: payload.ObservationPoint{
			Interface: payload.Interface{
				Index: aggFlow.OutputInterface,
			},
		},
		Namespace: aggFlow.Namespace,
		Host:      hostname,
		TCPFlags:  enrichment.FormatFCPFlags(aggFlow.TCPFlags),
		NextHop: payload.NextHop{
			IP: aggFlow.NextHop,
		},
	}
}
