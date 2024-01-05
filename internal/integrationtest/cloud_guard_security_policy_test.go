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
	CloudGuardCloudGuardSecurityPolicySingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: securityPolicyId1},
	}

	CloudGuardCloudGuardSecurityPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		//"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		//"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_security_policy.test_security_policy.id}`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	CloudGuardSecurityPolicyResourceConfig = ""
	securityPolicyId1                      = `${data.oci_cloud_guard_security_policies.test_security_policies.security_policy_collection.0.items.0.id}`
	SecurityPolicyResourceDependency       = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_policies", "test_security_policies", acctest.Required, acctest.Create, CloudGuardCloudGuardSecurityPolicyDataSourceRepresentation)
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardSecurityPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardSecurityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_cloud_guard_security_policies.test_security_policies"
	singularDatasourceName := "data.oci_cloud_guard_security_policy.test_security_policy"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_policies", "test_security_policies", acctest.Optional, acctest.Create, CloudGuardCloudGuardSecurityPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				//resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + SecurityPolicyResourceDependency +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_security_policy", "test_security_policy", acctest.Required, acctest.Create, CloudGuardCloudGuardSecurityPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardSecurityPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "category"),
				//For oracle managed Policies we return null.
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "friendly_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
				//Number of services can vary based on the Policy
				//resource.TestCheckResourceAttr(singularDatasourceName, "services.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
