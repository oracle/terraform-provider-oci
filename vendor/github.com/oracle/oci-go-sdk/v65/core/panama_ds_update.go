// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PanamaDsUpdate Panama data plane update.
type PanamaDsUpdate interface {

	// Unique identifier for the primary resource affected by this update,
	// such as its OCID.
	GetId() *string

	// True iff this update signals deletion of the identified resource.
	// If true, the type-specific fields of this object may be null.
	GetIsDelete() *bool

	// The date and time that the API call was made that led to this Update.
	// The date and time format is defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	GetTimeUpdated() *common.SDKTime
}

type panamadsupdate struct {
	JsonData    []byte
	Id          *string         `mandatory:"false" json:"id"`
	IsDelete    *bool           `mandatory:"false" json:"isDelete"`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
	Type        string          `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *panamadsupdate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpanamadsupdate panamadsupdate
	s := struct {
		Model Unmarshalerpanamadsupdate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.IsDelete = s.Model.IsDelete
	m.TimeUpdated = s.Model.TimeUpdated
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *panamadsupdate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DrgAttachmentUpdate":
		mm := DrgAttachmentUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgImportPolicyUpdate":
		mm := DrgImportPolicyUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NatGatewayUpdate":
		mm := NatGatewayUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgExportPolicyUpdate":
		mm := DrgExportPolicyUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgInetAttachUpdate":
		mm := DrgInetAttachUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgImportPolicyUpdateV2":
		mm := DrgImportPolicyUpdateV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgStaticRouteUpdate":
		mm := DrgStaticRouteUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgRouteLimitUpdate":
		mm := DrgRouteLimitUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DrgNatRuleUpdate":
		mm := DrgNatRuleUpdate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PanamaDSUpdate: %s.", m.Type)
		return *m, nil
	}
}

// GetId returns Id
func (m panamadsupdate) GetId() *string {
	return m.Id
}

// GetIsDelete returns IsDelete
func (m panamadsupdate) GetIsDelete() *bool {
	return m.IsDelete
}

// GetTimeUpdated returns TimeUpdated
func (m panamadsupdate) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m panamadsupdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m panamadsupdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
