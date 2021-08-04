// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	publicVantagePointSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  Representation{repType: Optional, create: `US East (Ashburn)`},
		"name":          Representation{repType: Optional, create: `OraclePublic-us-ashburn-1`},
	}

	publicVantagePointDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": Representation{repType: Required, create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  Representation{repType: Optional, create: `US East (Ashburn)`},
		"name":          Representation{repType: Optional, create: `OraclePublic-us-ashburn-1`},
	}

	PublicVantagePointResourceConfig = ""
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsPublicVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsPublicVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apm_synthetics_public_vantage_points.test_public_vantage_points"
	singularDatasourceName := "data.oci_apm_synthetics_public_vantage_point.test_public_vantage_point"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_public_vantage_points", "test_public_vantage_points", Optional, Create, publicVantagePointDataSourceRepresentation) +
					compartmentIdVariableStr + PublicVantagePointResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "name", "OraclePublic-us-ashburn-1"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "US East (Ashburn)"),

					resource.TestCheckResourceAttrSet(datasourceName, "public_vantage_point_collection.#"),
				),
			},
			// verify singular datasource
			{
				Config: config + generateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", Required, Create, apmDomainRepresentation) +
					generateDataSourceFromRepresentationMap("oci_apm_synthetics_public_vantage_point", "test_public_vantage_point", Optional, Create, publicVantagePointSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PublicVantagePointResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "OraclePublic-us-ashburn-1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "US East (Ashburn)"),
				),
			},
		},
	})
}
