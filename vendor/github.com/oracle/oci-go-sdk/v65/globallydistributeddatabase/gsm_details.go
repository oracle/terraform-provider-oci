// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GsmDetails Details of global service manager(GSM also known as shard director) instances for sharded database.
type GsmDetails struct {

	// Name of the GSM instance
	Name *string `mandatory:"true" json:"name"`

	// The compute count for the GSM instance.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The time the GSM instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the GSM instance was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Status of shard or catalog or gsm for the sharded database.
	Status GsmDetailsStatusEnum `mandatory:"true" json:"status"`

	// The time the ssl certificate associated with GSM expires. An RFC3339 formatted datetime string
	TimeSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeSslCertificateExpires"`

	// Identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// Additional metadata related to GSM's underlying supporting resource.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`
}

func (m GsmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GsmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGsmDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetGsmDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GsmDetailsStatusEnum Enum with underlying type: string
type GsmDetailsStatusEnum string

// Set of constants representing the allowable values for GsmDetailsStatusEnum
const (
	GsmDetailsStatusFailed                GsmDetailsStatusEnum = "FAILED"
	GsmDetailsStatusDeleting              GsmDetailsStatusEnum = "DELETING"
	GsmDetailsStatusDeleted               GsmDetailsStatusEnum = "DELETED"
	GsmDetailsStatusUpdating              GsmDetailsStatusEnum = "UPDATING"
	GsmDetailsStatusCreating              GsmDetailsStatusEnum = "CREATING"
	GsmDetailsStatusCreated               GsmDetailsStatusEnum = "CREATED"
	GsmDetailsStatusReadyForConfiguration GsmDetailsStatusEnum = "READY_FOR_CONFIGURATION"
	GsmDetailsStatusConfigured            GsmDetailsStatusEnum = "CONFIGURED"
	GsmDetailsStatusNeedsAttention        GsmDetailsStatusEnum = "NEEDS_ATTENTION"
)

var mappingGsmDetailsStatusEnum = map[string]GsmDetailsStatusEnum{
	"FAILED":                  GsmDetailsStatusFailed,
	"DELETING":                GsmDetailsStatusDeleting,
	"DELETED":                 GsmDetailsStatusDeleted,
	"UPDATING":                GsmDetailsStatusUpdating,
	"CREATING":                GsmDetailsStatusCreating,
	"CREATED":                 GsmDetailsStatusCreated,
	"READY_FOR_CONFIGURATION": GsmDetailsStatusReadyForConfiguration,
	"CONFIGURED":              GsmDetailsStatusConfigured,
	"NEEDS_ATTENTION":         GsmDetailsStatusNeedsAttention,
}

var mappingGsmDetailsStatusEnumLowerCase = map[string]GsmDetailsStatusEnum{
	"failed":                  GsmDetailsStatusFailed,
	"deleting":                GsmDetailsStatusDeleting,
	"deleted":                 GsmDetailsStatusDeleted,
	"updating":                GsmDetailsStatusUpdating,
	"creating":                GsmDetailsStatusCreating,
	"created":                 GsmDetailsStatusCreated,
	"ready_for_configuration": GsmDetailsStatusReadyForConfiguration,
	"configured":              GsmDetailsStatusConfigured,
	"needs_attention":         GsmDetailsStatusNeedsAttention,
}

// GetGsmDetailsStatusEnumValues Enumerates the set of values for GsmDetailsStatusEnum
func GetGsmDetailsStatusEnumValues() []GsmDetailsStatusEnum {
	values := make([]GsmDetailsStatusEnum, 0)
	for _, v := range mappingGsmDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGsmDetailsStatusEnumStringValues Enumerates the set of values in String for GsmDetailsStatusEnum
func GetGsmDetailsStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"CREATED",
		"READY_FOR_CONFIGURATION",
		"CONFIGURED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingGsmDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGsmDetailsStatusEnum(val string) (GsmDetailsStatusEnum, bool) {
	enum, ok := mappingGsmDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
