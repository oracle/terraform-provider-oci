// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

const (
	displayName  = "-tf-public-ip"
	displayName2 = displayName + "-updated"
	privateIpId  = "private_ips.0.id"
	privateIpId2 = "private_ips.1.id"
)

var (
	PublicIpRequiredOnlyResource = PublicIpResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Required, acctest.Create, publicIpRepresentation)

	PublicIpResourceConfig = PublicIpResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Update, publicIpRepresentation)

	publicIpSingularDataSourceRepresentation = map[string]interface{}{
		"private_ip_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_public_ip.test_public_ip2.private_ip_id}`},
	}

	publicIpDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: string(oci_core.PublicIpScopeRegion), Update: string(oci_core.PublicIpScopeAvailabilityDomain)},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"public_ip_pool_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: publicIpDataSourceFilterRepresentation}}
	publicIpDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_public_ip.test_public_ip2.id}`}, Update: []string{string(oci_core.PublicIpLifetimeReserved)}},
	}
	publicIpDataSourceNameFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{string(oci_core.PublicIpLifetimeReserved)}},
	}

	publicIpRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"lifetime":          acctest.Representation{RepType: acctest.Required, Create: `RESERVED`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `-tf-public-ip`, Update: `-tf-public-ip-updated`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"private_ip_id":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_core_private_ips.test_private_ips.` + privateIpId + `}`, Update: `${data.oci_core_private_ips.test_private_ips.` + privateIpId2 + `}`},
		"public_ip_pool_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
		"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesRepresentation},
	}

	ignoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `private_ip_id`}},
	}

	publicUnassignedIpRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"lifetime":          acctest.Representation{RepType: acctest.Required, Create: `RESERVED`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `-tf-public-ip`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"public_ip_pool_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_public_ip_pool.test_public_ip_pool.id}`},
	}

	PublicIpResourceDependencies = DefinedTagsDependencies + byoipRangeIdVariableStr + publicIpPoolCidrBlockVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool_capacity", "test_public_ip_pool_capacity", acctest.Required, acctest.Create, publicIpPoolCapacityRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip_pool", "test_public_ip_pool", acctest.Required, acctest.Create, publicIpPoolRepresentation) + `
	variable "InstanceImageOCID" {
		type = "map"
		default = {
			// See https://docs.us-phoenix-1.oraclecloud.com/images/
			// Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
			us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a"
			us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va"
			eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna"
			uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq"
		}
	}
	data "oci_identity_availability_domains" "test_availability_domains" {
		compartment_id = "${var.compartment_id}"
	}
	resource "oci_core_virtual_network" "test_vcn" {
		cidr_block      = "10.0.0.0/16"
		compartment_id  = "${var.compartment_id}"
		display_name    = "-tf-vcn"
		dns_label       = "testvcn"
	}
	resource "oci_core_subnet" "test_subnet" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.test_vcn.id}"
		route_table_id      = "${oci_core_virtual_network.test_vcn.default_route_table_id}"
		security_list_ids   = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.test_vcn.default_dhcp_options_id}"
		dns_label           = "testsubnet"
	}
	resource "oci_core_instance" "test_instance" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		compartment_id      = "${var.compartment_id}"
		display_name        = "-tf-instance"
		image               = "${var.InstanceImageOCID[var.region]}"
		shape               = "VM.Standard2.1"
		create_vnic_details {
			assign_public_ip = false
			subnet_id        = "${oci_core_subnet.test_subnet.id}"
			hostname_label   = "testinstance"
			display_name     = "-tf-instance-vnic"
		}
		metadata = {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}
	resource "oci_core_vnic_attachment" "test_vnic_attachments" {
		instance_id  = "${oci_core_instance.test_instance.id}"
		display_name = "-tf-vnic-attachment-2"
		create_vnic_details {
			assign_public_ip = false
			subnet_id        = "${oci_core_subnet.test_subnet.id}"
			display_name     = "-tf-vnic-2"
			hostname_label   = "testinstance2"
		}
	}
	data "oci_core_private_ips" "test_private_ips" {
		subnet_id        = "${oci_core_subnet.test_subnet.id}"
	}`
)

// issue-routing-tag: core/virtualNetwork
func TestCorePublicIpResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCorePublicIpResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_public_ip.test_public_ip"
	datasourceName := "data.oci_core_public_ips.test_public_ips"
	sDatasourceNameById := "data.oci_core_public_ip.test_oci_core_public_ip_by_id"
	sDatasourceNameByIp := "data.oci_core_public_ip.test_oci_core_public_ip_by_ip"
	sDatasourceNameByPrivateIpId := "data.oci_core_public_ip.test_oci_core_public_ip_by_private_ip_id"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PublicIpResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Create, publicIpRepresentation), "core", "publicIp", t)

	acctest.ResourceTest(t, testAccCheckCorePublicIpDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Required, acctest.Create, publicIpRepresentation) + `
					data "oci_core_public_ip" "test_oci_core_public_ip_by_id" {
						id = "${oci_core_public_ip.test_public_ip.id}"
					}

					data "oci_core_public_ip" "test_oci_core_public_ip_by_ip" {
						ip_address = "${oci_core_public_ip.test_public_ip.ip_address}"
					}`,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckNoResourceAttr(resourceName, "private_ip_id"),
				resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},

				// check oci_core_public_ip by id
				resource.TestCheckResourceAttr(sDatasourceNameById, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(sDatasourceNameById, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckResourceAttrSet(sDatasourceNameById, "id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameById, "ip_address"),
				resource.TestCheckNoResourceAttr(sDatasourceNameById, "private_ip_id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameById, "display_name"),
				resource.TestCheckNoResourceAttr(sDatasourceNameById, "availability_domain"),
				resource.TestCheckResourceAttrSet(sDatasourceNameById, "time_created"),
				resource.TestCheckResourceAttr(sDatasourceNameById, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(sDatasourceNameById, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

				// check oci_core_public_ip by public ip
				resource.TestCheckResourceAttr(sDatasourceNameByIp, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(sDatasourceNameByIp, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "ip_address"),
				resource.TestCheckNoResourceAttr(sDatasourceNameByIp, "private_ip_id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "display_name"),
				resource.TestCheckNoResourceAttr(sDatasourceNameByIp, "availability_domain"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "time_created"),
				resource.TestCheckResourceAttr(sDatasourceNameByIp, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(sDatasourceNameByIp, "state", string(oci_core.PublicIpLifecycleStateAvailable)),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Create, publicIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-public-ip"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_pool_id"),
				acctest.TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId),
				resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(publicIpRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-public-ip"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_pool_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters (partial Update)
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Create,
					acctest.GetUpdatedRepresentationCopy("display_name", acctest.Representation{RepType: acctest.Optional, Create: displayName2}, publicIpRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName2),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				acctest.TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId),
				resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters (full Update)
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Update, publicIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-public-ip-updated"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
				resource.TestCheckResourceAttrSet(resourceName, "public_ip_pool_id"),
				acctest.TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId2),
				resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters (unassign private ip id)
		{
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Optional, acctest.Update, publicUnassignedIpRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckResourceAttr(resourceName, "private_ip_id", ""), // Still defined, but now empty.
				resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
				resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

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
			Config: config + compartmentIdVariableStr + PublicIpResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip2", acctest.Optional, acctest.Update, publicIpRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_public_ips", "test_public_ips", acctest.Required, acctest.Create, publicIpDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_public_ip", "test_oci_core_public_ip_by_private_ip_id", acctest.Optional, acctest.Update, publicIpSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.assigned_entity_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.assigned_entity_type"),
				resource.TestCheckResourceAttr(datasourceName, "public_ips.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "public_ips.0.display_name", "-tf-public-ip-updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.ip_address"),
				resource.TestCheckResourceAttr(datasourceName, "public_ips.0.lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.private_ip_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.public_ip_pool_id"),

				// check oci_core_public_ip by private ip id
				resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
				resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "ip_address"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "private_ip_id"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "display_name"),
				resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "freeform_tags.%", "1"),
				resource.TestCheckNoResourceAttr(sDatasourceNameByPrivateIpId, "availability_domain"),
				resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "time_created"),
				resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "scope", string(oci_core.PublicIpScopeRegion)),
				resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "state", string(oci_core.PublicIpLifecycleStateAssigned)),
			),
		},
		// Test client-side filtering.
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_public_ips", "test_public_ips", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("filter", acctest.RepresentationGroup{RepType: acctest.Required, Group: publicIpDataSourceNameFilterRepresentation}, publicIpDataSourceRepresentation)) +
				compartmentIdVariableStr + PublicIpResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
				resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "0"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCorePublicIpDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_public_ip" {
			noResourceFound = false
			request := oci_core.GetPublicIpRequest{}

			tmp := rs.Primary.ID
			request.PublicIpId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetPublicIp(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.PublicIpLifecycleStateUnassigned): true, string(oci_core.PublicIpLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CorePublicIp") {
		resource.AddTestSweepers("CorePublicIp", &resource.Sweeper{
			Name:         "CorePublicIp",
			Dependencies: acctest.DependencyGraph["publicIp"],
			F:            sweepCorePublicIpResource,
		})
	}
}

func sweepCorePublicIpResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	publicIpIds, err := getPublicIpIds(compartment)
	if err != nil {
		return err
	}
	for _, publicIpId := range publicIpIds {
		if ok := acctest.SweeperDefaultResourceId[publicIpId]; !ok {
			deletePublicIpRequest := oci_core.DeletePublicIpRequest{}

			deletePublicIpRequest.PublicIpId = &publicIpId

			deletePublicIpRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeletePublicIp(context.Background(), deletePublicIpRequest)
			if error != nil {
				fmt.Printf("Error deleting PublicIp %s %s, It is possible that the resource is already deleted. Please verify manually \n", publicIpId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &publicIpId, publicIpSweepWaitCondition, time.Duration(3*time.Minute),
				publicIpSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getPublicIpIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PublicIpId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	var publicIps []oci_core.PublicIp
	var error error
	compartmentId := compartment

	listPublicIpsRequest := oci_core.ListPublicIpsRequest{}
	listPublicIpsRequest.CompartmentId = &compartmentId

	scopes := oci_core.GetListPublicIpsScopeEnumValues()
	for _, scope := range scopes {
		listPublicIpsRequest.Scope = scope
		if scope == oci_core.ListPublicIpsScopeRegion {
			publicIps, error = getPublicIpIdsForRegionScope(compartmentId, listPublicIpsRequest)
			if error != nil {
				return resourceIds, error
			}
		}
		if scope == oci_core.ListPublicIpsScopeAvailabilityDomain {
			publicIps, error = getPublicIpIdsForADScope(compartmentId, listPublicIpsRequest)
			if error != nil {
				return resourceIds, error
			}
		}
		for _, publicIp := range publicIps {
			id := *publicIp.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PublicIpId", id)
		}

	}
	return resourceIds, nil
}

func publicIpSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if publicIpResponse, ok := response.Response.(oci_core.GetPublicIpResponse); ok {
		return publicIpResponse.LifecycleState != oci_core.PublicIpLifecycleStateTerminated
	}
	return false
}

func publicIpSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetPublicIp(context.Background(), oci_core.GetPublicIpRequest{
		PublicIpId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
