// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystemSourceFromDbSystem The source DB system identifier (OCID) from which the cloned DB system was created.
type DbSystemSourceFromDbSystem struct {

	// The OCID of the DB system used as the source for the new DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`
}

func (m DbSystemSourceFromDbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSourceFromDbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceFromDbSystem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceFromDbSystem DbSystemSourceFromDbSystem
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceFromDbSystem
	}{
		"DBSYSTEM",
		(MarshalTypeDbSystemSourceFromDbSystem)(m),
	}

	return json.Marshal(&s)
}
