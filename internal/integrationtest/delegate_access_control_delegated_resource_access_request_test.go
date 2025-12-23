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
	DelegateAccessControlDelegatedResourceAccessRequestSingularDataSourceRepresentation = map[string]interface{}{
		"delegated_resource_access_request_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_delegate_access_control_delegated_resource_access_requests.test_delegated_resource_access_requests.delegated_resource_access_request_summary_collection.0.items.0.id}`},
	}

	DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	DelegateAccessControlDelegatedResourceAccessRequestResourceConfig = ""
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegatedResourceAccessRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegatedResourceAccessRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_delegated_resource_access_requests.test_delegated_resource_access_requests"
	singularDatasourceName := "data.oci_delegate_access_control_delegated_resource_access_request.test_delegated_resource_access_request"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_requests", "test_delegated_resource_access_requests", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegatedResourceAccessRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "delegated_resource_access_request_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_requests", "test_delegated_resource_access_requests", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_request", "test_delegated_resource_access_request", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlDelegatedResourceAccessRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delegated_resource_access_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "duration_in_hours"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "extend_duration_in_hours"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_approved"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reason_for_request"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_access_requested"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
