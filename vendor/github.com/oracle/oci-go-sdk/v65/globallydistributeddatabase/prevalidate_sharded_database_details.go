// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrevalidateShardedDatabaseDetails Input for prevalidate sharded database API to validate various operations payload.
type PrevalidateShardedDatabaseDetails struct {
	PrevalidateShardedDatabaseDetails PrevalidatePayload `mandatory:"true" json:"prevalidateShardedDatabaseDetails"`
}

func (m PrevalidateShardedDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrevalidateShardedDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *PrevalidateShardedDatabaseDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PrevalidateShardedDatabaseDetails prevalidatepayload `json:"prevalidateShardedDatabaseDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PrevalidateShardedDatabaseDetails.UnmarshalPolymorphicJSON(model.PrevalidateShardedDatabaseDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PrevalidateShardedDatabaseDetails = nn.(PrevalidatePayload)
	} else {
		m.PrevalidateShardedDatabaseDetails = nil
	}

	return
}
