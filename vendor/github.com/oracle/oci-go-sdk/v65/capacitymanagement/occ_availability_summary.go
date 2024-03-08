// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccAvailabilitySummary The details about the available capacity and constraints for different resource types present in the availability catalog.
type OccAvailabilitySummary struct {

	// The OCID of the availability catalog.
	CatalogId *string `mandatory:"true" json:"catalogId"`

	// The name of the OCI service in consideration. For example, Compute, Exadata, and so on.
	Namespace NamespaceEnum `mandatory:"true" json:"namespace"`

	// The date by which the customer must place the order to have their capacity requirements met by the customer handover date.
	DateFinalCustomerOrder *common.SDKTime `mandatory:"true" json:"dateFinalCustomerOrder"`

	// The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	DateExpectedCapacityHandover *common.SDKTime `mandatory:"true" json:"dateExpectedCapacityHandover"`

	// The different types of resources against which customers can place capacity requests.
	ResourceType OccAvailabilitySummaryResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The type of workload (Generic/ROW).
	WorkloadType OccAvailabilitySummaryWorkloadTypeEnum `mandatory:"true" json:"workloadType"`

	// The name of the resource that the customer can request.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The quantity of available resource that the customer can request.
	AvailableQuantity *int64 `mandatory:"true" json:"availableQuantity"`

	// The unit in which the resource available is measured.
	Unit *string `mandatory:"true" json:"unit"`
}

func (m OccAvailabilitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccAvailabilitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNamespaceEnum(string(m.Namespace)); !ok && m.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", m.Namespace, strings.Join(GetNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilitySummaryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetOccAvailabilitySummaryResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilitySummaryWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetOccAvailabilitySummaryWorkloadTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccAvailabilitySummaryResourceTypeEnum Enum with underlying type: string
type OccAvailabilitySummaryResourceTypeEnum string

// Set of constants representing the allowable values for OccAvailabilitySummaryResourceTypeEnum
const (
	OccAvailabilitySummaryResourceTypeServerHw           OccAvailabilitySummaryResourceTypeEnum = "SERVER_HW"
	OccAvailabilitySummaryResourceTypeCapacityConstraint OccAvailabilitySummaryResourceTypeEnum = "CAPACITY_CONSTRAINT"
)

var mappingOccAvailabilitySummaryResourceTypeEnum = map[string]OccAvailabilitySummaryResourceTypeEnum{
	"SERVER_HW":           OccAvailabilitySummaryResourceTypeServerHw,
	"CAPACITY_CONSTRAINT": OccAvailabilitySummaryResourceTypeCapacityConstraint,
}

var mappingOccAvailabilitySummaryResourceTypeEnumLowerCase = map[string]OccAvailabilitySummaryResourceTypeEnum{
	"server_hw":           OccAvailabilitySummaryResourceTypeServerHw,
	"capacity_constraint": OccAvailabilitySummaryResourceTypeCapacityConstraint,
}

// GetOccAvailabilitySummaryResourceTypeEnumValues Enumerates the set of values for OccAvailabilitySummaryResourceTypeEnum
func GetOccAvailabilitySummaryResourceTypeEnumValues() []OccAvailabilitySummaryResourceTypeEnum {
	values := make([]OccAvailabilitySummaryResourceTypeEnum, 0)
	for _, v := range mappingOccAvailabilitySummaryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccAvailabilitySummaryResourceTypeEnumStringValues Enumerates the set of values in String for OccAvailabilitySummaryResourceTypeEnum
func GetOccAvailabilitySummaryResourceTypeEnumStringValues() []string {
	return []string{
		"SERVER_HW",
		"CAPACITY_CONSTRAINT",
	}
}

// GetMappingOccAvailabilitySummaryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccAvailabilitySummaryResourceTypeEnum(val string) (OccAvailabilitySummaryResourceTypeEnum, bool) {
	enum, ok := mappingOccAvailabilitySummaryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccAvailabilitySummaryWorkloadTypeEnum Enum with underlying type: string
type OccAvailabilitySummaryWorkloadTypeEnum string

// Set of constants representing the allowable values for OccAvailabilitySummaryWorkloadTypeEnum
const (
	OccAvailabilitySummaryWorkloadTypeGeneric OccAvailabilitySummaryWorkloadTypeEnum = "GENERIC"
	OccAvailabilitySummaryWorkloadTypeRow     OccAvailabilitySummaryWorkloadTypeEnum = "ROW"
	OccAvailabilitySummaryWorkloadTypeUsProd  OccAvailabilitySummaryWorkloadTypeEnum = "US_PROD"
)

var mappingOccAvailabilitySummaryWorkloadTypeEnum = map[string]OccAvailabilitySummaryWorkloadTypeEnum{
	"GENERIC": OccAvailabilitySummaryWorkloadTypeGeneric,
	"ROW":     OccAvailabilitySummaryWorkloadTypeRow,
	"US_PROD": OccAvailabilitySummaryWorkloadTypeUsProd,
}

var mappingOccAvailabilitySummaryWorkloadTypeEnumLowerCase = map[string]OccAvailabilitySummaryWorkloadTypeEnum{
	"generic": OccAvailabilitySummaryWorkloadTypeGeneric,
	"row":     OccAvailabilitySummaryWorkloadTypeRow,
	"us_prod": OccAvailabilitySummaryWorkloadTypeUsProd,
}

// GetOccAvailabilitySummaryWorkloadTypeEnumValues Enumerates the set of values for OccAvailabilitySummaryWorkloadTypeEnum
func GetOccAvailabilitySummaryWorkloadTypeEnumValues() []OccAvailabilitySummaryWorkloadTypeEnum {
	values := make([]OccAvailabilitySummaryWorkloadTypeEnum, 0)
	for _, v := range mappingOccAvailabilitySummaryWorkloadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOccAvailabilitySummaryWorkloadTypeEnumStringValues Enumerates the set of values in String for OccAvailabilitySummaryWorkloadTypeEnum
func GetOccAvailabilitySummaryWorkloadTypeEnumStringValues() []string {
	return []string{
		"GENERIC",
		"ROW",
		"US_PROD",
	}
}

// GetMappingOccAvailabilitySummaryWorkloadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccAvailabilitySummaryWorkloadTypeEnum(val string) (OccAvailabilitySummaryWorkloadTypeEnum, bool) {
	enum, ok := mappingOccAvailabilitySummaryWorkloadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
