// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	PsaPsaWorkRequestLogsDataSourceRepresentation = map[string]interface{}{
		"work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_psa_psa_work_request.test_psa_work_request.work_request_id}`},
	}
	PsaPsaWorkRequestLogsResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_request_logs", "test_psa_psa_work_request_logs", acctest.Required, acctest.Create, PsaPsaWorkRequestLogsDataSourceRepresentation)
)

// issue-routing-tag: psa/default
func TestPsaPsaWorkRequestLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsaPsaWorkRequestLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	acctest.SaveConfigContent("", "", "", t)
	datasourceName := "data.oci_psa_psa_work_request_logs.test_psa_psa_work_request_logs"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + PsaPsaWorkRequestResourcConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_request", "test_psa_work_request", acctest.Required, acctest.Create, PsaPsaWorkRequestSingularDataSourceRepresentation) +
				PsaPsaWorkRequestLogsResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.message"),
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_log_entries.0.timestamp")),
		},
	})
}
