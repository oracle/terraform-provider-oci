// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"reflect"
	"testing"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"
)

func TestUnitAssignedSubscriptionSummaryToMapIncludesCommonFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		summary  oci_tenantmanagercontrolplane.AssignedSubscriptionSummary
		expected map[string]interface{}
	}{
		{
			name: "classic assigned subscription",
			summary: oci_tenantmanagercontrolplane.ClassicAssignedSubscriptionSummary{
				Id:                    oci_common.String("classic-id"),
				CompartmentId:         oci_common.String("classic-compartment"),
				ServiceName:           oci_common.String("UGBUCCS"),
				ClassicSubscriptionId: oci_common.String("classic-subscription-id"),
				FreeformTags:          map[string]string{"environment": "test"},
				LifecycleState:        oci_tenantmanagercontrolplane.ClassicSubscriptionLifecycleStateActive,
			},
			expected: map[string]interface{}{
				"entity_version":          "V1",
				"classic_subscription_id": "classic-subscription-id",
				"state":                   "ACTIVE",
				"managed_by":              "",
				"id":                      "classic-id",
				"compartment_id":          "classic-compartment",
				"service_name":            "UGBUCCS",
				"freeform_tags":           map[string]string{"environment": "test"},
			},
		},
		{
			name: "cloud assigned subscription",
			summary: oci_tenantmanagercontrolplane.CloudAssignedSubscriptionSummary{
				Id:                 oci_common.String("cloud-id"),
				CompartmentId:      oci_common.String("cloud-compartment"),
				ServiceName:        oci_common.String("CLOUDCM"),
				SubscriptionNumber: oci_common.String("cloud-subscription-number"),
				CurrencyCode:       oci_common.String("USD"),
				FreeformTags:       map[string]string{},
				LifecycleState:     oci_tenantmanagercontrolplane.SubscriptionLifecycleStateActive,
			},
			expected: map[string]interface{}{
				"entity_version":      "V2",
				"currency_code":       "USD",
				"state":               "ACTIVE",
				"subscription_number": "cloud-subscription-number",
				"id":                  "cloud-id",
				"compartment_id":      "cloud-compartment",
				"service_name":        "CLOUDCM",
				"freeform_tags":       map[string]string{},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := AssignedSubscriptionSummaryToMap(test.summary)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("AssignedSubscriptionSummaryToMap() = %#v, expected %#v", actual, test.expected)
			}
		})
	}
}
