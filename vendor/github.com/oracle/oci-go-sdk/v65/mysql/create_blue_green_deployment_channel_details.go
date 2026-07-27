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

// CreateBlueGreenDeploymentChannelDetails Replication channel details for a blue/green deployment.
type CreateBlueGreenDeploymentChannelDetails struct {

	// The username on the source DB system used by the blue/green workflow to configure the replication channel.
	// The username has a maximum length of 96 characters. For more information,
	// please see the MySQL documentation (https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html)
	SourceUsername *string `mandatory:"true" json:"sourceUsername"`

	// The password for the source DB system user used by the blue/green workflow to configure the replication channel.
	// The password must be between 8 and 32 characters long, and must contain at least 1
	// numeric character, 1 lowercase character, 1 uppercase character, and 1 special
	// (nonalphanumeric) character.
	SourcePassword *string `mandatory:"true" json:"sourcePassword"`

	// The SSL mode of the replication channel created by the blue/green workflow.
	// `VERIFY_CA` and `VERIFY_IDENTITY` require `sslCaCertificate`.
	// `REQUIRED` and `DISABLED` must not include `sslCaCertificate`.
	SslMode SslModeEnum `mandatory:"true" json:"sslMode"`

	// The username for the replication applier of the target MySQL DB System.
	ApplierUsername *string `mandatory:"false" json:"applierUsername"`

	SslCaCertificate CaCertificate `mandatory:"false" json:"sslCaCertificate"`
}

func (m CreateBlueGreenDeploymentChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBlueGreenDeploymentChannelDetails) ValidateEnumValue() (bool, error) {
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
func (m *CreateBlueGreenDeploymentChannelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ApplierUsername  *string       `json:"applierUsername"`
		SslCaCertificate cacertificate `json:"sslCaCertificate"`
		SourceUsername   *string       `json:"sourceUsername"`
		SourcePassword   *string       `json:"sourcePassword"`
		SslMode          SslModeEnum   `json:"sslMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ApplierUsername = model.ApplierUsername

	nn, e = model.SslCaCertificate.UnmarshalPolymorphicJSON(model.SslCaCertificate.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SslCaCertificate = nn.(CaCertificate)
	} else {
		m.SslCaCertificate = nil
	}

	m.SourceUsername = model.SourceUsername

	m.SourcePassword = model.SourcePassword

	m.SslMode = model.SslMode

	return
}
