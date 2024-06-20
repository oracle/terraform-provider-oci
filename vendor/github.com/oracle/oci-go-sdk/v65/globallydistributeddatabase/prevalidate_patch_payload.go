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

// PrevalidatePatchPayload Payload to prevalidate patch sharded database operation.
type PrevalidatePatchPayload struct {
	PrevalidatePayload *PatchShardedDatabaseDetails `mandatory:"true" json:"prevalidatePayload"`

	// Sharded database identifier
	ShardedDatabaseId *string `mandatory:"true" json:"shardedDatabaseId"`
}

func (m PrevalidatePatchPayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrevalidatePatchPayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PrevalidatePatchPayload) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrevalidatePatchPayload PrevalidatePatchPayload
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypePrevalidatePatchPayload
	}{
		"PATCH",
		(MarshalTypePrevalidatePatchPayload)(m),
	}

	return json.Marshal(&s)
}
