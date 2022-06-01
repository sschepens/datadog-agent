package flowaggregator

import (
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/netflow/enrichment"

	"github.com/DataDog/datadog-agent/pkg/netflow/common"
	"github.com/DataDog/datadog-agent/pkg/netflow/payload"
)

func buildPayload(aggFlow *common.Flow, hostname string) payload.FlowPayload {
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
			IP: common.IPBytesToString(aggFlow.ExporterAddr),
		},
		Start:      aggFlow.StartTimestamp,
		End:        aggFlow.EndTimestamp,
		Bytes:      aggFlow.Bytes,
		Packets:    aggFlow.Packets,
		EtherType:  etherType,
		IPProtocol: ipProtocol,
		Source: payload.Endpoint{
			IP:   common.IPBytesToString(aggFlow.SrcAddr),
			Port: aggFlow.SrcPort,
			Mac:  enrichment.FormatMacAddress(aggFlow.SrcMac),
			Mask: enrichment.FormatMask(aggFlow.SrcAddr, aggFlow.SrcMask),
		},
		Destination: payload.Endpoint{
			IP:   common.IPBytesToString(aggFlow.DstAddr),
			Port: aggFlow.DstPort,
			Mac:  enrichment.FormatMacAddress(aggFlow.DstMac),
			Mask: enrichment.FormatMask(aggFlow.DstAddr, aggFlow.DstMask),
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
			IP: common.IPBytesToString(aggFlow.NextHop),
		},
	}
}
