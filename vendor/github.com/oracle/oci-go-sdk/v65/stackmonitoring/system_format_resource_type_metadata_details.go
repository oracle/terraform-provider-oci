// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SystemFormatResourceTypeMetadataDetails The resource type metadata is defined in machine friendly format.
type SystemFormatResourceTypeMetadataDetails struct {

	// List of required properties for resource type.
	RequiredProperties []string `mandatory:"false" json:"requiredProperties"`

	// List of properties needed by the agent for monitoring the resource.
	// Valid only if resource type is OCI management agent based. When specified,
	// these properties are passed to the management agent during resource create or update.
	AgentProperties []string `mandatory:"false" json:"agentProperties"`

	// List of valid properties for resource type while creating the monitored resource.
	// If resources of this type specifies any other properties during create operation,
	// the operation will fail.
	ValidPropertiesForCreate []string `mandatory:"false" json:"validPropertiesForCreate"`

	// List of valid properties for resource type while updating the monitored resource.
	// If resources of this type specifies any other properties during update operation,
	// the operation will fail.
	ValidPropertiesForUpdate []string `mandatory:"false" json:"validPropertiesForUpdate"`

	// List of property sets used to uniquely identify the resources.
	// This check is made during create or update of stack monitoring resource.
	// The resource has to pass unique check for each set in the list.
	// For example, database can have user, password and SID as one unique set.
	// Another unique set would be user, password and service name.
	UniquePropertySets []UniquePropertySet `mandatory:"false" json:"uniquePropertySets"`

	// List of valid values for the properties. This is useful when resource type wants to
	// restrict only certain values for some properties. For instance for 'osType' property,
	// supported values can be restricted to be either Linux or Windows.
	// Example: `{ "osType": ["Linux","Windows","Solaris"]}`
	ValidPropertyValues map[string][]string `mandatory:"false" json:"validPropertyValues"`
}

func (m SystemFormatResourceTypeMetadataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SystemFormatResourceTypeMetadataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SystemFormatResourceTypeMetadataDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSystemFormatResourceTypeMetadataDetails SystemFormatResourceTypeMetadataDetails
	s := struct {
		DiscriminatorParam string `json:"format"`
		MarshalTypeSystemFormatResourceTypeMetadataDetails
	}{
		"SYSTEM_FORMAT",
		(MarshalTypeSystemFormatResourceTypeMetadataDetails)(m),
	}

	return json.Marshal(&s)
}
