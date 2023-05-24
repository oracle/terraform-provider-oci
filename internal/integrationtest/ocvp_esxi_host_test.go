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
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpEsxiHostRequiredOnlyResource = EsxiHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostRepresentation)

	OcvpEsxiHostResourceConfig = EsxiHostOptionalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Update, OcvpEsxiHostRepresentation)

	ReplacementEsxiHostResourceConfig = EsxiHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation)

	OcvpOcvpEsxiHostSingularDataSourceRepresentation = map[string]interface{}{
		"esxi_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_esxi_host.test_esxi_host.id}`},
	}
	OcvpOcvpEsxiHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compute_instance_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_esxi_host.test_esxi_host.compute_instance_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_billing_donors_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_swap_billing_only":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"sddc_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpEsxiHostDataSourceFilterRepresentation}}
	OcvpEsxiHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_esxi_host.test_esxi_host.id}`}},
	}
	OcvpSwapBillingOnlyEsxiHostsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_swap_billing_only": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	// for replace node
	failedEsxiHostDataSourceRepresentation = map[string]interface{}{
		"sddc_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.display_name}-1`},
	}
	//for upgrade ESXi host
	nonUpgradedEsxiHostDataSourceRepresentation = map[string]interface{}{
		"sddc_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_sddc.test_sddc.display_name}-2`},
	}

	OcvpSddcX7StandardRepresentation = acctest.RepresentationCopyWithNewProperties(OcvpSddcRepresentation, map[string]interface{}{
		"initial_host_shape_name": acctest.Representation{RepType: acctest.Required, Create: `BM.Standard2.52`},
		"esxi_hosts_count":        acctest.Representation{RepType: acctest.Required, Create: `2`},
	})
	OcvpSddcX7StandardResource              = OcvpSddcResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpSddcX7StandardRepresentation)
	OcvpSddcX7StandardUpgradeRepresentation = acctest.RepresentationCopyWithNewProperties(OcvpSddcX7StandardRepresentation, map[string]interface{}{
		"vmware_software_version": acctest.Representation{RepType: acctest.Required, Create: `7.0 test-L2`},
		"provisioning_vlan_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_provisioning_vlan.id}`},
		"replication_vlan_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vlan.test_replication_vlan.id}`},
	})
	OcvpSddcX7StandardUpgradeResource = OcvpSddcResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", acctest.Required, acctest.Create, OcvpSddcX7StandardUpgradeRepresentation)

	OcvpEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"billing_donor_host_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.donor_host_id}`, Update: `${var.donor_host_id_update}`},
		"capacity_reservation_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_compute_capacity_reservations.test_esxi_host_compute_capacity_reservations.compute_capacity_reservations.0.id}`},
		"compute_availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}`},
		"current_sku":                 acctest.Representation{RepType: acctest.Optional, Create: `THREE_YEARS`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"next_sku":                    acctest.Representation{RepType: acctest.Optional, Create: `ONE_YEAR`, Update: `THREE_YEARS`},
		"host_ocpu_count":             acctest.Representation{RepType: acctest.Optional, Create: `52`},
		"host_shape_name":             acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard2.52`},
	}
	replacementEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"failed_esxi_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_esxi_hosts.failed_esxi_hosts.esxi_host_collection[0].id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `replacement`},
	}
	upgradedEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
		"non_upgraded_esxi_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_esxi_hosts.non_upgraded_esxi_hosts.esxi_host_collection[0].id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `upgrade`},
	}

	donorHostsDataSourceRepresentation = map[string]interface{}{
		"is_billing_donors_only": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: map[string]interface{}{
				"name":   acctest.Representation{RepType: acctest.Required, Create: `host_shape_name`},
				"values": acctest.Representation{RepType: acctest.Required, Create: []string{`BM.Standard2.52`}},
			}},
			{RepType: acctest.Required, Group: map[string]interface{}{
				"name":   acctest.Representation{RepType: acctest.Required, Create: `current_sku`},
				"values": acctest.Representation{RepType: acctest.Required, Create: []string{`THREE_YEARS`}},
			}},
		},
	}

	donorHostIds [2]string

	donorHostsDataSourceDependencies = `
locals {
  sorted_billing_contract_end_date = distinct(sort(data.oci_ocvp_esxi_hosts.test_esxi_hosts.esxi_host_collection[*].billing_contract_end_date))
  donor_hosts = [ for host in data.oci_ocvp_esxi_hosts.test_esxi_hosts.esxi_host_collection: host.id if host.billing_contract_end_date == local.sorted_billing_contract_end_date[0]]
  donor_hosts_update = [ for host in data.oci_ocvp_esxi_hosts.test_esxi_hosts.esxi_host_collection: host.id if host.billing_contract_end_date == local.sorted_billing_contract_end_date[1]]
}

data "oci_ocvp_esxi_host" "donor_host" {
  esxi_host_id = local.donor_hosts[0]
}

data "oci_ocvp_esxi_host" "donor_host_update" {
  esxi_host_id = local.donor_hosts_update[0]
}
`
	ocvpEsxiHostCapacityReservationDataSource = `
data "oci_core_compute_capacity_reservations" "test_esxi_host_compute_capacity_reservations" {
	compartment_id = var.compartment_id
	availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
	state = "ACTIVE"
	
	filter {
		name   = "display_name"
		values = ["tf-esxi-host-test-capacity-reservation"]
	}
}
`

	EsxiHostResourceDependencies = OcvpSddcX7StandardResource + acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "failed_esxi_hosts", acctest.Optional, acctest.Create, failedEsxiHostDataSourceRepresentation)

	EsxiHostUpgradeResourceDependencies  = OcvpSddcX7StandardUpgradeResource + acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "non_upgraded_esxi_hosts", acctest.Optional, acctest.Create, nonUpgradedEsxiHostDataSourceRepresentation)
	EsxiHostOptionalResourceDependencies = OcvpSddcX7StandardUpgradeResource + ocvpEsxiHostCapacityReservationDataSource
)

// issue-routing-tag: ocvp/default
func TestOcvpEsxiHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpEsxiHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_ocvp_esxi_host.test_esxi_host"
	datasourceName := "data.oci_ocvp_esxi_hosts.test_esxi_hosts"
	singularDatasourceName := "data.oci_ocvp_esxi_host.test_esxi_host"

	var resId, resId2, donorHostIdVariableStr, donorHostIdUpdateVariableStr, donorHostId, donorHostIdUpdate string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify donor hosts data source and get donor host ids. Once a donor_host_id is used, the data source will not return it again. To prevent state drift, store donor host ids to variables here
		{
			Config: config + compartmentIdVariableStr + donorHostsDataSourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "test_esxi_hosts", acctest.Optional, acctest.Create, donorHostsDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_billing_donors_only", "true"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.state", "DELETED"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				func(s *terraform.State) (err error) {
					donorHostId, err = acctest.FromInstanceState(s, "data.oci_ocvp_esxi_host.donor_host", "id")
					if err != nil {
						return err
					}
					fmt.Println("Donor host1: ", donorHostId)
					donorHostIdVariableStr = fmt.Sprintf("variable \"donor_host_id\" { default = \"%s\" }\n", donorHostId)

					donorHostIdUpdate, err = acctest.FromInstanceState(s, "data.oci_ocvp_esxi_host.donor_host_update", "id")
					fmt.Println("Donor host2: ", donorHostIdUpdate)
					donorHostIdUpdateVariableStr = fmt.Sprintf("variable \"donor_host_id_update\" { default = \"%s\" }\n", donorHostIdUpdate)

					// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
					acctest.SaveConfigContent(config+compartmentIdVariableStr+EsxiHostUpgradeResourceDependencies+donorHostIdVariableStr+donorHostIdUpdateVariableStr+EsxiHostOptionalResourceDependencies+
						acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostRepresentation), "ocvp", "esxiHost", t)

					return err
				},
			),
		},
	})

	acctest.ResourceTest(t, testAccCheckOcvpEsxiHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "current_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "host_ocpu_count"),
				resource.TestCheckResourceAttrSet(resourceName, "host_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//verify swap billing data source
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "swap_billing_esxi_hosts", acctest.Optional, acctest.Update, OcvpSwapBillingOnlyEsxiHostsDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "compartment_id"),
				resource.TestCheckResourceAttr("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "is_swap_billing_only", `true`),
				resource.TestCheckResourceAttr("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "esxi_host_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "esxi_host_collection.0.compartment_id"),
				resource.TestCheckResourceAttrSet("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "esxi_host_collection.0.sddc_id"),
				resource.TestCheckResourceAttrSet("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "esxi_host_collection.0.id"),
				resource.TestCheckResourceAttrSet("data.oci_ocvp_esxi_hosts.swap_billing_esxi_hosts", "id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies,
		},
		// verify replace node
		{
			Config: config + compartmentIdVariableStr + EsxiHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "current_sku"),
				resource.TestCheckResourceAttr(resourceName, "display_name", `replacement`),
				resource.TestCheckResourceAttrSet(resourceName, "failed_esxi_host_id"),
				resource.TestCheckResourceAttrSet(resourceName, "host_ocpu_count"),
				resource.TestCheckResourceAttrSet(resourceName, "host_shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "next_sku"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
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
		// verify singular datasource for replace node
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpOcvpEsxiHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ReplacementEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_sku"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", `replacement`),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grace_period_end_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_ocpu_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_shape_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "next_sku"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// delete replace node and upgrade SDDC before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostUpgradeResourceDependencies,
		},
		// verify upgrade node
		{
			Config: config + compartmentIdVariableStr + EsxiHostUpgradeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, upgradedEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "non_upgraded_esxi_host_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 test-L2"),

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
		// verify singular datasource for upgrade node
		{
			Config: config + compartmentIdVariableStr + EsxiHostUpgradeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, upgradedEsxiHostRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpOcvpEsxiHostSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "non_upgraded_esxi_host_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_software_version", "7.0 test-L2"),
			),
		},
		// delete upgraded node before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostUpgradeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + donorHostIdVariableStr + donorHostIdUpdateVariableStr + EsxiHostOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(resourceName, "billing_donor_host_id", donorHostId),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "THREE_YEARS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", "BM.Standard2.52"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "ONE_YEAR"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 test-L2"),
				resource.TestCheckResourceAttrSet(resourceName, "is_billing_continuation_in_progress"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + donorHostIdVariableStr + donorHostIdUpdateVariableStr + OcvpEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(resourceName, "billing_donor_host_id", donorHostIdUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "THREE_YEARS"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", "BM.Standard2.52"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "THREE_YEARS"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 test-L2"),

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
			Config: config + donorHostIdVariableStr + donorHostIdUpdateVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "test_esxi_hosts", acctest.Optional, acctest.Update, OcvpOcvpEsxiHostDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "is_billing_donors_only", "false"),
				resource.TestCheckResourceAttr(datasourceName, "is_swap_billing_only", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_shape_name", "BM.Standard2.52"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config + donorHostIdVariableStr + donorHostIdUpdateVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpOcvpEsxiHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpEsxiHostResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(singularDatasourceName, "billing_donor_host_id", donorHostIdUpdate),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_sku", "THREE_YEARS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_ocpu_count", "52"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_shape_name", "BM.Standard2.52"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_billing_continuation_in_progress"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_billing_swapping_in_progress"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_sku", "THREE_YEARS"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vmware_software_version"),
			),
		},

		// verify resource import
		{
			Config:                  config + OcvpEsxiHostRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpEsxiHostDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EsxiHostClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_esxi_host" {
			noResourceFound = false
			request := oci_ocvp.GetEsxiHostRequest{}

			tmp := rs.Primary.ID
			request.EsxiHostId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetEsxiHost(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpEsxiHost") {
		resource.AddTestSweepers("OcvpEsxiHost", &resource.Sweeper{
			Name:         "OcvpEsxiHost",
			Dependencies: acctest.DependencyGraph["esxiHost"],
			F:            sweepOcvpEsxiHostResource,
		})
	}
}

func sweepOcvpEsxiHostResource(compartment string) error {
	esxiHostClient := acctest.GetTestClients(&schema.ResourceData{}).EsxiHostClient()
	esxiHostIds, err := getOcvpEsxiHostIds(compartment)
	if err != nil {
		return err
	}
	for _, esxiHostId := range esxiHostIds {
		if ok := acctest.SweeperDefaultResourceId[esxiHostId]; !ok {
			deleteEsxiHostRequest := oci_ocvp.DeleteEsxiHostRequest{}

			deleteEsxiHostRequest.EsxiHostId = &esxiHostId

			deleteEsxiHostRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := esxiHostClient.DeleteEsxiHost(context.Background(), deleteEsxiHostRequest)
			if error != nil {
				fmt.Printf("Error deleting EsxiHost %s %s, It is possible that the resource is already deleted. Please verify manually \n", esxiHostId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &esxiHostId, OcvpEsxiHostSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpEsxiHostSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpEsxiHostIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EsxiHostId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	esxiHostClient := acctest.GetTestClients(&schema.ResourceData{}).EsxiHostClient()

	listEsxiHostsRequest := oci_ocvp.ListEsxiHostsRequest{}
	listEsxiHostsRequest.LifecycleState = oci_ocvp.ListEsxiHostsLifecycleStateActive
	listEsxiHostsResponse, err := esxiHostClient.ListEsxiHosts(context.Background(), listEsxiHostsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EsxiHost list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, esxiHost := range listEsxiHostsResponse.Items {
		id := *esxiHost.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EsxiHostId", id)
	}
	return resourceIds, nil
}

func OcvpEsxiHostSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if esxiHostResponse, ok := response.Response.(oci_ocvp.GetEsxiHostResponse); ok {
		return esxiHostResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func OcvpEsxiHostSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EsxiHostClient().GetEsxiHost(context.Background(), oci_ocvp.GetEsxiHostRequest{
		EsxiHostId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
