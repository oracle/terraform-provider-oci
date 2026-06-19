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
	PsqlDbSystemPitrDetailSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_ocid}`},
	}

	PsqlDbSystemPitrDetailResourceConfig = ""
)

// issue-routing-tag: psql/default
func TestPsqlDbSystemPitrDetailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDbSystemPitrDetailResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("db_ocid")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"db_ocid\" { default = \"%s\" }\n", dbSystemId)

	singularDatasourceName := "data.oci_psql_db_system_pitr_detail.test_db_system_pitr_detail"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + dbSystemIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_system_pitr_detail", "test_db_system_pitr_detail", acctest.Required, acctest.Create, PsqlDbSystemPitrDetailSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlDbSystemPitrDetailResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pitr_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recovery_time_windows.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_time_windows.0.time_recovery_window_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recovery_time_windows.0.time_recovery_window_end"),
			),
		},
	})
}
