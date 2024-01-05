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
	DataSafeSecurityPolicyDeploymentRequiredOnlyResource = DataSafeSecurityPolicyDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentRepresentation)

	DataSafeSecurityPolicyDeploymentResourceConfig = DataSafeSecurityPolicyDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentRepresentation)

	DataSafeSecurityPolicyDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_deployment_id}`},
	}

	DataSafeSecurityPolicyDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
	}

	DataSafeSecurityPolicyDeploymentRepresentation = map[string]interface{}{
		"security_policy_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_deployment_id}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Updated security policy deployment description`, Update: `description2`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `SecurityPolicyDeployment_updated`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyDeploymentSystemTagsChangesRep},
	}

	IgnoreSecurityPolicyDeploymentSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyDeploymentChangeCompartmentRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"security_policy_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_deployment_id}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `Updated security policy deployment description`, Update: `description2`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `SecurityPolicyDeployment_updated`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: IgnoreSecurityPolicyDeploymentSystemTagsChangesRep},
	}

	DataSafeSecurityPolicyDeploymentResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyDeploymentResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the security policy deployment ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyDeploymentId := utils.GetEnvSettingWithBlankDefault("security_policy_deployment_id")
	securityPolicyDeploymentIdVariableStr := fmt.Sprintf("variable \"security_policy_deployment_id\" { default = \"%s\" }\n", securityPolicyDeploymentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_security_policy_deployment.test_security_policy_deployment"
	datasourceName := "data.oci_data_safe_security_policy_deployments.test_security_policy_deployments"
	singularDatasourceName := "data.oci_data_safe_security_policy_deployment.test_security_policy_deployment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + securityPolicyDeploymentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Updated security policy deployment description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "SecurityPolicyDeployment_updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_deployment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		//verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentChangeCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_deployment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployments", "test_security_policy_deployments", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_deployment", "test_security_policy_deployment", acctest.Required, acctest.Create, DataSafeSecurityPolicyDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr + DataSafeSecurityPolicyDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_deployment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + securityPolicyDeploymentIdVariableStr + DataSafeSecurityPolicyDeploymentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`security_policy_deployment_id`},
			ResourceName:            resourceName,
		},
	})
}
