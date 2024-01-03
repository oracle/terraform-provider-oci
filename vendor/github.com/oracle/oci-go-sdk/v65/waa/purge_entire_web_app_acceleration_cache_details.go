// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration (WAA) API
//
// API for the Web Application Acceleration service.
// Use this API to manage regional Web App Acceleration policies such as Caching and Compression
// for accelerating HTTP services.
//

package waa

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PurgeEntireWebAppAccelerationCacheDetails Clears all resources from the cache of the WebAppAcceleration.
type PurgeEntireWebAppAccelerationCacheDetails struct {
}

func (m PurgeEntireWebAppAccelerationCacheDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PurgeEntireWebAppAccelerationCacheDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PurgeEntireWebAppAccelerationCacheDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePurgeEntireWebAppAccelerationCacheDetails PurgeEntireWebAppAccelerationCacheDetails
	s := struct {
		DiscriminatorParam string `json:"purgeType"`
		MarshalTypePurgeEntireWebAppAccelerationCacheDetails
	}{
		"ENTIRE_CACHE",
		(MarshalTypePurgeEntireWebAppAccelerationCacheDetails)(m),
	}

	return json.Marshal(&s)
}
