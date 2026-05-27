// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CapabilityDetails Details about the supported type of resource.
type CapabilityDetails struct {

	// The service that the resource is part of.
	Service *string `mandatory:"true" json:"service"`

	// The type of resource.
	Name *string `mandatory:"true" json:"name"`

	AdditionalDetails AdditionalCapabilityDetails `mandatory:"false" json:"additionalDetails"`
}

func (m CapabilityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CapabilityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CapabilityDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AdditionalDetails additionalcapabilitydetails `json:"additionalDetails"`
		Service           *string                     `json:"service"`
		Name              *string                     `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.AdditionalDetails.UnmarshalPolymorphicJSON(model.AdditionalDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AdditionalDetails = nn.(AdditionalCapabilityDetails)
	} else {
		m.AdditionalDetails = nil
	}

	m.Service = model.Service

	m.Name = model.Name

	return
}
