// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DbSystemSourceFromNone A DB System created from no particular external source.
type DbSystemSourceFromNone struct {
}

func (m DbSystemSourceFromNone) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceFromNone) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceFromNone DbSystemSourceFromNone
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceFromNone
	}{
		"NONE",
		(MarshalTypeDbSystemSourceFromNone)(m),
	}

	return json.Marshal(&s)
}
