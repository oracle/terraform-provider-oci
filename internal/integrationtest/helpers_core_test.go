// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

const (
	shape           = "VM.Standard2.1"
	cidrBlockSubnet = "10.0.0.0/24"
	cidrBlockVcn    = "10.0.0.0/16"
)

const (
	instanceDnsConfig = `
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "t" {
	cidr_block      = "10.0.0.0/16"
	compartment_id  = "${var.compartment_id}"
	display_name    = "-tf-vcn"
	dns_label		= "testvcn"
}

resource "oci_core_subnet" "t" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.1.0/24"
  display_name        = "-tf-subnet"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.t.id}"
  route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
  security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
  dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
  dns_label			  = "testsubnet"
}

variable "InstanceImageOCID" {
  type = "map"
  default = {
    // Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
  }
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${var.InstanceImageOCID[var.region]}"
	shape = "VM.Standard2.1"
	create_vnic_details {
        subnet_id = "${oci_core_subnet.t.id}"
        hostname_label = "testinstance"
        display_name = "-tf-instance-vnic"
		defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
		freeform_tags = { "Department" = "Accounting" }
  	}
	metadata = {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}
	timeouts {
		create = "15m"
	}
}
resource "oci_core_network_security_group" "test_network_security_group1" {
	compartment_id = "${var.compartment_id}"
	vcn_id         = "${oci_core_virtual_network.t.id}"
	display_name = "testNetworkSecurityGroup1"
}
resource "oci_core_network_security_group" "test_network_security_group2" {
	compartment_id = "${var.compartment_id}"
	vcn_id         = "${oci_core_virtual_network.t.id}"
	display_name = "testNetworkSecurityGroup2"
}` + DefinedTagsDependencies
)

var subnetRegionalRepresentation = map[string]interface{}{
	"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
	"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
	"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
	"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
	"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
	"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
	"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
}

var ignoreDefinedTagsChangesRep = map[string]interface{}{
	"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
}

var imageIdMap = map[string]string{
	"us-phoenix-1":   "ocid1.image.oc1.phx.aaaaaaaadjnj3da72bztpxinmqpih62c2woscbp6l3wjn36by2cvmdhjub6a",
	"us-ashburn-1":   "ocid1.image.oc1.iad.aaaaaaaawufnve5jxze4xf7orejupw5iq3pms6cuadzjc7klojix6vmk42va",
	"eu-frankfurt-1": "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaagbrvhganmn7awcr7plaaf5vhabmzhx763z5afiitswjwmzh7upna",
	"uk-london-1":    "ocid1.image.oc1.uk-london-1.aaaaaaaajwtut4l7fo3cvyraate6erdkyf2wdk5vpk6fp6ycng3dv2y3ymvq",
}

var subnetConfig = `
data "oci_identity_availability_domains" "ADs" {
	compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "t" {
	cidr_block = "10.0.0.0/16"
	compartment_id = "${var.compartment_id}"
	display_name = "network_name"
}

resource "oci_core_subnet" "WebSubnetAD1" {
	availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
	cidr_block          = "10.0.1.0/24"
	display_name        = "WebSubnetAD1"
	compartment_id      = "${var.compartment_id}"
	vcn_id              = "${oci_core_virtual_network.t.id}"
	route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
	security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
	dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
}`

var instanceConfig = subnetConfig + `
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

data "oci_identity_policies" "policies" {
	compartment_id = "${var.compartment_id}"
}

data "oci_load_balancer_protocols" "protocols" {
	compartment_id = "${var.compartment_id}"
}

data "oci_core_shape" "shapes" {
	compartment_id = "${var.compartment_id}"
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	image_id =  "${var.InstanceImageOCID[var.region]}"
}

resource "oci_core_instance" "t" {
	availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
	compartment_id = "${var.compartment_id}"
	display_name = "-tf-instance"
	image = "${var.InstanceImageOCID[var.region]}"
	shape = "VM.Standard2.1"
	subnet_id = "${oci_core_subnet.WebSubnetAD1.id}"
	metadata = {
		ssh_authorized_keys = "${var.ssh_public_key}"
	}

	timeouts {
		create = "15m"
	}
}
`

func createVolumeInRegion(clients *tf_client.OracleClients, region string) (string, error) {
	compartment := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.IdentityClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&identityClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)
	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name

	createVolumeResponse, err := blockStorageClient.CreateVolume(context.Background(), oci_core.CreateVolumeRequest{
		CreateVolumeDetails: oci_core.CreateVolumeDetails{
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "core"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source Volume with the error %v", err)
	}
	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), volumeAvailableWaitCondition, "core", false)

	_, err = blockStorageClient.GetVolume(context.Background(), oci_core.GetVolumeRequest{
		VolumeId: createVolumeResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeAvailableWaitCondition failed for %s resource with error %v", *createVolumeResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createVolumeResponse.Id)
	}

	return *createVolumeResponse.Id, nil
}

func createVolumeGroupInRegion(clients *tf_client.OracleClients, region string, volumeId *string) (string, error) {
	compartment := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.IdentityClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&identityClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)
	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name
	sourceDetails := map[string]interface{}{
		"type":      "volumeIds",
		"volumeIds": [1]*string{volumeId},
	}

	createVolumeGroupResponse, err := blockStorageClient.CreateVolumeGroup(context.Background(), oci_core.CreateVolumeGroupRequest{
		CreateVolumeGroupDetails: oci_core.CreateVolumeGroupDetails{
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
			SourceDetails:      sourceDetails,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "core"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source Volume Group with the error %v", err)
	}
	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), volumeGroupAvailableWaitCondition, "core", false)

	_, err = blockStorageClient.GetVolumeGroup(context.Background(), oci_core.GetVolumeGroupRequest{
		VolumeGroupId: createVolumeGroupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeGroupAvailableWaitCondition failed for %s resource with error %v", *createVolumeGroupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createVolumeGroupResponse.Id)
	}

	return *createVolumeGroupResponse.Id, nil
}

func createVolumeBackupInRegion(clients *tf_client.OracleClients, region string, volumeId *string) (string, error) {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	createVolumeBackupResponse, err := blockStorageClient.CreateVolumeBackup(context.Background(), oci_core.CreateVolumeBackupRequest{
		CreateVolumeBackupDetails: oci_core.CreateVolumeBackupDetails{
			VolumeId: volumeId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source VolumeBackup with the error %v", err)

	}

	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), volumeBackupAvailableWaitCondition, "core", false)
	_, err = blockStorageClient.GetVolumeBackup(context.Background(), oci_core.GetVolumeBackupRequest{
		VolumeBackupId: createVolumeBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeBackupAvailableWaitCondition failed for %s resource with error %v", *createVolumeBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createVolumeBackupResponse.Id)
	}
	return *createVolumeBackupResponse.Id, nil

}

func createVolumeGroupBackupInRegion(clients *tf_client.OracleClients, region string, volumeGroupId *string) (string, error) {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	createVolumeGroupBackupResponse, err := blockStorageClient.CreateVolumeGroupBackup(context.Background(), oci_core.CreateVolumeGroupBackupRequest{
		CreateVolumeGroupBackupDetails: oci_core.CreateVolumeGroupBackupDetails{
			VolumeGroupId: volumeGroupId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source VolumeGroupBackup with the error %v", err)

	}

	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), volumeGroupBackupAvailableWaitCondition, "core", false)
	_, err = blockStorageClient.GetVolumeGroupBackup(context.Background(), oci_core.GetVolumeGroupBackupRequest{
		VolumeGroupBackupId: createVolumeGroupBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for volumeGroupBackupAvailableWaitCondition failed for %s resource with error %v", *createVolumeGroupBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createVolumeGroupBackupResponse.Id)
	}
	return *createVolumeGroupBackupResponse.Id, nil

}

func deleteVolumeInRegion(clients *tf_client.OracleClients, region string, volumeId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeId != "" {
		deleteVolumeRequest := oci_core.DeleteVolumeRequest{}
		deleteVolumeRequest.VolumeId = &volumeId

		_, err := blockStorageClient.DeleteVolume(context.Background(), deleteVolumeRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source Volume %s resource with error %v", *deleteVolumeRequest.VolumeId, err)
		}
	}

	return nil
}

func deleteVolumeGroupInRegion(clients *tf_client.OracleClients, region string, volumeGroupId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeGroupId != "" {
		deleteVolumeGroupRequest := oci_core.DeleteVolumeGroupRequest{}
		deleteVolumeGroupRequest.VolumeGroupId = &volumeGroupId

		_, err := blockStorageClient.DeleteVolumeGroup(context.Background(), deleteVolumeGroupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source VolumeGroup %s resource with error %v", *deleteVolumeGroupRequest.VolumeGroupId, err)
		}
	}

	return nil
}

func deleteVolumeBackupInRegion(clients *tf_client.OracleClients, region string, volumeBackupId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeBackupId != "" {
		deleteVolumeBackupRequest := oci_core.DeleteVolumeBackupRequest{}
		deleteVolumeBackupRequest.VolumeBackupId = &volumeBackupId

		_, err := blockStorageClient.DeleteVolumeBackup(context.Background(), deleteVolumeBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source VolumeBackup %s resource with error %v", *deleteVolumeBackupRequest.VolumeBackupId, err)
		}
	}

	return nil
}

func deleteVolumeGroupBackupInRegion(clients *tf_client.OracleClients, region string, volumeGroupBackupId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if volumeGroupBackupId != "" {
		deleteVolumeGroupBackupRequest := oci_core.DeleteVolumeGroupBackupRequest{}
		deleteVolumeGroupBackupRequest.VolumeGroupBackupId = &volumeGroupBackupId

		_, err := blockStorageClient.DeleteVolumeGroupBackup(context.Background(), deleteVolumeGroupBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source VolumeGroupBackup %s resource with error %v", *deleteVolumeGroupBackupRequest.VolumeGroupBackupId, err)
		}
	}

	return nil
}

func volumeAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeResponse, ok := response.Response.(oci_core.GetVolumeResponse); ok {
		return volumeResponse.LifecycleState != oci_core.VolumeLifecycleStateAvailable
	}

	return false
}

func volumeGroupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeGroupResponse, ok := response.Response.(oci_core.GetVolumeGroupResponse); ok {
		return volumeGroupResponse.LifecycleState != oci_core.VolumeGroupLifecycleStateAvailable
	}

	return false
}

func volumeBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeBackupResponse, ok := response.Response.(oci_core.GetVolumeBackupResponse); ok {
		return volumeBackupResponse.LifecycleState != oci_core.VolumeBackupLifecycleStateAvailable
	}

	return false
}

func volumeGroupBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if volumeGroupBackupResponse, ok := response.Response.(oci_core.GetVolumeGroupBackupResponse); ok {
		return volumeGroupBackupResponse.LifecycleState != oci_core.VolumeGroupBackupLifecycleStateAvailable
	}

	return false
}

func createBootVolumeInRegion(clients *tf_client.OracleClients, region string) (string, string, error) {
	compartment := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.IdentityClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&identityClient.BaseClient)
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)
	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name

	// Create subnet
	networkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(*clients.VirtualNetworkClient().ConfigurationProvider())

	if err != nil {
		return "", "", fmt.Errorf("[WARN] cannot configure client for the source region %v", err)
	}

	cidrBlockVcn := cidrBlockVcn
	networkClient.SetRegion(region)
	createVcnResponse, err := networkClient.CreateVcn(context.Background(), oci_core.CreateVcnRequest{
		CreateVcnDetails: oci_core.CreateVcnDetails{
			CidrBlock:     &cidrBlockVcn,
			CompartmentId: &compartment,
		}})

	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to Create source VCN with the error %v", err)
	}

	cidrBlockSubnet := cidrBlockSubnet

	createSubnetResponse, err := networkClient.CreateSubnet(context.Background(), oci_core.CreateSubnetRequest{
		CreateSubnetDetails: oci_core.CreateSubnetDetails{
			CompartmentId: &compartment,
			CidrBlock:     &cidrBlockSubnet,
			VcnId:         createVcnResponse.Id,
		},
	})

	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(*clients.ComputeClient().ConfigurationProvider())
	if err != nil {
		return "", "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	shape := shape
	computeClient.SetRegion(region)
	imageId := imageIdMap[region]
	createInstanceResponse, err := computeClient.LaunchInstance(context.Background(), oci_core.LaunchInstanceRequest{
		LaunchInstanceDetails: oci_core.LaunchInstanceDetails{
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
			Shape:              &shape,
			SubnetId:           createSubnetResponse.Id,
			ImageId:            &imageId,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "core"),
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to Create source Instance with the error %v", err)
	}
	instanceId := createInstanceResponse.Id

	retryPolicyInstance := tfresource.GetRetryPolicy(false, "core")
	retryPolicyInstance.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), instanceAvailableWaitCondition, "core", false)

	_, err = computeClient.GetInstance(context.Background(), oci_core.GetInstanceRequest{
		InstanceId: instanceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicyInstance,
		},
	})

	listBootVolumeReq, err := computeClient.ListBootVolumeAttachments(context.Background(), oci_core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: domain,
		CompartmentId:      &compartment,
		InstanceId:         instanceId,
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to ListBootVolumeAttachments with the error %v", err)
	}

	bootVolumeId := listBootVolumeReq.Items[0].BootVolumeId

	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), bootVolumeAvailableWaitCondition, "core", false)

	_, err = blockStorageClient.GetBootVolume(context.Background(), oci_core.GetBootVolumeRequest{
		BootVolumeId: bootVolumeId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] wait for bootVolumeAvailableWaitCondition failed for %s resource with error %v", *bootVolumeId, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *bootVolumeId)
	}

	return *instanceId, *bootVolumeId, nil
}

func createBootVolumeBackupInRegion(clients *tf_client.OracleClients, region string, bootVolumeId *string) (string, error) {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	createBootVolumeBackupResponse, err := blockStorageClient.CreateBootVolumeBackup(context.Background(), oci_core.CreateBootVolumeBackupRequest{
		CreateBootVolumeBackupDetails: oci_core.CreateBootVolumeBackupDetails{
			BootVolumeId: bootVolumeId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source BootVolumeBackup with the error %v", err)

	}

	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), bootVolumeBackupAvailableWaitCondition, "core", false)
	_, err = blockStorageClient.GetBootVolumeBackup(context.Background(), oci_core.GetBootVolumeBackupRequest{
		BootVolumeBackupId: createBootVolumeBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for bootVolumeBackupAvailableWaitCondition failed for %s resource with error %v", *createBootVolumeBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createBootVolumeBackupResponse.Id)
	}
	return *createBootVolumeBackupResponse.Id, nil

}

func deleteBootVolumeInRegion(clients *tf_client.OracleClients, region string, bootVolumeId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if bootVolumeId != "" {
		deleteBootVolumeRequest := oci_core.DeleteBootVolumeRequest{}
		deleteBootVolumeRequest.BootVolumeId = &bootVolumeId

		_, err := blockStorageClient.DeleteBootVolume(context.Background(), deleteBootVolumeRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source BootVolume %s resource with error %v", *deleteBootVolumeRequest.BootVolumeId, err)
		}
	}

	return nil
}

func deleteBootVolumeBackupInRegion(clients *tf_client.OracleClients, region string, bootVolumeBackupId string) error {
	blockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*clients.BlockstorageClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&blockStorageClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	blockStorageClient.SetRegion(region)

	if bootVolumeBackupId != "" {
		deleteBootVolumeBackupRequest := oci_core.DeleteBootVolumeBackupRequest{}
		deleteBootVolumeBackupRequest.BootVolumeBackupId = &bootVolumeBackupId

		_, err := blockStorageClient.DeleteBootVolumeBackup(context.Background(), deleteBootVolumeBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source BootVolumeBackup %s resource with error %v", *deleteBootVolumeBackupRequest.BootVolumeBackupId, err)
		}
	}

	return nil
}

func terminateInstanceInRegion(clients *tf_client.OracleClients, region string, instanceId string) error {
	computeClient, err := oci_core.NewComputeClientWithConfigurationProvider(*clients.ComputeClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&computeClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the source region: %v", err)
	}

	computeClient.SetRegion(region)

	if instanceId != "" {
		terminateInstanceRequest := oci_core.TerminateInstanceRequest{}
		terminateInstanceRequest.InstanceId = &instanceId

		_, err := computeClient.TerminateInstance(context.Background(), terminateInstanceRequest)
		if err != nil {
			return fmt.Errorf("failed to terminate instance %s resource with error %v", *terminateInstanceRequest.InstanceId, err)
		}
	}

	return nil
}

func bootVolumeAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if bootVolumeResponse, ok := response.Response.(oci_core.GetBootVolumeResponse); ok {
		return bootVolumeResponse.LifecycleState != oci_core.BootVolumeLifecycleStateAvailable
	}

	return false
}

func instanceAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if instanceResponse, ok := response.Response.(oci_core.GetInstanceResponse); ok {
		return instanceResponse.LifecycleState != oci_core.InstanceLifecycleStateRunning
	}

	return false
}

func bootVolumeBackupAvailableWaitCondition(response oci_common.OCIOperationResponse) bool {
	if bootVolumeBackupResponse, ok := response.Response.(oci_core.GetBootVolumeBackupResponse); ok {
		return bootVolumeBackupResponse.LifecycleState != oci_core.BootVolumeBackupLifecycleStateAvailable
	}

	return false
}
