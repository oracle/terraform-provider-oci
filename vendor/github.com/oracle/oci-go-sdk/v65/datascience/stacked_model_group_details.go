// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StackedModelGroupDetails Stacked model group type.
type StackedModelGroupDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model in the group that represents the base model for stacked deployment.
	BaseModelId *string `mandatory:"true" json:"baseModelId"`

	// An array of custom metadata details for the model group.
	CustomMetadataList []CustomMetadata `mandatory:"false" json:"customMetadataList"`
}

// GetCustomMetadataList returns CustomMetadataList
func (m StackedModelGroupDetails) GetCustomMetadataList() []CustomMetadata {
	return m.CustomMetadataList
}

func (m StackedModelGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackedModelGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StackedModelGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStackedModelGroupDetails StackedModelGroupDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeStackedModelGroupDetails
	}{
		"STACKED",
		(MarshalTypeStackedModelGroupDetails)(m),
	}

	return json.Marshal(&s)
}
