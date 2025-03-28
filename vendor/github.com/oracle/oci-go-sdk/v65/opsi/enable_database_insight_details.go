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
	case "MDS_MYSQL_DATABASE_SYSTEM":
		mm := EnableMdsMySqlDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXTERNAL_MYSQL_DATABASE_SYSTEM":
		mm := EnableExternalMysqlDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTONOMOUS_DATABASE":
		mm := EnableAutonomousDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_CLOUD_DATABASE":
		mm := EnableMacsManagedCloudDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PE_COMANAGED_DATABASE":
		mm := EnablePeComanagedDatabaseInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for EnableDatabaseInsightDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

func (m enabledatabaseinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m enabledatabaseinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
