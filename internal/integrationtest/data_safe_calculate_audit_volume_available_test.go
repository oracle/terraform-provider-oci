// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeCalculateAuditVolumeAvailableRepresentation1 = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile_management.test_audit_profile_management.id}`},
		"trail_locations":  acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	DataSafeCalculateAuditVolumeAvailableRepresentation2 = map[string]interface{}{
		"audit_profile_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile_management.test_audit_profile_management.id}`},
		"audit_collection_start_time": acctest.Representation{RepType: acctest.Optional, Create: `2023-05-17T15:05:28Z`},
		"database_unique_name":        acctest.Representation{RepType: acctest.Optional, Create: `databaseUniqueName`},
		"trail_locations":             acctest.Representation{RepType: acctest.Optional, Create: []string{`UNIFIED_AUDIT_TRAIL`}},
	}

	DataSafeAuditProfileManagementRepresentation = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `$(var.compartment_id)`},
		"target_id":                            acctest.Representation{RepType: acctest.Optional, Create: `$(var.target_id)`},
		"description":                          acctest.Representation{RepType: acctest.Optional, Create: `updated-description`, Update: `description2`},
		"change_retention_trigger":             acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"offline_months":                       acctest.Representation{RepType: acctest.Optional, Create: `50`, Update: `50`},
		"online_months":                        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `10`},
		"is_paid_usage_enabled":                acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"is_override_global_retention_setting": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}

	DataSafeCalculateAuditVolumeAvailableResourceDependencies = DataSafeAuditProfileManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Create, DataSafeAuditProfileManagementRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeCalculateAuditVolumeAvailableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCalculateAuditVolumeAvailableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartmentId")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_calculate_audit_volume_available.test_calculate_audit_volume_available"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeCalculateAuditVolumeAvailableResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_available", "test_calculate_audit_volume_available", acctest.Optional, acctest.Create, DataSafeCalculateAuditVolumeAvailableRepresentation2), "datasafe", "calculateAuditVolumeAvailable", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeAvailableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_available", "test_calculate_audit_volume_available", acctest.Optional, acctest.Create, DataSafeCalculateAuditVolumeAvailableRepresentation1),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(resourceName, "available_audit_volumes.#"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeAvailableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeAvailableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_available", "test_calculate_audit_volume_available", acctest.Optional, acctest.Create, DataSafeCalculateAuditVolumeAvailableRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "audit_collection_start_time", "2023-05-17T15:05:28Z"),
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "database_unique_name", "databaseUniqueName"),
				resource.TestCheckResourceAttr(resourceName, "trail_locations.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "available_audit_volumes.#"),

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
