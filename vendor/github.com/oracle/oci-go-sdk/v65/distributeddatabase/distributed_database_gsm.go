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

// DistributedDatabaseGsm Details of global service manager(GSM also known as shard director) instances for the Globally distributed database.
type DistributedDatabaseGsm struct {

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

	// Status of the gsm.
	Status DistributedDatabaseGsmStatusEnum `mandatory:"true" json:"status"`

	// The time the ssl certificate associated with Global service manager expires. An RFC3339 formatted datetime string
	TimeSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeSslCertificateExpires"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	GsmImageDetails *DistributedDbGsmImage `mandatory:"false" json:"gsmImageDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m DistributedDatabaseGsm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseGsm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseGsmStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseGsmStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DistributedDatabaseGsmStatusEnum Enum with underlying type: string
type DistributedDatabaseGsmStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseGsmStatusEnum
const (
	DistributedDatabaseGsmStatusFailed                DistributedDatabaseGsmStatusEnum = "FAILED"
	DistributedDatabaseGsmStatusDeleting              DistributedDatabaseGsmStatusEnum = "DELETING"
	DistributedDatabaseGsmStatusDeleted               DistributedDatabaseGsmStatusEnum = "DELETED"
	DistributedDatabaseGsmStatusUpdating              DistributedDatabaseGsmStatusEnum = "UPDATING"
	DistributedDatabaseGsmStatusCreating              DistributedDatabaseGsmStatusEnum = "CREATING"
	DistributedDatabaseGsmStatusCreated               DistributedDatabaseGsmStatusEnum = "CREATED"
	DistributedDatabaseGsmStatusReadyForConfiguration DistributedDatabaseGsmStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseGsmStatusConfigured            DistributedDatabaseGsmStatusEnum = "CONFIGURED"
	DistributedDatabaseGsmStatusNeedsAttention        DistributedDatabaseGsmStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseGsmStatusEnum = map[string]DistributedDatabaseGsmStatusEnum{
	"FAILED":                  DistributedDatabaseGsmStatusFailed,
	"DELETING":                DistributedDatabaseGsmStatusDeleting,
	"DELETED":                 DistributedDatabaseGsmStatusDeleted,
	"UPDATING":                DistributedDatabaseGsmStatusUpdating,
	"CREATING":                DistributedDatabaseGsmStatusCreating,
	"CREATED":                 DistributedDatabaseGsmStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseGsmStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseGsmStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseGsmStatusNeedsAttention,
}

var mappingDistributedDatabaseGsmStatusEnumLowerCase = map[string]DistributedDatabaseGsmStatusEnum{
	"failed":                  DistributedDatabaseGsmStatusFailed,
	"deleting":                DistributedDatabaseGsmStatusDeleting,
	"deleted":                 DistributedDatabaseGsmStatusDeleted,
	"updating":                DistributedDatabaseGsmStatusUpdating,
	"creating":                DistributedDatabaseGsmStatusCreating,
	"created":                 DistributedDatabaseGsmStatusCreated,
	"ready_for_configuration": DistributedDatabaseGsmStatusReadyForConfiguration,
	"configured":              DistributedDatabaseGsmStatusConfigured,
	"needs_attention":         DistributedDatabaseGsmStatusNeedsAttention,
}

// GetDistributedDatabaseGsmStatusEnumValues Enumerates the set of values for DistributedDatabaseGsmStatusEnum
func GetDistributedDatabaseGsmStatusEnumValues() []DistributedDatabaseGsmStatusEnum {
	values := make([]DistributedDatabaseGsmStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseGsmStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseGsmStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseGsmStatusEnum
func GetDistributedDatabaseGsmStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseGsmStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseGsmStatusEnum(val string) (DistributedDatabaseGsmStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseGsmStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
