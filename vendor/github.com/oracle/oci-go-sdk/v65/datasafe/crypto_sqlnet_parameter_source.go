// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoSqlnetParameterSource Source details for SQLNET parameter values.
type CryptoSqlnetParameterSource struct {

	// Source type of SQLNET parameter data.
	Type CryptoSqlnetParameterSourceTypeEnum `mandatory:"true" json:"type"`
}

func (m CryptoSqlnetParameterSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoSqlnetParameterSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoSqlnetParameterSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCryptoSqlnetParameterSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoSqlnetParameterSourceTypeEnum Enum with underlying type: string
type CryptoSqlnetParameterSourceTypeEnum string

// Set of constants representing the allowable values for CryptoSqlnetParameterSourceTypeEnum
const (
	CryptoSqlnetParameterSourceTypeSqlnetOra CryptoSqlnetParameterSourceTypeEnum = "SQLNET_ORA"
)

var mappingCryptoSqlnetParameterSourceTypeEnum = map[string]CryptoSqlnetParameterSourceTypeEnum{
	"SQLNET_ORA": CryptoSqlnetParameterSourceTypeSqlnetOra,
}

var mappingCryptoSqlnetParameterSourceTypeEnumLowerCase = map[string]CryptoSqlnetParameterSourceTypeEnum{
	"sqlnet_ora": CryptoSqlnetParameterSourceTypeSqlnetOra,
}

// GetCryptoSqlnetParameterSourceTypeEnumValues Enumerates the set of values for CryptoSqlnetParameterSourceTypeEnum
func GetCryptoSqlnetParameterSourceTypeEnumValues() []CryptoSqlnetParameterSourceTypeEnum {
	values := make([]CryptoSqlnetParameterSourceTypeEnum, 0)
	for _, v := range mappingCryptoSqlnetParameterSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoSqlnetParameterSourceTypeEnumStringValues Enumerates the set of values in String for CryptoSqlnetParameterSourceTypeEnum
func GetCryptoSqlnetParameterSourceTypeEnumStringValues() []string {
	return []string{
		"SQLNET_ORA",
	}
}

// GetMappingCryptoSqlnetParameterSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoSqlnetParameterSourceTypeEnum(val string) (CryptoSqlnetParameterSourceTypeEnum, bool) {
	enum, ok := mappingCryptoSqlnetParameterSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
