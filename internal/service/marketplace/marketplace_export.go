package marketplace

import (
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("marketplace", marketplaceResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportMarketplaceAcceptedAgreementHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_marketplace_accepted_agreement",
	DatasourceClass:        "oci_marketplace_accepted_agreements",
	DatasourceItemsAttr:    "accepted_agreements",
	ResourceAbbreviation:   "accepted_agreement",
	RequireResourceRefresh: true,
}

var exportMarketplacePublicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_marketplace_publication",
	DatasourceClass:        "oci_marketplace_publications",
	DatasourceItemsAttr:    "publications",
	ResourceAbbreviation:   "publication",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_marketplace.PublicationLifecycleStateActive),
	},
}

var marketplaceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMarketplaceAcceptedAgreementHints},
	},
}
