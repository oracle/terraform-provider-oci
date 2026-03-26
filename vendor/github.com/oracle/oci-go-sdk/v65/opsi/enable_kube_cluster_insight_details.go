// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// EnableKubeClusterInsightDetails The information about the Kubernetes cluster to be analyzed.
type EnableKubeClusterInsightDetails interface {
}

type enablekubeclusterinsightdetails struct {
	JsonData     []byte
	EntitySource string `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *enablekubeclusterinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerenablekubeclusterinsightdetails enablekubeclusterinsightdetails
	s := struct {
		Model Unmarshalerenablekubeclusterinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *enablekubeclusterinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "OKE_CLUSTER":
		mm := EnableOkeKubeClusterInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for EnableKubeClusterInsightDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

func (m enablekubeclusterinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m enablekubeclusterinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
