// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DevopsCodeRepositoryFilter The filter for GitLab events.
type DevopsCodeRepositoryFilter struct {
	Include *DevopsCodeRepositoryFilterAttributes `mandatory:"false" json:"include"`

	// The events only support PUSH.
	Events []DevopsCodeRepositoryFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m DevopsCodeRepositoryFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DevopsCodeRepositoryFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Events {
		if _, ok := GetMappingDevopsCodeRepositoryFilterEventsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Events: %s. Supported values are: %s.", val, strings.Join(GetDevopsCodeRepositoryFilterEventsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DevopsCodeRepositoryFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDevopsCodeRepositoryFilter DevopsCodeRepositoryFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeDevopsCodeRepositoryFilter
	}{
		"DEVOPS_CODE_REPOSITORY",
		(MarshalTypeDevopsCodeRepositoryFilter)(m),
	}

	return json.Marshal(&s)
}

// DevopsCodeRepositoryFilterEventsEnum Enum with underlying type: string
type DevopsCodeRepositoryFilterEventsEnum string

// Set of constants representing the allowable values for DevopsCodeRepositoryFilterEventsEnum
const (
	DevopsCodeRepositoryFilterEventsPush DevopsCodeRepositoryFilterEventsEnum = "PUSH"
)

var mappingDevopsCodeRepositoryFilterEventsEnum = map[string]DevopsCodeRepositoryFilterEventsEnum{
	"PUSH": DevopsCodeRepositoryFilterEventsPush,
}

// GetDevopsCodeRepositoryFilterEventsEnumValues Enumerates the set of values for DevopsCodeRepositoryFilterEventsEnum
func GetDevopsCodeRepositoryFilterEventsEnumValues() []DevopsCodeRepositoryFilterEventsEnum {
	values := make([]DevopsCodeRepositoryFilterEventsEnum, 0)
	for _, v := range mappingDevopsCodeRepositoryFilterEventsEnum {
		values = append(values, v)
	}
	return values
}

// GetDevopsCodeRepositoryFilterEventsEnumStringValues Enumerates the set of values in String for DevopsCodeRepositoryFilterEventsEnum
func GetDevopsCodeRepositoryFilterEventsEnumStringValues() []string {
	return []string{
		"PUSH",
	}
}

// GetMappingDevopsCodeRepositoryFilterEventsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDevopsCodeRepositoryFilterEventsEnum(val string) (DevopsCodeRepositoryFilterEventsEnum, bool) {
	mappingDevopsCodeRepositoryFilterEventsEnumIgnoreCase := make(map[string]DevopsCodeRepositoryFilterEventsEnum)
	for k, v := range mappingDevopsCodeRepositoryFilterEventsEnum {
		mappingDevopsCodeRepositoryFilterEventsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDevopsCodeRepositoryFilterEventsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
