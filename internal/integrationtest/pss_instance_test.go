// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_pss "github.com/oracle/oci-go-sdk/v65/pss"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PssInstanceIdentityDomainUrl         = utils.GetEnvSettingWithBlankDefault("pss_identity_domain_url")
	PssInstanceAdminIdentityDomainUserId = utils.GetEnvSettingWithBlankDefault("pss_admin_identity_domain_user_id")

	PssInstanceRequiredOnlyResource = PssInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Required, acctest.Create, PssInstanceRepresentation)

	PssInstanceResourceConfig = PssInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Update, PssInstanceRepresentation)

	PssInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_pss_instance.test_instance.id}`},
	}

	PssInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_ocid}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_pss_instance.test_instance.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: PssInstanceDataSourceFilterRepresentation}}
	PssInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_pss_instance.test_instance.id}`}},
	}

	PssInstanceRepresentation = map[string]interface{}{
		"admin_identity_domain_user_id": acctest.Representation{RepType: acctest.Required, Create: `${var.pss_admin_identity_domain_user_id}`},
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_ocid}`},
		"identity_domain_url":           acctest.Representation{RepType: acctest.Required, Create: `${var.pss_identity_domain_url}`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${var.home_region_defined_tag_namespace_name}.${var.home_region_defined_tag_name}" = "value"})}`, Update: `${tomap({"${var.home_region_defined_tag_namespace_name}.${var.home_region_defined_tag_name}" = "updatedValue"})}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	// Use a service-owned identity domain and an existing administrator in that domain.
	// As in the DDFS tests, use tags pre-created in the tenancy home region.
	PssInstanceResourceDependencies = fmt.Sprintf(`
variable "home_region_defined_tag_namespace_name" {}
variable "home_region_defined_tag_name" {}
variable "pss_identity_domain_url" {
  type    = string
  default = %q
}
variable "pss_admin_identity_domain_user_id" {
  type    = string
  default = %q
}
`, PssInstanceIdentityDomainUrl, PssInstanceAdminIdentityDomainUserId)
)

// issue-routing-tag: pss/default
func TestPssInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPssInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() + `
provider "oci" {
  ignore_defined_tags = [
    "Oracle-Tags.CreatedBy",
    "Oracle-Tags.CreatedOn",
  ]
}
`

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	// Keep Terraform's input aligned with the environment value used by the checks.
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_ocid\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_pss_instance.test_instance"
	datasourceName := "data.oci_pss_instances.test_instances"
	singularDatasourceName := "data.oci_pss_instance.test_instance"
	definedTagAttribute := "defined_tags." + utils.GetEnvSettingWithBlankDefault("home_region_defined_tag_namespace_name") + "." + utils.GetEnvSettingWithBlankDefault("home_region_defined_tag_name")

	var resId, resId2 string
	// Save the optional configuration for generated config/example coverage.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PssInstanceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Create, PssInstanceRepresentation), "pss", "instance", t)

	acctest.ResourceTest(t, testAccCheckPssInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PssInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Required, acctest.Create, PssInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_identity_domain_user_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "identity_domain_url", PssInstanceIdentityDomainUrl),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// As in DDFS, apply optional properties to the existing instance.
		{
			Config: config + compartmentIdVariableStr + PssInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Create, PssInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_identity_domain_user_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, definedTagAttribute, "value"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "identity_domain_url", PssInstanceIdentityDomainUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return nil
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PssInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PssInstanceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_identity_domain_user_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "identity_domain_url", PssInstanceIdentityDomainUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return nil
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PssInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Update, PssInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "admin_identity_domain_user_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, definedTagAttribute, "updatedValue"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "identity_domain_url", PssInstanceIdentityDomainUrl),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return nil
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_pss_instances", "test_instances", acctest.Optional, acctest.Update, PssInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + PssInstanceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Optional, acctest.Update, PssInstanceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "instance_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instance_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrPair(datasourceName, "instance_collection.0.items.0.id", resourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_pss_instance", "test_instance", acctest.Required, acctest.Create, PssInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PssInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "identity_domain_url", PssInstanceIdentityDomainUrl),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_domain_web_auth_client_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "identity_domain_web_auth_client_secret_vault_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ui_endpoint"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + PssInstanceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckPssInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PssInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_pss_instance" {
			noResourceFound = false
			request := oci_pss.GetInstanceRequest{}

			tmp := rs.Primary.ID
			request.InstanceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "pss")

			response, err := client.GetInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_pss.InstanceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("PssInstance") {
		resource.AddTestSweepers("PssInstance", &resource.Sweeper{
			Name:         "PssInstance",
			Dependencies: acctest.DependencyGraph["instance"],
			F:            sweepPssInstanceResource,
		})
	}
}

func sweepPssInstanceResource(compartment string) error {
	instanceClient := acctest.GetTestClients(&schema.ResourceData{}).PssInstanceClient()
	instanceIds, err := getPssInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, instanceId := range instanceIds {
		if ok := acctest.SweeperDefaultResourceId[instanceId]; !ok {
			deleteInstanceRequest := oci_pss.DeleteInstanceRequest{}

			deleteInstanceRequest.InstanceId = &instanceId

			deleteInstanceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "pss")
			_, error := instanceClient.DeleteInstance(context.Background(), deleteInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting Instance %s %s, It is possible that the resource is already deleted. Please verify manually \n", instanceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &instanceId, PssInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				PssInstanceSweepResponseFetchOperation, "pss", true)
		}
	}
	return nil
}

func getPssInstanceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	instanceClient := acctest.GetTestClients(&schema.ResourceData{}).PssInstanceClient()

	listInstancesRequest := oci_pss.ListInstancesRequest{}
	listInstancesRequest.CompartmentId = &compartmentId
	listInstancesRequest.LifecycleState = oci_pss.InstanceLifecycleStateActive
	listInstancesResponse, err := instanceClient.ListInstances(context.Background(), listInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Instance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, instance := range listInstancesResponse.Items {
		id := *instance.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InstanceId", id)
	}
	return resourceIds, nil
}

func PssInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if instanceResponse, ok := response.Response.(oci_pss.GetInstanceResponse); ok {
		return instanceResponse.LifecycleState != oci_pss.InstanceLifecycleStateDeleted
	}
	return false
}

func PssInstanceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PssInstanceClient().GetInstance(context.Background(), oci_pss.GetInstanceRequest{
		InstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
