// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SightingSummary Sighting summary Definition.
type SightingSummary struct {

	// Unique identifier for finding event
	Id *string `mandatory:"true" json:"id"`

	// Compartment Id where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier of the rule
	DetectorRuleId *string `mandatory:"true" json:"detectorRuleId"`

	// ClassificationStatus of the sighting event
	ClassificationStatus ClassificationStatusEnum `mandatory:"true" json:"classificationStatus"`

	// Identifier for the sighting type
	SightingType *string `mandatory:"true" json:"sightingType"`

	// Name of the sighting type
	SightingTypeDisplayName *string `mandatory:"true" json:"sightingTypeDisplayName"`

	// Name of the Mitre att&ck tactic
	TacticName *string `mandatory:"true" json:"tacticName"`

	// Name of the Mitre att&ck technique
	TechniqueName *string `mandatory:"true" json:"techniqueName"`

	// Score for the sighting
	SightingScore *int `mandatory:"true" json:"sightingScore"`

	// Severity of the sighting
	Severity SeverityEnum `mandatory:"true" json:"severity"`

	// Confidence of the sighting
	Confidence ConfidenceEnum `mandatory:"true" json:"confidence"`

	// The date and time the sighting was first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"true" json:"timeFirstDetected"`

	// The date and time the sighting was last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"true" json:"timeLastDetected"`

	// Regions involved in the sighting
	Regions []string `mandatory:"true" json:"regions"`

	// Problem Id to which the Sighting is associated
	ProblemId *string `mandatory:"false" json:"problemId"`

	// Unique identifier for principal actor
	ActorPrincipalId *string `mandatory:"false" json:"actorPrincipalId"`

	// Name of the principal actor
	ActorPrincipalName *string `mandatory:"false" json:"actorPrincipalName"`

	// Type of the principal actor
	ActorPrincipalType *string `mandatory:"false" json:"actorPrincipalType"`
}

func (m SightingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SightingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClassificationStatusEnum(string(m.ClassificationStatus)); !ok && m.ClassificationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClassificationStatus: %s. Supported values are: %s.", m.ClassificationStatus, strings.Join(GetClassificationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingConfidenceEnum(string(m.Confidence)); !ok && m.Confidence != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Confidence: %s. Supported values are: %s.", m.Confidence, strings.Join(GetConfidenceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
