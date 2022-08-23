package kms

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportKmsKeyHints.GetIdFn = getKmsKeyId
	exportKmsKeyVersionHints.GetIdFn = getKmsKeyVersionId
	exportKmsKeyHints.ProcessDiscoveredResourcesFn = processKmsKey
	exportKmsKeyVersionHints.ProcessDiscoveredResourcesFn = processKmsKeyVersion
	tf_export.RegisterCompartmentGraphs("kms", kmsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func processKmsKey(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["management_endpoint"] = resource.Parent.SourceAttributes["management_endpoint"].(string)
		var resourceSchema *schema.ResourceData = resource.RawResource.(*schema.ResourceData)
		resource.SourceAttributes["id"] = resourceSchema.Id()
	}
	return resources, nil
}

func processKmsKeyVersion(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["management_endpoint"] = resource.Parent.SourceAttributes["management_endpoint"].(string)
		resource.ImportId = resource.Id
	}
	return resources, nil
}

func getKmsKeyId(resource *tf_export.OCIResource) (string, error) {

	managementEndpoint, ok := resource.Parent.SourceAttributes["management_endpoint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find management_endpoint for Index id")
	}
	var keyId string
	// observed that Id is not always available in sourceAttributes - refer export_compartment.go->findResourcesGeneric() to visualize below docs
	// resource.sourceAttributes has the id in the cases where getKmsKeyId is called with LIST data source response, because list SetData() sets the Id, but this is only done temporarily to populate compositeID
	// When getKmsKeyId is called for resource, resource.sourceAttributes is not set yet,(so far we used LIST response to get composite Id) but we can get the real ocid after Read because Id was set in the method kms_key_resource.go->readKmsKey()
	switch resource.RawResource.(type) {
	case *schema.ResourceData:
		// 	rawResource from resource read response
		var resourceSchema *schema.ResourceData = resource.RawResource.(*schema.ResourceData)
		keyId = resourceSchema.Id()
	case map[string]interface{}:
		// 	rawResource from LIST data source read response
		var resourceMap map[string]interface{} = resource.RawResource.(map[string]interface{})
		keyId = resourceMap["id"].(string)
	}
	return GetCompositeKeyId(managementEndpoint, keyId), nil
}

func getKmsKeyVersionId(resource *tf_export.OCIResource) (string, error) {

	managementEndpoint, ok := resource.Parent.SourceAttributes["management_endpoint"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find management_endpoint for Kms KeyVersion")
	}
	keyId := resource.Parent.SourceAttributes["id"].(string)
	keyVersionId, ok := resource.SourceAttributes["key_version_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find keyVersionId for Kms KeyVersion")
	}
	return GetCompositeKeyVersionId(managementEndpoint, keyId, keyVersionId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportKmsKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_kms_key",
	DatasourceClass:        "oci_kms_keys",
	DatasourceItemsAttr:    "keys",
	ResourceAbbreviation:   "key",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_kms.KeyLifecycleStateEnabled),
	},
}

var exportKmsKeyVersionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_key_version",
	DatasourceClass:      "oci_kms_key_versions",
	DatasourceItemsAttr:  "key_versions",
	ResourceAbbreviation: "key_version",
	DiscoverableLifecycleStates: []string{
		string(oci_kms.KeyVersionLifecycleStateEnabled),
	},
}

var exportKmsVaultHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_vault",
	DatasourceClass:      "oci_kms_vaults",
	DatasourceItemsAttr:  "vaults",
	ResourceAbbreviation: "vault",
	DiscoverableLifecycleStates: []string{
		string(oci_kms.VaultLifecycleStateActive),
	},
}

var exportKmsSignHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_sign",
	ResourceAbbreviation: "sign",
}

var exportKmsVerifyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_verify",
	ResourceAbbreviation: "verify",
}

var exportKmsCreateReplicaHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_vault_replication",
	ResourceAbbreviation: "vault_replication",
}

var exportKmsDeleteReplicaHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_kms_vault_replication",
	ResourceAbbreviation: "vault_replication",
}

var kmsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportKmsVaultHints},
	},
	"oci_kms_key": {
		{
			TerraformResourceHints: exportKmsKeyVersionHints,
			DatasourceQueryParams: map[string]string{
				"key_id":              "id",
				"management_endpoint": "management_endpoint",
			},
		},
	},
	"oci_kms_vault": {
		{
			TerraformResourceHints: exportKmsKeyHints,
			DatasourceQueryParams: map[string]string{
				"management_endpoint": "management_endpoint",
			},
		},
	},
}
