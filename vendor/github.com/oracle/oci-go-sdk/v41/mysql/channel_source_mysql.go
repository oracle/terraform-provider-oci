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

	SslCaCertificate CaCertificate `mandatory:"false" json:"sslCaCertificate"`

	// The SSL mode of the Channel.
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

// UnmarshalJSON unmarshals from json
func (m *ChannelSourceMysql) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SslCaCertificate cacertificate                 `json:"sslCaCertificate"`
		Hostname         *string                       `json:"hostname"`
		Port             *int                          `json:"port"`
		Username         *string                       `json:"username"`
		SslMode          ChannelSourceMysqlSslModeEnum `json:"sslMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.SslCaCertificate.UnmarshalPolymorphicJSON(model.SslCaCertificate.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SslCaCertificate = nn.(CaCertificate)
	} else {
		m.SslCaCertificate = nil
	}

	m.Hostname = model.Hostname

	m.Port = model.Port

	m.Username = model.Username

	m.SslMode = model.SslMode

	return
}

// ChannelSourceMysqlSslModeEnum Enum with underlying type: string
type ChannelSourceMysqlSslModeEnum string

// Set of constants representing the allowable values for ChannelSourceMysqlSslModeEnum
const (
	ChannelSourceMysqlSslModeVerifyIdentity ChannelSourceMysqlSslModeEnum = "VERIFY_IDENTITY"
	ChannelSourceMysqlSslModeVerifyCa       ChannelSourceMysqlSslModeEnum = "VERIFY_CA"
	ChannelSourceMysqlSslModeRequired       ChannelSourceMysqlSslModeEnum = "REQUIRED"
	ChannelSourceMysqlSslModeDisabled       ChannelSourceMysqlSslModeEnum = "DISABLED"
)

var mappingChannelSourceMysqlSslMode = map[string]ChannelSourceMysqlSslModeEnum{
	"VERIFY_IDENTITY": ChannelSourceMysqlSslModeVerifyIdentity,
	"VERIFY_CA":       ChannelSourceMysqlSslModeVerifyCa,
	"REQUIRED":        ChannelSourceMysqlSslModeRequired,
	"DISABLED":        ChannelSourceMysqlSslModeDisabled,
}

// GetChannelSourceMysqlSslModeEnumValues Enumerates the set of values for ChannelSourceMysqlSslModeEnum
func GetChannelSourceMysqlSslModeEnumValues() []ChannelSourceMysqlSslModeEnum {
	values := make([]ChannelSourceMysqlSslModeEnum, 0)
	for _, v := range mappingChannelSourceMysqlSslMode {
		values = append(values, v)
	}
	return values
}
