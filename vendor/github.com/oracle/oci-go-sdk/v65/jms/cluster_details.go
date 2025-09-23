// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClusterDetails The state of the cluster during a specified time period.
type ClusterDetails struct {

	// The type of the cluster.
	ClusterType ClusterDetailsClusterTypeEnum `mandatory:"false" json:"clusterType,omitempty"`

	// Number of nodes as reported in the last cluster scan
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// Number of pods as reported in the last cluster scan
	PodCount *int `mandatory:"false" json:"podCount"`

	// Number of containers as reported in the last cluster scan
	ContainerCount *int `mandatory:"false" json:"containerCount"`
}

func (m ClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClusterDetailsClusterTypeEnum(string(m.ClusterType)); !ok && m.ClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterType: %s. Supported values are: %s.", m.ClusterType, strings.Join(GetClusterDetailsClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterDetailsClusterTypeEnum Enum with underlying type: string
type ClusterDetailsClusterTypeEnum string

// Set of constants representing the allowable values for ClusterDetailsClusterTypeEnum
const (
	ClusterDetailsClusterTypeOke   ClusterDetailsClusterTypeEnum = "OKE"
	ClusterDetailsClusterTypeOther ClusterDetailsClusterTypeEnum = "OTHER"
)

var mappingClusterDetailsClusterTypeEnum = map[string]ClusterDetailsClusterTypeEnum{
	"OKE":   ClusterDetailsClusterTypeOke,
	"OTHER": ClusterDetailsClusterTypeOther,
}

var mappingClusterDetailsClusterTypeEnumLowerCase = map[string]ClusterDetailsClusterTypeEnum{
	"oke":   ClusterDetailsClusterTypeOke,
	"other": ClusterDetailsClusterTypeOther,
}

// GetClusterDetailsClusterTypeEnumValues Enumerates the set of values for ClusterDetailsClusterTypeEnum
func GetClusterDetailsClusterTypeEnumValues() []ClusterDetailsClusterTypeEnum {
	values := make([]ClusterDetailsClusterTypeEnum, 0)
	for _, v := range mappingClusterDetailsClusterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterDetailsClusterTypeEnumStringValues Enumerates the set of values in String for ClusterDetailsClusterTypeEnum
func GetClusterDetailsClusterTypeEnumStringValues() []string {
	return []string{
		"OKE",
		"OTHER",
	}
}

// GetMappingClusterDetailsClusterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterDetailsClusterTypeEnum(val string) (ClusterDetailsClusterTypeEnum, bool) {
	enum, ok := mappingClusterDetailsClusterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
