// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseDetails The details of the Autonomous Database to be registered as a target database in Data Safe.
type AutonomousDatabaseDetails struct {

	// The OCID of the Autonomous Database registered as a target database in Data Safe.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The infrastructure type the database is running on.
	InfrastructureType InfrastructureTypeEnum `mandatory:"true" json:"infrastructureType"`
}

// GetInfrastructureType returns InfrastructureType
func (m AutonomousDatabaseDetails) GetInfrastructureType() InfrastructureTypeEnum {
	return m.InfrastructureType
}

func (m AutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInfrastructureTypeEnum(string(m.InfrastructureType)); !ok && m.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", m.InfrastructureType, strings.Join(GetInfrastructureTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutonomousDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutonomousDatabaseDetails AutonomousDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"databaseType"`
		MarshalTypeAutonomousDatabaseDetails
	}{
		"AUTONOMOUS_DATABASE",
		(MarshalTypeAutonomousDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
