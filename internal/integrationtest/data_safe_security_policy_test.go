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
	DataSafeSecurityPolicyRequiredOnlyResource = DataSafeSecurityPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Required, acctest.Create, DataSafeSecurityPolicyRepresentation)

	DataSafeSecurityPolicyResourceConfig = DataSafeSecurityPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyRepresentation)

	DataSafeSecurityPolicySingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
	}

	DataSafeSecurityPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"security_policy_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.security_policy_id}`},
	}

	DataSafeSecurityPolicyRepresentation = map[string]interface{}{
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `Updated security policy description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `SecurityPolicy_updated`, Update: `displayName2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyTagsChangesRep},
	}
	ignoreSecurityPolicyTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}

	DataSafeSecurityPolicyChangeCompartmentRepresentation = map[string]interface{}{
		"security_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `SecurityPolicy_updated`, Update: `displayName2`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSecurityPolicyTagsChangesRep},
	}

	DataSafeSecurityPolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the security policy ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeSecurityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	securityPolicyId := utils.GetEnvSettingWithBlankDefault("security_policy_ocid")
	securityPolicyIdVariableStr := fmt.Sprintf("variable \"security_policy_id\" { default = \"%s\" }\n", securityPolicyId)

	resourceName := "oci_data_safe_security_policy.test_security_policy"
	datasourceName := "data.oci_data_safe_security_policies.test_security_policies"
	singularDatasourceName := "data.oci_data_safe_security_policy.test_security_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+securityPolicyIdVariableStr+DataSafeSecurityPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyRepresentation, map[string]interface{}{
				"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
			})), "datasafe", "securityPolicy", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify change compartment
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + compartmentIdUVariableStr + DataSafeSecurityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyChangeCompartmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(DataSafeSecurityPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "security_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policies", "test_security_policies", acctest.Optional, acctest.Update, DataSafeSecurityPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Optional, acctest.Update, DataSafeSecurityPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "security_policy_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy", "test_security_policy", acctest.Required, acctest.Create, DataSafeSecurityPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyIdVariableStr + DataSafeSecurityPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_id"),

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
			Config:                  config + DataSafeSecurityPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`security_policy_id`},
			ResourceName:            resourceName,
		},
	})
}
