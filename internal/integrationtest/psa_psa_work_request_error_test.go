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
	PsaPsaWorkRequestErrorDataSourceRepresentation = map[string]interface{}{
		"work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_psa_psa_work_request.test_psa_work_request.work_request_id}`},
	}
	PsaPsaWorkRequestErrorResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_request_errors", "test_psa_psa_work_request_errors", acctest.Required, acctest.Create, PsaPsaWorkRequestErrorDataSourceRepresentation)
)

// issue-routing-tag: psa/default
func TestPsaPsaWorkRequestErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsaPsaWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	datasourceName := "data.oci_psa_psa_work_request_errors.test_psa_psa_work_request_errors"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + PsaPsaWorkRequestResourcConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psa_psa_work_request", "test_psa_work_request", acctest.Required, acctest.Create, PsaPsaWorkRequestSingularDataSourceRepresentation) +
				PsaPsaWorkRequestErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "work_request_errors.#"),
			),
		},
	})
}
