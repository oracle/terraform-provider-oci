// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/lb"
)

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"tenancy_ocid": "(Required) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Bare Metal console.",
		"user_ocid":    "(Required) The user OCID. This can be found in user settings in the Bare Metal console.",
		"fingerprint":  "(Required) The fingerprint for the user's RSA key. This can be found in user settings in the Bare Metal console.",
		"private_key": "(Optional) A PEM formatted RSA private key for the user.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_path": "(Optional) The path to the user's PEM formatted private key.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_password": "(Optional) The password used to secure the private key.",
	}
}

// Provider is the adapter for terraform, that gives access to all the resources
func Provider(configfn schema.ConfigureFunc) terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: dataSourcesMap(),
		Schema:         schemaMap(),
		ResourcesMap:   resourcesMap(),
		ConfigureFunc:  configfn,
	}
}

func schemaMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tenancy_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["tenancy_ocid"],
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["user_ocid"],
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["fingerprint"],
		},
		// Mostly used for testing. Don't put keys in your .tf files
		"private_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions["private_key"],
		},
		"private_key_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["private_key_path"],
		},
		"private_key_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions["private_key_password"],
		},
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history_data":       ConsoleHistoryDataDatasource(),
		"baremetal_core_cpes":                       CpeDatasource(),
		"baremetal_core_dhcp_options":               DHCPOptionsDatasource(),
		"baremetal_core_drg_attachments":            DrgAttachmentDatasource(),
		"baremetal_core_drgs":                       DrgDatasource(),
		"baremetal_core_images":                     ImageDatasource(),
		"baremetal_core_instance_credentials":       InstanceCredentialsDatasource(),
		"baremetal_core_instances":                  InstanceDatasource(),
		"baremetal_core_internet_gateways":          InternetGatewayDatasource(),
		"baremetal_core_ipsec_config":               IPSecConnectionConfigDatasource(),
		"baremetal_core_ipsec_connections":          IPSecConnectionsDatasource(),
		"baremetal_core_ipsec_status":               IPSecConnectionStatusDatasource(),
		"baremetal_core_route_tables":               RouteTableDatasource(),
		"baremetal_core_security_lists":             SecurityListDatasource(),
		"baremetal_core_shape":                      ShapeDatasource(),
		"baremetal_core_subnets":                    SubnetDatasource(),
		"baremetal_core_virtual_networks":           VirtualNetworkDatasource(),
		"baremetal_core_vnic":                       VnicDatasource(),
		"baremetal_core_vnic_attachments":           DatasourceCoreVnicAttachments(),
		"baremetal_core_volume_attachments":         VolumeAttachmentDatasource(),
		"baremetal_core_volume_backups":             VolumeBackupDatasource(),
		"baremetal_core_volumes":                    VolumeDatasource(),
		"baremetal_database_database":               DatabaseDatasource(),
		"baremetal_database_databases":              DatabasesDatasource(),
		"baremetal_database_db_home":                DBHomeDatasource(),
		"baremetal_database_db_homes":               DBHomesDatasource(),
		"baremetal_database_db_node":                DBNodeDatasource(),
		"baremetal_database_db_nodes":               DBNodesDatasource(),
		"baremetal_database_db_system_shapes":       DBSystemShapeDatasource(),
		"baremetal_database_db_systems":             DBSystemDatasource(),
		"baremetal_database_db_versions":            DBVersionDatasource(),
		"baremetal_database_supported_operations":   SupportedOperationDatasource(),
		"baremetal_identity_api_keys":               APIKeyDatasource(),
		"baremetal_identity_availability_domains":   AvailabilityDomainDatasource(),
		"baremetal_identity_compartments":           CompartmentDatasource(),
		"baremetal_identity_groups":                 GroupDatasource(),
		"baremetal_identity_policies":               PolicyDatasource(),
		"baremetal_identity_swift_passwords":        SwiftPasswordDatasource(),
		"baremetal_identity_user_group_memberships": UserGroupMembershipDatasource(),
		"baremetal_identity_users":                  UserDatasource(),
		"baremetal_load_balancer_backends":          lb.BackendDatasource(),
		"baremetal_load_balancer_backendsets":       lb.BackendSetDatasource(),
		"baremetal_load_balancer_certificates":      lb.CertificateDatasource(),
		"baremetal_load_balancer_policies":          lb.PolicyDatasource(),
		"baremetal_load_balancer_protocols":         lb.ProtocolDatasource(),
		"baremetal_load_balancer_shapes":            lb.ShapeDatasource(),
		"baremetal_load_balancers":                  lb.LoadBalancerDatasource(),
		"baremetal_objectstorage_bucket_summaries":  BucketSummaryDatasource(),
		"baremetal_objectstorage_namespace":         NamespaceDatasource(),
		"baremetal_objectstorage_object_head":       ObjectHeadDatasource(),
		"baremetal_objectstorage_objects":           ObjectDatasource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history":           ConsoleHistoryResource(),
		"baremetal_core_cpe":                       CpeResource(),
		"baremetal_core_dhcp_options":              DHCPOptionsResource(),
		"baremetal_core_drg":                       DrgResource(),
		"baremetal_core_drg_attachment":            DrgAttachmentResource(),
		"baremetal_core_image":                     ImageResource(),
		"baremetal_core_instance":                  InstanceResource(),
		"baremetal_core_internet_gateway":          InternetGatewayResource(),
		"baremetal_core_ipsec":                     IPSecConnectionResource(),
		"baremetal_core_route_table":               RouteTableResource(),
		"baremetal_core_security_list":             SecurityListResource(),
		"baremetal_core_subnet":                    SubnetResource(),
		"baremetal_core_virtual_network":           VirtualNetworkResource(),
		"baremetal_core_volume":                    VolumeResource(),
		"baremetal_core_volume_attachment":         VolumeAttachmentResource(),
		"baremetal_core_volume_backup":             VolumeBackupResource(),
		"baremetal_database_db_system":             DBSystemResource(),
		"baremetal_identity_api_key":               APIKeyResource(),
		"baremetal_identity_compartment":           CompartmentResource(),
		"baremetal_identity_group":                 GroupResource(),
		"baremetal_identity_policy":                PolicyResource(),
		"baremetal_identity_swift_password":        SwiftPasswordResource(),
		"baremetal_identity_ui_password":           UIPasswordResource(),
		"baremetal_identity_user":                  UserResource(),
		"baremetal_identity_user_group_membership": UserGroupMembershipResource(),
		"baremetal_load_balancer":                  lb.LoadBalancerResource(),
		"baremetal_load_balancer_backend":          lb.LoadBalancerBackendResource(),
		"baremetal_load_balancer_backendset":       lb.LoadBalancerBackendSetResource(),
		"baremetal_load_balancer_certificate":      lb.LoadBalancerCertificateResource(),
		"baremetal_load_balancer_listener":         lb.LoadBalancerListenerResource(),
		"baremetal_objectstorage_bucket":           BucketResource(),
		"baremetal_objectstorage_object":           ObjectResource(),
	}
}

func providerConfig(d *schema.ResourceData) (client interface{}, err error) {
	tenancyOCID := d.Get("tenancy_ocid").(string)
	userOCID := d.Get("user_ocid").(string)
	fingerprint := d.Get("fingerprint").(string)
	privateKeyBuffer, hasKey := d.Get("private_key").(string)
	privateKeyPath, hasKeyPath := d.Get("private_key_path").(string)
	privateKeyPassword, hasKeyPass := d.Get("private_key_password").(string)

	clientOpts := []baremetal.NewClientOptionsFunc{
		func(o *baremetal.NewClientOptions) {
			o.UserAgent = fmt.Sprintf("baremetal-terraform-v%s", baremetal.SDKVersion)
		},
		func(o *baremetal.NewClientOptions) {
			o.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
		},
	}

	if hasKey && privateKeyBuffer != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyBytes([]byte(privateKeyBuffer)))
	} else if hasKeyPath && privateKeyPath != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyFilePath(privateKeyPath))
	} else {
		err = errors.New("One of private_key or private_key_path is required")
		return
	}

	if hasKeyPass && privateKeyPassword != "" {
		clientOpts = append(clientOpts, baremetal.PrivateKeyPassword(privateKeyPassword))
	}

	client, err = baremetal.NewClient(userOCID, tenancyOCID, fingerprint, clientOpts...)
	return
}
