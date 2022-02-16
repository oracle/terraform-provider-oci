// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ImportableAgentEntitySummary An agent entity that can be imported into Operations Insights.
type ImportableAgentEntitySummary interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	GetManagementAgentId() *string

	// The Display Name (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Display) of the Management Agent
	GetManagementAgentDisplayName() *string
}

type importableagententitysummary struct {
	JsonData                   []byte
	ManagementAgentId          *string `mandatory:"true" json:"managementAgentId"`
	ManagementAgentDisplayName *string `mandatory:"true" json:"managementAgentDisplayName"`
	EntitySource               string  `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *importableagententitysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerimportableagententitysummary importableagententitysummary
	s := struct {
		Model Unmarshalerimportableagententitysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ManagementAgentId = s.Model.ManagementAgentId
	m.ManagementAgentDisplayName = s.Model.ManagementAgentDisplayName
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *importableagententitysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := HostImportableAgentEntitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetManagementAgentId returns ManagementAgentId
func (m importableagententitysummary) GetManagementAgentId() *string {
	return m.ManagementAgentId
}

//GetManagementAgentDisplayName returns ManagementAgentDisplayName
func (m importableagententitysummary) GetManagementAgentDisplayName() *string {
	return m.ManagementAgentDisplayName
}

func (m importableagententitysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m importableagententitysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
