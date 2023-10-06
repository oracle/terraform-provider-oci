// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ContainerengineClusterNamespaceProfileVersionRequiredOnlyResource = ContainerengineClusterNamespaceProfileVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileVersionRepresentation)

	ContainerengineClusterNamespaceProfileVersionResourceConfig = ContainerengineClusterNamespaceProfileVersionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileVersionRepresentation)

	ContainerengineClusterNamespaceProfileVersionSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_namespace_profile_version_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version.id}`},
	}

	ContainerengineClusterNamespaceProfileVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	ContainerengineClusterNamespaceProfileVersionRepresentation = map[string]interface{}{
		"admin_cluster_role_name":        acctest.Representation{RepType: acctest.Required, Create: `cluster-admin`},
		"cluster_namespace_profile_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id}`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `testClusterNamespaceProfileVersion`},
		"allowed_namespace_annotations":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: ContainerengineClusterNamespaceProfileVersionAllowedNamespaceAnnotationsRepresentation1}, {RepType: acctest.Required, Group: ContainerengineClusterNamespaceProfileVersionAllowedNamespaceAnnotationsRepresentation2}},
		"allowed_namespace_labels":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ContainerengineClusterNamespaceProfileVersionAllowedNamespaceLabelsRepresentation},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"fixed_namespace_annotations":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceProfileVersionFixedNamespaceAnnotationsRepresentation},
		"fixed_namespace_labels":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceProfileVersionFixedNamespaceLabelsRepresentation},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_deprecated":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"required_namespace_annotations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceProfileVersionRequiredNamespaceAnnotationsRepresentation},
		"required_namespace_labels":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ContainerengineClusterNamespaceProfileVersionRequiredNamespaceLabelsRepresentation},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreContainerengineClusterNamespaceProfileDefinedTagsChangesRepresentation},
		"depends_on":                     acctest.Representation{RepType: acctest.Required, Create: []string{`oci_containerengine_cluster.test_cluster`, `oci_containerengine_cluster_attachment.test_cluster_attachment`}},
	}
	ContainerengineClusterNamespaceProfileVersionAllowedNamespaceAnnotationsRepresentation1 = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `allowed-annotation-1`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowed-annotation-value-1`, `allowed-annotation-value-2`}},
	}
	ContainerengineClusterNamespaceProfileVersionAllowedNamespaceAnnotationsRepresentation2 = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `allowed-annotation-2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowed-annotation-value-3`, `allowed-annotation-value-4`, `allowed-annotation-value-5`}},
	}
	ContainerengineClusterNamespaceProfileVersionAllowedNamespaceLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `allowed-label`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: []string{`allowed-label-value`}},
	}
	ContainerengineClusterNamespaceProfileVersionFixedNamespaceAnnotationsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `fixed-annotation`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `fixed-annotation-value`},
	}
	ContainerengineClusterNamespaceProfileVersionFixedNamespaceLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `fixed-label`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `fixed-label-value`},
	}
	ContainerengineClusterNamespaceProfileVersionRequiredNamespaceAnnotationsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `required-annotation`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: []string{`required-annotation-value`}},
	}
	ContainerengineClusterNamespaceProfileVersionRequiredNamespaceLabelsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Required, Create: `required-label`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: []string{`required-label-value`}},
	}

	ContainerengineClusterNamespaceProfileVersionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_attachment", "test_cluster_attachment", acctest.Required, acctest.Create, ContainerengineClusterAttachmentRepresentation) +
		ContainerengineClusterAttachmentResourceDependencies
)

// issue-routing-tag: containerengine/default
func TestContainerengineClusterNamespaceProfileVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineClusterNamespaceProfileVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version"
	datasourceName := "data.oci_containerengine_cluster_namespace_profile_versions.test_cluster_namespace_profile_versions"
	singularDatasourceName := "data.oci_containerengine_cluster_namespace_profile_version.test_cluster_namespace_profile_version"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ContainerengineClusterNamespaceProfileVersionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceProfileVersionRepresentation), "containerengine", "clusterNamespaceProfileVersion", t)

	acctest.ResourceTest(t, testAccCheckContainerengineClusterNamespaceProfileVersionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_cluster_role_name", "cluster-admin"),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "testClusterNamespaceProfileVersion"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Create, ContainerengineClusterNamespaceProfileVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_cluster_role_name", "cluster-admin"),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_annotations.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_labels", map[string]string{
					"key": "allowed-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_annotations", map[string]string{
					"key":   "fixed-annotation",
					"value": "fixed-annotation-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_labels", map[string]string{
					"key":   "fixed-label",
					"value": "fixed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_deprecated", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "testClusterNamespaceProfileVersion"),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithPropertiesContainingNestedSets(resourceName, "required_namespace_annotations", map[string]interface{}{
					"key":   "required-annotation",
					"value": [1]string{"required-annotation-value"},
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithPropertiesContainingNestedSets(resourceName, "required_namespace_labels", map[string]interface{}{
					"key":   "required-label",
					"value": [1]string{"required-label-value"},
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ContainerengineClusterNamespaceProfileVersionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_cluster_role_name", "cluster-admin"),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_annotations.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_labels", map[string]string{
					"key": "allowed-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_annotations", map[string]string{
					"key":   "fixed-annotation",
					"value": "fixed-annotation-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_labels", map[string]string{
					"key":   "fixed-label",
					"value": "fixed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_deprecated", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "testClusterNamespaceProfileVersion"),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_annotations", map[string]string{
					"key": "required-annotation",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_labels", map[string]string{
					"key": "required-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "admin_cluster_role_name", "cluster-admin"),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_annotations.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "allowed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "allowed_namespace_labels", map[string]string{
					"key": "allowed-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_namespace_profile_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_annotations", map[string]string{
					"key":   "fixed-annotation",
					"value": "fixed-annotation-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_labels", map[string]string{
					"key":   "fixed-label",
					"value": "fixed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_deprecated", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "testClusterNamespaceProfileVersion"),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_annotations", map[string]string{
					"key": "required-annotation",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_labels", map[string]string{
					"key": "required-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_versions", "test_cluster_namespace_profile_versions", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileVersionDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Optional, acctest.Update, ContainerengineClusterNamespaceProfileVersionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_profile_version_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cluster_namespace_profile_version_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_containerengine_cluster_namespace_profile_version", "test_cluster_namespace_profile_version", acctest.Required, acctest.Create, ContainerengineClusterNamespaceProfileVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ContainerengineClusterNamespaceProfileVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_namespace_profile_version_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "admin_cluster_role_name", "cluster-admin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_namespace_annotations.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-1",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "allowed_namespace_annotations", map[string]string{
					"key": "allowed-annotation-2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "allowed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "allowed_namespace_labels", map[string]string{
					"key": "allowed-label",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fixed_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_annotations", map[string]string{
					"key":   "fixed-annotation",
					"value": "fixed-annotation-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "fixed_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "fixed_namespace_labels", map[string]string{
					"key":   "fixed-label",
					"value": "fixed-label-value",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_deprecated", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "testClusterNamespaceProfileVersion"),
				resource.TestCheckResourceAttr(singularDatasourceName, "required_namespace_annotations.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_annotations", map[string]string{
					"key": "required-annotation",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "required_namespace_labels.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "required_namespace_labels", map[string]string{
					"key": "required-label",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ContainerengineClusterNamespaceProfileVersionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckContainerengineClusterNamespaceProfileVersionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ContainerEngineClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_containerengine_cluster_namespace_profile_version" {
			noResourceFound = false
			request := oci_containerengine.GetClusterNamespaceProfileVersionRequest{}

			tmp := rs.Primary.ID
			request.ClusterNamespaceProfileVersionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")

			response, err := client.GetClusterNamespaceProfileVersion(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateDeleted): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ContainerengineClusterNamespaceProfileVersion") {
		resource.AddTestSweepers("ContainerengineClusterNamespaceProfileVersion", &resource.Sweeper{
			Name:         "ContainerengineClusterNamespaceProfileVersion",
			Dependencies: acctest.DependencyGraph["clusterNamespaceProfileVersion"],
			F:            sweepContainerengineClusterNamespaceProfileVersionResource,
		})
	}
}

func sweepContainerengineClusterNamespaceProfileVersionResource(compartment string) error {
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()
	clusterNamespaceProfileVersionIds, err := getContainerengineClusterNamespaceProfileVersionIds(compartment)
	if err != nil {
		return err
	}
	for _, clusterNamespaceProfileVersionId := range clusterNamespaceProfileVersionIds {
		if ok := acctest.SweeperDefaultResourceId[clusterNamespaceProfileVersionId]; !ok {
			deleteClusterNamespaceProfileVersionRequest := oci_containerengine.DeleteClusterNamespaceProfileVersionRequest{}

			deleteClusterNamespaceProfileVersionRequest.ClusterNamespaceProfileVersionId = &clusterNamespaceProfileVersionId

			deleteClusterNamespaceProfileVersionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "containerengine")
			_, error := containerEngineClient.DeleteClusterNamespaceProfileVersion(context.Background(), deleteClusterNamespaceProfileVersionRequest)
			if error != nil {
				fmt.Printf("Error deleting ClusterNamespaceProfileVersion %s %s, It is possible that the resource is already deleted. Please verify manually \n", clusterNamespaceProfileVersionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &clusterNamespaceProfileVersionId, ContainerengineClusterNamespaceProfileVersionSweepWaitCondition, time.Duration(3*time.Minute),
				ContainerengineClusterNamespaceProfileVersionSweepResponseFetchOperation, "containerengine", true)
		}
	}
	return nil
}

func getContainerengineClusterNamespaceProfileVersionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ClusterNamespaceProfileVersionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	containerEngineClient := acctest.GetTestClients(&schema.ResourceData{}).ContainerEngineClient()

	listClusterNamespaceProfileVersionsRequest := oci_containerengine.ListClusterNamespaceProfileVersionsRequest{}
	listClusterNamespaceProfileVersionsRequest.CompartmentId = &compartmentId
	listClusterNamespaceProfileVersionsRequest.LifecycleState = oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateActive
	listClusterNamespaceProfileVersionsResponse, err := containerEngineClient.ListClusterNamespaceProfileVersions(context.Background(), listClusterNamespaceProfileVersionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ClusterNamespaceProfileVersion list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, clusterNamespaceProfileVersion := range listClusterNamespaceProfileVersionsResponse.Items {
		id := *clusterNamespaceProfileVersion.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ClusterNamespaceProfileVersionId", id)
	}
	return resourceIds, nil
}

func ContainerengineClusterNamespaceProfileVersionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if clusterNamespaceProfileVersionResponse, ok := response.Response.(oci_containerengine.GetClusterNamespaceProfileVersionResponse); ok {
		return clusterNamespaceProfileVersionResponse.LifecycleState != oci_containerengine.ClusterNamespaceProfileVersionLifecycleStateDeleted
	}
	return false
}

func ContainerengineClusterNamespaceProfileVersionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ContainerEngineClient().GetClusterNamespaceProfileVersion(context.Background(), oci_containerengine.GetClusterNamespaceProfileVersionRequest{
		ClusterNamespaceProfileVersionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
