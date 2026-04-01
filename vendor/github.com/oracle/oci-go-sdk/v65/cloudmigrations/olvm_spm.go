// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmSpm The host storage pool manager (SPM) status and definition.
type OlvmSpm struct {

	// Priority of this SPM.
	Priority *int `mandatory:"false" json:"priority"`

	// Status of this SPM.
	SpmStatus OlvmSpmSpmStatusEnum `mandatory:"false" json:"spmStatus,omitempty"`
}

func (m OlvmSpm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmSpm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmSpmSpmStatusEnum(string(m.SpmStatus)); !ok && m.SpmStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SpmStatus: %s. Supported values are: %s.", m.SpmStatus, strings.Join(GetOlvmSpmSpmStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmSpmSpmStatusEnum Enum with underlying type: string
type OlvmSpmSpmStatusEnum string

// Set of constants representing the allowable values for OlvmSpmSpmStatusEnum
const (
	OlvmSpmSpmStatusContending OlvmSpmSpmStatusEnum = "CONTENDING"
	OlvmSpmSpmStatusNone       OlvmSpmSpmStatusEnum = "NONE"
	OlvmSpmSpmStatusSpm        OlvmSpmSpmStatusEnum = "SPM"
)

var mappingOlvmSpmSpmStatusEnum = map[string]OlvmSpmSpmStatusEnum{
	"CONTENDING": OlvmSpmSpmStatusContending,
	"NONE":       OlvmSpmSpmStatusNone,
	"SPM":        OlvmSpmSpmStatusSpm,
}

var mappingOlvmSpmSpmStatusEnumLowerCase = map[string]OlvmSpmSpmStatusEnum{
	"contending": OlvmSpmSpmStatusContending,
	"none":       OlvmSpmSpmStatusNone,
	"spm":        OlvmSpmSpmStatusSpm,
}

// GetOlvmSpmSpmStatusEnumValues Enumerates the set of values for OlvmSpmSpmStatusEnum
func GetOlvmSpmSpmStatusEnumValues() []OlvmSpmSpmStatusEnum {
	values := make([]OlvmSpmSpmStatusEnum, 0)
	for _, v := range mappingOlvmSpmSpmStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmSpmSpmStatusEnumStringValues Enumerates the set of values in String for OlvmSpmSpmStatusEnum
func GetOlvmSpmSpmStatusEnumStringValues() []string {
	return []string{
		"CONTENDING",
		"NONE",
		"SPM",
	}
}

// GetMappingOlvmSpmSpmStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmSpmSpmStatusEnum(val string) (OlvmSpmSpmStatusEnum, bool) {
	enum, ok := mappingOlvmSpmSpmStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
