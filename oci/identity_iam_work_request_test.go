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
	iamWorkRequestSingularDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": Representation{RepType: Required, Create: `${data.oci_identity_iam_work_requests.test_iam_work_requests.iam_work_requests.0.id}`},
	}

	iamWorkRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${oci_identity_domain.test_domain.compartment_id}`},
		"resource_identifier": Representation{RepType: Required, Create: `${oci_identity_domain.test_domain.id}`},
	}

	IamWorkRequestResourceConfig = GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", Required, Create,
		RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
			"state": Representation{RepType: Required, Create: `inactive`},
		})) + GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", Required, Create, iamWorkRequestDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_requests.test_iam_work_requests"
	singularDatasourceName := "data.oci_identity_iam_work_request.test_iam_work_request"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + IamWorkRequestResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_identifier"),

				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.operation_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.percent_complete"),
				resource.TestCheckResourceAttr(datasourceName, "iam_work_requests.0.resources.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.status"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.time_accepted"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.time_finished"),
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_requests.0.time_started"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", Required, Create, iamWorkRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IamWorkRequestResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "iam_work_request_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operation_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "percent_complete"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}
