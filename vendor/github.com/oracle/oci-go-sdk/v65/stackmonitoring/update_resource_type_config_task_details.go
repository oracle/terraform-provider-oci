// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateResourceTypeConfigTaskDetails Update the availability metrics and the handler configuration for
// the telegraf/collectd resource types.
type UpdateResourceTypeConfigTaskDetails struct {

	// A collection of resource type configuration details. User can provide
	// availability proxy metrics list for resource types along with the
	// telegraf/collectd handler configuration for the resource types.
	ResourceTypesConfiguration []ResourceTypeConfigDetails `mandatory:"true" json:"resourceTypesConfiguration"`

	// Type of the handler.
	HandlerType HandlerTypeEnum `mandatory:"true" json:"handlerType"`
}

func (m UpdateResourceTypeConfigTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateResourceTypeConfigTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHandlerTypeEnum(string(m.HandlerType)); !ok && m.HandlerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HandlerType: %s. Supported values are: %s.", m.HandlerType, strings.Join(GetHandlerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateResourceTypeConfigTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateResourceTypeConfigTaskDetails UpdateResourceTypeConfigTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateResourceTypeConfigTaskDetails
	}{
		"UPDATE_RESOURCE_TYPE_CONFIGS",
		(MarshalTypeUpdateResourceTypeConfigTaskDetails)(m),
	}

	return json.Marshal(&s)
}
