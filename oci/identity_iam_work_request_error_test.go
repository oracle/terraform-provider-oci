// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	iamWorkRequestErrorDataSourceRepresentation = map[string]interface{}{
		"iam_work_request_id": Representation{RepType: Required, Create: `${data.oci_identity_iam_work_request.test_iam_work_request.id}`},
	}

	IamWorkRequestErrorResourceConfig = GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", Required, Create,
		RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
			"state": Representation{RepType: Required, Create: `inactive`},
		})) +
		GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_requests", "test_iam_work_requests", Required, Create, iamWorkRequestDataSourceRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request", "test_iam_work_request", Required, Create, iamWorkRequestSingularDataSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityIamWorkRequestErrorResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityIamWorkRequestErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_iam_work_request_errors.test_iam_work_request_errors"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_iam_work_request_errors", "test_iam_work_request_errors", Required, Create, iamWorkRequestErrorDataSourceRepresentation) +
				compartmentIdVariableStr + IamWorkRequestErrorResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "iam_work_request_errors.#"),
				func(s *terraform.State) (err error) {
					// Mostly errors will be nil. Check error data only if it exists.
					errorMessageCountStr, err := FromInstanceState(s, datasourceName, "iam_work_request_errors.#")
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
