// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["network_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(NetworkFirewallNetworkFirewallPolicyDecryptionProfileResource(), fieldMap, readSingularNetworkFirewallNetworkFirewallPolicyDecryptionProfile)
}

func readSingularNetworkFirewallNetworkFirewallPolicyDecryptionProfile(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

type NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_firewall.NetworkFirewallClient
	Res    *oci_network_firewall.GetDecryptionProfileResponse
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSourceCrud) Get() error {
	request := oci_network_firewall.GetDecryptionProfileRequest{}

	if decryptionProfileName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionProfileName.(string)
		request.DecryptionProfileName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_firewall")

	response, err := s.Client.GetDecryptionProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSource-", NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSource(), s.D))
	switch v := (s.Res.DecryptionProfile).(type) {
	case oci_network_firewall.SslForwardProxyProfile:
		s.D.Set("type", "SSL_FORWARD_PROXY")

		if v.AreCertificateExtensionsRestricted != nil {
			s.D.Set("are_certificate_extensions_restricted", *v.AreCertificateExtensionsRestricted)
		}

		if v.IsAutoIncludeAltName != nil {
			s.D.Set("is_auto_include_alt_name", *v.IsAutoIncludeAltName)
		}

		if v.IsExpiredCertificateBlocked != nil {
			s.D.Set("is_expired_certificate_blocked", *v.IsExpiredCertificateBlocked)
		}

		if v.IsOutOfCapacityBlocked != nil {
			s.D.Set("is_out_of_capacity_blocked", *v.IsOutOfCapacityBlocked)
		}

		if v.IsRevocationStatusTimeoutBlocked != nil {
			s.D.Set("is_revocation_status_timeout_blocked", *v.IsRevocationStatusTimeoutBlocked)
		}

		if v.IsUnknownRevocationStatusBlocked != nil {
			s.D.Set("is_unknown_revocation_status_blocked", *v.IsUnknownRevocationStatusBlocked)
		}

		if v.IsUnsupportedCipherBlocked != nil {
			s.D.Set("is_unsupported_cipher_blocked", *v.IsUnsupportedCipherBlocked)
		}

		if v.IsUnsupportedVersionBlocked != nil {
			s.D.Set("is_unsupported_version_blocked", *v.IsUnsupportedVersionBlocked)
		}

		if v.IsUntrustedIssuerBlocked != nil {
			s.D.Set("is_untrusted_issuer_blocked", *v.IsUntrustedIssuerBlocked)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	case oci_network_firewall.SslInboundInspectionProfile:
		s.D.Set("type", "SSL_INBOUND_INSPECTION")

		if v.IsOutOfCapacityBlocked != nil {
			s.D.Set("is_out_of_capacity_blocked", *v.IsOutOfCapacityBlocked)
		}

		if v.IsUnsupportedCipherBlocked != nil {
			s.D.Set("is_unsupported_cipher_blocked", *v.IsUnsupportedCipherBlocked)
		}

		if v.IsUnsupportedVersionBlocked != nil {
			s.D.Set("is_unsupported_version_blocked", *v.IsUnsupportedVersionBlocked)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DecryptionProfile)
		return nil
	}

	return nil
}
