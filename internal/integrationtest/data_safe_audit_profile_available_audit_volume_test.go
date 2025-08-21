// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeauditProfileAvailableAuditVolumeSingularDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.auditProfileId}`},
		"work_request_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.workId}`},
		"month_in_consideration_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationGreaterThan`},
		"month_in_consideration_less_than":    acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationLessThan`},
		"trail_location":                      acctest.Representation{RepType: acctest.Optional, Create: `trailLocation`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileAvailableAuditVolumeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditProfileAvailableAuditVolumeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	workId := utils.GetEnvSettingWithBlankDefault("work_request_id")
	workIdVariableStr := fmt.Sprintf("variable \"workId\" { default = \"%s\" }\n", workId)

	auditId := utils.GetEnvSettingWithBlankDefault("auditProfileId")
	auditIdVariableStr := fmt.Sprintf("variable \"auditProfileId\" { default = \"%s\" }\n", auditId)

	datasourceName := "data.oci_data_safe_audit_profile_available_audit_volumes.test_audit_profile_available_audit_volumes"
	singularDatasourceName := "data.oci_data_safe_audit_profile_available_audit_volume.test_audit_profile_available_audit_volume"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + workIdVariableStr + auditIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile_available_audit_volumes", "test_audit_profile_available_audit_volumes", acctest.Required, acctest.Create, DataSafeauditProfileAvailableAuditVolumeSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_audit_volume_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + workIdVariableStr + auditIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile_available_audit_volume", "test_audit_profile_available_audit_volume", acctest.Required, acctest.Create, DataSafeauditProfileAvailableAuditVolumeSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
