// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_oce "github.com/oracle/oci-go-sdk/v31/oce"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OceInstanceRequiredOnlyResource = OceInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Required, Create, oceInstanceRepresentation)

	OceInstanceResourceConfig = OceInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Optional, Update, oceInstanceRepresentation)

	oceInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"oce_instance_id": Representation{repType: Required, create: `${oci_oce_oce_instance.test_oce_instance.id}`},
	}

	oceInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		//"display_name":   Representation{repType: Optional, create: `displayName`},
		"state":  Representation{repType: Optional, create: `Active`},
		"filter": RepresentationGroup{Required, oceInstanceDataSourceFilterRepresentation}}
	oceInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_oce_oce_instance.test_oce_instance.id}`}},
	}

	instanceName              = randomString(15, charsetWithoutDigits)
	oceInstanceRepresentation = map[string]interface{}{
		"admin_email":              Representation{repType: Required, create: `${var.admin_email}`},
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"idcs_access_token":        Representation{repType: Required, create: `${var.idcs_access_token}`},
		"name":                     Representation{repType: Required, create: instanceName},
		"object_storage_namespace": Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}`},
		"tenancy_id":               Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.id}`},
		"tenancy_name":             Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}`},
		"defined_tags":             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":              Representation{repType: Optional, create: `description`, update: `description2`},
		"freeform_tags":            Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
		"instance_access_type":     Representation{repType: Optional, create: `PUBLIC`},
		"instance_license_type":    Representation{repType: Optional, create: `NEW`, update: `BYOL`},
		"instance_usage_type":      Representation{repType: Optional, create: `NONPRIMARY`},
		"upgrade_schedule":         Representation{repType: Optional, create: `UPGRADE_IMMEDIATELY`},
		"waf_primary_domain":       Representation{repType: Optional, create: `oracle.com`, update: `java.com`},
	}

	OceInstanceResourceDependencies = generateDataSourceFromRepresentationMap("oci_identity_tenancy", "test_tenancy", Required, Create, tenancySingularDataSourceRepresentation) +
		DefinedTagsDependencies
)

func TestOceOceInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOceOceInstanceResource_basic")
	defer httpreplay.SaveScenario()

	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestOceOceInstanceResource_basic") {
		t.Skip("Skipping suppressed TestOceOceInstanceResource_basic")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	adminEmail := getEnvSettingWithBlankDefault("admin_email")
	adminEmailVariableStr := fmt.Sprintf("variable \"admin_email\" { default = \"%s\" }\n", adminEmail)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_oce_oce_instance.test_oce_instance"
	datasourceName := "data.oci_oce_oce_instances.test_oce_instances"
	singularDatasourceName := "data.oci_oce_oce_instance.test_oce_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOceOceInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Required, Create, oceInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "name", instanceName),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Optional, Create, oceInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "guid"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
					resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "instance_license_type", "NEW"),
					resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "NONPRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "name", instanceName),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),
					resource.TestCheckResourceAttr(resourceName, "waf_primary_domain", "oracle.com"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Optional, Create,
						representationCopyWithNewProperties(oceInstanceRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "guid"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
					resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "instance_license_type", "NEW"),
					resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "NONPRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "name", instanceName),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),
					resource.TestCheckResourceAttr(resourceName, "waf_primary_domain", "oracle.com"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Optional, Update, oceInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "admin_email"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "guid"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_tenancy"),
					resource.TestCheckResourceAttr(resourceName, "instance_access_type", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "instance_license_type", "BYOL"),
					resource.TestCheckResourceAttr(resourceName, "instance_usage_type", "NONPRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "name", instanceName),
					resource.TestCheckResourceAttrSet(resourceName, "object_storage_namespace"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "tenancy_name"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),
					resource.TestCheckResourceAttr(resourceName, "waf_primary_domain", "java.com"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_oce_oce_instances", "test_oce_instances", Optional, Update, oceInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Optional, Update, oceInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(datasourceName, "state", "Active"),

					resource.TestCheckResourceAttr(datasourceName, "oce_instances.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.admin_email"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.guid"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.idcs_tenancy"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_access_type", "PUBLIC"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_license_type", "BYOL"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.instance_usage_type", "NONPRIMARY"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.name", instanceName),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.object_storage_namespace"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.state_message"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.tenancy_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.tenancy_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "oce_instances.0.time_updated"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.upgrade_schedule", "UPGRADE_IMMEDIATELY"),
					resource.TestCheckResourceAttr(datasourceName, "oce_instances.0.waf_primary_domain", "java.com"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_oce_oce_instance", "test_oce_instance", Required, Create, oceInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "oce_instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "admin_email"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "guid"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_tenancy"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_access_type", "PUBLIC"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_license_type", "BYOL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "instance_usage_type", "NONPRIMARY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", instanceName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "object_storage_namespace"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state_message"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "upgrade_schedule", "UPGRADE_IMMEDIATELY"),
					resource.TestCheckResourceAttr(singularDatasourceName, "waf_primary_domain", "java.com"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + adminEmailVariableStr + idcsAccessTokenVariableStr + OceInstanceResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"idcs_access_token",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckOceOceInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).oceInstanceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oce_oce_instance" {
			noResourceFound = false
			request := oci_oce.GetOceInstanceRequest{}

			tmp := rs.Primary.ID
			request.OceInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "oce")

			response, err := client.GetOceInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oce.OceInstanceLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("OceOceInstance") {
		resource.AddTestSweepers("OceOceInstance", &resource.Sweeper{
			Name:         "OceOceInstance",
			Dependencies: DependencyGraph["oceInstance"],
			F:            sweepOceOceInstanceResource,
		})
	}
}

func sweepOceOceInstanceResource(compartment string) error {
	oceInstanceClient := GetTestClients(&schema.ResourceData{}).oceInstanceClient()
	oceInstanceIds, err := getOceInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, oceInstanceId := range oceInstanceIds {
		if ok := SweeperDefaultResourceId[oceInstanceId]; !ok {
			deleteOceInstanceRequest := oci_oce.DeleteOceInstanceRequest{}

			deleteOceInstanceRequest.OceInstanceId = &oceInstanceId

			deleteOceInstanceRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "oce")
			_, error := oceInstanceClient.DeleteOceInstance(context.Background(), deleteOceInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting OceInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", oceInstanceId, error)
				continue
			}
			waitTillCondition(testAccProvider, &oceInstanceId, oceInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				oceInstanceSweepResponseFetchOperation, "oce", true)
		}
	}
	return nil
}

func getOceInstanceIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "OceInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oceInstanceClient := GetTestClients(&schema.ResourceData{}).oceInstanceClient()

	listOceInstancesRequest := oci_oce.ListOceInstancesRequest{}
	listOceInstancesRequest.CompartmentId = &compartmentId
	listOceInstancesRequest.LifecycleState = oci_oce.ListOceInstancesLifecycleStateActive
	listOceInstancesResponse, err := oceInstanceClient.ListOceInstances(context.Background(), listOceInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OceInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oceInstance := range listOceInstancesResponse.Items {
		id := *oceInstance.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "OceInstanceId", id)
	}
	return resourceIds, nil
}

func oceInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oceInstanceResponse, ok := response.Response.(oci_oce.GetOceInstanceResponse); ok {
		return oceInstanceResponse.LifecycleState != oci_oce.OceInstanceLifecycleStateDeleted
	}
	return false
}

func oceInstanceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.oceInstanceClient().GetOceInstance(context.Background(), oci_oce.GetOceInstanceRequest{
		OceInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
