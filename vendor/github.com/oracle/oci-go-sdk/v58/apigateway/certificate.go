// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Certificate A certificate contains information to be installed on a gateway to secure the traffic going
// through it.
// For more information, see
// API Gateway Concepts (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayconcepts.htm).
type Certificate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The entity to be secured by the certificate and additional host names.
	SubjectNames []string `mandatory:"true" json:"subjectNames"`

	// The date and time the certificate will expire.
	TimeNotValidAfter *common.SDKTime `mandatory:"true" json:"timeNotValidAfter"`

	// The data of the leaf certificate in pem format.
	Certificate *string `mandatory:"true" json:"certificate"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The intermediate certificate data associated with the certificate in pem format.
	IntermediateCertificates *string `mandatory:"false" json:"intermediateCertificates"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the certificate.
	LifecycleState CertificateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Certificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Certificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCertificateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCertificateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CertificateLifecycleStateEnum Enum with underlying type: string
type CertificateLifecycleStateEnum string

// Set of constants representing the allowable values for CertificateLifecycleStateEnum
const (
	CertificateLifecycleStateCreating CertificateLifecycleStateEnum = "CREATING"
	CertificateLifecycleStateActive   CertificateLifecycleStateEnum = "ACTIVE"
	CertificateLifecycleStateUpdating CertificateLifecycleStateEnum = "UPDATING"
	CertificateLifecycleStateDeleting CertificateLifecycleStateEnum = "DELETING"
	CertificateLifecycleStateDeleted  CertificateLifecycleStateEnum = "DELETED"
	CertificateLifecycleStateFailed   CertificateLifecycleStateEnum = "FAILED"
)

var mappingCertificateLifecycleStateEnum = map[string]CertificateLifecycleStateEnum{
	"CREATING": CertificateLifecycleStateCreating,
	"ACTIVE":   CertificateLifecycleStateActive,
	"UPDATING": CertificateLifecycleStateUpdating,
	"DELETING": CertificateLifecycleStateDeleting,
	"DELETED":  CertificateLifecycleStateDeleted,
	"FAILED":   CertificateLifecycleStateFailed,
}

// GetCertificateLifecycleStateEnumValues Enumerates the set of values for CertificateLifecycleStateEnum
func GetCertificateLifecycleStateEnumValues() []CertificateLifecycleStateEnum {
	values := make([]CertificateLifecycleStateEnum, 0)
	for _, v := range mappingCertificateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateLifecycleStateEnumStringValues Enumerates the set of values in String for CertificateLifecycleStateEnum
func GetCertificateLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCertificateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateLifecycleStateEnum(val string) (CertificateLifecycleStateEnum, bool) {
	mappingCertificateLifecycleStateEnumIgnoreCase := make(map[string]CertificateLifecycleStateEnum)
	for k, v := range mappingCertificateLifecycleStateEnum {
		mappingCertificateLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCertificateLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
