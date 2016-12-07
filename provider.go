package main

import (
	"crypto/rsa"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/core"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/identity"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
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
		"private_key_password": "(Required) The password used to secure the private key.",
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
		"tenancy_ocid": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["tenancy_ocid"],
		},
		"user_ocid": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["user_ocid"],
		},
		"fingerprint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: descriptions["fingerprint"],
		},
		"private_key": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: descriptions["private_key"],
		},
		"private_key_path": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: descriptions["private_key_path"],
		},
		"private_key_password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: descriptions["private_key_password"],
		},
	}
}

func dataSourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history_data": core.ConsoleHistoryDataDatasource(),
		"baremetal_core_cpes":                 core.CpeDatasource(),
		"baremetal_core_dhcp_options":         core.DHCPOptionsDatasource(),
		"baremetal_core_drg_attachments":      core.DrgAttachmentDatasource(),
		"baremetal_core_drgs":                 core.DrgDatasource(),
		"baremetal_core_images":               core.ImageDatasource(),
		"baremetal_core_instances":            core.InstanceDatasource(),
		"baremetal_core_internet_gateways":    core.InternetGatewayDatasource(),
		"baremetal_core_ipsec_config":         core.IPSecConfigDatasource(),
		"baremetal_core_ipsec_connections":    core.IPSecConnectionsDatasource(),
		"baremetal_core_ipsec_status":         core.IPSecStatusDatasource(),
		"baremetal_core_route_tables":         core.RouteTableDatasource(),
		"baremetal_core_shape":                core.ShapeDatasource(),
		"baremetal_core_subnets":              core.SubnetDatasource(),
		"baremetal_core_virtual_networks":     core.VirtualNetworkDatasource(),
		"baremetal_core_vnic_attachments":     core.DatasourceCoreVnicAttachments(),
		"baremetal_core_vnic":                 core.VnicDatasource(),
		"baremetal_core_volume_attachments":   core.VolumeAttachmentDatasource(),
		"baremetal_core_volume_backups":       core.VolumeBackupDatasource(),
		"baremetal_core_volumes":              core.VolumeDatasource(),
		"baremetal_identity_api_keys":         identity.APIKeyDatasource(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_core_console_history":   core.ConsoleHistoryResource(),
		"baremetal_core_cpe":               core.CpeResource(),
		"baremetal_core_dhcp_options":      core.DHCPOptionsResource(),
		"baremetal_core_drg_attachment":    core.DrgAttachmentResource(),
		"baremetal_core_drg":               core.DrgResource(),
		"baremetal_core_image":             core.ImageResource(),
		"baremetal_core_instance":          core.InstanceResource(),
		"baremetal_core_internet_gateway":  core.InternetGatewayResource(),
		"baremetal_core_ipsec":             core.IPSecResource(),
		"baremetal_core_route_table":       core.RouteTableResource(),
		"baremetal_core_subnet":            core.SubnetResource(),
		"baremetal_core_virtual_network":   core.VirtualNetworkResource(),
		"baremetal_core_volume_attachment": core.VolumeAttachmentResource(),
		"baremetal_core_volume_backup":     core.VolumeBackupResource(),
		"baremetal_core_volume":            core.VolumeResource(),
		"baremetal_identity_api_key":       identity.APIKeyResource(),
		"baremetal_identity_compartment":   identity.CompartmentResource(),
		"baremetal_identity_group":         identity.GroupResource(),
		"baremetal_identity_policy":        identity.PolicyResource(),
		"baremetal_identity_ui_password":   identity.UIPasswordResource(),
		"baremetal_identity_user":          identity.UserResource(),
	}
}

func providerConfig(d *schema.ResourceData) (client interface{}, err error) {
	tenancyOCID := d.Get("tenancy_ocid").(string)
	userOCID := d.Get("user_ocid").(string)
	fingerprint := d.Get("fingerprint").(string)
	privateKeyBuffer := d.Get("private_key").(string)
	privateKeyPath := d.Get("private_key_path").(string)
	privateKeyPassword := d.Get("private_key_password").(string)

	var privateKey *rsa.PrivateKey

	if privateKeyBuffer != "" {
		if privateKey, err = baremetal.PrivateKeyFromBytes([]byte(privateKeyBuffer), privateKeyPassword); err != nil {
			return nil, err
		}
	}

	if privateKeyPath != "" {
		if privateKey, err = baremetal.PrivateKeyFromFile(privateKeyPath, privateKeyPassword); err != nil {
			return nil, err
		}
	}

	client = baremetal.New(userOCID, tenancyOCID, fingerprint, privateKey, "")
	return
}
