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
	oci_analytics "github.com/oracle/oci-go-sdk/analytics"
	"github.com/oracle/oci-go-sdk/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AnalyticsInstanceRequiredOnlyResource = AnalyticsInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Required, Create, analyticsInstanceRepresentation)

	AnalyticsInstanceResourceConfig = AnalyticsInstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, analyticsInstanceRepresentation)

	analyticsInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"analytics_instance_id": Representation{repType: Required, create: `${oci_analytics_analytics_instance.test_analytics_instance.id}`},
	}

	analyticsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"capacity_type":  Representation{repType: Optional, create: `OLPU_COUNT`},
		"feature_set":    Representation{repType: Optional, create: `ENTERPRISE_ANALYTICS`},
		"name":           Representation{repType: Optional, create: analyticsinstanceOptionalName},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, analyticsInstanceDataSourceFilterRepresentation}}
	analyticsInstanceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_analytics_analytics_instance.test_analytics_instance.id}`}},
	}

	analyticsinstanceName           = randomString(15, charsetWithoutDigits)
	analyticsinstanceOptionalName   = randomString(15, charsetWithoutDigits)
	analyticsInstanceRepresentation = map[string]interface{}{
		"capacity":       RepresentationGroup{Required, analyticsInstanceCapacityRepresentation},
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"feature_set":    Representation{repType: Required, create: `ENTERPRISE_ANALYTICS`},
		"license_type":   Representation{repType: Required, create: `LICENSE_INCLUDED`, update: `BRING_YOUR_OWN_LICENSE`},
		"name":           Representation{repType: Required, create: analyticsinstanceOptionalName},
		//"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        Representation{repType: Optional, create: `description`, update: `description2`},
		"email_notification": Representation{repType: Optional, create: `emailNotification`, update: `emailNotification2`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"idcs_access_token":  Representation{repType: Required, create: `${var.idcs_access_token}`},
	}
	analyticsInstanceCapacityRepresentation = map[string]interface{}{
		"capacity_type":  Representation{repType: Required, create: `OLPU_COUNT`},
		"capacity_value": Representation{repType: Required, create: `2`},
	}

	analyticsInstanceCapacityUpdateRepresentation = map[string]interface{}{
		"capacity_type":  Representation{repType: Required, create: `OLPU_COUNT`},
		"capacity_value": Representation{repType: Required, create: `4`},
	}

	AnalyticsInstanceResourceDependencies = ""
)

func TestAnalyticsAnalyticsInstanceResource_basic(t *testing.T) {
	if strings.Contains(getEnvSettingWithBlankDefault("suppressed_tests"), "TestAnalyticsAnalyticsInstanceResource_basic") {
		t.Skip("Skipping suppressed TestAnalyticsAnalyticsInstanceResource_basic")
	}

	httpreplay.SetScenario("TestAnalyticsAnalyticsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	idcsAccessToken := getEnvSettingWithBlankDefault("idcs_access_token")
	idcsAccessTokenVariableStr := fmt.Sprintf("variable \"idcs_access_token\" { default = \"%s\" }\n", idcsAccessToken)

	resourceName := "oci_analytics_analytics_instance.test_analytics_instance"
	datasourceName := "data.oci_analytics_analytics_instances.test_analytics_instances"
	singularDatasourceName := "data.oci_analytics_analytics_instance.test_analytics_instance"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckAnalyticsAnalyticsInstanceDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Required, Create, getUpdatedRepresentationCopy("name", Representation{repType: Required, create: analyticsinstanceName}, analyticsInstanceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceName),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Create, analyticsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Create,
						representationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "LICENSE_INCLUDED"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, analyticsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify up scaling
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update,
						representationCopyWithNewProperties(representationCopyWithRemovedProperties(analyticsInstanceRepresentation, []string{"capacity"}), map[string]interface{}{
							"capacity": RepresentationGroup{Required, analyticsInstanceCapacityUpdateRepresentation},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "4"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify down scaling
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, analyticsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify stop
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, representationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"state": Representation{repType: Required, create: `INACTIVE`},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify start
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, representationCopyWithNewProperties(analyticsInstanceRepresentation, map[string]interface{}{
						"state": Representation{repType: Required, create: `ACTIVE`},
					})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(resourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(resourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "idcs_access_token"),
					resource.TestCheckResourceAttr(resourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_analytics_analytics_instances", "test_analytics_instances", Optional, Update, analyticsInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceDependencies +
					generateResourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Optional, Update, analyticsInstanceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(datasourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.description", "description2"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(datasourceName, "analytics_instances.0.name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.service_url"),
					resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "analytics_instances.0.time_updated"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_analytics_analytics_instance", "test_analytics_instance", Required, Create, analyticsInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "analytics_instance_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "capacity.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_type", "OLPU_COUNT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "capacity.0.capacity_value", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "email_notification", "emailNotification2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "feature_set", "ENTERPRISE_ANALYTICS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "license_type", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", analyticsinstanceOptionalName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "service_url"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + idcsAccessTokenVariableStr + AnalyticsInstanceResourceConfig,
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

func testAccCheckAnalyticsAnalyticsInstanceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).analyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_analytics_analytics_instance" {
			noResourceFound = false
			request := oci_analytics.GetAnalyticsInstanceRequest{}

			tmp := rs.Primary.ID
			request.AnalyticsInstanceId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "analytics")

			response, err := client.GetAnalyticsInstance(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_analytics.AnalyticsInstanceLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("AnalyticsAnalyticsInstance") {
		resource.AddTestSweepers("AnalyticsAnalyticsInstance", &resource.Sweeper{
			Name:         "AnalyticsAnalyticsInstance",
			Dependencies: DependencyGraph["analyticsInstance"],
			F:            sweepAnalyticsAnalyticsInstanceResource,
		})
	}
}

func sweepAnalyticsAnalyticsInstanceResource(compartment string) error {
	analyticsClient := GetTestClients(&schema.ResourceData{}).analyticsClient()
	analyticsInstanceIds, err := getAnalyticsInstanceIds(compartment)
	if err != nil {
		return err
	}
	for _, analyticsInstanceId := range analyticsInstanceIds {
		if ok := SweeperDefaultResourceId[analyticsInstanceId]; !ok {
			deleteAnalyticsInstanceRequest := oci_analytics.DeleteAnalyticsInstanceRequest{}

			deleteAnalyticsInstanceRequest.AnalyticsInstanceId = &analyticsInstanceId

			deleteAnalyticsInstanceRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "analytics")
			_, error := analyticsClient.DeleteAnalyticsInstance(context.Background(), deleteAnalyticsInstanceRequest)
			if error != nil {
				fmt.Printf("Error deleting AnalyticsInstance %s %s, It is possible that the resource is already deleted. Please verify manually \n", analyticsInstanceId, error)
				continue
			}
			waitTillCondition(testAccProvider, &analyticsInstanceId, analyticsInstanceSweepWaitCondition, time.Duration(3*time.Minute),
				analyticsInstanceSweepResponseFetchOperation, "analytics", true)
		}
	}
	return nil
}

func getAnalyticsInstanceIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "AnalyticsInstanceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	analyticsClient := GetTestClients(&schema.ResourceData{}).analyticsClient()

	listAnalyticsInstancesRequest := oci_analytics.ListAnalyticsInstancesRequest{}
	listAnalyticsInstancesRequest.CompartmentId = &compartmentId
	listAnalyticsInstancesRequest.LifecycleState = oci_analytics.ListAnalyticsInstancesLifecycleStateActive
	listAnalyticsInstancesResponse, err := analyticsClient.ListAnalyticsInstances(context.Background(), listAnalyticsInstancesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AnalyticsInstance list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, analyticsInstance := range listAnalyticsInstancesResponse.Items {
		id := *analyticsInstance.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "AnalyticsInstanceId", id)
	}
	return resourceIds, nil
}

func analyticsInstanceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if analyticsInstanceResponse, ok := response.Response.(oci_analytics.GetAnalyticsInstanceResponse); ok {
		return analyticsInstanceResponse.LifecycleState != oci_analytics.AnalyticsInstanceLifecycleStateDeleted
	}
	return false
}

func analyticsInstanceSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.analyticsClient().GetAnalyticsInstance(context.Background(), oci_analytics.GetAnalyticsInstanceRequest{
		AnalyticsInstanceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
