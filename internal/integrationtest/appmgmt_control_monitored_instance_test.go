// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MonitoredInstanceResourceConfig                   = MonitoredInstanceResourceDependencies
	monitoredInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	monitoredInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	appmgmtControlInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_management_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `false`},
	}

	privateVnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	MonitoredInstanceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"create_vnic_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: privateVnicDetailsRepresentation},
		"source_details":      acctest.RepresentationGroup{RepType: acctest.Required, Group: sourceDetailsRepresentation},
		"agent_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: appmgmtControlInstanceAgentConfigRepresentation},
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.OsManagedImageOCID[var.region]}`}, //variable defined in helpers.go
	})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":  acctest.Representation{RepType: acctest.Required, Create: `testvcn`},
			"cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.1.0.0/16`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, internetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", acctest.Required, acctest.Create, routeTablesRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": acctest.Representation{RepType: acctest.Required, Create: `10.1.20.0/24`}, "dns_label": acctest.Representation{RepType: acctest.Required, Create: `testsubnet`}, "route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`}})) +
		AvailabilityDomainConfig + utils.OsManagedImageIdsVariable
)

// issue-routing-tag: appmgmt_control/default
func TestAppmgmtControlMonitoredInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAppmgmtControlMonitoredInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_appmgmt_control_monitored_instances.test_monitored_instances"
	singularDatasourceName := "data.oci_appmgmt_control_monitored_instance.test_monitored_instance"

	acctest.SaveConfigContent("", "", "", t) //original, disabled for dynamic dependency creation

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Create dependencies
		{
			Config: config + compartmentIdVariableStr + MonitoredInstanceResourceDependencies,
			Check: func(s *terraform.State) (err error) {
				log.Printf("[DEBUG] Appmgmt Control Monitored Resource should be created")
				return nil
			},
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_appmgmt_control_monitored_instances", "test_monitored_instances", acctest.Required, acctest.Create, monitoredInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoredInstanceResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "monitored_instance_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_appmgmt_control_monitored_instance", "test_monitored_instance", acctest.Required, acctest.Create, monitoredInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoredInstanceResourceDependencies,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitoring_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
	})
}
