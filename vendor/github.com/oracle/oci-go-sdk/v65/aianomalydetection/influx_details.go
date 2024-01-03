// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InfluxDetails Possible data sources
type InfluxDetails interface {
}

type influxdetails struct {
	JsonData      []byte
	InfluxVersion string `json:"influxVersion"`
}

// UnmarshalJSON unmarshals json
func (m *influxdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinfluxdetails influxdetails
	s := struct {
		Model Unmarshalerinfluxdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InfluxVersion = s.Model.InfluxVersion

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *influxdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InfluxVersion {
	case "V_1_8":
		mm := InfluxDetailsV1v8{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "V_2_0":
		mm := InfluxDetailsV2v0{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for InfluxDetails: %s.", m.InfluxVersion)
		return *m, nil
	}
}

func (m influxdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m influxdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
