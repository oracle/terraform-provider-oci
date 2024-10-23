// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateModelDetails Details for updating a model.
type UpdateModelDetails struct {

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	//  Example: `My Model`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the model.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// An array of custom metadata details for the model.
	CustomMetadataList []Metadata `mandatory:"false" json:"customMetadataList"`

	// An array of defined metadata details for the model.
	DefinedMetadataList []Metadata `mandatory:"false" json:"definedMetadataList"`

	// The OCID of the model version set that the model is associated to.
	ModelVersionSetId *string `mandatory:"false" json:"modelVersionSetId"`

	// The version label can add an additional description of the lifecycle state of the model or the application using/training the model.
	VersionLabel *string `mandatory:"false" json:"versionLabel"`

	RetentionSetting *RetentionSetting `mandatory:"false" json:"retentionSetting"`

	BackupSetting *BackupSetting `mandatory:"false" json:"backupSetting"`
}

func (m UpdateModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
