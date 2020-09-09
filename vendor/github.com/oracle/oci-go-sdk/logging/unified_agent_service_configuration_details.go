// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UnifiedAgentServiceConfigurationDetails Top level Unified Agent service configuration object.
type UnifiedAgentServiceConfigurationDetails interface {
}

type unifiedagentserviceconfigurationdetails struct {
	JsonData          []byte
	ConfigurationType string `json:"configurationType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentserviceconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentserviceconfigurationdetails unifiedagentserviceconfigurationdetails
	s := struct {
		Model Unmarshalerunifiedagentserviceconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigurationType = s.Model.ConfigurationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentserviceconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigurationType {
	case "LOGGING":
		mm := UnifiedAgentLoggingConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m unifiedagentserviceconfigurationdetails) String() string {
	return common.PointerString(m)
}
