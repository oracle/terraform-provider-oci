// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSubsettingAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                              acctest.Representation{RepType: acctest.Optional, Create: `targetId`},
		"subsetting_policy_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_subsetting_policy.test_subsetting_policy.id}`},
		"target_database_group_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_target_database_group.test_target_database_group.id}`},
		"target_id":                             acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
	}

	DataSafeSubsettingAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_analytics", "test_subsetting_analytics", acctest.Required, acctest.Create, DataSafeSubsettingAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	datasourceName := "data.oci_data_safe_subsetting_analytics.test_subsetting_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + DataSafeSubsettingAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_analytics_collection.#"),
			),
		},
	})
}
