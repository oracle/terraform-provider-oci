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

// ThreatFeed The settings of the threat intelligence feed. You can block requests from IP addresses based on their reputations with various commercial and open source threat feeds.
type ThreatFeed struct {

	// The unique key of the threat intelligence feed.
	Key *string `mandatory:"false" json:"key"`

	// The name of the threat intelligence feed.
	Name *string `mandatory:"false" json:"name"`

	// The action to take when traffic is flagged as malicious by data from the threat intelligence feed. If unspecified, defaults to `OFF`.
	Action ThreatFeedActionEnum `mandatory:"false" json:"action,omitempty"`

	// The description of the threat intelligence feed.
	Description *string `mandatory:"false" json:"description"`
}

func (m ThreatFeed) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThreatFeed) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingThreatFeedActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetThreatFeedActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ThreatFeedActionEnum Enum with underlying type: string
type ThreatFeedActionEnum string

// Set of constants representing the allowable values for ThreatFeedActionEnum
const (
	ThreatFeedActionOff    ThreatFeedActionEnum = "OFF"
	ThreatFeedActionDetect ThreatFeedActionEnum = "DETECT"
	ThreatFeedActionBlock  ThreatFeedActionEnum = "BLOCK"
)

var mappingThreatFeedActionEnum = map[string]ThreatFeedActionEnum{
	"OFF":    ThreatFeedActionOff,
	"DETECT": ThreatFeedActionDetect,
	"BLOCK":  ThreatFeedActionBlock,
}

var mappingThreatFeedActionEnumLowerCase = map[string]ThreatFeedActionEnum{
	"off":    ThreatFeedActionOff,
	"detect": ThreatFeedActionDetect,
	"block":  ThreatFeedActionBlock,
}

// GetThreatFeedActionEnumValues Enumerates the set of values for ThreatFeedActionEnum
func GetThreatFeedActionEnumValues() []ThreatFeedActionEnum {
	values := make([]ThreatFeedActionEnum, 0)
	for _, v := range mappingThreatFeedActionEnum {
		values = append(values, v)
	}
	return values
}

// GetThreatFeedActionEnumStringValues Enumerates the set of values in String for ThreatFeedActionEnum
func GetThreatFeedActionEnumStringValues() []string {
	return []string{
		"OFF",
		"DETECT",
		"BLOCK",
	}
}

// GetMappingThreatFeedActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThreatFeedActionEnum(val string) (ThreatFeedActionEnum, bool) {
	enum, ok := mappingThreatFeedActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
