// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Metric A metric. This a quantitative measurement of an entity.
type Metric struct {

	// The name of the metric
	Name *string `mandatory:"true" json:"name"`

	// Must be NULL at the moment, and "name" must be a known metric.
	ValueSource *string `mandatory:"false" json:"valueSource"`

	// The unit of the metric
	Unit *string `mandatory:"false" json:"unit"`

	// A description of the metric
	Description *string `mandatory:"false" json:"description"`
}

func (m Metric) String() string {
	return common.PointerString(m)
}
