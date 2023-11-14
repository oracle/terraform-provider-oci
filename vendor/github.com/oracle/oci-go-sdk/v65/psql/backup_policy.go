// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// A description of the PGSQL Control Plane API
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupPolicy Posgresql DB system backup policy
type BackupPolicy interface {

	// How many days the customers data should be stored after the db system deletion.
	GetRetentionDays() *int
}

type backuppolicy struct {
	JsonData      []byte
	RetentionDays *int   `mandatory:"false" json:"retentionDays"`
	Kind          string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *backuppolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbackuppolicy backuppolicy
	s := struct {
		Model Unmarshalerbackuppolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RetentionDays = s.Model.RetentionDays
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *backuppolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "DAILY":
		mm := DailyBackupPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEEKLY":
		mm := WeeklyBackupPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := NoneBackupPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MONTHLY":
		mm := MonthlyBackupPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BackupPolicy: %s.", m.Kind)
		return *m, nil
	}
}

// GetRetentionDays returns RetentionDays
func (m backuppolicy) GetRetentionDays() *int {
	return m.RetentionDays
}

func (m backuppolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m backuppolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupPolicyKindEnum Enum with underlying type: string
type BackupPolicyKindEnum string

// Set of constants representing the allowable values for BackupPolicyKindEnum
const (
	BackupPolicyKindDaily   BackupPolicyKindEnum = "DAILY"
	BackupPolicyKindWeekly  BackupPolicyKindEnum = "WEEKLY"
	BackupPolicyKindMonthly BackupPolicyKindEnum = "MONTHLY"
	BackupPolicyKindNone    BackupPolicyKindEnum = "NONE"
)

var mappingBackupPolicyKindEnum = map[string]BackupPolicyKindEnum{
	"DAILY":   BackupPolicyKindDaily,
	"WEEKLY":  BackupPolicyKindWeekly,
	"MONTHLY": BackupPolicyKindMonthly,
	"NONE":    BackupPolicyKindNone,
}

var mappingBackupPolicyKindEnumLowerCase = map[string]BackupPolicyKindEnum{
	"daily":   BackupPolicyKindDaily,
	"weekly":  BackupPolicyKindWeekly,
	"monthly": BackupPolicyKindMonthly,
	"none":    BackupPolicyKindNone,
}

// GetBackupPolicyKindEnumValues Enumerates the set of values for BackupPolicyKindEnum
func GetBackupPolicyKindEnumValues() []BackupPolicyKindEnum {
	values := make([]BackupPolicyKindEnum, 0)
	for _, v := range mappingBackupPolicyKindEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupPolicyKindEnumStringValues Enumerates the set of values in String for BackupPolicyKindEnum
func GetBackupPolicyKindEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"NONE",
	}
}

// GetMappingBackupPolicyKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupPolicyKindEnum(val string) (BackupPolicyKindEnum, bool) {
	enum, ok := mappingBackupPolicyKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
