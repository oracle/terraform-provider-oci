package generative_ai_agent

import (
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("generative_ai_agent", generativeAiAgentResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGenerativeAiAgentDataSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_agent_data_source",
	DatasourceClass:        "oci_generative_ai_agent_data_sources",
	DatasourceItemsAttr:    "data_source_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai_agent.DataSourceLifecycleStateActive),
	},
}

var exportGenerativeAiAgentAgentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_agent_agent",
	DatasourceClass:        "oci_generative_ai_agent_agents",
	DatasourceItemsAttr:    "agent_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "agent",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai_agent.AgentLifecycleStateActive),
	},
}

var exportGenerativeAiAgentDataIngestionJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_agent_data_ingestion_job",
	DatasourceClass:        "oci_generative_ai_agent_data_ingestion_jobs",
	DatasourceItemsAttr:    "data_ingestion_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "data_ingestion_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai_agent.DataIngestionJobLifecycleStateSucceeded),
	},
}

var exportGenerativeAiAgentKnowledgeBaseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_agent_knowledge_base",
	DatasourceClass:        "oci_generative_ai_agent_knowledge_bases",
	DatasourceItemsAttr:    "knowledge_base_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "knowledge_base",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai_agent.KnowledgeBaseLifecycleStateActive),
	},
}

var exportGenerativeAiAgentAgentEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_generative_ai_agent_agent_endpoint",
	DatasourceClass:        "oci_generative_ai_agent_agent_endpoints",
	DatasourceItemsAttr:    "agent_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "agent_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_generative_ai_agent.AgentEndpointLifecycleStateActive),
	},
}

var generativeAiAgentResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGenerativeAiAgentDataSourceHints},
		{TerraformResourceHints: exportGenerativeAiAgentAgentHints},
		{TerraformResourceHints: exportGenerativeAiAgentDataIngestionJobHints},
		{TerraformResourceHints: exportGenerativeAiAgentKnowledgeBaseHints},
		{TerraformResourceHints: exportGenerativeAiAgentAgentEndpointHints},
	},
}
