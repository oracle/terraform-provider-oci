// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Model Models are mathematical representations of the relationships between data. Models are represented by their associated metadata and artifacts.
type Model struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Details about the lifecycle state of the model.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2019-08-25T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The OCID of the model version set that the model is associated to.
	ModelVersionSetId *string `mandatory:"true" json:"modelVersionSetId"`

	// The name of the model version set that the model is associated to.
	ModelVersionSetName *string `mandatory:"true" json:"modelVersionSetName"`

	// Unique identifier assigned to each version of the model.
	VersionId *int64 `mandatory:"true" json:"versionId"`

	// The version label can add an additional description of the lifecycle state of the model or the application using and training the model.
	VersionLabel *string `mandatory:"true" json:"versionLabel"`

	// The category of the model.
	Category ModelCategoryEnum `mandatory:"true" json:"category"`

	// Identifier to indicate whether a model artifact resides in the Service Tenancy or Customer Tenancy.
	IsModelByReference *bool `mandatory:"true" json:"isModelByReference"`

	RetentionSetting *RetentionSetting `mandatory:"true" json:"retentionSetting"`

	BackupSetting *BackupSetting `mandatory:"true" json:"backupSetting"`

	RetentionOperationDetails *RetentionOperationDetails `mandatory:"true" json:"retentionOperationDetails"`

	BackupOperationDetails *BackupOperationDetails `mandatory:"true" json:"backupOperationDetails"`

	// A short description of the model.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of custom metadata details for the model.
	CustomMetadataList []Metadata `mandatory:"false" json:"customMetadataList"`

	// An array of defined metadata details for the model.
	DefinedMetadataList []Metadata `mandatory:"false" json:"definedMetadataList"`

	// Input schema file content in String format
	InputSchema *string `mandatory:"false" json:"inputSchema"`

	// Output schema file content in String format
	OutputSchema *string `mandatory:"false" json:"outputSchema"`
}

func (m Model) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Model) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModelLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetModelLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingModelCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetModelCategoryEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
