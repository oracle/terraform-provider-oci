// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnifiedAgentMonitoringApplicationConfigurationDetails Unified Agent monitoing application configuration details
type UnifiedAgentMonitoringApplicationConfigurationDetails interface {
}

type unifiedagentmonitoringapplicationconfigurationdetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *unifiedagentmonitoringapplicationconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerunifiedagentmonitoringapplicationconfigurationdetails unifiedagentmonitoringapplicationconfigurationdetails
	s := struct {
		Model Unmarshalerunifiedagentmonitoringapplicationconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *unifiedagentmonitoringapplicationconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "KUBERNETES":
		mm := UnifiedAgentKubernetesConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "URL":
		mm := UnifiedAgentUrlConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TAIL":
		mm := UnifiedAgentOpenmetricsTailConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UnifiedAgentMonitoringApplicationConfigurationDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m unifiedagentmonitoringapplicationconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m unifiedagentmonitoringapplicationconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
