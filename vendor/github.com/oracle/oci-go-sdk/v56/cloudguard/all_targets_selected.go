// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AllTargetsSelected All Targets selected.
type AllTargetsSelected struct {
}

func (m AllTargetsSelected) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AllTargetsSelected) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllTargetsSelected AllTargetsSelected
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeAllTargetsSelected
	}{
		"ALL",
		(MarshalTypeAllTargetsSelected)(m),
	}

	return json.Marshal(&s)
}
