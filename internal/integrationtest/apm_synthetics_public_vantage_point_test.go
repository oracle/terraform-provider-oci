// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ApmSyntheticsApmSyntheticspublicVantagePointSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `US East (Ashburn)`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `OraclePublic-us-ashburn-1`},
	}

	ApmSyntheticsApmSyntheticspublicVantagePointDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `US East (Ashburn)`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `OraclePublic-us-ashburn-1`},
	}

	ApmSyntheticsPublicVantagePointResourceConfig = ""
)

// issue-routing-tag: apm_synthetics/default
func TestApmSyntheticsPublicVantagePointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmSyntheticsPublicVantagePointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apm_synthetics_public_vantage_points.test_public_vantage_points"
	singularDatasourceName := "data.oci_apm_synthetics_public_vantage_point.test_public_vantage_point"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_public_vantage_points", "test_public_vantage_points", acctest.Optional, acctest.Create, ApmSyntheticsApmSyntheticspublicVantagePointDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsPublicVantagePointResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "OraclePublic-us-ashburn-1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "US East (Ashburn)"),

				resource.TestCheckResourceAttrSet(datasourceName, "public_vantage_point_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_synthetics_public_vantage_point", "test_public_vantage_point", acctest.Optional, acctest.Create, ApmSyntheticsApmSyntheticspublicVantagePointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmSyntheticsPublicVantagePointResourceConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "OraclePublic-us-ashburn-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "US East (Ashburn)"),
			),
		},
	})
}
