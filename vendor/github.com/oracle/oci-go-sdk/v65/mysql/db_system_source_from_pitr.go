// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemSourceFromPitr DB System OCID to perform a point in time recovery to the current point in time.
// DB System OCID and recovery point to perform a point in time recovery to the
// specified recovery point.
type DbSystemSourceFromPitr struct {

	// The OCID of the DB System from which a backup shall be selected to be
	// restored when creating the new DB System. Use this together with
	// recovery point to perform a point in time recovery operation.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The date and time, as per RFC 3339, of the change up to which the
	// new DB System shall be restored to, using a backup and logs from the
	// original DB System. In case no point in time is specified, then this
	// new DB System shall be restored up to the latest change recorded for
	// the original DB System.
	RecoveryPoint *common.SDKTime `mandatory:"false" json:"recoveryPoint"`
}

func (m DbSystemSourceFromPitr) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemSourceFromPitr) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbSystemSourceFromPitr) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbSystemSourceFromPitr DbSystemSourceFromPitr
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbSystemSourceFromPitr
	}{
		"PITR",
		(MarshalTypeDbSystemSourceFromPitr)(m),
	}

	return json.Marshal(&s)
}
