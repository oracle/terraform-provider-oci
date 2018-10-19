// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstanceConfigurationVolumeSourceDetails The representation of InstanceConfigurationVolumeSourceDetails
type InstanceConfigurationVolumeSourceDetails interface {
}

type instanceconfigurationvolumesourcedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *instanceconfigurationvolumesourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceconfigurationvolumesourcedetails instanceconfigurationvolumesourcedetails
	s := struct {
		Model Unmarshalerinstanceconfigurationvolumesourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceconfigurationvolumesourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "volumeBackup":
		mm := InstanceConfigurationVolumeSourceFromVolumeBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "volume":
		mm := InstanceConfigurationVolumeSourceFromVolumeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m instanceconfigurationvolumesourcedetails) String() string {
	return common.PointerString(m)
}
