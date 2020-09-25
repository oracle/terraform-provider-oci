// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UpdateChannelSourceFromMysqlDetails Parameters detailing how to provision the source endpoint that is a MySQL Server.
// Typically a MySQL Server that is not managed by the MySQL Database Service.
type UpdateChannelSourceFromMysqlDetails struct {

	// The network address of the MySQL instance.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The port the source MySQL instance listens on.
	Port *int `mandatory:"false" json:"port"`

	// The name of the replication user on the source MySQL instance.
	// The username has a maximum length of 96 characters. For more information,
	// please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html)
	Username *string `mandatory:"false" json:"username"`

	// The password for the replication user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character,
	// and 1 special (nonalphanumeric) character.
	Password *string `mandatory:"false" json:"password"`

	// The SSL mode of the Channel.
	SslMode ChannelSourceMysqlSslModeEnum `mandatory:"false" json:"sslMode,omitempty"`
}

func (m UpdateChannelSourceFromMysqlDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateChannelSourceFromMysqlDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateChannelSourceFromMysqlDetails UpdateChannelSourceFromMysqlDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUpdateChannelSourceFromMysqlDetails
	}{
		"MYSQL",
		(MarshalTypeUpdateChannelSourceFromMysqlDetails)(m),
	}

	return json.Marshal(&s)
}
