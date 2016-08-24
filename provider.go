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
		"baremetal_core_shape":              core.ResourceCoreShape(),
		"baremetal_core_vnic_attachments":   core.ResourceCoreVnicAttachments(),
		"baremetal_core_volumes":            core.ResourceCoreVolumes(),
		"baremetal_core_volume_attachments": core.ResourceCoreVolumeAttachments(),
		"baremetal_core_instances":          core.ResourceCoreInstances(),
		"baremetal_core_subnets":            core.ResourceCoreSubnets(),
		"baremetal_core_drgs":               core.ResourceCoreDrgs(),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_identity_user":          identity.ResourceIdentityUser(),
		"baremetal_identity_group":         identity.ResourceIdentityGroup(),
		"baremetal_identity_policy":        identity.ResourceIdentityPolicy(),
		"baremetal_identity_compartment":   identity.ResourceIdentityCompartment(),
		"baremetal_core_cpe":               core.ResourceCoreCpe(),
		"baremetal_core_volume":            core.ResourceCoreVolume(),
		"baremetal_core_instance":          core.ResourceCoreInstance(),
		"baremetal_core_volume_attachment": core.ResourceCoreVolumeAttachment(),
		"baremetal_core_subnet":            core.ResourceCoreSubnet(),
		"baremetal_core_drg":               core.ResourceCoreDrg(),
		"baremetal_core_virtual_network":   core.ResourceCoreVirtualNetwork(),
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

	client = baremetal.New(userOCID, tenancyOCID, fingerprint, privateKey)
	return
}
