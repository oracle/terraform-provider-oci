// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_generative_ai_agent_agent", GenerativeAiAgentAgentResource())
	tfresource.RegisterResource("oci_generative_ai_agent_agent_endpoint", GenerativeAiAgentAgentEndpointResource())
	tfresource.RegisterResource("oci_generative_ai_agent_data_ingestion_job", GenerativeAiAgentDataIngestionJobResource())
	tfresource.RegisterResource("oci_generative_ai_agent_data_source", GenerativeAiAgentDataSourceResource())
	tfresource.RegisterResource("oci_generative_ai_agent_knowledge_base", GenerativeAiAgentKnowledgeBaseResource())
}
