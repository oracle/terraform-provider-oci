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
	DataSafeCalculateAuditVolumeCollectedRequiredOnlyResource = DataSafeCalculateAuditVolumeCollectedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_collected", "test_calculate_audit_volume_collected", acctest.Required, acctest.Create, DataSafeCalculateAuditVolumeCollectedRepresentation)

	DataSafeCalculateAuditVolumeCollectedRepresentation = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile_management.test_audit_profile_management.id}`},
		"time_from_month":  acctest.Representation{RepType: acctest.Required, Create: `2024-05-17T15:05:28Z`},
		"time_to_month":    acctest.Representation{RepType: acctest.Optional, Create: `2024-06-17T15:05:28Z`},
	}

	DataSafeCalculateAuditVolumeCollectedResourceDependencies = DataSafeAuditProfileManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile_management", "test_audit_profile_management", acctest.Optional, acctest.Create, DataSafeAuditProfileManagementRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeCalculateAuditVolumeCollectedResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCalculateAuditVolumeCollectedResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_calculate_audit_volume_collected.test_calculate_audit_volume_collected"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeCalculateAuditVolumeCollectedResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_collected", "test_calculate_audit_volume_collected", acctest.Optional, acctest.Create, DataSafeCalculateAuditVolumeCollectedRepresentation), "datasafe", "calculateAuditVolumeCollected", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeCollectedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_collected", "test_calculate_audit_volume_collected", acctest.Required, acctest.Create, DataSafeCalculateAuditVolumeCollectedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "time_from_month", "2024-05-17T15:05:28Z"),
				resource.TestCheckResourceAttrSet(resourceName, "collected_audit_volumes.#"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeCollectedResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeCalculateAuditVolumeCollectedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_calculate_audit_volume_collected", "test_calculate_audit_volume_collected", acctest.Optional, acctest.Create, DataSafeCalculateAuditVolumeCollectedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "time_from_month", "2024-05-17T15:05:28Z"),
				resource.TestCheckResourceAttr(resourceName, "time_to_month", "2024-06-17T15:05:28Z"),
				resource.TestCheckResourceAttrSet(resourceName, "collected_audit_volumes.#"),

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
