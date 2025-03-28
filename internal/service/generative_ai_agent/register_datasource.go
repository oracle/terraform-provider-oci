// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_generative_ai_agent_agent", GenerativeAiAgentAgentDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_agent_endpoint", GenerativeAiAgentAgentEndpointDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_agent_endpoints", GenerativeAiAgentAgentEndpointsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_agents", GenerativeAiAgentAgentsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_data_ingestion_job", GenerativeAiAgentDataIngestionJobDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_data_ingestion_job_log_content", GenerativeAiAgentDataIngestionJobLogContentDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_data_ingestion_jobs", GenerativeAiAgentDataIngestionJobsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_data_source", GenerativeAiAgentDataSourceDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_data_sources", GenerativeAiAgentDataSourcesDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_knowledge_base", GenerativeAiAgentKnowledgeBaseDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_agent_knowledge_bases", GenerativeAiAgentKnowledgeBasesDataSource())
}
