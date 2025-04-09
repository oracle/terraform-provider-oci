package bds

import (
	"fmt"

	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportBdsBdsInstanceApiKeyHints.GetIdFn = getBdsBdsInstanceApiKeyId
	exportBdsBdsInstanceIdentityConfigurationHints.GetIdFn = getBdsBdsInstanceIdentityConfigurationId
	exportBdsBdsInstanceApiKeyHints.ProcessDiscoveredResourcesFn = processBdsInstanceApiKeys
	exportBdsBdsInstanceMetastoreConfigHints.ProcessDiscoveredResourcesFn = processBdsInstanceMetastoreConfigs
	exportBdsBdsInstanceIdentityConfigurationHints.ProcessDiscoveredResourcesFn = processBdsInstanceIdentityConfigurations
	exportBdsBdsInstanceResourcePrincipalConfigurationHints.GetIdFn = getBdsBdsInstanceResourcePrincipalConfigurationId
	tf_export.RegisterCompartmentGraphs("bds", bdsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processBdsInstanceMetastoreConfigs(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		metastoreConfigId := resource.Id
		bdsInstanceId := resource.Parent.Id
		resource.ImportId = GetBdsInstanceMetastoreConfigCompositeId(bdsInstanceId, metastoreConfigId)
	}
	return resources, nil
}

func processBdsInstanceApiKeys(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		apiKeyId := resource.Id
		bdsInstanceId := resource.SourceAttributes["bds_instance_id"].(string)
		resource.ImportId = GetBdsInstanceApiKeyCompositeId(apiKeyId, bdsInstanceId)
	}
	return resources, nil
}
func processBdsInstanceIdentityConfigurations(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		identityConfigurationId := resource.Id
		bdsInstanceId := resource.SourceAttributes["bds_instance_id"].(string)
		resource.ImportId = GetBdsInstanceIdentityConfigurationCompositeId(bdsInstanceId, identityConfigurationId)
	}
	return resources, nil
}

func getBdsBdsInstanceApiKeyId(resource *tf_export.OCIResource) (string, error) {

	apiKeyId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apiKeyId for Bds BdsInstanceApiKey")
	}
	bdsInstanceId := resource.Parent.Id
	return GetBdsInstanceApiKeyCompositeId(apiKeyId, bdsInstanceId), nil
}

func getBdsBdsInstanceMetastoreConfigId(resource *tf_export.OCIResource) (string, error) {

	bdsInstanceId := resource.Parent.Id
	metastoreConfigId, ok := resource.SourceAttributes["metastore_config_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find metastoreConfigId for Bds BdsInstanceMetastoreConfig")
	}
	return GetBdsInstanceMetastoreConfigCompositeId(bdsInstanceId, metastoreConfigId), nil
}

func getBdsBdsInstanceIdentityConfigurationId(resource *tf_export.OCIResource) (string, error) {

	bdsInstanceId := resource.Parent.Id
	identityConfigurationId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find identityConfigurationId for Bds BdsInstanceIdentityConfiguration")
	}
	return GetBdsInstanceIdentityConfigurationCompositeId(bdsInstanceId, identityConfigurationId), nil
}

func getBdsBdsInstanceResourcePrincipalConfigurationId(resource *tf_export.OCIResource) (string, error) {

	bdsInstanceId := resource.Parent.Id
	resourcePrincipalConfigurationId, ok := resource.SourceAttributes["resource_principal_configuration_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find resourcePrincipalConfigurationId for Bds BdsInstanceResourcePrincipalConfiguration")
	}
	return GetBdsInstanceResourcePrincipalConfigurationCompositeId(bdsInstanceId, resourcePrincipalConfigurationId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportBdsBdsInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bds_bds_instance",
	DatasourceClass:        "oci_bds_bds_instances",
	DatasourceItemsAttr:    "bds_instances",
	ResourceAbbreviation:   "bds_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bds.BdsInstanceLifecycleStateActive),
	},
}

var exportBdsBdsInstanceApiKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bds_bds_instance_api_key",
	DatasourceClass:        "oci_bds_bds_instance_api_keys",
	DatasourceItemsAttr:    "bds_api_keys",
	ResourceAbbreviation:   "bds_instance_api_key",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bds.BdsApiKeyLifecycleStateActive),
	},
}

var exportBdsBdsInstanceMetastoreConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_bds_bds_instance_metastore_config",
	DatasourceClass:      "oci_bds_bds_instance_metastore_configs",
	DatasourceItemsAttr:  "bds_metastore_configurations",
	ResourceAbbreviation: "bds_instance_metastore_config",
	DiscoverableLifecycleStates: []string{
		string(oci_bds.BdsMetastoreConfigurationLifecycleStateActive),
	},
}

var exportBdsBdsInstanceIdentityConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bds_bds_instance_identity_configuration",
	DatasourceClass:        "oci_bds_bds_instance_identity_configurations",
	DatasourceItemsAttr:    "identity_configurations",
	ResourceAbbreviation:   "bds_instance_identity_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bds.IdentityConfigurationLifecycleStateActive),
	},
}

var exportBdsBdsInstanceResourcePrincipalConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_bds_bds_instance_resource_principal_configuration",
	DatasourceClass:        "oci_bds_bds_instance_resource_principal_configurations",
	DatasourceItemsAttr:    "resource_principal_configurations",
	ResourceAbbreviation:   "bds_instance_resource_principal_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_bds.ResourcePrincipalConfigurationLifecycleStateActive),
	},
}

var bdsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBdsBdsInstanceHints},
	},
	"oci_bds_bds_instance": {
		{
			TerraformResourceHints: exportBdsBdsInstanceApiKeyHints,
			DatasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBdsBdsInstanceIdentityConfigurationHints,
			DatasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBdsBdsInstanceMetastoreConfigHints,
			DatasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportBdsBdsInstanceResourcePrincipalConfigurationHints,
			DatasourceQueryParams: map[string]string{
				"bds_instance_id": "id",
			},
		},
	},
}
