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

// UnifiedAgentKubernetesConfigurationDetails Unified Agent kubernetes service configuration object.
type UnifiedAgentKubernetesConfigurationDetails struct {
	Source *UnifiedAgentKubernetesSource `mandatory:"true" json:"source"`

	Destination *UnifiedAgentMonitoringDestination `mandatory:"true" json:"destination"`

	Filter *UnifiedAgentKubernetesFilter `mandatory:"false" json:"filter"`
}

func (m UnifiedAgentKubernetesConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentKubernetesConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentKubernetesConfigurationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentKubernetesConfigurationDetails UnifiedAgentKubernetesConfigurationDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUnifiedAgentKubernetesConfigurationDetails
	}{
		"KUBERNETES",
		(MarshalTypeUnifiedAgentKubernetesConfigurationDetails)(m),
	}

	return json.Marshal(&s)
}
