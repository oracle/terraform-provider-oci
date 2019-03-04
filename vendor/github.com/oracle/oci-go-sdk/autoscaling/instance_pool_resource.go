// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InstancePoolResource An Instance Pool resource
type InstancePoolResource struct {

	// The OCID of resource that the AutoScalingConfiguration will manage.
	Id *string `mandatory:"true" json:"id"`
}

//GetId returns Id
func (m InstancePoolResource) GetId() *string {
	return m.Id
}

func (m InstancePoolResource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstancePoolResource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstancePoolResource InstancePoolResource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstancePoolResource
	}{
		"instancePool",
		(MarshalTypeInstancePoolResource)(m),
	}

	return json.Marshal(&s)
}
