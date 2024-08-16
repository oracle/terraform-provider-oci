// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalPluggableDatabaseFeatureDetails The details required to enable the specified Database Management feature.
type ExternalPluggableDatabaseFeatureDetails interface {
	GetConnectorDetails() ConnectorDetails
}

type externalpluggabledatabasefeaturedetails struct {
	JsonData         []byte
	ConnectorDetails connectordetails `mandatory:"true" json:"connectorDetails"`
	Feature          string           `json:"feature"`
}

// UnmarshalJSON unmarshals json
func (m *externalpluggabledatabasefeaturedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternalpluggabledatabasefeaturedetails externalpluggabledatabasefeaturedetails
	s := struct {
		Model Unmarshalerexternalpluggabledatabasefeaturedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectorDetails = s.Model.ConnectorDetails
	m.Feature = s.Model.Feature

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externalpluggabledatabasefeaturedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Feature {
	case "DIAGNOSTICS_AND_MANAGEMENT":
		mm := ExternalPluggableDatabaseDiagnosticsAndManagementFeatureDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_LIFECYCLE_MANAGEMENT":
		mm := ExternalPluggableDatabaseLifecycleManagementFeatureDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQLWATCH":
		mm := ExternalPluggableDatabaseSqlWatchFeatureDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExternalPluggableDatabaseFeatureDetails: %s.", m.Feature)
		return *m, nil
	}
}

// GetConnectorDetails returns ConnectorDetails
func (m externalpluggabledatabasefeaturedetails) GetConnectorDetails() connectordetails {
	return m.ConnectorDetails
}

func (m externalpluggabledatabasefeaturedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externalpluggabledatabasefeaturedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
