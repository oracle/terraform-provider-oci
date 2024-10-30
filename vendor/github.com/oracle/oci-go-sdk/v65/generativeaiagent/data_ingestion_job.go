// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// **Generative AI Agents API**
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataIngestionJob **DataIngestionJob**
// When you create a data source, you specify the location of the data files. To make those files usable by an agent, you must download them into the agent's associated knowledge base, a process known as data ingestion. Data ingestion is a process that extracts data from data source documents, converts it into a structured format suitable for analysis, and then stores it in a knowledge base.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type DataIngestionJob struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DataIngestionJob.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent DataSource.
	DataSourceId *string `mandatory:"true" json:"dataSourceId"`

	DataIngestionJobStatistics *DataIngestionJobStatistics `mandatory:"true" json:"dataIngestionJobStatistics"`

	// The date and time the data ingestion job was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the data ingestion job.
	LifecycleState DataIngestionJobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the data ingestion job was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the data ingestion job in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DataIngestionJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataIngestionJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataIngestionJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataIngestionJobLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataIngestionJobLifecycleStateEnum Enum with underlying type: string
type DataIngestionJobLifecycleStateEnum string

// Set of constants representing the allowable values for DataIngestionJobLifecycleStateEnum
const (
	DataIngestionJobLifecycleStateAccepted   DataIngestionJobLifecycleStateEnum = "ACCEPTED"
	DataIngestionJobLifecycleStateInProgress DataIngestionJobLifecycleStateEnum = "IN_PROGRESS"
	DataIngestionJobLifecycleStateWaiting    DataIngestionJobLifecycleStateEnum = "WAITING"
	DataIngestionJobLifecycleStateFailed     DataIngestionJobLifecycleStateEnum = "FAILED"
	DataIngestionJobLifecycleStateSucceeded  DataIngestionJobLifecycleStateEnum = "SUCCEEDED"
	DataIngestionJobLifecycleStateDeleting   DataIngestionJobLifecycleStateEnum = "DELETING"
	DataIngestionJobLifecycleStateDeleted    DataIngestionJobLifecycleStateEnum = "DELETED"
)

var mappingDataIngestionJobLifecycleStateEnum = map[string]DataIngestionJobLifecycleStateEnum{
	"ACCEPTED":    DataIngestionJobLifecycleStateAccepted,
	"IN_PROGRESS": DataIngestionJobLifecycleStateInProgress,
	"WAITING":     DataIngestionJobLifecycleStateWaiting,
	"FAILED":      DataIngestionJobLifecycleStateFailed,
	"SUCCEEDED":   DataIngestionJobLifecycleStateSucceeded,
	"DELETING":    DataIngestionJobLifecycleStateDeleting,
	"DELETED":     DataIngestionJobLifecycleStateDeleted,
}

var mappingDataIngestionJobLifecycleStateEnumLowerCase = map[string]DataIngestionJobLifecycleStateEnum{
	"accepted":    DataIngestionJobLifecycleStateAccepted,
	"in_progress": DataIngestionJobLifecycleStateInProgress,
	"waiting":     DataIngestionJobLifecycleStateWaiting,
	"failed":      DataIngestionJobLifecycleStateFailed,
	"succeeded":   DataIngestionJobLifecycleStateSucceeded,
	"deleting":    DataIngestionJobLifecycleStateDeleting,
	"deleted":     DataIngestionJobLifecycleStateDeleted,
}

// GetDataIngestionJobLifecycleStateEnumValues Enumerates the set of values for DataIngestionJobLifecycleStateEnum
func GetDataIngestionJobLifecycleStateEnumValues() []DataIngestionJobLifecycleStateEnum {
	values := make([]DataIngestionJobLifecycleStateEnum, 0)
	for _, v := range mappingDataIngestionJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataIngestionJobLifecycleStateEnumStringValues Enumerates the set of values in String for DataIngestionJobLifecycleStateEnum
func GetDataIngestionJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingDataIngestionJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataIngestionJobLifecycleStateEnum(val string) (DataIngestionJobLifecycleStateEnum, bool) {
	enum, ok := mappingDataIngestionJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
