// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataIngestionJobType DataIngestionJob type.
type DataIngestionJobType struct {

	// Type of ingestionJob.
	Type DataIngestionJobTypeTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m DataIngestionJobType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataIngestionJobType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataIngestionJobTypeTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDataIngestionJobTypeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataIngestionJobTypeTypeEnum Enum with underlying type: string
type DataIngestionJobTypeTypeEnum string

// Set of constants representing the allowable values for DataIngestionJobTypeTypeEnum
const (
	DataIngestionJobTypeTypeCleanup   DataIngestionJobTypeTypeEnum = "CLEANUP"
	DataIngestionJobTypeTypeIngestion DataIngestionJobTypeTypeEnum = "INGESTION"
)

var mappingDataIngestionJobTypeTypeEnum = map[string]DataIngestionJobTypeTypeEnum{
	"CLEANUP":   DataIngestionJobTypeTypeCleanup,
	"INGESTION": DataIngestionJobTypeTypeIngestion,
}

var mappingDataIngestionJobTypeTypeEnumLowerCase = map[string]DataIngestionJobTypeTypeEnum{
	"cleanup":   DataIngestionJobTypeTypeCleanup,
	"ingestion": DataIngestionJobTypeTypeIngestion,
}

// GetDataIngestionJobTypeTypeEnumValues Enumerates the set of values for DataIngestionJobTypeTypeEnum
func GetDataIngestionJobTypeTypeEnumValues() []DataIngestionJobTypeTypeEnum {
	values := make([]DataIngestionJobTypeTypeEnum, 0)
	for _, v := range mappingDataIngestionJobTypeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataIngestionJobTypeTypeEnumStringValues Enumerates the set of values in String for DataIngestionJobTypeTypeEnum
func GetDataIngestionJobTypeTypeEnumStringValues() []string {
	return []string{
		"CLEANUP",
		"INGESTION",
	}
}

// GetMappingDataIngestionJobTypeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataIngestionJobTypeTypeEnum(val string) (DataIngestionJobTypeTypeEnum, bool) {
	enum, ok := mappingDataIngestionJobTypeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
