// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDigitalTwinRelationshipDetails The information about new digital twin relationship to be created.
type CreateDigitalTwinRelationshipDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
	IotDomainId *string `mandatory:"true" json:"iotDomainId"`

	// Its the name of the relationship that links two digital twin instances. Here, it is the relationship name of the source digital twin model.
	ContentPath *string `mandatory:"true" json:"contentPath"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance.
	SourceDigitalTwinInstanceId *string `mandatory:"true" json:"sourceDigitalTwinInstanceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.
	TargetDigitalTwinInstanceId *string `mandatory:"true" json:"targetDigitalTwinInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The value(s) of the relationship properties defined in the source digital twin model.
	Content map[string]interface{} `mandatory:"false" json:"content"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateDigitalTwinRelationshipDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDigitalTwinRelationshipDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
