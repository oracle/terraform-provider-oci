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
		"audit_profile_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile.test_audit_profile.id}`},
		"work_request_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_work_request.test_work_request.id}`},
		"month_in_consideration_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationGreaterThan`},
		"month_in_consideration_less_than":    acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationLessThan`},
		"trail_location":                      acctest.Representation{RepType: acctest.Optional, Create: `trailLocation`},
	}

	DataSafeauditProfileAvailableAuditVolumeDataSourceRepresentation = map[string]interface{}{
		"audit_profile_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_audit_profile.test_audit_profile.id}`},
		"work_request_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_work_request.test_work_request.id}`},
		"month_in_consideration_greater_than": acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationGreaterThan`},
		"month_in_consideration_less_than":    acctest.Representation{RepType: acctest.Optional, Create: `monthInConsiderationLessThan`},
		"trail_location":                      acctest.Representation{RepType: acctest.Optional, Create: `trailLocation`},
	}

	DataSafeAuditProfileAvailableAuditVolumeResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_work_requests", "test_work_requests", acctest.Required, acctest.Create, ContainerengineContainerengineWorkRequestDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_audit_profile", "test_audit_profile", acctest.Required, acctest.Create, auditProfileRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditProfileAvailableAuditVolumeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditProfileAvailableAuditVolumeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_audit_profile_available_audit_volumes.test_audit_profile_available_audit_volumes"
	singularDatasourceName := "data.oci_data_safe_audit_profile_available_audit_volume.test_audit_profile_available_audit_volume"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile_available_audit_volumes", "test_audit_profile_available_audit_volumes", acctest.Required, acctest.Create, DataSafeauditProfileAvailableAuditVolumeDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditProfileAvailableAuditVolumeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "month_in_consideration_greater_than"),
				resource.TestCheckResourceAttrSet(datasourceName, "month_in_consideration_less_than"),
				resource.TestCheckResourceAttr(datasourceName, "trail_location", "trailLocation"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_audit_volume_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_profile_available_audit_volume", "test_audit_profile_available_audit_volume", acctest.Required, acctest.Create, DataSafeauditProfileAvailableAuditVolumeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditProfileAvailableAuditVolumeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "audit_profile_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "month_in_consideration_greater_than", "monthInConsiderationGreaterThan"),
				resource.TestCheckResourceAttr(singularDatasourceName, "month_in_consideration_less_than", "monthInConsiderationLessThan"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trail_location", "trailLocation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
			),
		},
	})
}
