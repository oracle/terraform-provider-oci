// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityPolicyDeploymentManagementResourceConfig = DataSafeSecurityPolicyDeploymentManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create, DataSafeSecurityPolicyDeploymentManagementRepresentation)

	DataSafeSecurityPolicyDeploymentManagementRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"target_id":      acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyDeploymentManagementTagsChangesRep},
	}
	ignoreSecurityPolicyDeploymentManagementTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`, `system_tags`}},
	}

	DataSafeSecurityPolicyDeploymentManagementResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyDeploymentManagementResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the target ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyDeploymentManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_security_policy_deployment_management.test_security_policy_deployment_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+targetIdVariableStr+DataSafeSecurityPolicyDeploymentManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment_management", "test_security_policy_deployment_management", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyDeploymentManagementRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "securityPolicyDeploymentManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSecurityPolicyDeploymentManagementResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
