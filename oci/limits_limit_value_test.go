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
	limitValueDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"service_name":        Representation{repType: Required, create: `${data.oci_limits_services.test_services.services.0.name}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"name":                Representation{repType: Optional, create: `custom-image-count`},
		"scope_type":          Representation{repType: Optional, create: `AD`},
	}

	LimitValueResourceConfig = AvailabilityDomainConfig +
		generateDataSourceFromRepresentationMap("oci_limits_services", "test_services", Required, Create, limitsServiceDataSourceRepresentation)
)

// issue-routing-tag: limits/default
func TestLimitsLimitValueResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitValueResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_limit_values.test_limit_values"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_limits_limit_values", "test_limit_values", Required, Create, limitValueDataSourceRepresentation) +
				compartmentIdVariableStr + LimitValueResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "service_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "limit_values.#"),
			),
		},
	})
}
