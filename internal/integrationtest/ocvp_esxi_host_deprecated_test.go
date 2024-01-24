// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpEsxiHostDeprecatedRequiredOnlyResource = EsxiHostDeprecatedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedRepresentation)

	OcvpEsxiHostDeprecatedResourceConfig = EsxiHostDeprecatedOptionalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Update, OcvpEsxiHostOptionalRepresentation)

	ReplacementEsxiHostDeprecatedResourceConfig = EsxiHostDeprecatedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation)

	OcvpEsxiHostDeprecatedSingularDataSourceRepresentation = map[string]interface{}{
		"esxi_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_esxi_host.test_esxi_host.id}`},
	}

	OcvpOcvpEsxiHostDeprecatedDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compute_instance_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_esxi_host.test_esxi_host.compute_instance_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: esxiName, Update: esxiUpdateName},
		"is_billing_donors_only": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_swap_billing_only":   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"sddc_id":                acctest.Representation{RepType: acctest.Optional, Create: `${local.upgraded_sddc_id}`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpEsxiHostDeprecatedDataSourceFilterRepresentation}}
	OcvpEsxiHostDeprecatedDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_esxi_host.test_esxi_host.id}`}},
	}
	OcvpDeprecatedSwapBillingOnlyEsxiHostsDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"is_swap_billing_only": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
	// for replace node
	failedEsxiHostDataSourceRepresentation = map[string]interface{}{
		"sddc_id":      acctest.Representation{RepType: acctest.Optional, Create: `${local.v6_sddc_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `cluster-1-1`},
	}
	//for upgrade ESXi host
	nonUpgradedEsxiHostDataSourceRepresentation = map[string]interface{}{
		"sddc_id":      acctest.Representation{RepType: acctest.Optional, Create: `${local.upgraded_sddc_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `cluster-1-1`},
	}

	OcvpEsxiHostDeprecatedRepresentation = map[string]interface{}{
		"sddc_id":                     acctest.Representation{RepType: acctest.Required, Create: `${local.v6_sddc_id}`},
		"billing_donor_host_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.donor_host_id}`, Update: `${var.donor_host_id_update}`},
		"capacity_reservation_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_capacity_reservation.test_esxi_host_compute_capacity_reservation.id}`},
		"compute_availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}`},
		"current_sku":                 acctest.Representation{RepType: acctest.Optional, Create: `MONTH`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: esxiName, Update: esxiUpdateName},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"next_sku":                    acctest.Representation{RepType: acctest.Optional, Create: `ONE_YEAR`, Update: `MONTH`},
		"host_ocpu_count":             acctest.Representation{RepType: acctest.Optional, Create: esxiOcpuCount},
		"host_shape_name":             acctest.Representation{RepType: acctest.Optional, Create: esxiShapeName},
	}
	OcvpEsxiHostOptionalRepresentation = acctest.RepresentationCopyWithNewProperties(OcvpEsxiHostDeprecatedRepresentation, map[string]interface{}{
		"sddc_id": acctest.Representation{RepType: acctest.Required, Create: `${local.upgraded_sddc_id}`},
	})
	replacementEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":             acctest.Representation{RepType: acctest.Required, Create: `${local.v6_sddc_id}`},
		"failed_esxi_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_esxi_hosts.failed_esxi_hosts.esxi_host_collection[0].id}`},
	}
	upgradedEsxiHostRepresentation = map[string]interface{}{
		"sddc_id":                   acctest.Representation{RepType: acctest.Required, Create: `${local.upgraded_sddc_id}`},
		"non_upgraded_esxi_host_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_ocvp_esxi_hosts.non_upgraded_esxi_hosts.esxi_host_collection[0].id}`},
	}

	donorHostsDataSourceRepresentation = map[string]interface{}{
		"is_billing_donors_only": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"compartment_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"filter": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: map[string]interface{}{
				"name":   acctest.Representation{RepType: acctest.Required, Create: `host_shape_name`},
				"values": acctest.Representation{RepType: acctest.Required, Create: []string{esxiShapeName}},
			}},
			{RepType: acctest.Required, Group: map[string]interface{}{
				"name":   acctest.Representation{RepType: acctest.Required, Create: `current_sku`},
				"values": acctest.Representation{RepType: acctest.Required, Create: []string{`MONTH`}},
			}},
		},
	}

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
	sddcDataSourceDependencies = `
locals {
  v6_sddc_id = data.oci_ocvp_sddcs.v6_sddcs.sddc_collection[0].id
  upgraded_sddc_id = data.oci_ocvp_sddcs.upgraded_sddcs.sddc_collection[0].id
  v6_sddc_name = "esxi-test-sddc"
  upgraded_sddc_name = "sddc-upgraded"
}

data "oci_ocvp_sddcs" "v6_sddcs" {
  compartment_id = var.compartment_id
  display_name = local.v6_sddc_name
  state = "ACTIVE"
}

data "oci_ocvp_sddcs" "upgraded_sddcs" {
  compartment_id = var.compartment_id
  display_name = local.upgraded_sddc_name
  state = "ACTIVE"
}
`

	ocvpAvailabilityDomainDependency = `
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.compartment_id}"
}
`

	EsxiHostDeprecatedResourceDependencies = sddcDataSourceDependencies + acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "failed_esxi_hosts", acctest.Optional, acctest.Create, failedEsxiHostDataSourceRepresentation)

	EsxiHostDeprecatedUpgradeResourceDependencies  = sddcDataSourceDependencies + acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "non_upgraded_esxi_hosts", acctest.Optional, acctest.Create, nonUpgradedEsxiHostDataSourceRepresentation)
	EsxiHostDeprecatedOptionalResourceDependencies = sddcDataSourceDependencies + ocvpAvailabilityDomainDependency + ocvpEsxiHostCapacityReservationResource
)

// issue-routing-tag: ocvp/default
func TestOcvpEsxiHostDeprecatedResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpEsxiHostDeprecatedResource_basic")
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
					acctest.SaveConfigContent(config+compartmentIdVariableStr+EsxiHostDeprecatedUpgradeResourceDependencies+donorHostIdVariableStr+donorHostIdUpdateVariableStr+EsxiHostDeprecatedOptionalResourceDependencies+
						acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostDeprecatedRepresentation), "ocvp", "esxiHost", t)

					return err
				},
			),
		},
	})

	acctest.ResourceTest(t, testAccCheckOcvpEsxiHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedRepresentation),
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
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "swap_billing_esxi_hosts", acctest.Optional, acctest.Update, OcvpDeprecatedSwapBillingOnlyEsxiHostsDataSourceRepresentation),
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
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedResourceDependencies,
		},
		// verify replace node
		{
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, replacementEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(resourceName, "current_sku"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ReplacementEsxiHostDeprecatedResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_sku"),
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
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedUpgradeResourceDependencies,
		},
		// verify upgrade node
		{
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedUpgradeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, upgradedEsxiHostRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "non_upgraded_esxi_host_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),

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
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedUpgradeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, upgradedEsxiHostRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "non_upgraded_esxi_host_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_software_version", noInstanceVmwareVersionV7),
			),
		},
		// delete upgraded node before next Create
		{
			Config: config + compartmentIdVariableStr + EsxiHostDeprecatedUpgradeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + donorHostIdVariableStr + donorHostIdUpdateVariableStr + EsxiHostDeprecatedOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Optional, acctest.Create, OcvpEsxiHostOptionalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(resourceName, "billing_donor_host_id", donorHostId),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "display_name", esxiName),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", esxiOcpuCount),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", esxiShapeName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "ONE_YEAR"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),
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
			Config: config + compartmentIdVariableStr + donorHostIdVariableStr + donorHostIdUpdateVariableStr + OcvpEsxiHostDeprecatedResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(resourceName, "billing_donor_host_id", donorHostIdUpdate),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(resourceName, "display_name", esxiUpdateName),
				resource.TestCheckResourceAttrSet(resourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(resourceName, "host_ocpu_count", esxiOcpuCount),
				resource.TestCheckResourceAttr(resourceName, "host_shape_name", esxiShapeName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "next_sku", "MONTH"),
				resource.TestCheckResourceAttrSet(resourceName, "sddc_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", noInstanceVmwareVersionV7),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_hosts", "test_esxi_hosts", acctest.Optional, acctest.Update, OcvpOcvpEsxiHostDeprecatedDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpEsxiHostDeprecatedResourceConfig,
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
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.display_name", esxiUpdateName),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.sddc_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_ocpu_count", esxiOcpuCount),
				resource.TestCheckResourceAttr(datasourceName, "esxi_host_collection.0.host_shape_name", esxiShapeName),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.compute_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "esxi_host_collection.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config + donorHostIdVariableStr + donorHostIdUpdateVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_esxi_host", "test_esxi_host", acctest.Required, acctest.Create, OcvpEsxiHostDeprecatedSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpEsxiHostDeprecatedResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "esxi_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "billing_contract_end_date"),
				resource.TestCheckResourceAttr(singularDatasourceName, "billing_donor_host_id", donorHostIdUpdate),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "current_sku", "MONTH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", esxiUpdateName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "freeform_tags.%"),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_ocpu_count", esxiOcpuCount),
				resource.TestCheckResourceAttr(singularDatasourceName, "host_shape_name", esxiShapeName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_billing_continuation_in_progress"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_billing_swapping_in_progress"),
				resource.TestCheckResourceAttr(singularDatasourceName, "next_sku", "MONTH"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vmware_software_version"),
			),
		},

		// verify resource import
		{
			Config:                  config + OcvpEsxiHostDeprecatedRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OcvpEsxiHostDeprecated") {
		resource.AddTestSweepers("OcvpEsxiHostDeprecated", &resource.Sweeper{
			Name:         "OcvpEsxiHostDeprecated",
			Dependencies: acctest.DependencyGraph["esxiHost"],
			F:            sweepOcvpEsxiHostResource,
		})
	}
}
