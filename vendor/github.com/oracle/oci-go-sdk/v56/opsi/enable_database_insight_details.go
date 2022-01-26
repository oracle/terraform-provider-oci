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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EnableDatabaseInsightDetails The information about database to be analyzed.
type EnableDatabaseInsightDetails interface {
}

type enabledatabaseinsightdetails struct {
	JsonData     []byte
	EntitySource string `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *enabledatabaseinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerenabledatabaseinsightdetails enabledatabaseinsightdetails
	s := struct {
		Model Unmarshalerenabledatabaseinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *enabledatabaseinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_DATABASE":
		mm := EnableEmManagedExternalDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m enabledatabaseinsightdetails) String() string {
	return common.PointerString(m)
}
