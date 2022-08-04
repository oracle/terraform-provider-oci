package streaming

import (
	oci_streaming "github.com/oracle/oci-go-sdk/v65/streaming"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportStreamingStreamHints.ProcessDiscoveredResourcesFn = processStreamingStream
	tf_export.RegisterCompartmentGraphs("streaming", streamingResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func processStreamingStream(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, streamingStream := range resources {
		// compartment_id conflict with stream_pool_id
		if _, exists := streamingStream.SourceAttributes["compartment_id"]; exists {
			if _, ok := streamingStream.SourceAttributes["stream_pool_id"]; ok {
				delete(streamingStream.SourceAttributes, "stream_pool_id")
			}
		}
	}
	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportStreamingConnectHarnessHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_streaming_connect_harness",
	DatasourceClass:      "oci_streaming_connect_harnesses",
	DatasourceItemsAttr:  "connect_harness",
	ResourceAbbreviation: "connect_harness",
	DiscoverableLifecycleStates: []string{
		string(oci_streaming.ConnectHarnessLifecycleStateActive),
	},
}

var exportStreamingStreamPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_streaming_stream_pool",
	DatasourceClass:        "oci_streaming_stream_pools",
	DatasourceItemsAttr:    "stream_pools",
	ResourceAbbreviation:   "stream_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_streaming.StreamPoolLifecycleStateActive),
	},
}

var exportStreamingStreamHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_streaming_stream",
	DatasourceClass:        "oci_streaming_streams",
	DatasourceItemsAttr:    "streams",
	ResourceAbbreviation:   "stream",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_streaming.StreamLifecycleStateActive),
	},
}

var streamingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStreamingConnectHarnessHints},
		{TerraformResourceHints: exportStreamingStreamPoolHints},
		{TerraformResourceHints: exportStreamingStreamHints},
	},
}
