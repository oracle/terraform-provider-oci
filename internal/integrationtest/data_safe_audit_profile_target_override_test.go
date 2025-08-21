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
	DataSafeAuditProfileTargetOverrideDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile.test_audit_profile.id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	DataSafeAuditProfileTargetOverrideResourceConfig = DataSafeAuditProfileResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Required, acctest.Create, DataSafeAuditProfileRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileTargetOverrideResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditProfileTargetOverrideResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_audit_profile_target_overrides.test_audit_profile_target_overrides"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile_target_overrides", "test_audit_profile_target_overrides", acctest.Required, acctest.Create, DataSafeAuditProfileTargetOverrideDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditProfileTargetOverrideResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_profile_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_conforming_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_overriding_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_overriding_offline_months_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_overriding_online_months_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_override_collection.0.targets_overriding_paid_usage_count"),
			),
		},
	})
}
