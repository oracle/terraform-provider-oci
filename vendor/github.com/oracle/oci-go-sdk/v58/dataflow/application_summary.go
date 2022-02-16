// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ApplicationSummary A Data Flow application object used in bulk listings.
type ApplicationSummary struct {

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A user-friendly name. This name is not necessarily unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The application ID.
	Id *string `mandatory:"true" json:"id"`

	// The Spark language.
	Language ApplicationLanguageEnum `mandatory:"true" json:"language"`

	// The current state of this application.
	LifecycleState ApplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"true" json:"ownerPrincipalId"`

	// The Spark version utilized to run the application.
	SparkVersion *string `mandatory:"true" json:"sparkVersion"`

	// The date and time a application was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time a application was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The username of the user who created the resource.  If the username of the owner does not exist,
	// `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName *string `mandatory:"false" json:"ownerUserName"`

	// The Spark application processing type.
	Type ApplicationTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m ApplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationLanguageEnum(string(m.Language)); !ok && m.Language != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Language: %s. Supported values are: %s.", m.Language, strings.Join(GetApplicationLanguageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingApplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplicationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingApplicationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetApplicationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
