// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateEmailPrivateEndpointDetails Information about the new private endpoint resource.
type CreateEmailPrivateEndpointDetails struct {

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the customer subnet where the private endpoint will be created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Private endpoint type, to differentiate between requirements of each.
	PrivateEndpointType CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum `mandatory:"true" json:"privateEndpointType"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to.
	// For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - The nsgIds array is optional.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Optional IP within the provided subnet for the Private Endpoint to use.
	// If not provided, PECP will choose from the given subnet.
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`

	// Optional IP within the provided subnet to use for the reverse connection source IP to customer smarthost.
	// If not provided, PECP will choose from the given subnet.
	PrivateEndpointSourceIp *string `mandatory:"false" json:"privateEndpointSourceIp"`

	// Customer provided smart host that will route emails from Email Delivery in the case where delivery is needed in addition to submission.
	// Used when the privateEndpoint type is SUBMISSION
	SmartHost *string `mandatory:"false" json:"smartHost"`

	// Display name of the private endpoint resource being created.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A string that describes the details about the private endpoint.
	// It does not have to be unique, and you can change it.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateEmailPrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEmailPrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum(string(m.PrivateEndpointType)); !ok && m.PrivateEndpointType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivateEndpointType: %s. Supported values are: %s.", m.PrivateEndpointType, strings.Join(GetCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum Enum with underlying type: string
type CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum string

// Set of constants representing the allowable values for CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum
const (
	CreateEmailPrivateEndpointDetailsPrivateEndpointTypeSubmission CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum = "SUBMISSION"
	CreateEmailPrivateEndpointDetailsPrivateEndpointTypeBounce     CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum = "BOUNCE"
)

var mappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum = map[string]CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum{
	"SUBMISSION": CreateEmailPrivateEndpointDetailsPrivateEndpointTypeSubmission,
	"BOUNCE":     CreateEmailPrivateEndpointDetailsPrivateEndpointTypeBounce,
}

var mappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumLowerCase = map[string]CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum{
	"submission": CreateEmailPrivateEndpointDetailsPrivateEndpointTypeSubmission,
	"bounce":     CreateEmailPrivateEndpointDetailsPrivateEndpointTypeBounce,
}

// GetCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumValues Enumerates the set of values for CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum
func GetCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumValues() []CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum {
	values := make([]CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum, 0)
	for _, v := range mappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumStringValues Enumerates the set of values in String for CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum
func GetCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumStringValues() []string {
	return []string{
		"SUBMISSION",
		"BOUNCE",
	}
}

// GetMappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum(val string) (CreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnum, bool) {
	enum, ok := mappingCreateEmailPrivateEndpointDetailsPrivateEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
