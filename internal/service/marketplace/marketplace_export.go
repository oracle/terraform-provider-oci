package marketplace

import (
	"context"
	"fmt"

	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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
	results := []*tf_export.OCIResource{}
	request := oci_marketplace.ListPublicationsRequest{}
	request.ListingType = oci_marketplace.ListPublicationsListingTypeCommunity
	request.CompartmentId = &parent.CompartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "marketplace")

	response, err := ctx.Clients.MarketplaceClient().ListPublications(context.Background(), request)
	if err != nil {
		return nil, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.MarketplaceClient().ListPublications(context.Background(), request)
		if err != nil {
			return nil, err
		}
		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, publication := range response.Items {
		publicationResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

		d := publicationResource.TestResourceData()
		d.SetId(*publication.Id)

		if err := publicationResource.Read(d, ctx.Clients); err != nil {
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(publicationResource.Schema, d),
			RawResource:      publication,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, publicationResource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.Parent.TerraformName, *publication.Name)
			resource.TerraformName = tf_export.CheckDuplicateResourceName(resource.TerraformName)
		}

		results = append(results, resource)
	}

	return results, nil
}
