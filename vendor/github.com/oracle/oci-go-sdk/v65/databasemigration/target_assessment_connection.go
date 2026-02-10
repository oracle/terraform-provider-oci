// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetAssessmentConnection Target Assessment Connection object
type TargetAssessmentConnection struct {

	// The OCID of the resource being referenced.
	Id *string `mandatory:"false" json:"id"`

	// Defines the type of connection. For example, ORACLE.
	ConnectionType ConnectionTypeEnum `mandatory:"false" json:"connectionType,omitempty"`

	// The technology type.
	TechnologyType TechnologyTypeEnum `mandatory:"false" json:"technologyType,omitempty"`

	// Technology sub-type e.g. ADW_SHARED, ADW_DEDICATED, ATP_SHARED, ATP_DEDICATED
	TechnologySubType *string `mandatory:"false" json:"technologySubType"`

	// The database version
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`
}

func (m TargetAssessmentConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetAssessmentConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConnectionTypeEnum(string(m.ConnectionType)); !ok && m.ConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionType: %s. Supported values are: %s.", m.ConnectionType, strings.Join(GetConnectionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetTechnologyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
