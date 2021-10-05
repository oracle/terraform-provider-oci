// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v49/common"
	oci_identity "github.com/oracle/oci-go-sdk/v49/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DynamicGroupRequiredOnlyResource = DynamicGroupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Required, Create, dynamicGroupRepresentation)

	dynamicGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"name":           Representation{RepType: Optional, Create: `DevCompartmentDynamicGroup`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, dynamicGroupDataSourceFilterRepresentation}}
	dynamicGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_dynamic_group.test_dynamic_group.id}`}},
	}

	dynamicGroupRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"description":    Representation{RepType: Required, Create: `Instance Group for dev compartment`, Update: `description2`},
		"matching_rule":  Representation{RepType: Required, Create: `${var.dynamic_group_matching_rule}`, Update: `${var.dynamic_group_matching_rule}`},
		"name":           Representation{RepType: Required, Create: `DevCompartmentDynamicGroup`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DynamicGroupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityDynamicGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDynamicGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	matchingRuleValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRuleVariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRuleValueStr)

	matchingRule2ValueStr := fmt.Sprintf("instance.compartment_id='%s'", compartmentId)
	matchingRule2VariableStr := fmt.Sprintf("variable \"dynamic_group_matching_rule\" {default = \"%s\" }\n", matchingRule2ValueStr)
	resourceName := "oci_identity_dynamic_group.test_dynamic_group"
	datasourceName := "data.oci_identity_dynamic_groups.test_dynamic_groups"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+matchingRuleVariableStr+DynamicGroupResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Create, dynamicGroupRepresentation), "identity", "dynamicGroup", t)

	ResourceTest(t, testAccCheckIdentityDynamicGroupDestroy, []resource.TestStep{
		// verify matching rule syntax
		{
			Config: config + `
variable "dynamic_group_description" { default = "description2" }
variable "dynamic_group_matching_rule" { default = "bad_matching_rule" }
variable "dynamic_group_name" { default = "DevCompartmentDynamicGroup" }
` + compartmentIdVariableStr + GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Required, Create, dynamicGroupRepresentation),
			ExpectError: regexp.MustCompile("Unable to parse matching rule"),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + DynamicGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Required, Create, dynamicGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Instance Group for dev compartment"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + DynamicGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + matchingRuleVariableStr + DynamicGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Create, dynamicGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "Instance Group for dev compartment"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRuleValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + matchingRule2VariableStr + DynamicGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Update, dynamicGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "matching_rule", matchingRule2ValueStr),
				resource.TestCheckResourceAttr(resourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config + matchingRule2VariableStr +
				GenerateDataSourceFromRepresentationMap("oci_identity_dynamic_groups", "test_dynamic_groups", Optional, Update, dynamicGroupDataSourceRepresentation) +
				compartmentIdVariableStr + DynamicGroupResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_dynamic_group", "test_dynamic_group", Optional, Update, dynamicGroupRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.matching_rule", matchingRule2ValueStr),
				resource.TestCheckResourceAttr(datasourceName, "dynamic_groups.0.name", "DevCompartmentDynamicGroup"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "dynamic_groups.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckIdentityDynamicGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_dynamic_group" {
			noResourceFound = false
			request := oci_identity.GetDynamicGroupRequest{}

			tmp := rs.Primary.ID
			request.DynamicGroupId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetDynamicGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.DynamicGroupLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
