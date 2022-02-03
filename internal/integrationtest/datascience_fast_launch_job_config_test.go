// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	fastLaunchJobConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	FastLaunchJobConfigResourceConfig = ""
)

// issue-routing-tag: datascience/default
func TestDatascienceFastLaunchJobConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceFastLaunchJobConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_fast_launch_job_configs.test_fast_launch_job_configs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_fast_launch_job_configs", "test_fast_launch_job_configs", acctest.Required, acctest.Create, fastLaunchJobConfigDataSourceRepresentation) +
				compartmentIdVariableStr + FastLaunchJobConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.core_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.managed_egress_support"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "fast_launch_job_configs.0.shape_series"),
			),
		},
	})
}
