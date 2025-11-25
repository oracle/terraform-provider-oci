// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
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
	LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceRepresentation = map[string]interface{}{
		"log_analytics_entity_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_log_analytics_entity.test_log_analytics_entity_2.id}`},
		"namespace":                  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"direct_or_all_associations": acctest.Representation{RepType: acctest.Optional, Create: `DIRECT`},
	}

	LogAnalyticsLogAnalyticsEntityAssociationsListResourceConfig = LogAnalyticsLogAnalyticsEntityAssociationsAddResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_associations_add", "test_log_analytics_entity_associations_add", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityAssociationsAddRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityAssociationsListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityAssociationsListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}

	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	datasourceName := "data.oci_log_analytics_log_analytics_entity_associations_list.test_log_analytics_entity_associations_list"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsLogAnalyticsEntityAssociationsListResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_associations_list", "test_log_analytics_entity_associations_list", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "direct_or_all_associations", "DIRECT"),
				resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_entity_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

				// resource.TestCheckResourceAttrSet(datasourceName, "log_analytics_entity_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "log_analytics_entity_collection.0.items.#", "1"),
			),
		},
	})
}
