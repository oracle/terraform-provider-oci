// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	iamWorkRequestLogDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_iam_work_request.test_iam_work_request.id}`},
	}

	IamWorkRequestLogResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
			"state": acctest.Representation{RepType: acctest.Required, Create: `inactive`},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", acctest.Required, acctest.Create, iamWorkRequestDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", acctest.Required, acctest.Create, iamWorkRequestSingularDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestLogResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_request_logs.test_iam_work_request_logs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request_logs", "test_iam_work_request_logs", acctest.Required, acctest.Create, iamWorkRequestLogDataSourceRepresentation) +
				compartmentIdVariableStr + IamWorkRequestLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.0.message"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.0.timestamp"),
			),
		},
	})
}
