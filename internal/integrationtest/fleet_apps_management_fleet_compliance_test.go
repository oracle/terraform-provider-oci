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
	FleetAppsManagementFleetComplianceSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_fleet_apps_management_fleet.test_fleet.id}`},
	}

	FleetAppsManagementFleetComplianceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, DataflowPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_fleet_apps_management_fleet", "test_fleet", acctest.Required, acctest.Create, FleetAppsManagementFleetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, FunctionsFunctionRepresentation) +
		// TODO: FIX APP
		// acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listings", "test_pbf_listings", acctest.Required, acctest.Create, FunctionsPbfListingDataSourceRepresentation) +
		KeyResourceDependencyConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, KmsVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, VaultSecretRepresentation)
)

// issue-routing-tag: fleet_apps_management/default
func TestFleetAppsManagementFleetComplianceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFleetAppsManagementFleetComplianceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_fleet_apps_management_fleet_compliance.test_fleet_compliance"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_fleet_apps_management_fleet_compliance", "test_fleet_compliance", acctest.Required, acctest.Create, FleetAppsManagementFleetComplianceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FleetAppsManagementFleetComplianceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compliance_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "confirmed_target_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "non_compliant_target_count"),
			),
		},
	})
}
