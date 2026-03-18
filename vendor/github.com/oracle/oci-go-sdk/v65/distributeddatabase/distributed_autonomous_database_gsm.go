// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedAutonomousDatabaseGsm Details of global service manager(GSM also known as shard director) instances for the Globally distributed autonomous database.
type DistributedAutonomousDatabaseGsm struct {

	// Name of the Global service manager instance
	Name *string `mandatory:"true" json:"name"`

	// The compute count for the Global service manager instance.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the Global service manager instance.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The time the Global service manager instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Global service manager instance was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Status of the gsm for the Globally distributed autonomous database.
	Status DistributedAutonomousDatabaseGsmStatusEnum `mandatory:"true" json:"status"`

	// The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string
	TimeSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeSslCertificateExpires"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	GsmImageDetails *DistributedAutonomousDatabaseGsmImage `mandatory:"false" json:"gsmImageDetails"`

	Metadata *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`
}

func (m DistributedAutonomousDatabaseGsm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseGsm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseGsmStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedAutonomousDatabaseGsmStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedAutonomousDatabaseGsmStatusEnum Enum with underlying type: string
type DistributedAutonomousDatabaseGsmStatusEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseGsmStatusEnum
const (
	DistributedAutonomousDatabaseGsmStatusFailed                DistributedAutonomousDatabaseGsmStatusEnum = "FAILED"
	DistributedAutonomousDatabaseGsmStatusDeleting              DistributedAutonomousDatabaseGsmStatusEnum = "DELETING"
	DistributedAutonomousDatabaseGsmStatusDeleted               DistributedAutonomousDatabaseGsmStatusEnum = "DELETED"
	DistributedAutonomousDatabaseGsmStatusUpdating              DistributedAutonomousDatabaseGsmStatusEnum = "UPDATING"
	DistributedAutonomousDatabaseGsmStatusCreating              DistributedAutonomousDatabaseGsmStatusEnum = "CREATING"
	DistributedAutonomousDatabaseGsmStatusCreated               DistributedAutonomousDatabaseGsmStatusEnum = "CREATED"
	DistributedAutonomousDatabaseGsmStatusReadyForConfiguration DistributedAutonomousDatabaseGsmStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedAutonomousDatabaseGsmStatusConfigured            DistributedAutonomousDatabaseGsmStatusEnum = "CONFIGURED"
	DistributedAutonomousDatabaseGsmStatusNeedsAttention        DistributedAutonomousDatabaseGsmStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedAutonomousDatabaseGsmStatusEnum = map[string]DistributedAutonomousDatabaseGsmStatusEnum{
	"FAILED":                  DistributedAutonomousDatabaseGsmStatusFailed,
	"DELETING":                DistributedAutonomousDatabaseGsmStatusDeleting,
	"DELETED":                 DistributedAutonomousDatabaseGsmStatusDeleted,
	"UPDATING":                DistributedAutonomousDatabaseGsmStatusUpdating,
	"CREATING":                DistributedAutonomousDatabaseGsmStatusCreating,
	"CREATED":                 DistributedAutonomousDatabaseGsmStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedAutonomousDatabaseGsmStatusReadyForConfiguration,
	"CONFIGURED":              DistributedAutonomousDatabaseGsmStatusConfigured,
	"NEEDS_ATTENTION":         DistributedAutonomousDatabaseGsmStatusNeedsAttention,
}

var mappingDistributedAutonomousDatabaseGsmStatusEnumLowerCase = map[string]DistributedAutonomousDatabaseGsmStatusEnum{
	"failed":                  DistributedAutonomousDatabaseGsmStatusFailed,
	"deleting":                DistributedAutonomousDatabaseGsmStatusDeleting,
	"deleted":                 DistributedAutonomousDatabaseGsmStatusDeleted,
	"updating":                DistributedAutonomousDatabaseGsmStatusUpdating,
	"creating":                DistributedAutonomousDatabaseGsmStatusCreating,
	"created":                 DistributedAutonomousDatabaseGsmStatusCreated,
	"ready_for_configuration": DistributedAutonomousDatabaseGsmStatusReadyForConfiguration,
	"configured":              DistributedAutonomousDatabaseGsmStatusConfigured,
	"needs_attention":         DistributedAutonomousDatabaseGsmStatusNeedsAttention,
}

// GetDistributedAutonomousDatabaseGsmStatusEnumValues Enumerates the set of values for DistributedAutonomousDatabaseGsmStatusEnum
func GetDistributedAutonomousDatabaseGsmStatusEnumValues() []DistributedAutonomousDatabaseGsmStatusEnum {
	values := make([]DistributedAutonomousDatabaseGsmStatusEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseGsmStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseGsmStatusEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseGsmStatusEnum
func GetDistributedAutonomousDatabaseGsmStatusEnumStringValues() []string {
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

// GetMappingDistributedAutonomousDatabaseGsmStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseGsmStatusEnum(val string) (DistributedAutonomousDatabaseGsmStatusEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseGsmStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
