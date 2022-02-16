// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UnifiedAgentWindowsEventSource Windows events log source object.
type UnifiedAgentWindowsEventSource struct {

	// unique name for the source
	Name *string `mandatory:"true" json:"name"`

	Channels []string `mandatory:"false" json:"channels"`
}

//GetName returns Name
func (m UnifiedAgentWindowsEventSource) GetName() *string {
	return m.Name
}

func (m UnifiedAgentWindowsEventSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentWindowsEventSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentWindowsEventSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentWindowsEventSource UnifiedAgentWindowsEventSource
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUnifiedAgentWindowsEventSource
	}{
		"WINDOWS_EVENT_LOG",
		(MarshalTypeUnifiedAgentWindowsEventSource)(m),
	}

	return json.Marshal(&s)
}
