package management_agent

import (
	"fmt"

	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportManagementAgentManagementAgentDataSourceHints.GetIdFn = getManagementAgentManagementAgentDataSourceId
	exportManagementAgentManagementAgentDataSourceHints.ProcessDiscoveredResourcesFn = processManagementAgentManagementAgentDataSourceKeys

	tf_export.RegisterCompartmentGraphs("management_agent", managementAgentResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processManagementAgentManagementAgentDataSourceKeys(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {

	for _, resource := range resources {
		managementAgentId := resource.Id
		dataSourceId := resource.SourceAttributes["data_source_key"].(string)
		resource.ImportId = GetManagementAgentDataSourceCompositeId(managementAgentId, dataSourceId)
	}
	return resources, nil
}

func getManagementAgentManagementAgentDataSourceId(resource *tf_export.OCIResource) (string, error) {

	dataSourceKey, ok := resource.SourceAttributes["data_source_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find dataSourceKey for ManagementAgent ManagementAgentDataSource")
	}
	managementAgentId := resource.Parent.Id
	return GetManagementAgentDataSourceCompositeId(dataSourceKey, managementAgentId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportManagementAgentManagementAgentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_management_agent_management_agent",
	DatasourceClass:      "oci_management_agent_management_agents",
	DatasourceItemsAttr:  "management_agents",
	ResourceAbbreviation: "management_agent",
	DiscoverableLifecycleStates: []string{
		string(oci_management_agent.LifecycleStatesActive),
	},
}

var exportManagementAgentManagementAgentInstallKeyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_management_agent_management_agent_install_key",
	DatasourceClass:      "oci_management_agent_management_agent_install_keys",
	DatasourceItemsAttr:  "management_agent_install_keys",
	ResourceAbbreviation: "management_agent_install_key",
	DiscoverableLifecycleStates: []string{
		string(oci_management_agent.LifecycleStatesActive),
	},
}

var exportManagementAgentManagementAgentDataSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_management_agent_management_agent_data_source",
	DatasourceClass:        "oci_management_agent_management_agent_data_sources",
	DatasourceItemsAttr:    "data_sources",
	ResourceAbbreviation:   "management_agent_data_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_management_agent.LifecycleStatesActive),
	},
}

var managementAgentResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportManagementAgentManagementAgentHints},
		{TerraformResourceHints: exportManagementAgentManagementAgentInstallKeyHints},
	},
	"oci_management_agent_management_agent": {
		{
			TerraformResourceHints: exportManagementAgentManagementAgentDataSourceHints,
			DatasourceQueryParams: map[string]string{
				"management_agent_id": "id",
			},
		},
	},
}
