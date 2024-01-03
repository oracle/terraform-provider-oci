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

// UnifiedAgentLoggingGrepFilter Logging grep filter object greps events by the values of specified fields.
// Ref: https://docs.fluentd.org/filter/grep
type UnifiedAgentLoggingGrepFilter struct {

	// Unique name for the filter.
	Name *string `mandatory:"true" json:"name"`

	// A list of filtering rules to include logs
	AllowList []GrepFilterAllowRule `mandatory:"false" json:"allowList"`

	// A list of filtering rules to reject logs
	DenyList []GrepFilterDenyRule `mandatory:"false" json:"denyList"`
}

// GetName returns Name
func (m UnifiedAgentLoggingGrepFilter) GetName() *string {
	return m.Name
}

func (m UnifiedAgentLoggingGrepFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentLoggingGrepFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentLoggingGrepFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentLoggingGrepFilter UnifiedAgentLoggingGrepFilter
	s := struct {
		DiscriminatorParam string `json:"filterType"`
		MarshalTypeUnifiedAgentLoggingGrepFilter
	}{
		"GREP_FILTER",
		(MarshalTypeUnifiedAgentLoggingGrepFilter)(m),
	}

	return json.Marshal(&s)
}
