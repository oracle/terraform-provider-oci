// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

var (
	CrossConnectRequiredOnlyResource = CrossConnectResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation)

	CrossConnectResourceConfig = CrossConnectResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create, crossConnectRepresentation)

	crossConnectSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	crossConnectDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"cross_connect_group_id": Representation{repType: Optional, create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		//"state":                  Representation{repType: Optional, create: `AVAILABLE`},
		"filter": RepresentationGroup{Required, crossConnectDataSourceFilterRepresentation}}
	crossConnectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_cross_connect.test_cross_connect.id}`}},
	}

	crossConnectRepresentation = map[string]interface{}{
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"location_name":          Representation{repType: Required, create: `Fake Location, Phoenix, AZ`},
		"port_speed_shape_name":  Representation{repType: Required, create: `10 Gbps`},
		"cross_connect_group_id": Representation{repType: Optional, create: `${oci_core_cross_connect_group.test_cross_connect_group.id}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		//"far_cross_connect_or_cross_connect_group_id":  Representation{repType: Optional, create: `${oci_core_far_cross_connect_or_cross_connect_group.test_far_cross_connect_or_cross_connect_group.id}`},
		//"near_cross_connect_or_cross_connect_group_id": Representation{repType: Optional, create: `${oci_core_near_cross_connect_or_cross_connect_group.test_near_cross_connect_or_cross_connect_group.id}`},
		"is_active": Representation{repType: Optional, create: `true`},
	}

	CrossConnectResourceDependencies = CrossConnectGroupResourceConfig
)

func TestCoreCrossConnectResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_cross_connect.test_cross_connect"
	datasourceName := "data.oci_core_cross_connects.test_cross_connects"
	singularDatasourceName := "data.oci_core_cross_connect.test_cross_connect"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreCrossConnectDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					//resource.TestCheckResourceAttrSet(resourceName, "far_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "location_name", "Fake Location, Phoenix, AZ"),
					//resource.TestCheckResourceAttrSet(resourceName, "near_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PENDING_CUSTOMER"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Create, crossConnectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					//resource.TestCheckResourceAttrSet(resourceName, "far_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "location_name", "Fake Location, Phoenix, AZ"),
					//resource.TestCheckResourceAttrSet(resourceName, "near_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CrossConnectResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					//resource.TestCheckResourceAttrSet(resourceName, "far_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "location_name", "Fake Location, Phoenix, AZ"),
					//resource.TestCheckResourceAttrSet(resourceName, "near_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(resourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(resourceName, "state", "PROVISIONED"),

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
					generateDataSourceFromRepresentationMap("oci_core_cross_connects", "test_cross_connects", Optional, Update, crossConnectDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Optional, Update, crossConnectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					//resource.TestCheckResourceAttrSet(datasourceName, "far_cross_connect_or_cross_connect_group_id"),
					//resource.TestCheckResourceAttrSet(datasourceName, "near_cross_connect_or_cross_connect_group_id"),
					//resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

					resource.TestCheckResourceAttr(datasourceName, "cross_connects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.display_name", "displayName2"),
					//resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.far_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.location_name", "Fake Location, Phoenix, AZ"),
					//resource.TestCheckResourceAttrSet(datasourceName, "cross_connects.0.near_cross_connect_or_cross_connect_group_id"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(datasourceName, "cross_connects.0.state", "PROVISIONED"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CrossConnectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_group_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "far_cross_connect_or_cross_connect_group_id"),
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "near_cross_connect_or_cross_connect_group_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "location_name", "Fake Location, Phoenix, AZ"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "port_speed_shape_name", "10 Gbps"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + CrossConnectResourceConfig,
			},
			// verify resource import
			{
				Config:            config + compartmentIdVariableStr + CrossConnectResourceConfig,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"cross_connect_id",
					"is_active",
					"far_cross_connect_or_cross_connect_group_id",
					"near_cross_connect_or_cross_connect_group_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckCoreCrossConnectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_cross_connect" {
			noResourceFound = false
			request := oci_core.GetCrossConnectRequest{}

			tmp := rs.Primary.ID
			request.CrossConnectId = &tmp

			response, err := client.GetCrossConnect(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CrossConnectLifecycleStateTerminated): true,
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
