// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryPrecheckResult A single precheck result.
type DisasterRecoveryPrecheckResult struct {

	// UUID to uniquely identify the each check result.
	Key *string `mandatory:"true" json:"key"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"true" json:"description"`

	// Status of the DR precheck result.
	Status PrecheckStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource related to the corresponding check.
	RelatedResourceId *string `mandatory:"true" json:"relatedResourceId"`

	// Type of resource related to corresponding check.
	RelatedResourceType RelatedResourceTypeEnum `mandatory:"true" json:"relatedResourceType"`

	// The code returned when GoldenGate reports an error while running a step during pipeline initialization.
	// https://docs.oracle.com/en/middleware/goldengate/core/23/error-messages/ogg-00001-ogg-40000.html#GUID-97FF7AA7-7A5C-4AA7-B29F-3CC8D26761F2
	Code *string `mandatory:"true" json:"code"`

	// The corrective action for non-passing checks. Null for passed checks.
	CorrectiveAction *string `mandatory:"false" json:"correctiveAction"`
}

func (m DisasterRecoveryPrecheckResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryPrecheckResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrecheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPrecheckStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRelatedResourceTypeEnum(string(m.RelatedResourceType)); !ok && m.RelatedResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RelatedResourceType: %s. Supported values are: %s.", m.RelatedResourceType, strings.Join(GetRelatedResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
