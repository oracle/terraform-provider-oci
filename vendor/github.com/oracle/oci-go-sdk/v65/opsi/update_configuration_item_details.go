// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateConfigurationItemDetails Configuration item details for OPSI configuration update.
type UpdateConfigurationItemDetails interface {
}

type updateconfigurationitemdetails struct {
	JsonData       []byte
	ConfigItemType string `json:"configItemType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigurationitemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigurationitemdetails updateconfigurationitemdetails
	s := struct {
		Model Unmarshalerupdateconfigurationitemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigItemType = s.Model.ConfigItemType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigurationitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigItemType {
	case "BASIC":
		mm := UpdateBasicConfigurationItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateConfigurationItemDetails: %s.", m.ConfigItemType)
		return *m, nil
	}
}

func (m updateconfigurationitemdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconfigurationitemdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
