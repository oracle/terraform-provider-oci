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

// PrevalidateCreatePayload Payload to prevalidate create sharded database operation.
type PrevalidateCreatePayload struct {
	PrevalidatePayload CreateShardedDatabaseDetails `mandatory:"true" json:"prevalidatePayload"`
}

func (m PrevalidateCreatePayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrevalidateCreatePayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PrevalidateCreatePayload) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrevalidateCreatePayload PrevalidateCreatePayload
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePrevalidateCreatePayload
	}{
		"CREATE",
		(MarshalTypePrevalidateCreatePayload)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *PrevalidateCreatePayload) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		PrevalidatePayload createshardeddatabasedetails `json:"prevalidatePayload"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.PrevalidatePayload.UnmarshalPolymorphicJSON(model.PrevalidatePayload.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PrevalidatePayload = nn.(CreateShardedDatabaseDetails)
	} else {
		m.PrevalidatePayload = nil
	}

	return
}
