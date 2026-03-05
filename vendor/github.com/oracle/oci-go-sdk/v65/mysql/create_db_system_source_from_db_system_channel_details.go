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

// CreateDbSystemSourceFromDbSystemChannelDetails Properties to setup a replication channel with the source (cloned) DB system.
type CreateDbSystemSourceFromDbSystemChannelDetails struct {

	// The name of the replication user on the source DB system.
	// The username has a maximum length of 96 characters. For more information,
	// please see the MySQL documentation (https://dev.mysql.com/doc/en/change-replication-source-to.html)
	SourceUsername *string `mandatory:"true" json:"sourceUsername"`

	// The password for the replication user. The password must be
	// between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character,
	// and 1 special (nonalphanumeric) character.
	SourcePassword *string `mandatory:"true" json:"sourcePassword"`

	// The SSL mode of the Channel.
	SslMode SslModeEnum `mandatory:"true" json:"sslMode"`

	SslCaCertificate CaCertificate `mandatory:"false" json:"sslCaCertificate"`

	// The username for the replication applier of the created MySQL DB System.
	ApplierUsername *string `mandatory:"false" json:"applierUsername"`
}

func (m CreateDbSystemSourceFromDbSystemChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemSourceFromDbSystemChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSslModeEnum(string(m.SslMode)); !ok && m.SslMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SslMode: %s. Supported values are: %s.", m.SslMode, strings.Join(GetSslModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateDbSystemSourceFromDbSystemChannelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SslCaCertificate cacertificate `json:"sslCaCertificate"`
		ApplierUsername  *string       `json:"applierUsername"`
		SourceUsername   *string       `json:"sourceUsername"`
		SourcePassword   *string       `json:"sourcePassword"`
		SslMode          SslModeEnum   `json:"sslMode"`
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

	m.ApplierUsername = model.ApplierUsername

	m.SourceUsername = model.SourceUsername

	m.SourcePassword = model.SourcePassword

	m.SslMode = model.SslMode

	return
}
