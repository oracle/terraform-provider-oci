// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// NotificationsTargetDetails The notifications target.
type NotificationsTargetDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId *string `mandatory:"true" json:"topicId"`
}

func (m NotificationsTargetDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m NotificationsTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNotificationsTargetDetails NotificationsTargetDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeNotificationsTargetDetails
	}{
		"notifications",
		(MarshalTypeNotificationsTargetDetails)(m),
	}

	return json.Marshal(&s)
}
