// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

const (
	vmClusterNetworkNumDnsServers     = 2
	vmClusterNetworkNumNtpServers     = 1
	vmClusterNetworkNumVmNetworks     = 1
	vmClusterNetworkNumNodePerNetwork = 1
	vmClusterNetworkNumFreeformTags   = 1
	vmClusterNetworkIsScanEnabled     = false
)

var (
	DataccVmClusterNetworkRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DataccVmClusterNetworkRepresentation)

	DataccVmClusterNetworkResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DataccVmClusterNetworkRepresentation)

	DataccVmClusterNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_network_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datacc_vm_cluster_network.test_vm_cluster_network.id}`},
	}

	DataccVmClusterNetworkDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_cluster_network_display_name}`, Update: `${var.vm_cluster_network_display_name_for_update}`},
		"infrastructure_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_cluster_network_infrastructure_id}`},
		"is_scan_enabled":          acctest.Representation{RepType: acctest.Optional, Create: strconv.FormatBool(vmClusterNetworkIsScanEnabled)},
		"node_count":               acctest.Representation{RepType: acctest.Optional, Create: strconv.Itoa(vmClusterNetworkNumNodePerNetwork)},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`REQUIRES_VALIDATION`}, Update: []string{`REQUIRES_VALIDATION`}},
		"vm_network_consumer_type": acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_cluster_network_consumer_type}`},
	}

	DataccVmClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_display_name}`, Update: `${var.vm_cluster_network_display_name_for_update}`},
		"infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_infrastructure_id}`},
		"dns_servers":       acctest.Representation{RepType: acctest.Required, Create: []string{`${var.vm_cluster_network_dns_server_0}`, `${var.vm_cluster_network_dns_server_1}`}},
		"ntp_servers":       acctest.Representation{RepType: acctest.Required, Create: []string{`${var.vm_cluster_network_ntp_server_0}`}},
		"vm_networks": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DataccVmClusterNetworkPrimaryVmNetworkRepresentation},
		},
		"consumer_type": acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_cluster_network_consumer_type}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"node_count":    acctest.Representation{RepType: acctest.Optional, Create: strconv.Itoa(vmClusterNetworkNumNodePerNetwork)},
	}
	DataccVmClusterNetworkPrimaryVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_domain_name}`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_gateway}`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_netmask}`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_network_type}`},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_vlan_id}`},
		"nodes": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: DataccVmClusterNetworkPrimaryVmNetworkPrimaryNodeRepresentation},
		},
		"prefix": acctest.Representation{RepType: acctest.Optional, Create: `${var.vm_cluster_network_network_0_prefix}`},
	}
	DataccVmClusterNetworkPrimaryVmNetworkPrimaryNodeRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_node_0_hostname}`},
		"ip":       acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_network_0_node_0_ip}`},
	}
)

func GenerateTFVariableStrings(resourceType string) string {
	type tfVariable struct {
		name  string
		value string
	}
	var variables []tfVariable
	tfVarPrefix := globalvar.TfEnvPrefix
	resourcePrefix := tfVarPrefix + resourceType + "_"
	for _, envVar := range os.Environ() {
		key, value, ok := strings.Cut(envVar, "=")
		if !ok || !strings.HasPrefix(key, resourcePrefix) {
			continue
		}
		variableName := strings.TrimPrefix(key, tfVarPrefix)
		if variableName == "" {
			continue
		}
		variables = append(variables, tfVariable{
			name:  variableName,
			value: value,
		})
	}
	// Keep the generated HCL in a stable order so test configs are deterministic
	// and easier to read/debug across runs.
	sort.Slice(variables, func(i, j int) bool {
		return variables[i].name < variables[j].name
	})
	var builder strings.Builder
	for _, variable := range variables {
		builder.WriteString(fmt.Sprintf("variable %q { default = %q }\n", variable.name, variable.value))
	}
	return builder.String()
}

// issue-routing-tag: datacc/default
func TestDataccVmClusterNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataccVmClusterNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// override terraform-federation-test profile with our own user profile
	if overrideProfile := os.Getenv("datacc_custom_config_file_profile_override"); overrideProfile != "" {
		t.Setenv(globalvar.TfEnvPrefix+globalvar.ConfigFileProfileAttrName, overrideProfile)
		t.Setenv(globalvar.TfEnvPrefix+globalvar.AuthAttrName, "")
		t.Setenv(globalvar.AuthAttrName, globalvar.AuthSecurityToken)
	}

	const testResourceType = "vm_cluster_network"
	tfVariableStr := GenerateTFVariableStrings(testResourceType)
	getTFVar := func(variableName string) string {
		return os.Getenv(globalvar.TfEnvPrefix + testResourceType + "_" + variableName)
	}

	compartmentId := getTFVar("compartment_id")
	compartmentIdU := getTFVar("compartment_id_for_update")
	displayName := getTFVar("display_name")
	displayNameU := getTFVar("display_name_for_update")
	infrastructureId := getTFVar("infrastructure_id")
	consumerType := getTFVar("consumer_type")
	network0DomainName := getTFVar("network_0_domain_name")
	network0Gateway := getTFVar("network_0_gateway")
	network0Netmask := getTFVar("network_0_netmask")
	network0NetworkType := getTFVar("network_0_network_type")
	network0Prefix := getTFVar("network_0_prefix")
	network0VlanId := getTFVar("network_0_vlan_id")
	network0Node0Hostname := getTFVar("network_0_node_0_hostname")
	network0Node0Ip := getTFVar("network_0_node_0_ip")

	resourceName := "oci_datacc_vm_cluster_network.test_vm_cluster_network"
	datasourceName := "data.oci_datacc_vm_cluster_networks.test_vm_cluster_networks"
	singularDatasourceName := "data.oci_datacc_vm_cluster_network.test_vm_cluster_network"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+tfVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, DataccVmClusterNetworkRepresentation), "datacc", "vmClusterNetwork", t)

	acctest.ResourceTest(t, testAccCheckDataccVmClusterNetworkDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + tfVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DataccVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.domain_name", network0DomainName),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.gateway", network0Gateway),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.netmask", network0Netmask),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.network_type", network0NetworkType),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.#", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.hostname", network0Node0Hostname),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.ip", network0Node0Ip),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + tfVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + tfVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, DataccVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumer_type", consumerType),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(vmClusterNetworkNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmClusterNetworkNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "node_count", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(vmClusterNetworkNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.domain_name", network0DomainName),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.gateway", network0Gateway),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.netmask", network0Netmask),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.network_type", network0NetworkType),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.#", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.hostname", network0Node0Hostname),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.ip", network0Node0Ip),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.prefix", network0Prefix),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.vlan_id", network0VlanId),

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
			Config: config + tfVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(DataccVmClusterNetworkRepresentation, map[string]interface{}{
					"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vm_cluster_network_compartment_id_for_update}`},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "consumer_type", consumerType),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(vmClusterNetworkNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmClusterNetworkNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "node_count", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(vmClusterNetworkNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.domain_name", network0DomainName),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.gateway", network0Gateway),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.netmask", network0Netmask),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.network_type", network0NetworkType),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.#", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.hostname", network0Node0Hostname),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.ip", network0Node0Ip),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.prefix", network0Prefix),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.vlan_id", network0VlanId),

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
			Config: config + tfVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DataccVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "consumer_type", consumerType),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(resourceName, "dns_servers.#", strconv.Itoa(vmClusterNetworkNumDnsServers)),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", strconv.Itoa(vmClusterNetworkNumFreeformTags)),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(resourceName, "node_count", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "ntp_servers.#", strconv.Itoa(vmClusterNetworkNumNtpServers)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.domain_name", network0DomainName),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.gateway", network0Gateway),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.netmask", network0Netmask),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.network_type", network0NetworkType),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.#", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.hostname", network0Node0Hostname),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.nodes.0.ip", network0Node0Ip),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.prefix", network0Prefix),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.0.vlan_id", network0VlanId),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_vm_cluster_networks", "test_vm_cluster_networks", acctest.Optional, acctest.Update, DataccVmClusterNetworkDataSourceRepresentation) +
				tfVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DataccVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(datasourceName, "infrastructure_id", infrastructureId),
				resource.TestCheckResourceAttr(datasourceName, "is_scan_enabled", strconv.FormatBool(vmClusterNetworkIsScanEnabled)),
				resource.TestCheckResourceAttr(datasourceName, "node_count", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_network_collection.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_network_collection.0.items.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DataccVmClusterNetworkSingularDataSourceRepresentation) +
				tfVariableStr + DataccVmClusterNetworkResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumer_type", consumerType),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", displayNameU),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_servers.#", strconv.Itoa(vmClusterNetworkNumDnsServers)),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", strconv.Itoa(vmClusterNetworkNumFreeformTags)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_scan_enabled", strconv.FormatBool(vmClusterNetworkIsScanEnabled)),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_count", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp_servers.#", strconv.Itoa(vmClusterNetworkNumNtpServers)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.#", strconv.Itoa(vmClusterNetworkNumVmNetworks)),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.domain_name", network0DomainName),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.gateway", network0Gateway),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.netmask", network0Netmask),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.network_type", network0NetworkType),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.nodes.#", strconv.Itoa(vmClusterNetworkNumNodePerNetwork)),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.nodes.0.hostname", network0Node0Hostname),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.nodes.0.ip", network0Node0Ip),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.0.prefix", network0Prefix),
			),
		},
		// verify resource import
		{
			Config:                  config + DataccVmClusterNetworkRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataccVmClusterNetworkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BaseinfraClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacc_vm_cluster_network" {
			noResourceFound = false
			request := oci_datacc.GetVmClusterNetworkRequest{}

			tmp := rs.Primary.ID
			request.VmClusterNetworkId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")

			response, err := client.GetVmClusterNetwork(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacc.VmClusterNetworkLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataccVmClusterNetwork") {
		resource.AddTestSweepers("DataccVmClusterNetwork", &resource.Sweeper{
			Name:         "DataccVmClusterNetwork",
			Dependencies: acctest.DependencyGraph["vmClusterNetwork"],
			F:            sweepDataccVmClusterNetworkResource,
		})
	}
}

func sweepDataccVmClusterNetworkResource(compartment string) error {
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()
	vmClusterNetworkIds, err := getDataccVmClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterNetworkId := range vmClusterNetworkIds {
		if ok := acctest.SweeperDefaultResourceId[vmClusterNetworkId]; !ok {
			deleteVmClusterNetworkRequest := oci_datacc.DeleteVmClusterNetworkRequest{}

			deleteVmClusterNetworkRequest.VmClusterNetworkId = &vmClusterNetworkId

			deleteVmClusterNetworkRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datacc")
			_, error := baseinfraClient.DeleteVmClusterNetwork(context.Background(), deleteVmClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting VmClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterNetworkId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmClusterNetworkId, DataccVmClusterNetworkSweepWaitCondition, time.Duration(3*time.Minute),
				DataccVmClusterNetworkSweepResponseFetchOperation, "datacc", true)
		}
	}
	return nil
}

func getDataccVmClusterNetworkIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	baseinfraClient := acctest.GetTestClients(&schema.ResourceData{}).BaseinfraClient()

	listVmClusterNetworksRequest := oci_datacc.ListVmClusterNetworksRequest{}
	listVmClusterNetworksRequest.CompartmentId = &compartmentId
	listVmClusterNetworksRequest.LifecycleState = []oci_datacc.VmClusterNetworkLifecycleStateEnum{oci_datacc.VmClusterNetworkLifecycleStateRequiresValidation}
	listVmClusterNetworksResponse, err := baseinfraClient.ListVmClusterNetworks(context.Background(), listVmClusterNetworksRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VmClusterNetwork list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, vmClusterNetwork := range listVmClusterNetworksResponse.Items {
		id := *vmClusterNetwork.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterNetworkId", id)
	}
	return resourceIds, nil
}

func DataccVmClusterNetworkSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterNetworkResponse, ok := response.Response.(oci_datacc.GetVmClusterNetworkResponse); ok {
		return vmClusterNetworkResponse.LifecycleState != oci_datacc.VmClusterNetworkLifecycleStateDeleted
	}
	return false
}

func DataccVmClusterNetworkSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BaseinfraClient().GetVmClusterNetwork(context.Background(), oci_datacc.GetVmClusterNetworkRequest{
		VmClusterNetworkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
