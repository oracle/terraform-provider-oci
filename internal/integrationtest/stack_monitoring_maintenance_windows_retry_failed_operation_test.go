// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	StackMonitoringMaintenanceWindowsRetryFailedOperationRepresentation = map[string]interface{}{
		"maintenance_window_id": acctest.Representation{RepType: acctest.Required, Create: `${var.mw_id_to_retry}`},
	}

	StackMonitoringMaintenanceWindowsRetryFailedOperationResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMaintenanceWindowsRetryFailedOperationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMaintenanceWindowsRetryFailedOperationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	mwIdToRetry := utils.GetEnvSettingWithBlankDefault("mw_id_retry")
	mwIdToRetryVariableStr := fmt.Sprintf("variable \"mw_id_to_retry\" { default = \"%s\" }\n", mwIdToRetry)

	resourceName := "oci_stack_monitoring_maintenance_windows_retry_failed_operation.test_maintenance_windows_retry_failed_operation"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMaintenanceWindowsRetryFailedOperationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_windows_retry_failed_operation", "test_maintenance_windows_retry_failed_operation", acctest.Required, acctest.Create, StackMonitoringMaintenanceWindowsRetryFailedOperationRepresentation), "stackmonitoring", "maintenanceWindowsRetryFailedOperation", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + mwIdToRetryVariableStr + StackMonitoringMaintenanceWindowsRetryFailedOperationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_maintenance_windows_retry_failed_operation", "test_maintenance_windows_retry_failed_operation", acctest.Required, acctest.Create, StackMonitoringMaintenanceWindowsRetryFailedOperationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "maintenance_window_id"),

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
