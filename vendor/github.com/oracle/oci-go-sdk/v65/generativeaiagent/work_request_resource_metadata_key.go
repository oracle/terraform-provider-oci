// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// WorkRequestResourceMetadataKeyEnum Enum with underlying type: string
type WorkRequestResourceMetadataKeyEnum string

// Set of constants representing the allowable values for WorkRequestResourceMetadataKeyEnum
const (
	WorkRequestResourceMetadataKeyIsDryRun WorkRequestResourceMetadataKeyEnum = "IS_DRY_RUN"
)

var mappingWorkRequestResourceMetadataKeyEnum = map[string]WorkRequestResourceMetadataKeyEnum{
	"IS_DRY_RUN": WorkRequestResourceMetadataKeyIsDryRun,
}

var mappingWorkRequestResourceMetadataKeyEnumLowerCase = map[string]WorkRequestResourceMetadataKeyEnum{
	"is_dry_run": WorkRequestResourceMetadataKeyIsDryRun,
}

// GetWorkRequestResourceMetadataKeyEnumValues Enumerates the set of values for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumValues() []WorkRequestResourceMetadataKeyEnum {
	values := make([]WorkRequestResourceMetadataKeyEnum, 0)
	for _, v := range mappingWorkRequestResourceMetadataKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceMetadataKeyEnumStringValues Enumerates the set of values in String for WorkRequestResourceMetadataKeyEnum
func GetWorkRequestResourceMetadataKeyEnumStringValues() []string {
	return []string{
		"IS_DRY_RUN",
	}
}

// GetMappingWorkRequestResourceMetadataKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceMetadataKeyEnum(val string) (WorkRequestResourceMetadataKeyEnum, bool) {
	enum, ok := mappingWorkRequestResourceMetadataKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
