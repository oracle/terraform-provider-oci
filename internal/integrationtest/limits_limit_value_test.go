// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LimitsLimitsLimitValueDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"service_name":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_limits_services.test_services.services.0.name}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `custom-image-count`},
		"scope_type":          acctest.Representation{RepType: acctest.Optional, Create: `AD`},
	}

	LimitsLimitValueResourceConfig = AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_limits_services", "test_services", acctest.Required, acctest.Create, LimitsLimitsServiceDataSourceRepresentation)
)

// issue-routing-tag: limits/default
func TestLimitsLimitValueResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLimitsLimitValueResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	datasourceName := "data.oci_limits_limit_values.test_limit_values"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_limits_limit_values", "test_limit_values", acctest.Required, acctest.Create, LimitsLimitsLimitValueDataSourceRepresentation) +
				compartmentIdVariableStr + LimitsLimitValueResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttrSet(datasourceName, "service_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "limit_values.#"),
			),
		},
	})
}
