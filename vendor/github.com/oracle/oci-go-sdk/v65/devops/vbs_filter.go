// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VbsFilter The filter for VBS events.
type VbsFilter struct {
	Include *VbsFilterAttributes `mandatory:"false" json:"include"`

	Exclude *VbsFilterExclusionAttributes `mandatory:"false" json:"exclude"`

	// The events, for example, PUSH, PULL_REQUEST_MERGE.
	Events []VbsFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m VbsFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbsFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingVbsFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetVbsFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VbsFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVbsFilter VbsFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeVbsFilter
	}{
		"VBS",
		(MarshalTypeVbsFilter)(m),
	}

	return json.Marshal(&s)
}

// VbsFilterEventsEnum Enum with underlying type: string
type VbsFilterEventsEnum string

// Set of constants representing the allowable values for VbsFilterEventsEnum
const (
	VbsFilterEventsPush                VbsFilterEventsEnum = "PUSH"
	VbsFilterEventsMergeRequestCreated VbsFilterEventsEnum = "MERGE_REQUEST_CREATED"
	VbsFilterEventsMergeRequestUpdated VbsFilterEventsEnum = "MERGE_REQUEST_UPDATED"
	VbsFilterEventsMergeRequestMerged  VbsFilterEventsEnum = "MERGE_REQUEST_MERGED"
)

var mappingVbsFilterEventsEnum = map[string]VbsFilterEventsEnum{
	"PUSH":                  VbsFilterEventsPush,
	"MERGE_REQUEST_CREATED": VbsFilterEventsMergeRequestCreated,
	"MERGE_REQUEST_UPDATED": VbsFilterEventsMergeRequestUpdated,
	"MERGE_REQUEST_MERGED":  VbsFilterEventsMergeRequestMerged,
}

var mappingVbsFilterEventsEnumLowerCase = map[string]VbsFilterEventsEnum{
	"push":                  VbsFilterEventsPush,
	"merge_request_created": VbsFilterEventsMergeRequestCreated,
	"merge_request_updated": VbsFilterEventsMergeRequestUpdated,
	"merge_request_merged":  VbsFilterEventsMergeRequestMerged,
}

// GetVbsFilterEventsEnumValues Enumerates the set of values for VbsFilterEventsEnum
func GetVbsFilterEventsEnumValues() []VbsFilterEventsEnum {
	values := make([]VbsFilterEventsEnum, 0)
	for _, v := range mappingVbsFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetVbsFilterEventsEnumStringValues Enumerates the set of values in String for VbsFilterEventsEnum
func GetVbsFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
		"MERGE_REQUEST_CREATED",
		"MERGE_REQUEST_UPDATED",
		"MERGE_REQUEST_MERGED",
	}
}

// GetMappingVbsFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVbsFilterEventsEnum(val string) (VbsFilterEventsEnum, bool) {
	enum, ok := mappingVbsFilterEventsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
