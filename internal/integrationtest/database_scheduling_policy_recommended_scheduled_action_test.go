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
	DatabaseSchedulingPolicyRecommendedScheduledActionDataSourceRepresentation = map[string]interface{}{
		"plan_intent":                          acctest.Representation{RepType: acctest.Required, Create: `EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE`},
		"scheduling_policy_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_database_scheduling_policy.test_scheduling_policy.id}`},
		"scheduling_policy_target_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
	}

	DatabaseSchedulingPolicyRecommendedScheduledActionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy", "test_scheduling_policy", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_scheduling_policy_scheduling_window", "test_scheduling_policy_scheduling_window", acctest.Required, acctest.Create, DatabaseSchedulingPolicySchedulingWindowRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseSchedulingPolicyRecommendedScheduledActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseSchedulingPolicyRecommendedScheduledActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_scheduling_policy_recommended_scheduled_actions.test_scheduling_policy_recommended_scheduled_actions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduling_policy_recommended_scheduled_actions", "test_scheduling_policy_recommended_scheduled_actions", acctest.Required, acctest.Create, DatabaseSchedulingPolicyRecommendedScheduledActionDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseSchedulingPolicyRecommendedScheduledActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "plan_intent", "EXADATA_INFRASTRUCTURE_FULL_SOFTWARE_UPDATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "scheduling_policy_target_resource_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "recommended_scheduled_actions_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "recommended_scheduled_actions_collection.0.items.#", "3"),
			),
		},
	})
}
