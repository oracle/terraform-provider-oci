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
	ApmTracesApmTracesqueryQuickPickDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
	}

	ApmTracesQueryQuickPickResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_traces/default
func TestApmTracesQueryQuickPickResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesQueryQuickPickResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_apm_traces_query_quick_picks.test_query_quick_picks"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_query_quick_picks", "test_query_quick_picks", acctest.Required, acctest.Create, ApmTracesApmTracesqueryQuickPickDataSourceRepresentation) +
				compartmentIdVariableStr + ApmTracesQueryQuickPickResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "quick_picks.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "quick_picks.0.quick_pick_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "quick_picks.0.quick_pick_query"),
			),
		},
	})
}
