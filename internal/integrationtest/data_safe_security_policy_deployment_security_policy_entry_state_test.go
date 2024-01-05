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
	DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateSingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_deployment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_deployment_id}`},
		"security_policy_entry_state_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_entry_state_id}`},
	}

	DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceRepresentation = map[string]interface{}{
		"security_policy_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_deployment_id}`},
		"deployment_status":             acctest.Representation{RepType: acctest.Optional, Create: `CREATED`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the security policy deployment ocid and entry state id are hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyDeploymentId := utils.GetEnvSettingWithBlankDefault("security_policy_deployment_ocid")
	securityPolicyDeploymentIdVariableStr := fmt.Sprintf("variable \"security_policy_deployment_id\" { default = \"%s\" }\n", securityPolicyDeploymentId)

	securityPolicyEntryStateId := utils.GetEnvSettingWithBlankDefault("security_policy_entry_state_id")
	securityPolicyEntryStateIdVariableStr := fmt.Sprintf("variable \"security_policy_entry_state_id\" { default = \"%s\" }\n", securityPolicyEntryStateId)

	datasourceName := "data.oci_data_safe_security_policy_deployment_security_policy_entry_states.test_security_policy_deployment_security_policy_entry_states"
	singularDatasourceName := "data.oci_data_safe_security_policy_deployment_security_policy_entry_state.test_security_policy_deployment_security_policy_entry_state"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployment_security_policy_entry_states", "test_security_policy_deployment_security_policy_entry_states", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "deployment_status", "CREATED"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_deployment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_entry_state_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployment_security_policy_entry_state", "test_security_policy_deployment_security_policy_entry_state", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr + securityPolicyEntryStateIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_deployment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_entry_state_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "entry_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
			),
		},
	})
}
