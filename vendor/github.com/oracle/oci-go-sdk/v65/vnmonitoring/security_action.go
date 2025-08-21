// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityAction Defines the security action details taken on the traffic.
type SecurityAction interface {

	// Type of the `SecurityAction`.
	GetActionType() SecurityActionActionTypeEnum
}

type securityaction struct {
	JsonData   []byte
	ActionType SecurityActionActionTypeEnum `mandatory:"true" json:"actionType"`
	Action     string                       `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *securityaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecurityaction securityaction
	s := struct {
		Model Unmarshalersecurityaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionType = s.Model.ActionType
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *securityaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "ALLOWED":
		mm := AllowedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_MISSING_POLICY":
		mm := ZprMissingPolicySecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_DENIED":
		mm := ZprDeniedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_POLICY_NOT_EVALUATED_MISSING_ROUTE":
		mm := ZprPolicyNotEvaluatedMissingRouteSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_NSG_UNSUPPORTED":
		mm := ZprNsgUnsupportedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_ALLOWED":
		mm := ZprAllowedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DENIED":
		mm := DeniedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_CIDR_UNSUPPORTED":
		mm := ZprCidrUnsupportedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZPR_POLICY_NOT_EVALUATED_SL_NSG_DENIED":
		mm := ZprPolicyNotEvaluatedSlNsgDeniedSecurityAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SecurityAction: %s.", m.Action)
		return *m, nil
	}
}

// GetActionType returns ActionType
func (m securityaction) GetActionType() SecurityActionActionTypeEnum {
	return m.ActionType
}

func (m securityaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m securityaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityActionActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetSecurityActionActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityActionActionTypeEnum Enum with underlying type: string
type SecurityActionActionTypeEnum string

// Set of constants representing the allowable values for SecurityActionActionTypeEnum
const (
	SecurityActionActionTypeExplicit SecurityActionActionTypeEnum = "EXPLICIT"
	SecurityActionActionTypeImplicit SecurityActionActionTypeEnum = "IMPLICIT"
)

var mappingSecurityActionActionTypeEnum = map[string]SecurityActionActionTypeEnum{
	"EXPLICIT": SecurityActionActionTypeExplicit,
	"IMPLICIT": SecurityActionActionTypeImplicit,
}

var mappingSecurityActionActionTypeEnumLowerCase = map[string]SecurityActionActionTypeEnum{
	"explicit": SecurityActionActionTypeExplicit,
	"implicit": SecurityActionActionTypeImplicit,
}

// GetSecurityActionActionTypeEnumValues Enumerates the set of values for SecurityActionActionTypeEnum
func GetSecurityActionActionTypeEnumValues() []SecurityActionActionTypeEnum {
	values := make([]SecurityActionActionTypeEnum, 0)
	for _, v := range mappingSecurityActionActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityActionActionTypeEnumStringValues Enumerates the set of values in String for SecurityActionActionTypeEnum
func GetSecurityActionActionTypeEnumStringValues() []string {
	return []string{
		"EXPLICIT",
		"IMPLICIT",
	}
}

// GetMappingSecurityActionActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityActionActionTypeEnum(val string) (SecurityActionActionTypeEnum, bool) {
	enum, ok := mappingSecurityActionActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityActionActionEnum Enum with underlying type: string
type SecurityActionActionEnum string

// Set of constants representing the allowable values for SecurityActionActionEnum
const (
	SecurityActionActionAllowed                           SecurityActionActionEnum = "ALLOWED"
	SecurityActionActionDenied                            SecurityActionActionEnum = "DENIED"
	SecurityActionActionZprDenied                         SecurityActionActionEnum = "ZPR_DENIED"
	SecurityActionActionZprAllowed                        SecurityActionActionEnum = "ZPR_ALLOWED"
	SecurityActionActionZprCidrUnsupported                SecurityActionActionEnum = "ZPR_CIDR_UNSUPPORTED"
	SecurityActionActionZprNsgUnsupported                 SecurityActionActionEnum = "ZPR_NSG_UNSUPPORTED"
	SecurityActionActionZprPolicyNotEvaluatedMissingRoute SecurityActionActionEnum = "ZPR_POLICY_NOT_EVALUATED_MISSING_ROUTE"
	SecurityActionActionZprPolicyNotEvaluatedSlNsgDenied  SecurityActionActionEnum = "ZPR_POLICY_NOT_EVALUATED_SL_NSG_DENIED"
	SecurityActionActionZprMissingPolicy                  SecurityActionActionEnum = "ZPR_MISSING_POLICY"
)

var mappingSecurityActionActionEnum = map[string]SecurityActionActionEnum{
	"ALLOWED":                                SecurityActionActionAllowed,
	"DENIED":                                 SecurityActionActionDenied,
	"ZPR_DENIED":                             SecurityActionActionZprDenied,
	"ZPR_ALLOWED":                            SecurityActionActionZprAllowed,
	"ZPR_CIDR_UNSUPPORTED":                   SecurityActionActionZprCidrUnsupported,
	"ZPR_NSG_UNSUPPORTED":                    SecurityActionActionZprNsgUnsupported,
	"ZPR_POLICY_NOT_EVALUATED_MISSING_ROUTE": SecurityActionActionZprPolicyNotEvaluatedMissingRoute,
	"ZPR_POLICY_NOT_EVALUATED_SL_NSG_DENIED": SecurityActionActionZprPolicyNotEvaluatedSlNsgDenied,
	"ZPR_MISSING_POLICY":                     SecurityActionActionZprMissingPolicy,
}

var mappingSecurityActionActionEnumLowerCase = map[string]SecurityActionActionEnum{
	"allowed":                                SecurityActionActionAllowed,
	"denied":                                 SecurityActionActionDenied,
	"zpr_denied":                             SecurityActionActionZprDenied,
	"zpr_allowed":                            SecurityActionActionZprAllowed,
	"zpr_cidr_unsupported":                   SecurityActionActionZprCidrUnsupported,
	"zpr_nsg_unsupported":                    SecurityActionActionZprNsgUnsupported,
	"zpr_policy_not_evaluated_missing_route": SecurityActionActionZprPolicyNotEvaluatedMissingRoute,
	"zpr_policy_not_evaluated_sl_nsg_denied": SecurityActionActionZprPolicyNotEvaluatedSlNsgDenied,
	"zpr_missing_policy":                     SecurityActionActionZprMissingPolicy,
}

// GetSecurityActionActionEnumValues Enumerates the set of values for SecurityActionActionEnum
func GetSecurityActionActionEnumValues() []SecurityActionActionEnum {
	values := make([]SecurityActionActionEnum, 0)
	for _, v := range mappingSecurityActionActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityActionActionEnumStringValues Enumerates the set of values in String for SecurityActionActionEnum
func GetSecurityActionActionEnumStringValues() []string {
	return []string{
		"ALLOWED",
		"DENIED",
		"ZPR_DENIED",
		"ZPR_ALLOWED",
		"ZPR_CIDR_UNSUPPORTED",
		"ZPR_NSG_UNSUPPORTED",
		"ZPR_POLICY_NOT_EVALUATED_MISSING_ROUTE",
		"ZPR_POLICY_NOT_EVALUATED_SL_NSG_DENIED",
		"ZPR_MISSING_POLICY",
	}
}

// GetMappingSecurityActionActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityActionActionEnum(val string) (SecurityActionActionEnum, bool) {
	enum, ok := mappingSecurityActionActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
