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

// UnifiedAgentMonitoringConfigurationDetails Unified Agent monitoring service configuration object.
type UnifiedAgentMonitoringConfigurationDetails struct {

	// Unified Agent monitoring application configuration details.
	ApplicationConfigurations []UnifiedAgentMonitoringApplicationConfigurationDetails `mandatory:"true" json:"applicationConfigurations"`
}

func (m UnifiedAgentMonitoringConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentMonitoringConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentMonitoringConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentMonitoringConfigurationDetails UnifiedAgentMonitoringConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"configurationType"`
		MarshalTypeUnifiedAgentMonitoringConfigurationDetails
	}{
		"MONITORING",
		(MarshalTypeUnifiedAgentMonitoringConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAgentMonitoringConfigurationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ApplicationConfigurations []unifiedagentmonitoringapplicationconfigurationdetails `json:"applicationConfigurations"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ApplicationConfigurations = make([]UnifiedAgentMonitoringApplicationConfigurationDetails, len(model.ApplicationConfigurations))
	for i, n := range model.ApplicationConfigurations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ApplicationConfigurations[i] = nn.(UnifiedAgentMonitoringApplicationConfigurationDetails)
		} else {
			m.ApplicationConfigurations[i] = nil
		}
	}
	return
}
