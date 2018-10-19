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

// InstanceConfigurationInstanceDetails The representation of InstanceConfigurationInstanceDetails
type InstanceConfigurationInstanceDetails interface {
}

type instanceconfigurationinstancedetails struct {
	JsonData     []byte
	InstanceType string `json:"instanceType"`
}

// UnmarshalJSON unmarshals json
func (m *instanceconfigurationinstancedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceconfigurationinstancedetails instanceconfigurationinstancedetails
	s := struct {
		Model Unmarshalerinstanceconfigurationinstancedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InstanceType = s.Model.InstanceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceconfigurationinstancedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.InstanceType {
	case "compute":
		mm := ComputeInstanceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m instanceconfigurationinstancedetails) String() string {
	return common.PointerString(m)
}
