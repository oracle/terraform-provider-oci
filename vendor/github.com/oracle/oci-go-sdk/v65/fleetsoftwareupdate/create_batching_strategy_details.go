// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBatchingStrategyDetails Batching strategy details to use during PRECHECK and APPLY Cycle Actions.
type CreateBatchingStrategyDetails interface {
}

type createbatchingstrategydetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createbatchingstrategydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebatchingstrategydetails createbatchingstrategydetails
	s := struct {
		Model Unmarshalercreatebatchingstrategydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbatchingstrategydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SEQUENTIAL":
		mm := CreateSequentialBatchingStrategyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIFTY_FIFTY":
		mm := CreateFiftyFiftyBatchingStrategyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SERVICE_AVAILABILITY_FACTOR":
		mm := CreateServiceAvailabilityFactorBatchingStrategyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NON_ROLLING":
		mm := CreateNonRollingBatchingStrategyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateBatchingStrategyDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createbatchingstrategydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createbatchingstrategydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
