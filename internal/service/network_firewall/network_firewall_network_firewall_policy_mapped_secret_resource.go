// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyMappedSecretResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyMappedSecret,
		Read:     readNetworkFirewallNetworkFirewallPolicyMappedSecret,
		Update:   updateNetworkFirewallNetworkFirewallPolicyMappedSecret,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyMappedSecret,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"OCI_VAULT",
				}, true),
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"SSL_INBOUND_INSPECTION",
					"SSL_FORWARD_PROXY",
				}, false),
			},
			"vault_secret_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_number": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyMappedSecret(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyMappedSecret(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyMappedSecret(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyMappedSecret(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.MappedSecret
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "mappedSecrets")
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) Create() error {
	request := oci_network_firewall.CreateMappedSecretRequest{}
	err := s.populateTopLevelPolymorphicCreateMappedSecretRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateMappedSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedSecret
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) Get() error {
	request := oci_network_firewall.GetMappedSecretRequest{}

	if mappedSecretName, ok := s.D.GetOkExists("name"); ok {
		tmp := mappedSecretName.(string)
		request.MappedSecretName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	mappedSecretName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "mappedSecrets")
	if err == nil {
		request.MappedSecretName = &mappedSecretName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetMappedSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedSecret
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) Update() error {
	request := oci_network_firewall.UpdateMappedSecretRequest{}
	err := s.populateTopLevelPolymorphicUpdateMappedSecretRequest(&request)
	if err != nil {
		return err
	}

	if mappedSecretName, ok := s.D.GetOkExists("name"); ok {
		tmp := mappedSecretName.(string)
		request.MappedSecretName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateMappedSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MappedSecret
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteMappedSecretRequest{}

	if mappedSecretName, ok := s.D.GetOkExists("name"); ok {
		tmp := mappedSecretName.(string)
		request.MappedSecretName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteMappedSecret(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) SetData() error {

	mappedSecretName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "mappedSecrets")
	if err == nil {
		s.D.Set("name", &mappedSecretName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_network_firewall.VaultMappedSecret:
		s.D.Set("source", "OCI_VAULT")

		if v.VaultSecretId != nil {
			s.D.Set("vault_secret_id", *v.VaultSecretId)
		}

		if v.VersionNumber != nil {
			s.D.Set("version_number", *v.VersionNumber)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}

		s.D.Set("type", v.Type)
	default:
		log.Printf("[WARN] Received 'source' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func MappedSecretSummaryToMap(obj oci_network_firewall.MappedSecretSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) populateTopLevelPolymorphicCreateMappedSecretRequest(request *oci_network_firewall.CreateMappedSecretRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("OCI_VAULT"):
		details := oci_network_firewall.CreateVaultMappedSecretDetails{}
		if vaultSecretId, ok := s.D.GetOkExists("vault_secret_id"); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
			tmp := versionNumber.(int)
			details.VersionNumber = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_network_firewall.InspectionTypeEnum(type_.(string))
		}
		request.CreateMappedSecretDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyMappedSecretResourceCrud) populateTopLevelPolymorphicUpdateMappedSecretRequest(request *oci_network_firewall.UpdateMappedSecretRequest) error {
	//discriminator
	sourceRaw, ok := s.D.GetOkExists("source")
	var source string
	if ok {
		source = sourceRaw.(string)
	} else {
		source = "" // default value
	}
	switch strings.ToLower(source) {
	case strings.ToLower("OCI_VAULT"):
		details := oci_network_firewall.UpdateVaultMappedSecretDetails{}
		if vaultSecretId, ok := s.D.GetOkExists("vault_secret_id"); ok {
			tmp := vaultSecretId.(string)
			details.VaultSecretId = &tmp
		}
		if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
			tmp := versionNumber.(int)
			details.VersionNumber = &tmp
		}
		if type_, ok := s.D.GetOkExists("type"); ok {
			details.Type = oci_network_firewall.InspectionTypeEnum(type_.(string))
		}
		request.UpdateMappedSecretDetails = details
	default:
		return fmt.Errorf("unknown source '%v' was specified", source)
	}
	return nil
}
