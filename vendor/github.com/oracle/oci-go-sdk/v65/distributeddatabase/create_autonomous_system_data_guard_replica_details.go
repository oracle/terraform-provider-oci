// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAutonomousSystemDataGuardReplicaDetails Details required to create a system data guard replica.
type CreateAutonomousSystemDataGuardReplicaDetails struct {

	// The name of the data guard replica.
	// It must start with a letter and contain only letters, digits, and underscores.
	// Maximum length is 40 characters.
	Name *string `mandatory:"true" json:"name"`

	// Details of the databases associated with the system data guard replica.
	Databases []CreateDistributedAutonomousDatabaseShardDatabaseDetails `mandatory:"true" json:"databases"`
}

func (m CreateAutonomousSystemDataGuardReplicaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousSystemDataGuardReplicaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateAutonomousSystemDataGuardReplicaDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name      *string                                                   `json:"name"`
		Databases []createdistributedautonomousdatabasesharddatabasedetails `json:"databases"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Databases = make([]CreateDistributedAutonomousDatabaseShardDatabaseDetails, len(model.Databases))
	for i, n := range model.Databases {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Databases[i] = nn.(CreateDistributedAutonomousDatabaseShardDatabaseDetails)
		} else {
			m.Databases[i] = nil
		}
	}
	return
}
