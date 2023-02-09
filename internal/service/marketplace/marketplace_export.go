package marketplace

import (
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportMarketplacePublicationHints.FindResourcesOverrideFn = findPublications
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
		string(oci_marketplace.PublicationLifecycleStateCreating),
		string(oci_marketplace.PublicationLifecycleStateActive),
	},
}

var marketplaceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMarketplaceAcceptedAgreementHints},
		{TerraformResourceHints: exportMarketplacePublicationHints},
	},
}

func findPublications(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) (resources []*tf_export.OCIResource, err error) {
	if tfMeta.DatasourceQueryParams == nil {
		tfMeta.DatasourceQueryParams = map[string]string{}
	}
	// Setting the "listing_type" field to the special value "COMMUNITY" will
	// result in terraform fetching community listings (i.e. publications)
	// when populating the "oci_marketplace_publications" data source
	tfMeta.DatasourceQueryParams["listing_type"] = "'COMMUNITY'"
	return tf_export.FindResourcesGeneric(ctx, tfMeta, parent, resourceGraph)
}
