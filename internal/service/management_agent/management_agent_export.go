package management_agent

import (
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("management_agent", managementAgentResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

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

var managementAgentResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportManagementAgentManagementAgentHints},
		{TerraformResourceHints: exportManagementAgentManagementAgentInstallKeyHints},
	},
}
