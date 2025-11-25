// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsLogAnalyticsEntityAssociationsRemoveRepresentation = map[string]interface{}{
		"association_entities":    acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_log_analytics_entity.test_log_analytics_entity.id}`}},
		"log_analytics_entity_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_log_analytics_entity.test_log_analytics_entity_2.id}`},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceDependencies = LogAnalyticsLogAnalyticsEntityAssociationsListResourceConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_associations_list", "test_log_analytics_entity_associations_list", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityAssociationsListDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsEntityAssociationsRemoveResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsEntityAssociationsRemoveResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := utils.GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}

	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_log_analytics_log_analytics_entity_associations_remove.test_log_analytics_entity_associations_remove"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_associations_remove", "test_log_analytics_entity_associations_remove", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityAssociationsRemoveRepresentation), "loganalytics", "logAnalyticsEntityAssociationsRemove", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + LogAnalyticsLogAnalyticsEntityAssociationsRemoveResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_entity_associations_remove", "test_log_analytics_entity_associations_remove", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsEntityAssociationsRemoveRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association_entities.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "log_analytics_entity_id"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
