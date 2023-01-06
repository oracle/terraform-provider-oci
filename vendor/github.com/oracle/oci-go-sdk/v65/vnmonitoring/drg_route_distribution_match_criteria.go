// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgRouteDistributionMatchCriteria The match criteria in a route distribution statement. The match criteria outlines which routes
// should be imported or exported.
type DrgRouteDistributionMatchCriteria interface {
}

type drgroutedistributionmatchcriteria struct {
	JsonData  []byte
	MatchType string `json:"matchType"`
}

// UnmarshalJSON unmarshals json
func (m *drgroutedistributionmatchcriteria) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrgroutedistributionmatchcriteria drgroutedistributionmatchcriteria
	s := struct {
		Model Unmarshalerdrgroutedistributionmatchcriteria
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MatchType = s.Model.MatchType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drgroutedistributionmatchcriteria) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MatchType {
	case "DRG_ATTACHMENT_ID":
		mm := DrgAttachmentIdDrgRouteDistributionMatchCriteria{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DRG_ATTACHMENT_TYPE":
		mm := DrgAttachmentTypeDrgRouteDistributionMatchCriteria{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m drgroutedistributionmatchcriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drgroutedistributionmatchcriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgRouteDistributionMatchCriteriaMatchTypeEnum Enum with underlying type: string
type DrgRouteDistributionMatchCriteriaMatchTypeEnum string

// Set of constants representing the allowable values for DrgRouteDistributionMatchCriteriaMatchTypeEnum
const (
	DrgRouteDistributionMatchCriteriaMatchTypeType DrgRouteDistributionMatchCriteriaMatchTypeEnum = "DRG_ATTACHMENT_TYPE"
	DrgRouteDistributionMatchCriteriaMatchTypeId   DrgRouteDistributionMatchCriteriaMatchTypeEnum = "DRG_ATTACHMENT_ID"
)

var mappingDrgRouteDistributionMatchCriteriaMatchTypeEnum = map[string]DrgRouteDistributionMatchCriteriaMatchTypeEnum{
	"DRG_ATTACHMENT_TYPE": DrgRouteDistributionMatchCriteriaMatchTypeType,
	"DRG_ATTACHMENT_ID":   DrgRouteDistributionMatchCriteriaMatchTypeId,
}

var mappingDrgRouteDistributionMatchCriteriaMatchTypeEnumLowerCase = map[string]DrgRouteDistributionMatchCriteriaMatchTypeEnum{
	"drg_attachment_type": DrgRouteDistributionMatchCriteriaMatchTypeType,
	"drg_attachment_id":   DrgRouteDistributionMatchCriteriaMatchTypeId,
}

// GetDrgRouteDistributionMatchCriteriaMatchTypeEnumValues Enumerates the set of values for DrgRouteDistributionMatchCriteriaMatchTypeEnum
func GetDrgRouteDistributionMatchCriteriaMatchTypeEnumValues() []DrgRouteDistributionMatchCriteriaMatchTypeEnum {
	values := make([]DrgRouteDistributionMatchCriteriaMatchTypeEnum, 0)
	for _, v := range mappingDrgRouteDistributionMatchCriteriaMatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgRouteDistributionMatchCriteriaMatchTypeEnumStringValues Enumerates the set of values in String for DrgRouteDistributionMatchCriteriaMatchTypeEnum
func GetDrgRouteDistributionMatchCriteriaMatchTypeEnumStringValues() []string {
	return []string{
		"DRG_ATTACHMENT_TYPE",
		"DRG_ATTACHMENT_ID",
	}
}

// GetMappingDrgRouteDistributionMatchCriteriaMatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgRouteDistributionMatchCriteriaMatchTypeEnum(val string) (DrgRouteDistributionMatchCriteriaMatchTypeEnum, bool) {
	enum, ok := mappingDrgRouteDistributionMatchCriteriaMatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
