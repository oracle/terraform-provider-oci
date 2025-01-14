// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryDetails Disaster Recovery Configuration Details.
type DisasterRecoveryDetails struct {

	// The OCID of the replicated primary repository.
	PrimaryRepositoryId *string `mandatory:"true" json:"primaryRepositoryId"`

	// Region identifier of the Disaster Recovery Primary Repository. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	PrimaryRegion *string `mandatory:"true" json:"primaryRegion"`

	// Type of the primary repository
	PrimaryRepositoryKind DisasterRecoveryDetailsPrimaryRepositoryKindEnum `mandatory:"false" json:"primaryRepositoryKind,omitempty"`
}

func (m DisasterRecoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisasterRecoveryDetailsPrimaryRepositoryKindEnum(string(m.PrimaryRepositoryKind)); !ok && m.PrimaryRepositoryKind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrimaryRepositoryKind: %s. Supported values are: %s.", m.PrimaryRepositoryKind, strings.Join(GetDisasterRecoveryDetailsPrimaryRepositoryKindEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisasterRecoveryDetailsPrimaryRepositoryKindEnum Enum with underlying type: string
type DisasterRecoveryDetailsPrimaryRepositoryKindEnum string

// Set of constants representing the allowable values for DisasterRecoveryDetailsPrimaryRepositoryKindEnum
const (
	DisasterRecoveryDetailsPrimaryRepositoryKindHosted DisasterRecoveryDetailsPrimaryRepositoryKindEnum = "HOSTED"
	DisasterRecoveryDetailsPrimaryRepositoryKindForked DisasterRecoveryDetailsPrimaryRepositoryKindEnum = "FORKED"
	DisasterRecoveryDetailsPrimaryRepositoryKindNone   DisasterRecoveryDetailsPrimaryRepositoryKindEnum = "NONE"
)

var mappingDisasterRecoveryDetailsPrimaryRepositoryKindEnum = map[string]DisasterRecoveryDetailsPrimaryRepositoryKindEnum{
	"HOSTED": DisasterRecoveryDetailsPrimaryRepositoryKindHosted,
	"FORKED": DisasterRecoveryDetailsPrimaryRepositoryKindForked,
	"NONE":   DisasterRecoveryDetailsPrimaryRepositoryKindNone,
}

var mappingDisasterRecoveryDetailsPrimaryRepositoryKindEnumLowerCase = map[string]DisasterRecoveryDetailsPrimaryRepositoryKindEnum{
	"hosted": DisasterRecoveryDetailsPrimaryRepositoryKindHosted,
	"forked": DisasterRecoveryDetailsPrimaryRepositoryKindForked,
	"none":   DisasterRecoveryDetailsPrimaryRepositoryKindNone,
}

// GetDisasterRecoveryDetailsPrimaryRepositoryKindEnumValues Enumerates the set of values for DisasterRecoveryDetailsPrimaryRepositoryKindEnum
func GetDisasterRecoveryDetailsPrimaryRepositoryKindEnumValues() []DisasterRecoveryDetailsPrimaryRepositoryKindEnum {
	values := make([]DisasterRecoveryDetailsPrimaryRepositoryKindEnum, 0)
	for _, v := range mappingDisasterRecoveryDetailsPrimaryRepositoryKindEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryDetailsPrimaryRepositoryKindEnumStringValues Enumerates the set of values in String for DisasterRecoveryDetailsPrimaryRepositoryKindEnum
func GetDisasterRecoveryDetailsPrimaryRepositoryKindEnumStringValues() []string {
	return []string{
		"HOSTED",
		"FORKED",
		"NONE",
	}
}

// GetMappingDisasterRecoveryDetailsPrimaryRepositoryKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryDetailsPrimaryRepositoryKindEnum(val string) (DisasterRecoveryDetailsPrimaryRepositoryKindEnum, bool) {
	enum, ok := mappingDisasterRecoveryDetailsPrimaryRepositoryKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
