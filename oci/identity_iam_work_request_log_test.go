// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	iamWorkRequestLogDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": Representation{RepType: Required, Create: `${data.oci_identity_iam_work_request.test_iam_work_request.id}`},
	}

	IamWorkRequestLogResourceConfig = GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", Required, Create,
		RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
			"state": Representation{RepType: Required, Create: `inactive`},
		})) +
		GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", Required, Create, iamWorkRequestDataSourceRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", Required, Create, iamWorkRequestSingularDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestLogResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestLogResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_request_logs.test_iam_work_request_logs"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request_logs", "test_iam_work_request_logs", Required, Create, iamWorkRequestLogDataSourceRepresentation) +
				compartmentIdVariableStr + IamWorkRequestLogResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.0.message"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_logs.0.timestamp"),
			),
		},
	})
}
