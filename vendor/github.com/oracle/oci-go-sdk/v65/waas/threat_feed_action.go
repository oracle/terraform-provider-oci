// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ThreatFeedAction The action to take for a request that has been determined to be potentially malicious.
type ThreatFeedAction struct {

	// The unique key of the object for which the action applies.
	Key *string `mandatory:"true" json:"key"`

	// The selected action. If unspecified, defaults to `OFF`.
	Action ThreatFeedActionActionEnum `mandatory:"true" json:"action"`
}

func (m ThreatFeedAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThreatFeedAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingThreatFeedActionActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetThreatFeedActionActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ThreatFeedActionActionEnum Enum with underlying type: string
type ThreatFeedActionActionEnum string

// Set of constants representing the allowable values for ThreatFeedActionActionEnum
const (
	ThreatFeedActionActionOff    ThreatFeedActionActionEnum = "OFF"
	ThreatFeedActionActionDetect ThreatFeedActionActionEnum = "DETECT"
	ThreatFeedActionActionBlock  ThreatFeedActionActionEnum = "BLOCK"
)

var mappingThreatFeedActionActionEnum = map[string]ThreatFeedActionActionEnum{
	"OFF":    ThreatFeedActionActionOff,
	"DETECT": ThreatFeedActionActionDetect,
	"BLOCK":  ThreatFeedActionActionBlock,
}

var mappingThreatFeedActionActionEnumLowerCase = map[string]ThreatFeedActionActionEnum{
	"off":    ThreatFeedActionActionOff,
	"detect": ThreatFeedActionActionDetect,
	"block":  ThreatFeedActionActionBlock,
}

// GetThreatFeedActionActionEnumValues Enumerates the set of values for ThreatFeedActionActionEnum
func GetThreatFeedActionActionEnumValues() []ThreatFeedActionActionEnum {
	values := make([]ThreatFeedActionActionEnum, 0)
	for _, v := range mappingThreatFeedActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetThreatFeedActionActionEnumStringValues Enumerates the set of values in String for ThreatFeedActionActionEnum
func GetThreatFeedActionActionEnumStringValues() []string {
	return []string{
		"OFF",
		"DETECT",
		"BLOCK",
	}
}

// GetMappingThreatFeedActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThreatFeedActionActionEnum(val string) (ThreatFeedActionActionEnum, bool) {
	enum, ok := mappingThreatFeedActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
