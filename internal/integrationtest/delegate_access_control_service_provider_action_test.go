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
	DelegateAccessControlServiceProviderActionSingularDataSourceRepresentation = map[string]interface{}{
		//"service_provider_action_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_delegate_access_control_service_provider_action.test_service_provider_action.id}`},
		"service_provider_action_id": acctest.Representation{RepType: acctest.Required, Create: `${var.spActionId}`},
	}

	DelegateAccessControlServiceProviderActionDataSourceRepresentation = map[string]interface{}{
		// can be only in root compartment, which is the tenancy
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"resource_type":                 acctest.Representation{RepType: acctest.Optional, Create: `VMCLUSTER`},
		"service_provider_service_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`serviceProviderServiceType`}},
		"state":                         acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DelegateAccessControlServiceProviderActionResourceConfig = ""
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlServiceProviderActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlServiceProviderActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_service_provider_actions.test_service_provider_actions"
	singularDatasourceName := "data.oci_delegate_access_control_service_provider_action.test_service_provider_action"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_provider_actions", "test_service_provider_actions", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderActionDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlServiceProviderActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "VMCLUSTER"),
				resource.TestCheckResourceAttr(datasourceName, "service_provider_service_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),*/

				resource.TestCheckResourceAttrSet(datasourceName, "service_provider_action_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_service_provider_action", "test_service_provider_action", acctest.Required, acctest.Create, DelegateAccessControlServiceProviderActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DelegateAccessControlServiceProviderActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_provider_action_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "component"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "properties.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_provider_service_types.#", "2"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
