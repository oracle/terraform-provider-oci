// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceEvent The details of a maintenance event.
type MaintenanceEvent struct {

	// The OCID of the DB System this maintenance event is associated with.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID of the compartment the maintenance event belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the record was created,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the maintenance event is expected to start,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// The MySQL version prior to the maintenance.
	MysqlVersionPre *string `mandatory:"true" json:"mysqlVersionPre"`

	// The MySQL version after the maintenance.
	MysqlVersionPost *string `mandatory:"true" json:"mysqlVersionPost"`

	// How the maintenance event was triggered.
	// AUTOMATIC:  maintenance event triggered as part of scheduled maintenance.
	// MANUAL:  maintenance event triggered manually.
	// SHAPE:     maintenance event triggered by a shape update.
	MaintenanceType MaintenanceTypeEnum `mandatory:"true" json:"maintenanceType"`

	// type: string
	// description: |
	//   The nature of the maintenance event.
	//   DATABASE:  maintenance event causing a MySQL version upgrade.
	//   OS_UPDATE: maintenance event causing an OS update.
	MaintenanceAction MaintenanceActionEnum `mandatory:"true" json:"maintenanceAction"`

	// The last status of the maintenance event.
	MaintenanceStatus MaintenanceEventMaintenanceStatusEnum `mandatory:"true" json:"maintenanceStatus"`

	// Information regarding what was performed during that maintenance.
	MaintenanceNotes *string `mandatory:"false" json:"maintenanceNotes"`

	// The date and time the maintenance event started,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the maintenance event ended,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The date and time the DB System was initially down during the maintenance,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeMysqlSwitchOverStarted *common.SDKTime `mandatory:"false" json:"timeMysqlSwitchOverStarted"`

	// The date and time the DB System came back online during the maintenance,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeMysqlSwitchOverEnded *common.SDKTime `mandatory:"false" json:"timeMysqlSwitchOverEnded"`
}

func (m MaintenanceEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceTypeEnum(string(m.MaintenanceType)); !ok && m.MaintenanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceType: %s. Supported values are: %s.", m.MaintenanceType, strings.Join(GetMaintenanceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceActionEnum(string(m.MaintenanceAction)); !ok && m.MaintenanceAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceAction: %s. Supported values are: %s.", m.MaintenanceAction, strings.Join(GetMaintenanceActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMaintenanceEventMaintenanceStatusEnum(string(m.MaintenanceStatus)); !ok && m.MaintenanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceStatus: %s. Supported values are: %s.", m.MaintenanceStatus, strings.Join(GetMaintenanceEventMaintenanceStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceEventMaintenanceStatusEnum Enum with underlying type: string
type MaintenanceEventMaintenanceStatusEnum string

// Set of constants representing the allowable values for MaintenanceEventMaintenanceStatusEnum
const (
	MaintenanceEventMaintenanceStatusSucceeded MaintenanceEventMaintenanceStatusEnum = "SUCCEEDED"
	MaintenanceEventMaintenanceStatusFailed    MaintenanceEventMaintenanceStatusEnum = "FAILED"
	MaintenanceEventMaintenanceStatusCanceled  MaintenanceEventMaintenanceStatusEnum = "CANCELED"
)

var mappingMaintenanceEventMaintenanceStatusEnum = map[string]MaintenanceEventMaintenanceStatusEnum{
	"SUCCEEDED": MaintenanceEventMaintenanceStatusSucceeded,
	"FAILED":    MaintenanceEventMaintenanceStatusFailed,
	"CANCELED":  MaintenanceEventMaintenanceStatusCanceled,
}

var mappingMaintenanceEventMaintenanceStatusEnumLowerCase = map[string]MaintenanceEventMaintenanceStatusEnum{
	"succeeded": MaintenanceEventMaintenanceStatusSucceeded,
	"failed":    MaintenanceEventMaintenanceStatusFailed,
	"canceled":  MaintenanceEventMaintenanceStatusCanceled,
}

// GetMaintenanceEventMaintenanceStatusEnumValues Enumerates the set of values for MaintenanceEventMaintenanceStatusEnum
func GetMaintenanceEventMaintenanceStatusEnumValues() []MaintenanceEventMaintenanceStatusEnum {
	values := make([]MaintenanceEventMaintenanceStatusEnum, 0)
	for _, v := range mappingMaintenanceEventMaintenanceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceEventMaintenanceStatusEnumStringValues Enumerates the set of values in String for MaintenanceEventMaintenanceStatusEnum
func GetMaintenanceEventMaintenanceStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

// GetMappingMaintenanceEventMaintenanceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceEventMaintenanceStatusEnum(val string) (MaintenanceEventMaintenanceStatusEnum, bool) {
	enum, ok := mappingMaintenanceEventMaintenanceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
