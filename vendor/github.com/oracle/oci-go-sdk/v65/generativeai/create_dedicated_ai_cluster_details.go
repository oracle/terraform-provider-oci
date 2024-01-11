// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service API
//
// **Generative AI Service**
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable LLMs that cover a wide range of use cases for text generation. Use the playground to try out the models out-of-the-box or create and host your own fine-tuned custom models based on your own data on dedicated AI clusters.
//

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDedicatedAiClusterDetails **CreateDedicatedAiClusterDetails**
// The data to create a dedicated AI cluster.
type CreateDedicatedAiClusterDetails struct {

	// dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.
	Type DedicatedAiClusterTypeEnum `mandatory:"true" json:"type"`

	// The compartment OCID to create the dedicated AI cluster in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of dedicated units in this AI cluster.
	UnitCount *int `mandatory:"true" json:"unitCount"`

	// The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.
	UnitShape DedicatedAiClusterUnitShapeEnum `mandatory:"true" json:"unitShape"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description of the dedicated AI cluster.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDedicatedAiClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDedicatedAiClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedAiClusterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDedicatedAiClusterTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDedicatedAiClusterUnitShapeEnum(string(m.UnitShape)); !ok && m.UnitShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnitShape: %s. Supported values are: %s.", m.UnitShape, strings.Join(GetDedicatedAiClusterUnitShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
