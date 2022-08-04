package apigateway

import (
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("apigateway", apigatewayResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportApigatewayApiHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_api",
	DatasourceClass:        "oci_apigateway_apis",
	DatasourceItemsAttr:    "api_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "api",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.ApiLifecycleStateActive),
	},
}

var exportApigatewayGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_gateway",
	DatasourceClass:        "oci_apigateway_gateways",
	DatasourceItemsAttr:    "gateway_collection",
	ResourceAbbreviation:   "gateway",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.GatewayLifecycleStateActive),
	},
}

var exportApigatewayDeploymentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_deployment",
	DatasourceClass:        "oci_apigateway_deployments",
	DatasourceItemsAttr:    "deployment_collection",
	ResourceAbbreviation:   "deployment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.DeploymentLifecycleStateActive),
	},
}

var exportApigatewayCertificateHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_certificate",
	DatasourceClass:        "oci_apigateway_certificates",
	DatasourceItemsAttr:    "certificate_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "certificate",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.CertificateLifecycleStateActive),
	},
}

var exportApigatewaySubscriberHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_subscriber",
	DatasourceClass:        "oci_apigateway_subscribers",
	DatasourceItemsAttr:    "subscriber_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "subscriber",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.SubscriberLifecycleStateActive),
	},
}

var exportApigatewayUsagePlanHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apigateway_usage_plan",
	DatasourceClass:        "oci_apigateway_usage_plans",
	DatasourceItemsAttr:    "usage_plan_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "usage_plan",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apigateway.UsagePlanLifecycleStateActive),
	},
}

var apigatewayResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApigatewayApiHints},
		{TerraformResourceHints: exportApigatewayGatewayHints},
		{TerraformResourceHints: exportApigatewayDeploymentHints},
		{TerraformResourceHints: exportApigatewayCertificateHints},
		{TerraformResourceHints: exportApigatewaySubscriberHints},
		{TerraformResourceHints: exportApigatewayUsagePlanHints},
	},
}
