package oci_product_catalog

import (
	oci_oci_product_catalog "github.com/oracle/oci-go-sdk/v65/ociproductcatalog"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportOciProductCatalogInternalAdminProductHints.GetIdFn = getOciProductCatalogInternalAdminProductId
	exportOciProductCatalogInternalProductHints.GetIdFn = getOciProductCatalogInternalProductId
	tf_export.RegisterCompartmentGraphs("oci_product_catalog", ociProductCatalogResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getOciProductCatalogInternalAdminProductId(resource *tf_export.OCIResource) (string, error) {

	productId := resource.Parent.Id
	return GetInternalAdminProductCompositeId(productId), nil
}

func getOciProductCatalogInternalProductId(resource *tf_export.OCIResource) (string, error) {

	productId := resource.Parent.Id
	return GetInternalProductCompositeId(productId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportOciProductCatalogInternalAdminProductHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_oci_product_catalog_internal_admin_product",
	DatasourceClass:      "oci_oci_product_catalog_internal_admin_product",
	ResourceAbbreviation: "internal_admin_product",
	DiscoverableLifecycleStates: []string{
		string(oci_oci_product_catalog.ProductLifecycleStateActive),
		string(oci_oci_product_catalog.ProductLifecycleStateInactive),
		string(oci_oci_product_catalog.ProductLifecycleStateNeedsAttention),
	},
}

var exportOciProductCatalogInternalProductHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_oci_product_catalog_internal_product",
	DatasourceClass:      "oci_oci_product_catalog_internal_product",
	ResourceAbbreviation: "internal_product",
	DiscoverableLifecycleStates: []string{
		string(oci_oci_product_catalog.ProductLifecycleStateActive),
		string(oci_oci_product_catalog.ProductLifecycleStateInactive),
		string(oci_oci_product_catalog.ProductLifecycleStateNeedsAttention),
	},
}

var ociProductCatalogResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
