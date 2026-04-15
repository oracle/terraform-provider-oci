// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadataResourceConfiguration Configuration for EXADATA
type ExadataResourceConfiguration struct {

	// The type of node in the Exadata system.
	NodeType *string `mandatory:"true" json:"nodeType"`

	// The type of usage for the resource.
	UsageType *string `mandatory:"true" json:"usageType"`

	// The hardware generation of the Exadata system.
	HwGeneration *string `mandatory:"false" json:"hwGeneration"`
}

func (m ExadataResourceConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataResourceConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExadataResourceConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExadataResourceConfiguration ExadataResourceConfiguration
	s := struct {
		DiscriminatorParam string `json:"resource"`
		MarshalTypeExadataResourceConfiguration
	}{
		"EXADATA",
		(MarshalTypeExadataResourceConfiguration)(m),
	}

	return json.Marshal(&s)
}
