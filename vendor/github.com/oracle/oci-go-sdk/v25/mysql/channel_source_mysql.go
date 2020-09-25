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

// ChannelSourceMysql Core properties of a Mysql Channel source.
type ChannelSourceMysql struct {

	// The network address of the MySQL instance.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The port the source MySQL instance listens on.
	Port *int `mandatory:"true" json:"port"`

	// The name of the replication user on the source MySQL instance.
	// The username has a maximum length of 96 characters. For more information,
	// please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html)
	Username *string `mandatory:"true" json:"username"`

	// The state of the Channel.
	SslMode ChannelSourceMysqlSslModeEnum `mandatory:"true" json:"sslMode"`
}

func (m ChannelSourceMysql) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ChannelSourceMysql) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeChannelSourceMysql ChannelSourceMysql
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeChannelSourceMysql
	}{
		"MYSQL",
		(MarshalTypeChannelSourceMysql)(m),
	}

	return json.Marshal(&s)
}

// ChannelSourceMysqlSslModeEnum Enum with underlying type: string
type ChannelSourceMysqlSslModeEnum string

// Set of constants representing the allowable values for ChannelSourceMysqlSslModeEnum
const (
	ChannelSourceMysqlSslModeRequired ChannelSourceMysqlSslModeEnum = "REQUIRED"
	ChannelSourceMysqlSslModeDisabled ChannelSourceMysqlSslModeEnum = "DISABLED"
)

var mappingChannelSourceMysqlSslMode = map[string]ChannelSourceMysqlSslModeEnum{
	"REQUIRED": ChannelSourceMysqlSslModeRequired,
	"DISABLED": ChannelSourceMysqlSslModeDisabled,
}

// GetChannelSourceMysqlSslModeEnumValues Enumerates the set of values for ChannelSourceMysqlSslModeEnum
func GetChannelSourceMysqlSslModeEnumValues() []ChannelSourceMysqlSslModeEnum {
	values := make([]ChannelSourceMysqlSslModeEnum, 0)
	for _, v := range mappingChannelSourceMysqlSslMode {
		values = append(values, v)
	}
	return values
}
