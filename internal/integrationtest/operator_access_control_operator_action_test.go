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
	OperatorAccessControlOperatorAccessControlOperatorActionSingularDataSourceRepresentation = map[string]interface{}{
		"operator_action_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_operator_access_control_operator_actions.test_operator_actions.operator_action_collection.0.items.0.id}`},
	}

	OperatorAccessControlOperatorAccessControlOperatorActionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `EXADATAINFRASTRUCTURE`},
	}

	OperatorAccessControlOperatorActionResourceConfig = ""
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlOperatorActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOperatorAccessControlOperatorActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_operator_actions.test_operator_actions"
	singularDatasourceName := "data.oci_operator_access_control_operator_action.test_operator_action"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_actions", "test_operator_actions", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlOperatorActionDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlOperatorActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "operator_action_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "operator_action_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_actions", "test_operator_actions", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlOperatorActionDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_operator_access_control_operator_action", "test_operator_action", acctest.Required, acctest.Create, OperatorAccessControlOperatorAccessControlOperatorActionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperatorAccessControlOperatorActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "component"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "customer_display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "properties.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
			),
		},
	})
}
