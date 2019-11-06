// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScaleAnalyticsInstanceDetails Input payload to scale an Analytics instance up or down.
type ScaleAnalyticsInstanceDetails struct {
	Capacity *Capacity `mandatory:"true" json:"capacity"`
}

func (m ScaleAnalyticsInstanceDetails) String() string {
	return common.PointerString(m)
}
