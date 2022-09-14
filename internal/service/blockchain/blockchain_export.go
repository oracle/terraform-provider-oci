package blockchain

import (
	"fmt"

	oci_blockchain "github.com/oracle/oci-go-sdk/v65/blockchain"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportBlockchainPeerHints.GetIdFn = getBlockchainPeerId
	exportBlockchainOsnHints.GetIdFn = getBlockchainOsnId
	tf_export.RegisterCompartmentGraphs("blockchain", blockchainResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getBlockchainPeerId(resource *tf_export.OCIResource) (string, error) {

	blockchainPlatformId := resource.Parent.Id
	peerId, ok := resource.SourceAttributes["peer_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerId for Blockchain Peer")
	}
	return GetPeerCompositeId(blockchainPlatformId, peerId), nil
}

func getBlockchainOsnId(resource *tf_export.OCIResource) (string, error) {

	blockchainPlatformId := resource.Parent.Id
	peerId, ok := resource.SourceAttributes["osn_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find peerId for Blockchain Peer")
	}
	return GetOsnCompositeId(blockchainPlatformId, peerId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportBlockchainBlockchainPlatformHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_blockchain_blockchain_platform",
	DatasourceClass:        "oci_blockchain_blockchain_platforms",
	DatasourceItemsAttr:    "blockchain_platform_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "blockchain_platform",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_blockchain.BlockchainPlatformLifecycleStateActive),
	},
}

var exportBlockchainPeerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_blockchain_peer",
	DatasourceClass:        "oci_blockchain_peers",
	DatasourceItemsAttr:    "peer_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "peer",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_blockchain.PeerLifecycleStateActive),
	},
}

var exportBlockchainOsnHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_blockchain_osn",
	DatasourceClass:        "oci_blockchain_osns",
	DatasourceItemsAttr:    "osn_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "osn",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_blockchain.OsnLifecycleStateActive),
	},
}

var blockchainResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBlockchainBlockchainPlatformHints},
	},
	"oci_blockchain_blockchain_platform": {
		{
			TerraformResourceHints: exportBlockchainOsnHints,
			DatasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBlockchainPeerHints,
			DatasourceQueryParams: map[string]string{
				"blockchain_platform_id": "id",
			},
		},
	},
}
