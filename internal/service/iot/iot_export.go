package iot

import (
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("iot", iotResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportIotDigitalTwinModelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_digital_twin_model",
	DatasourceClass:        "oci_iot_digital_twin_models",
	DatasourceItemsAttr:    "digital_twin_model_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "digital_twin_model",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.LifecycleStateActive),
	},
	DefaultValuesForMissingAttributes: map[string]interface{}{
		"spec": "{}",
	},
}

var exportIotIotDomainGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_iot_domain_group",
	DatasourceClass:        "oci_iot_iot_domain_groups",
	DatasourceItemsAttr:    "iot_domain_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "iot_domain_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.IotDomainGroupLifecycleStateActive),
	},
}

var exportIotIotDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_iot_domain",
	DatasourceClass:        "oci_iot_iot_domains",
	DatasourceItemsAttr:    "iot_domain_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "iot_domain",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.IotDomainLifecycleStateActive),
	},
}

var exportIotDigitalTwinRelationshipHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_digital_twin_relationship",
	DatasourceClass:        "oci_iot_digital_twin_relationships",
	DatasourceItemsAttr:    "digital_twin_relationship_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "digital_twin_relationship",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.LifecycleStateActive),
	},
}

var exportIotDigitalTwinInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_digital_twin_instance",
	DatasourceClass:        "oci_iot_digital_twin_instances",
	DatasourceItemsAttr:    "digital_twin_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "digital_twin_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.LifecycleStateActive),
	},
}

var exportIotDigitalTwinAdapterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_iot_digital_twin_adapter",
	DatasourceClass:        "oci_iot_digital_twin_adapters",
	DatasourceItemsAttr:    "digital_twin_adapter_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "digital_twin_adapter",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_iot.LifecycleStateActive),
	},
}

var iotResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportIotIotDomainGroupHints},
		{TerraformResourceHints: exportIotIotDomainHints},
	},
	"oci_iot_iot_domain": {
		{
			TerraformResourceHints: exportIotDigitalTwinAdapterHints,
			DatasourceQueryParams: map[string]string{
				"iot_domain_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIotDigitalTwinInstanceHints,
			DatasourceQueryParams: map[string]string{
				"iot_domain_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIotDigitalTwinModelHints,
			DatasourceQueryParams: map[string]string{
				"iot_domain_id": "id",
			},
		},
		{
			TerraformResourceHints: exportIotDigitalTwinRelationshipHints,
			DatasourceQueryParams: map[string]string{
				"iot_domain_id": "id",
			},
		},
	},
}
