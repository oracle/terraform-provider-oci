// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityIdentityIamWorkRequestErrorDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_iam_work_request.test_iam_work_request.id}`},
	}

	IdentityIamWorkRequestErrorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(IdentityDomainRepresentation, map[string]interface{}{
			"state": acctest.Representation{RepType: acctest.Required, Create: `inactive`},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", acctest.Required, acctest.Create, IdentityIdentityIamWorkRequestDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", acctest.Required, acctest.Create, IdentityIdentityIamWorkRequestSingularDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestErrorResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_request_errors.test_iam_work_request_errors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request_errors", "test_iam_work_request_errors", acctest.Required, acctest.Create, IdentityIdentityIamWorkRequestErrorDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityIamWorkRequestErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_errors.#"),
				func(s *terraform.State) (err error) {
					// Mostly errors will be nil. Check error data only if it exists.
					errorMessageCountStr, err := acctest.FromInstanceState(s, datasourceName, "iam_work_request_errors.#")
					if err != nil {
						return err
					}
					errorMessageCount, err := strconv.Atoi(errorMessageCountStr)
					if err != nil {
						return err
					}
					if errorMessageCount > 0 {
						resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_errors.0.code")
						resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_errors.0.message")
						resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_errors.0.timestamp")
					}
					return nil
				},
			),
		},
	})
}
