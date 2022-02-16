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

// Filter The filters for the trigger.
type Filter interface {
}

type filter struct {
	JsonData      []byte
	TriggerSource string `json:"triggerSource"`
}

// UnmarshalJSON unmarshals json
func (m *filter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfilter filter
	s := struct {
		Model Unmarshalerfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TriggerSource = s.Model.TriggerSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *filter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerSource {
	case "DEVOPS_CODE_REPOSITORY":
		mm := DevopsCodeRepositoryFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB":
		mm := GitlabFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITHUB":
		mm := GithubFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m filter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m filter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
