// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ThinAssociationAuthorizationResponse The representation of ThinAssociationAuthorizationResponse
type ThinAssociationAuthorizationResponse struct {

	// The authorization responses.
	Responses []ThinAuthorizationResponse `mandatory:"true" json:"responses"`

	// The association verification result.
	AssociationResult ThinAssociationAuthorizationResponseAssociationResultEnum `mandatory:"true" json:"associationResult"`

	// The decision cache duration.
	DecisionCacheDuration *string `mandatory:"false" json:"decisionCacheDuration"`
}

func (m ThinAssociationAuthorizationResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThinAssociationAuthorizationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingThinAssociationAuthorizationResponseAssociationResultEnum(string(m.AssociationResult)); !ok && m.AssociationResult != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationResult: %s. Supported values are: %s.", m.AssociationResult, strings.Join(GetThinAssociationAuthorizationResponseAssociationResultEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ThinAssociationAuthorizationResponseAssociationResultEnum Enum with underlying type: string
type ThinAssociationAuthorizationResponseAssociationResultEnum string

// Set of constants representing the allowable values for ThinAssociationAuthorizationResponseAssociationResultEnum
const (
	ThinAssociationAuthorizationResponseAssociationResultFailUnknown        ThinAssociationAuthorizationResponseAssociationResultEnum = "FAIL_UNKNOWN"
	ThinAssociationAuthorizationResponseAssociationResultFailBadRequest     ThinAssociationAuthorizationResponseAssociationResultEnum = "FAIL_BAD_REQUEST"
	ThinAssociationAuthorizationResponseAssociationResultFailMissingEndorse ThinAssociationAuthorizationResponseAssociationResultEnum = "FAIL_MISSING_ENDORSE"
	ThinAssociationAuthorizationResponseAssociationResultFailMissingAdmit   ThinAssociationAuthorizationResponseAssociationResultEnum = "FAIL_MISSING_ADMIT"
	ThinAssociationAuthorizationResponseAssociationResultSuccess            ThinAssociationAuthorizationResponseAssociationResultEnum = "SUCCESS"
)

var mappingThinAssociationAuthorizationResponseAssociationResultEnum = map[string]ThinAssociationAuthorizationResponseAssociationResultEnum{
	"FAIL_UNKNOWN":         ThinAssociationAuthorizationResponseAssociationResultFailUnknown,
	"FAIL_BAD_REQUEST":     ThinAssociationAuthorizationResponseAssociationResultFailBadRequest,
	"FAIL_MISSING_ENDORSE": ThinAssociationAuthorizationResponseAssociationResultFailMissingEndorse,
	"FAIL_MISSING_ADMIT":   ThinAssociationAuthorizationResponseAssociationResultFailMissingAdmit,
	"SUCCESS":              ThinAssociationAuthorizationResponseAssociationResultSuccess,
}

// GetThinAssociationAuthorizationResponseAssociationResultEnumValues Enumerates the set of values for ThinAssociationAuthorizationResponseAssociationResultEnum
func GetThinAssociationAuthorizationResponseAssociationResultEnumValues() []ThinAssociationAuthorizationResponseAssociationResultEnum {
	values := make([]ThinAssociationAuthorizationResponseAssociationResultEnum, 0)
	for _, v := range mappingThinAssociationAuthorizationResponseAssociationResultEnum {
		values = append(values, v)
	}
	return values
}

// GetThinAssociationAuthorizationResponseAssociationResultEnumStringValues Enumerates the set of values in String for ThinAssociationAuthorizationResponseAssociationResultEnum
func GetThinAssociationAuthorizationResponseAssociationResultEnumStringValues() []string {
	return []string{
		"FAIL_UNKNOWN",
		"FAIL_BAD_REQUEST",
		"FAIL_MISSING_ENDORSE",
		"FAIL_MISSING_ADMIT",
		"SUCCESS",
	}
}

// GetMappingThinAssociationAuthorizationResponseAssociationResultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingThinAssociationAuthorizationResponseAssociationResultEnum(val string) (ThinAssociationAuthorizationResponseAssociationResultEnum, bool) {
	mappingThinAssociationAuthorizationResponseAssociationResultEnumIgnoreCase := make(map[string]ThinAssociationAuthorizationResponseAssociationResultEnum)
	for k, v := range mappingThinAssociationAuthorizationResponseAssociationResultEnum {
		mappingThinAssociationAuthorizationResponseAssociationResultEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingThinAssociationAuthorizationResponseAssociationResultEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
