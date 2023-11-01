// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlUpdateQueryProperties Query Properties applicable to SQL type of collection method
type SqlUpdateQueryProperties struct {
	SqlDetails *SqlDetails `mandatory:"false" json:"sqlDetails"`

	// List of values and position of PL/SQL procedure IN parameters
	InParamDetails []SqlInParamDetails `mandatory:"false" json:"inParamDetails"`

	OutParamDetails *SqlOutParamDetails `mandatory:"false" json:"outParamDetails"`

	// Type of SQL data collection method i.e. either a Statement or SQL Script File
	SqlType SqlQueryTypesEnum `mandatory:"false" json:"sqlType,omitempty"`
}

func (m SqlUpdateQueryProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlUpdateQueryProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlQueryTypesEnum(string(m.SqlType)); !ok && m.SqlType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlType: %s. Supported values are: %s.", m.SqlType, strings.Join(GetSqlQueryTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlUpdateQueryProperties) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlUpdateQueryProperties SqlUpdateQueryProperties
	s := struct {
		DiscriminatorParam string `json:"collectionMethod"`
		MarshalTypeSqlUpdateQueryProperties
	}{
		"SQL",
		(MarshalTypeSqlUpdateQueryProperties)(m),
	}

	return json.Marshal(&s)
}
