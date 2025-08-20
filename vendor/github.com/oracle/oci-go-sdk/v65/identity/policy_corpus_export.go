// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PolicyCorpusExport Resource object containing details of the Policy Corpus Export task.
type PolicyCorpusExport struct {

	// Id of the Policy Corpus Export task.
	Id *string `mandatory:"true" json:"id"`

	// CompartmentId whose policy corpus are to be exported for Policy Simulation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Status of Policy Corpus Export task.
	LifecycleState PolicyCorpusExportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	PolicyCorpusUploadPath *ObjectStorageFileUploadLocation `mandatory:"true" json:"policyCorpusUploadPath"`

	// Time when Policy Corpus export task was initiated.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when Policy Corpus export task was last modified.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Message detailing failure in Policy Corpus export task.
	FailureMessage *string `mandatory:"false" json:"failureMessage"`
}

func (m PolicyCorpusExport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PolicyCorpusExport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPolicyCorpusExportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPolicyCorpusExportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PolicyCorpusExportLifecycleStateEnum Enum with underlying type: string
type PolicyCorpusExportLifecycleStateEnum string

// Set of constants representing the allowable values for PolicyCorpusExportLifecycleStateEnum
const (
	PolicyCorpusExportLifecycleStateAccepted   PolicyCorpusExportLifecycleStateEnum = "ACCEPTED"
	PolicyCorpusExportLifecycleStateWaiting    PolicyCorpusExportLifecycleStateEnum = "WAITING"
	PolicyCorpusExportLifecycleStateInProgress PolicyCorpusExportLifecycleStateEnum = "IN_PROGRESS"
	PolicyCorpusExportLifecycleStateFailed     PolicyCorpusExportLifecycleStateEnum = "FAILED"
	PolicyCorpusExportLifecycleStateSucceeded  PolicyCorpusExportLifecycleStateEnum = "SUCCEEDED"
	PolicyCorpusExportLifecycleStateCanceling  PolicyCorpusExportLifecycleStateEnum = "CANCELING"
)

var mappingPolicyCorpusExportLifecycleStateEnum = map[string]PolicyCorpusExportLifecycleStateEnum{
	"ACCEPTED":    PolicyCorpusExportLifecycleStateAccepted,
	"WAITING":     PolicyCorpusExportLifecycleStateWaiting,
	"IN_PROGRESS": PolicyCorpusExportLifecycleStateInProgress,
	"FAILED":      PolicyCorpusExportLifecycleStateFailed,
	"SUCCEEDED":   PolicyCorpusExportLifecycleStateSucceeded,
	"CANCELING":   PolicyCorpusExportLifecycleStateCanceling,
}

var mappingPolicyCorpusExportLifecycleStateEnumLowerCase = map[string]PolicyCorpusExportLifecycleStateEnum{
	"accepted":    PolicyCorpusExportLifecycleStateAccepted,
	"waiting":     PolicyCorpusExportLifecycleStateWaiting,
	"in_progress": PolicyCorpusExportLifecycleStateInProgress,
	"failed":      PolicyCorpusExportLifecycleStateFailed,
	"succeeded":   PolicyCorpusExportLifecycleStateSucceeded,
	"canceling":   PolicyCorpusExportLifecycleStateCanceling,
}

// GetPolicyCorpusExportLifecycleStateEnumValues Enumerates the set of values for PolicyCorpusExportLifecycleStateEnum
func GetPolicyCorpusExportLifecycleStateEnumValues() []PolicyCorpusExportLifecycleStateEnum {
	values := make([]PolicyCorpusExportLifecycleStateEnum, 0)
	for _, v := range mappingPolicyCorpusExportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPolicyCorpusExportLifecycleStateEnumStringValues Enumerates the set of values in String for PolicyCorpusExportLifecycleStateEnum
func GetPolicyCorpusExportLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
	}
}

// GetMappingPolicyCorpusExportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPolicyCorpusExportLifecycleStateEnum(val string) (PolicyCorpusExportLifecycleStateEnum, bool) {
	enum, ok := mappingPolicyCorpusExportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
