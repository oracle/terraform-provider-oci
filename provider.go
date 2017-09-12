// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"tenancy_ocid": "(Required) The tenancy OCID for a user. The tenancy OCID can be found at the bottom of user settings in the Oracle Cloud Infrastructure console.",
		"user_ocid":    "(Required) The user OCID. This can be found in user settings in the Oracle Cloud Infrastructure console.",
		"fingerprint":  "(Required) The fingerprint for the user's RSA key. This can be found in user settings in the Oracle Cloud Infrastructure console.",
		"region":       "(Required) The region for API connections (e.g. us-ashburn-1).",
		"private_key": "(Optional) A PEM formatted RSA private key for the user.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_path": "(Optional) The path to the user's PEM formatted private key.\n" +
			"A private_key or a private_key_path must be provided.",
		"private_key_password": "(Optional) The password used to secure the private key.",
		"disable_auto_retries": "(Optional) Disable Automatic retries for retriable errors.\n" +
			"Auto retries were introduced to solve some eventual consistency problems but it also introduced performance issues on destroy operations.",
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
			DefaultFunc: schema.EnvDefaultFunc("OCI_TENANCY_OCID", nil),
		},
		"user_ocid": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["user_ocid"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_USER_OCID", nil),
		},
		"fingerprint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["fingerprint"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_FINGERPRINT", nil),
		},
		// Mostly used for testing. Don't put keys in your .tf files
		"private_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Sensitive:   true,
			Description: descriptions["private_key"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY", nil),
		},
		"private_key_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: descriptions["private_key_path"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY_PATH", nil),
		},
		"private_key_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Default:     "",
			Description: descriptions["private_key_password"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_PRIVATE_KEY_PASSWORD", nil),
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["region"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_REGION", nil),
		},
		"disable_auto_retries": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: descriptions["disable_auto_retries"],
			DefaultFunc: schema.EnvDefaultFunc("OCI_DISABLE_AUTO_RETRIES", nil),
		},
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_core_console_history_data":       ConsoleHistoryDataDatasource(),
		"oci_core_cpes":                       CpeDatasource(),
		"oci_core_dhcp_options":               DHCPOptionsDatasource(),
		"oci_core_drg_attachments":            DrgAttachmentDatasource(),
		"oci_core_drgs":                       DrgDatasource(),
		"oci_core_images":                     ImageDatasource(),
		"oci_core_instance_credentials":       InstanceCredentialsDatasource(),
		"oci_core_instances":                  InstanceDatasource(),
		"oci_core_internet_gateways":          InternetGatewayDatasource(),
		"oci_core_ipsec_config":               IPSecConnectionConfigDatasource(),
		"oci_core_ipsec_connections":          IPSecConnectionsDatasource(),
		"oci_core_ipsec_status":               IPSecConnectionStatusDatasource(),
		"oci_core_route_tables":               RouteTableDatasource(),
		"oci_core_security_lists":             SecurityListDatasource(),
		"oci_core_shape":                      InstanceShapeDatasource(),
		"oci_core_subnets":                    SubnetDatasource(),
		"oci_core_virtual_networks":           VirtualNetworkDatasource(),
		"oci_core_vnic":                       VnicDatasource(),
		"oci_core_vnic_attachments":           DatasourceCoreVnicAttachments(),
		"oci_core_volume_attachments":         VolumeAttachmentDatasource(),
		"oci_core_volume_backups":             VolumeBackupDatasource(),
		"oci_core_volumes":                    VolumeDatasource(),
		"oci_database_database":               DatabaseDatasource(),
		"oci_database_databases":              DatabasesDatasource(),
		"oci_database_db_home":                DBHomeDatasource(),
		"oci_database_db_homes":               DBHomesDatasource(),
		"oci_database_db_node":                DBNodeDatasource(),
		"oci_database_db_nodes":               DBNodesDatasource(),
		"oci_database_db_system_shapes":       DBSystemShapeDatasource(),
		"oci_database_db_systems":             DBSystemDatasource(),
		"oci_database_db_versions":            DBVersionDatasource(),
		"oci_identity_api_keys":               APIKeyDatasource(),
		"oci_identity_availability_domains":   AvailabilityDomainDatasource(),
		"oci_identity_compartments":           CompartmentDatasource(),
		"oci_identity_groups":                 GroupDatasource(),
		"oci_identity_policies":               IdentityPolicyDatasource(),
		"oci_identity_swift_passwords":        SwiftPasswordDatasource(),
		"oci_identity_user_group_memberships": UserGroupMembershipDatasource(),
		"oci_identity_users":                  UserDatasource(),
		"oci_load_balancer_backends":          BackendDatasource(),
		"oci_load_balancer_backendsets":       BackendSetDatasource(),
		"oci_load_balancer_certificates":      CertificateDatasource(),
		"oci_load_balancer_policies":          LoadBalancerPolicyDatasource(),
		"oci_load_balancer_protocols":         ProtocolDatasource(),
		"oci_load_balancer_shapes":            LoadBalancerShapeDatasource(),
		"oci_load_balancers":                  LoadBalancerDatasource(),
		"oci_objectstorage_bucket_summaries":  BucketSummaryDatasource(),
		"oci_objectstorage_namespace":         NamespaceDatasource(),
		"oci_objectstorage_object_head":       ObjectHeadDatasource(),
		"oci_objectstorage_objects":           ObjectDatasource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"oci_core_console_history":           ConsoleHistoryResource(),
		"oci_core_cpe":                       CpeResource(),
		"oci_core_dhcp_options":              DHCPOptionsResource(),
		"oci_core_drg":                       DrgResource(),
		"oci_core_drg_attachment":            DrgAttachmentResource(),
		"oci_core_image":                     ImageResource(),
		"oci_core_instance":                  InstanceResource(),
		"oci_core_internet_gateway":          InternetGatewayResource(),
		"oci_core_ipsec":                     IPSecConnectionResource(),
		"oci_core_route_table":               RouteTableResource(),
		"oci_core_security_list":             SecurityListResource(),
		"oci_core_subnet":                    SubnetResource(),
		"oci_core_virtual_network":           VirtualNetworkResource(),
		"oci_core_vnic_attachment":           VnicAttachmentResource(),
		"oci_core_volume":                    VolumeResource(),
		"oci_core_volume_attachment":         VolumeAttachmentResource(),
		"oci_core_volume_backup":             VolumeBackupResource(),
		"oci_database_db_system":             DBSystemResource(),
		"oci_identity_api_key":               APIKeyResource(),
		"oci_identity_compartment":           CompartmentResource(),
		"oci_identity_group":                 GroupResource(),
		"oci_identity_policy":                PolicyResource(),
		"oci_identity_swift_password":        SwiftPasswordResource(),
		"oci_identity_ui_password":           UIPasswordResource(),
		"oci_identity_user":                  UserResource(),
		"oci_identity_user_group_membership": UserGroupMembershipResource(),
		"oci_load_balancer":                  LoadBalancerResource(),
		"oci_load_balancer_backend":          LoadBalancerBackendResource(),
		"oci_load_balancer_backendset":       LoadBalancerBackendSetResource(),
		"oci_load_balancer_certificate":      LoadBalancerCertificateResource(),
		"oci_load_balancer_listener":         LoadBalancerListenerResource(),
		"oci_objectstorage_bucket":           BucketResource(),
		"oci_objectstorage_object":           ObjectResource(),
		"oci_objectstorage_preauthrequest":   PreauthenticatedRequestResource(),
	}
}

func getEnvSetting(s string, dv string) string {
	v := os.Getenv("TF_VAR_" + s)
	if v != "" {
		return v
	}
	v = os.Getenv("OCI_" + s)
	if v != "" {
		return v
	}
	v = os.Getenv(s)
	if v != "" {
		return v
	}
	return dv
}

func getRequiredEnvSetting(s string) string {
	v := getEnvSetting(s, "")
	if v == "" {
		panic(fmt.Sprintf("Required env setting %s is missing", s))
	}
	return v
}

func providerConfig(d *schema.ResourceData) (client interface{}, err error) {
	tenancyOCID := d.Get("tenancy_ocid").(string)
	userOCID := d.Get("user_ocid").(string)
	fingerprint := d.Get("fingerprint").(string)
	privateKeyBuffer, hasKey := d.Get("private_key").(string)
	privateKeyPath, hasKeyPath := d.Get("private_key_path").(string)
	privateKeyPassword, hasKeyPass := d.Get("private_key_password").(string)
	region, hasRegion := d.Get("region").(string)
	disableAutoRetries, hasDisableRetries := d.Get("disable_auto_retries").(bool)

	// for internal use
	urlTemplate := getEnvSetting("url_template", "")
	allowInsecureTls := getEnvSetting("allow_insecure_tls", "")

	clientOpts := []baremetal.NewClientOptionsFunc{
		func(o *baremetal.NewClientOptions) {
			o.UserAgent = fmt.Sprintf("oci-terraform-v%s", baremetal.SDKVersion)
		},
	}

	if allowInsecureTls == "true" {
		log.Println("[WARN] USING INSECURE TLS")
		clientOpts = append(clientOpts, baremetal.CustomTransport(
			&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		))
	} else {
		clientOpts = append(clientOpts, baremetal.CustomTransport(
			&http.Transport{Proxy: http.ProxyFromEnvironment}),
		)
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

	if hasRegion && region != "" {
		clientOpts = append(clientOpts, baremetal.Region(region))
	}

	if hasDisableRetries {
		clientOpts = append(clientOpts, baremetal.DisableAutoRetries(disableAutoRetries))
	}

	if urlTemplate != "" {
		clientOpts = append(clientOpts, baremetal.UrlTemplate(urlTemplate))
	}

	client, err = baremetal.NewClient(userOCID, tenancyOCID, fingerprint, clientOpts...)
	return
}
