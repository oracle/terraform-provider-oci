// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateAgent            OperationTypeEnum = "CREATE_AGENT"
	OperationTypeUpdateAgent            OperationTypeEnum = "UPDATE_AGENT"
	OperationTypeDeleteAgent            OperationTypeEnum = "DELETE_AGENT"
	OperationTypeMoveAgent              OperationTypeEnum = "MOVE_AGENT"
	OperationTypeCreateAgentEndpoint    OperationTypeEnum = "CREATE_AGENT_ENDPOINT"
	OperationTypeUpdateAgentEndpoint    OperationTypeEnum = "UPDATE_AGENT_ENDPOINT"
	OperationTypeDeleteAgentEndpoint    OperationTypeEnum = "DELETE_AGENT_ENDPOINT"
	OperationTypeMoveAgentEndpoint      OperationTypeEnum = "MOVE_AGENT_ENDPOINT"
	OperationTypeCreateDataSource       OperationTypeEnum = "CREATE_DATA_SOURCE"
	OperationTypeUpdateDataSource       OperationTypeEnum = "UPDATE_DATA_SOURCE"
	OperationTypeDeleteDataSource       OperationTypeEnum = "DELETE_DATA_SOURCE"
	OperationTypeCreateKnowledgeBase    OperationTypeEnum = "CREATE_KNOWLEDGE_BASE"
	OperationTypeUpdateKnowledgeBase    OperationTypeEnum = "UPDATE_KNOWLEDGE_BASE"
	OperationTypeDeleteKnowledgeBase    OperationTypeEnum = "DELETE_KNOWLEDGE_BASE"
	OperationTypeMoveKnowledgeBase      OperationTypeEnum = "MOVE_KNOWLEDGE_BASE"
	OperationTypeCreateDataIngestionJob OperationTypeEnum = "CREATE_DATA_INGESTION_JOB"
	OperationTypeDeleteDataIngestionJob OperationTypeEnum = "DELETE_DATA_INGESTION_JOB"
	OperationTypeCreateTool             OperationTypeEnum = "CREATE_TOOL"
	OperationTypeUpdateTool             OperationTypeEnum = "UPDATE_TOOL"
	OperationTypeDeleteTool             OperationTypeEnum = "DELETE_TOOL"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_AGENT":              OperationTypeCreateAgent,
	"UPDATE_AGENT":              OperationTypeUpdateAgent,
	"DELETE_AGENT":              OperationTypeDeleteAgent,
	"MOVE_AGENT":                OperationTypeMoveAgent,
	"CREATE_AGENT_ENDPOINT":     OperationTypeCreateAgentEndpoint,
	"UPDATE_AGENT_ENDPOINT":     OperationTypeUpdateAgentEndpoint,
	"DELETE_AGENT_ENDPOINT":     OperationTypeDeleteAgentEndpoint,
	"MOVE_AGENT_ENDPOINT":       OperationTypeMoveAgentEndpoint,
	"CREATE_DATA_SOURCE":        OperationTypeCreateDataSource,
	"UPDATE_DATA_SOURCE":        OperationTypeUpdateDataSource,
	"DELETE_DATA_SOURCE":        OperationTypeDeleteDataSource,
	"CREATE_KNOWLEDGE_BASE":     OperationTypeCreateKnowledgeBase,
	"UPDATE_KNOWLEDGE_BASE":     OperationTypeUpdateKnowledgeBase,
	"DELETE_KNOWLEDGE_BASE":     OperationTypeDeleteKnowledgeBase,
	"MOVE_KNOWLEDGE_BASE":       OperationTypeMoveKnowledgeBase,
	"CREATE_DATA_INGESTION_JOB": OperationTypeCreateDataIngestionJob,
	"DELETE_DATA_INGESTION_JOB": OperationTypeDeleteDataIngestionJob,
	"CREATE_TOOL":               OperationTypeCreateTool,
	"UPDATE_TOOL":               OperationTypeUpdateTool,
	"DELETE_TOOL":               OperationTypeDeleteTool,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_agent":              OperationTypeCreateAgent,
	"update_agent":              OperationTypeUpdateAgent,
	"delete_agent":              OperationTypeDeleteAgent,
	"move_agent":                OperationTypeMoveAgent,
	"create_agent_endpoint":     OperationTypeCreateAgentEndpoint,
	"update_agent_endpoint":     OperationTypeUpdateAgentEndpoint,
	"delete_agent_endpoint":     OperationTypeDeleteAgentEndpoint,
	"move_agent_endpoint":       OperationTypeMoveAgentEndpoint,
	"create_data_source":        OperationTypeCreateDataSource,
	"update_data_source":        OperationTypeUpdateDataSource,
	"delete_data_source":        OperationTypeDeleteDataSource,
	"create_knowledge_base":     OperationTypeCreateKnowledgeBase,
	"update_knowledge_base":     OperationTypeUpdateKnowledgeBase,
	"delete_knowledge_base":     OperationTypeDeleteKnowledgeBase,
	"move_knowledge_base":       OperationTypeMoveKnowledgeBase,
	"create_data_ingestion_job": OperationTypeCreateDataIngestionJob,
	"delete_data_ingestion_job": OperationTypeDeleteDataIngestionJob,
	"create_tool":               OperationTypeCreateTool,
	"update_tool":               OperationTypeUpdateTool,
	"delete_tool":               OperationTypeDeleteTool,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_AGENT",
		"UPDATE_AGENT",
		"DELETE_AGENT",
		"MOVE_AGENT",
		"CREATE_AGENT_ENDPOINT",
		"UPDATE_AGENT_ENDPOINT",
		"DELETE_AGENT_ENDPOINT",
		"MOVE_AGENT_ENDPOINT",
		"CREATE_DATA_SOURCE",
		"UPDATE_DATA_SOURCE",
		"DELETE_DATA_SOURCE",
		"CREATE_KNOWLEDGE_BASE",
		"UPDATE_KNOWLEDGE_BASE",
		"DELETE_KNOWLEDGE_BASE",
		"MOVE_KNOWLEDGE_BASE",
		"CREATE_DATA_INGESTION_JOB",
		"DELETE_DATA_INGESTION_JOB",
		"CREATE_TOOL",
		"UPDATE_TOOL",
		"DELETE_TOOL",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
