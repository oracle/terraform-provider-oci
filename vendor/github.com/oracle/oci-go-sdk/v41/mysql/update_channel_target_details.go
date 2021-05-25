// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
)

// UpdateChannelTargetDetails Parameters detailing how to provision the target for the given Channel.
type UpdateChannelTargetDetails interface {
}

type updatechanneltargetdetails struct {
	JsonData   []byte
	TargetType string `json:"targetType"`
}

// UnmarshalJSON unmarshals json
func (m *updatechanneltargetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatechanneltargetdetails updatechanneltargetdetails
	s := struct {
		Model Unmarshalerupdatechanneltargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TargetType = s.Model.TargetType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatechanneltargetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TargetType {
	case "DBSYSTEM":
		mm := UpdateChannelTargetFromDbSystemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updatechanneltargetdetails) String() string {
	return common.PointerString(m)
}
