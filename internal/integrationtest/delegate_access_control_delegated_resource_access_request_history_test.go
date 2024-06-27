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
	DelegateAccessControlDelegatedResourceAccessRequestHistoryDataSourceRepresentation = map[string]interface{}{
		"delegated_resource_access_request_id": acctest.Representation{RepType: acctest.Required, Create: `${var.histAccReqId}`},
	}

	DelegateAccessControlDelegatedResourceAccessRequestHistoryResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_requests", "test_delegated_resource_access_requests", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation)
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegatedResourceAccessRequestHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegatedResourceAccessRequestHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_delegated_resource_access_request_histories.test_delegated_resource_access_request_histories"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_request_histories", "test_delegated_resource_access_request_histories", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegatedResourceAccessRequestHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "delegated_resource_access_request_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "delegated_resource_access_request_history_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "delegated_resource_access_request_history_collection.0.items.#", "4"),
			),
		},
	})
}
