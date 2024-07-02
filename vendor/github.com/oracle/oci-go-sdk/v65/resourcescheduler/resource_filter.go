// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceFilter This is a generic filter used to decide which resources that the schedule be applied to.
type ResourceFilter interface {
}

type resourcefilter struct {
	JsonData  []byte
	Attribute string `json:"attribute"`
}

// UnmarshalJSON unmarshals json
func (m *resourcefilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresourcefilter resourcefilter
	s := struct {
		Model Unmarshalerresourcefilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Attribute = s.Model.Attribute

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resourcefilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Attribute {
	case "TIME_CREATED":
		mm := TimeCreatedResourceFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_TYPE":
		mm := ResourceTypeResourceFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIFECYCLE_STATE":
		mm := LifecycleStateResourceFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_ID":
		mm := CompartmentIdResourceFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEFINED_TAGS":
		mm := DefinedTagsResourceFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResourceFilter: %s.", m.Attribute)
		return *m, nil
	}
}

func (m resourcefilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resourcefilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceFilterAttributeEnum Enum with underlying type: string
type ResourceFilterAttributeEnum string

// Set of constants representing the allowable values for ResourceFilterAttributeEnum
const (
	ResourceFilterAttributeCompartmentId  ResourceFilterAttributeEnum = "COMPARTMENT_ID"
	ResourceFilterAttributeResourceType   ResourceFilterAttributeEnum = "RESOURCE_TYPE"
	ResourceFilterAttributeDefinedTags    ResourceFilterAttributeEnum = "DEFINED_TAGS"
	ResourceFilterAttributeTimeCreated    ResourceFilterAttributeEnum = "TIME_CREATED"
	ResourceFilterAttributeLifecycleState ResourceFilterAttributeEnum = "LIFECYCLE_STATE"
)

var mappingResourceFilterAttributeEnum = map[string]ResourceFilterAttributeEnum{
	"COMPARTMENT_ID":  ResourceFilterAttributeCompartmentId,
	"RESOURCE_TYPE":   ResourceFilterAttributeResourceType,
	"DEFINED_TAGS":    ResourceFilterAttributeDefinedTags,
	"TIME_CREATED":    ResourceFilterAttributeTimeCreated,
	"LIFECYCLE_STATE": ResourceFilterAttributeLifecycleState,
}

var mappingResourceFilterAttributeEnumLowerCase = map[string]ResourceFilterAttributeEnum{
	"compartment_id":  ResourceFilterAttributeCompartmentId,
	"resource_type":   ResourceFilterAttributeResourceType,
	"defined_tags":    ResourceFilterAttributeDefinedTags,
	"time_created":    ResourceFilterAttributeTimeCreated,
	"lifecycle_state": ResourceFilterAttributeLifecycleState,
}

// GetResourceFilterAttributeEnumValues Enumerates the set of values for ResourceFilterAttributeEnum
func GetResourceFilterAttributeEnumValues() []ResourceFilterAttributeEnum {
	values := make([]ResourceFilterAttributeEnum, 0)
	for _, v := range mappingResourceFilterAttributeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceFilterAttributeEnumStringValues Enumerates the set of values in String for ResourceFilterAttributeEnum
func GetResourceFilterAttributeEnumStringValues() []string {
	return []string{
		"COMPARTMENT_ID",
		"RESOURCE_TYPE",
		"DEFINED_TAGS",
		"TIME_CREATED",
		"LIFECYCLE_STATE",
	}
}

// GetMappingResourceFilterAttributeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceFilterAttributeEnum(val string) (ResourceFilterAttributeEnum, bool) {
	enum, ok := mappingResourceFilterAttributeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
