// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	resourceAvailabilitySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"limit_name":          Representation{repType: Required, create: `adb-free-count`},
		"service_name":        Representation{repType: Required, create: `database`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	ResourceAvailabilityResourceConfig = AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_limits_services", "test_services", Required, Create, limitsServiceDataSourceRepresentation)
)

// issue-routing-tag: limits/default
func TestLimitsResourceAvailabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsResourceAvailabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	singularDatasourceName := "data.oci_limits_resource_availability.test_resource_availability"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_limits_resource_availability", "test_resource_availability", Required, Create, resourceAvailabilitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ResourceAvailabilityResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "limit_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fractional_usage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "used"),
			),
		},
	})
}
