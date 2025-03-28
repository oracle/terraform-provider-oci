// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableExternalNonContainerDatabaseManagementFeatureDetails The details required to enable a Database Management feature for an external non-container database.
type EnableExternalNonContainerDatabaseManagementFeatureDetails struct {
	FeatureDetails ExternalDatabaseFeatureDetails `mandatory:"true" json:"featureDetails"`
}

func (m EnableExternalNonContainerDatabaseManagementFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableExternalNonContainerDatabaseManagementFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EnableExternalNonContainerDatabaseManagementFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FeatureDetails externaldatabasefeaturedetails `json:"featureDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.FeatureDetails.UnmarshalPolymorphicJSON(model.FeatureDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FeatureDetails = nn.(ExternalDatabaseFeatureDetails)
	} else {
		m.FeatureDetails = nil
	}

	return
}
