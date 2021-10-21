// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MonitoredInstanceResourceConfig                   = MonitoredInstanceResourceDependencies
	monitoredInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_instance_id": Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	monitoredInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	appmgmtControlInstanceAgentConfigRepresentation = map[string]interface{}{
		"are_all_plugins_disabled": Representation{RepType: Required, Create: `false`, Update: `false`},
		"is_management_disabled":   Representation{RepType: Required, Create: `false`, Update: `false`},
		"is_monitoring_disabled":   Representation{RepType: Required, Create: `false`, Update: `false`},
	}

	privateVnicDetailsRepresentation = map[string]interface{}{
		"assign_public_ip": Representation{RepType: Required, Create: `false`},
		"subnet_id":        Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	MonitoredInstanceResourceDependencies = GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"create_vnic_details": RepresentationGroup{Required, privateVnicDetailsRepresentation},
		"source_details":      RepresentationGroup{Required, sourceDetailsRepresentation},
		"agent_config":        RepresentationGroup{Required, appmgmtControlInstanceAgentConfigRepresentation},
		"image":               Representation{RepType: Required, Create: `${var.OsManagedImageOCID[var.region]}`}, //variable defined in helpers.go
	})) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label":  Representation{RepType: Required, Create: `testvcn`},
			"cidr_block": Representation{RepType: Required, Create: `10.1.0.0/16`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", Required, Create, internetGatewayRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", Required, Create, routeTablesRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{"availability_domain": Representation{RepType: Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`}, "cidr_block": Representation{RepType: Required, Create: `10.1.20.0/24`}, "dns_label": Representation{RepType: Required, Create: `testsubnet`}, "route_table_id": Representation{RepType: Required, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`}})) +
		AvailabilityDomainConfig + OsManagedImageIdsVariable
)

// issue-routing-tag: appmgmt_control/default
func TestAppmgmtControlMonitoredInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAppmgmtControlMonitoredInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_appmgmt_control_monitored_instances.test_monitored_instances"
	singularDatasourceName := "data.oci_appmgmt_control_monitored_instance.test_monitored_instance"

	SaveConfigContent("", "", "", t) //original, disabled for dynamic dependency creation

	ResourceTest(t, nil, []resource.TestStep{
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
				GenerateDataSourceFromRepresentationMap("oci_appmgmt_control_monitored_instances", "test_monitored_instances", Required, Create, monitoredInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoredInstanceResourceDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "monitored_instance_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_appmgmt_control_monitored_instance", "test_monitored_instance", Required, Create, monitoredInstanceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MonitoredInstanceResourceDependencies,

			Check: ComposeAggregateTestCheckFuncWrapper(
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
