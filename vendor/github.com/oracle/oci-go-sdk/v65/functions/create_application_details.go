// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateApplicationDetails Properties for a new application.
type CreateApplicationDetails struct {

	// The OCID of the compartment to create the application within.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the application. The display name must be unique within the compartment containing the application. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application.
	SubnetIds []string `mandatory:"true" json:"subnetIds"`

	// Application configuration. These values are passed on to the function as environment variables, functions may override application configuration.
	// Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.
	// Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	// The maximum size for all configuration keys and values is limited to 4KB. This is measured as the sum of octets necessary to represent each key and value in UTF-8.
	Config map[string]string `mandatory:"false" json:"config"`

	// Valid values are `GENERIC_X86`, `GENERIC_ARM` and `GENERIC_X86_ARM`. Default is `GENERIC_X86`. Setting this to `GENERIC_X86`, will run the functions in the application on X86 processor architecture.
	// Setting this to `GENERIC_ARM`, will run the functions in the application on ARM processor architecture.
	// When set to `GENERIC_X86_ARM`, functions in the application are run on either X86 or ARM processor architecture.
	// Accepted values are:
	// `GENERIC_X86`, `GENERIC_ARM`, `GENERIC_X86_ARM`
	Shape CreateApplicationDetailsShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the Network Security Groups to add the application to.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	// A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls.
	// The syslog URL must be reachable from all of the subnets configured for the application.
	// Note: If you enable the OCI Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the OCI Logging service, and not to the syslog URL.
	// Example: `tcp://logserver.myserver:1234`
	SyslogUrl *string `mandatory:"false" json:"syslogUrl"`

	TraceConfig *ApplicationTraceConfig `mandatory:"false" json:"traceConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	ImagePolicyConfig *ImagePolicyConfig `mandatory:"false" json:"imagePolicyConfig"`
}

func (m CreateApplicationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateApplicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateApplicationDetailsShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetCreateApplicationDetailsShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateApplicationDetailsShapeEnum Enum with underlying type: string
type CreateApplicationDetailsShapeEnum string

// Set of constants representing the allowable values for CreateApplicationDetailsShapeEnum
const (
	CreateApplicationDetailsShapeX86    CreateApplicationDetailsShapeEnum = "GENERIC_X86"
	CreateApplicationDetailsShapeArm    CreateApplicationDetailsShapeEnum = "GENERIC_ARM"
	CreateApplicationDetailsShapeX86Arm CreateApplicationDetailsShapeEnum = "GENERIC_X86_ARM"
)

var mappingCreateApplicationDetailsShapeEnum = map[string]CreateApplicationDetailsShapeEnum{
	"GENERIC_X86":     CreateApplicationDetailsShapeX86,
	"GENERIC_ARM":     CreateApplicationDetailsShapeArm,
	"GENERIC_X86_ARM": CreateApplicationDetailsShapeX86Arm,
}

var mappingCreateApplicationDetailsShapeEnumLowerCase = map[string]CreateApplicationDetailsShapeEnum{
	"generic_x86":     CreateApplicationDetailsShapeX86,
	"generic_arm":     CreateApplicationDetailsShapeArm,
	"generic_x86_arm": CreateApplicationDetailsShapeX86Arm,
}

// GetCreateApplicationDetailsShapeEnumValues Enumerates the set of values for CreateApplicationDetailsShapeEnum
func GetCreateApplicationDetailsShapeEnumValues() []CreateApplicationDetailsShapeEnum {
	values := make([]CreateApplicationDetailsShapeEnum, 0)
	for _, v := range mappingCreateApplicationDetailsShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateApplicationDetailsShapeEnumStringValues Enumerates the set of values in String for CreateApplicationDetailsShapeEnum
func GetCreateApplicationDetailsShapeEnumStringValues() []string {
	return []string{
		"GENERIC_X86",
		"GENERIC_ARM",
		"GENERIC_X86_ARM",
	}
}

// GetMappingCreateApplicationDetailsShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateApplicationDetailsShapeEnum(val string) (CreateApplicationDetailsShapeEnum, bool) {
	enum, ok := mappingCreateApplicationDetailsShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
