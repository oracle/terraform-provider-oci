// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityIdentityIamWorkRequestSingularDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_iam_work_requests.test_iam_work_requests.iam_work_requests.0.id}`},
	}

	IdentityIdentityIamWorkRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domain.test_domain.compartment_id}`},
		"resource_identifier": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domain.test_domain.id}`},
	}

	IdentityIamWorkRequestResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(IdentityDomainRepresentation, map[string]interface{}{
			"state": acctest.Representation{RepType: acctest.Required, Create: `inactive`},
		})) + acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", acctest.Required, acctest.Create, IdentityIdentityIamWorkRequestDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_requests.test_iam_work_requests"
	singularDatasourceName := "data.oci_identity_iam_work_request.test_iam_work_request"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + IdentityIamWorkRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", acctest.Required, acctest.Create, IdentityIdentityIamWorkRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityIamWorkRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
