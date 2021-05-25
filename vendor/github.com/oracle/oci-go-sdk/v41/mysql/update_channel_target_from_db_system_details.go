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

// UpdateChannelTargetFromDbSystemDetails Parameters detailing how to provision the target endpoint that is a DB System.
type UpdateChannelTargetFromDbSystemDetails struct {

	// The case-insensitive name that identifies the replication channel. Channel names
	// must follow the rules defined for MySQL identifiers (https://dev.mysql.com/doc/refman/8.0/en/identifiers.html).
	// The names of non-Deleted Channels must be unique for each DB System.
	ChannelName *string `mandatory:"false" json:"channelName"`

	// The username for the replication applier of the target MySQL DB System.
	ApplierUsername *string `mandatory:"false" json:"applierUsername"`
}

func (m UpdateChannelTargetFromDbSystemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateChannelTargetFromDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateChannelTargetFromDbSystemDetails UpdateChannelTargetFromDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeUpdateChannelTargetFromDbSystemDetails
	}{
		"DBSYSTEM",
		(MarshalTypeUpdateChannelTargetFromDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}
