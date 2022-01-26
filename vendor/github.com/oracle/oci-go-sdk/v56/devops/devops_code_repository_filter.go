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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingDevopsCodeRepositoryFilterEvents = map[string]DevopsCodeRepositoryFilterEventsEnum{
	"PUSH": DevopsCodeRepositoryFilterEventsPush,
}

// GetDevopsCodeRepositoryFilterEventsEnumValues Enumerates the set of values for DevopsCodeRepositoryFilterEventsEnum
func GetDevopsCodeRepositoryFilterEventsEnumValues() []DevopsCodeRepositoryFilterEventsEnum {
	values := make([]DevopsCodeRepositoryFilterEventsEnum, 0)
	for _, v := range mappingDevopsCodeRepositoryFilterEvents {
		values = append(values, v)
	}
	return values
}
