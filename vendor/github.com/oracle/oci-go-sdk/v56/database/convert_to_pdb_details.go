// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ConvertToPdbDetails Details for converting a non-container database to pluggable database.
type ConvertToPdbDetails struct {

	// The operations used to convert a non-container database to a pluggable database.
	// - Use `PRECHECK` to run a pre-check operation on non-container database prior to converting it into a pluggable database.
	// - Use `CONVERT` to convert a non-container database into a pluggable database.
	// - Use `SYNC` if the non-container database was manually converted into a pluggable database using the dbcli command-line utility. Databases may need to be converted manually if the CONVERT action fails when converting a non-container database using the API.
	// - Use `SYNC_ROLLBACK` if the conversion of a non-container database into a pluggable database was manually rolled back using the dbcli command line utility. Conversions may need to be manually rolled back if the CONVERT action fails when converting a non-container database using the API.
	Action ConvertToPdbDetailsActionEnum `mandatory:"true" json:"action"`

	ConvertToPdbTargetDetails ConvertToPdbTargetBase `mandatory:"false" json:"convertToPdbTargetDetails"`
}

func (m ConvertToPdbDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ConvertToPdbDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConvertToPdbTargetDetails converttopdbtargetbase        `json:"convertToPdbTargetDetails"`
		Action                    ConvertToPdbDetailsActionEnum `json:"action"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConvertToPdbTargetDetails.UnmarshalPolymorphicJSON(model.ConvertToPdbTargetDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConvertToPdbTargetDetails = nn.(ConvertToPdbTargetBase)
	} else {
		m.ConvertToPdbTargetDetails = nil
	}

	m.Action = model.Action

	return
}

// ConvertToPdbDetailsActionEnum Enum with underlying type: string
type ConvertToPdbDetailsActionEnum string

// Set of constants representing the allowable values for ConvertToPdbDetailsActionEnum
const (
	ConvertToPdbDetailsActionPrecheck     ConvertToPdbDetailsActionEnum = "PRECHECK"
	ConvertToPdbDetailsActionConvert      ConvertToPdbDetailsActionEnum = "CONVERT"
	ConvertToPdbDetailsActionSync         ConvertToPdbDetailsActionEnum = "SYNC"
	ConvertToPdbDetailsActionSyncRollback ConvertToPdbDetailsActionEnum = "SYNC_ROLLBACK"
)

var mappingConvertToPdbDetailsAction = map[string]ConvertToPdbDetailsActionEnum{
	"PRECHECK":      ConvertToPdbDetailsActionPrecheck,
	"CONVERT":       ConvertToPdbDetailsActionConvert,
	"SYNC":          ConvertToPdbDetailsActionSync,
	"SYNC_ROLLBACK": ConvertToPdbDetailsActionSyncRollback,
}

// GetConvertToPdbDetailsActionEnumValues Enumerates the set of values for ConvertToPdbDetailsActionEnum
func GetConvertToPdbDetailsActionEnumValues() []ConvertToPdbDetailsActionEnum {
	values := make([]ConvertToPdbDetailsActionEnum, 0)
	for _, v := range mappingConvertToPdbDetailsAction {
		values = append(values, v)
	}
	return values
}
