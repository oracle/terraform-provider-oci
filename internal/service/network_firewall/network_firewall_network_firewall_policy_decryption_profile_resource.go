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

func NetworkFirewallNetworkFirewallPolicyDecryptionProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyDecryptionProfile,
		Read:     readNetworkFirewallNetworkFirewallPolicyDecryptionProfile,
		Update:   updateNetworkFirewallNetworkFirewallPolicyDecryptionProfile,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyDecryptionProfile,
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
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"SSL_FORWARD_PROXY",
					"SSL_INBOUND_INSPECTION",
				}, true),
			},

			// Optional
			"are_certificate_extensions_restricted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_auto_include_alt_name": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_expired_certificate_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_out_of_capacity_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_revocation_status_timeout_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_unknown_revocation_status_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_unsupported_cipher_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_unsupported_version_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_untrusted_issuer_blocked": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyDecryptionProfile(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyDecryptionProfile(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyDecryptionProfile(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyDecryptionProfile(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.DecryptionProfile
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "decryptionProfiles")
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) Create() error {
	request := oci_network_firewall.CreateDecryptionProfileRequest{}
	err := s.populateTopLevelPolymorphicCreateDecryptionProfileRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateDecryptionProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionProfile
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) Get() error {
	request := oci_network_firewall.GetDecryptionProfileRequest{}

	if decryptionProfileName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionProfileName.(string)
		request.DecryptionProfileName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	decryptionProfileName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "decryptionProfiles")
	if err == nil {
		request.DecryptionProfileName = &decryptionProfileName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetDecryptionProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionProfile
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) Update() error {
	request := oci_network_firewall.UpdateDecryptionProfileRequest{}
	err := s.populateTopLevelPolymorphicUpdateDecryptionProfileRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if decryptionProfileName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionProfileName.(string)
		request.DecryptionProfileName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateDecryptionProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionProfile
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteDecryptionProfileRequest{}

	if decryptionProfileName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionProfileName.(string)
		request.DecryptionProfileName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteDecryptionProfile(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) SetData() error {

	decryptionProfileName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "decryptionProfiles")
	if err == nil {
		s.D.Set("name", &decryptionProfileName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
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
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func DecryptionProfileSummaryToMap(obj oci_network_firewall.DecryptionProfileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) populateTopLevelPolymorphicCreateDecryptionProfileRequest(request *oci_network_firewall.CreateDecryptionProfileRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SSL_FORWARD_PROXY"):
		details := oci_network_firewall.CreateSslForwardProxyProfileDetails{}
		if areCertificateExtensionsRestricted, ok := s.D.GetOkExists("are_certificate_extensions_restricted"); ok {
			tmp := areCertificateExtensionsRestricted.(bool)
			details.AreCertificateExtensionsRestricted = &tmp
		}
		if isAutoIncludeAltName, ok := s.D.GetOkExists("is_auto_include_alt_name"); ok {
			tmp := isAutoIncludeAltName.(bool)
			details.IsAutoIncludeAltName = &tmp
		}
		if isExpiredCertificateBlocked, ok := s.D.GetOkExists("is_expired_certificate_blocked"); ok {
			tmp := isExpiredCertificateBlocked.(bool)
			details.IsExpiredCertificateBlocked = &tmp
		}
		if isOutOfCapacityBlocked, ok := s.D.GetOkExists("is_out_of_capacity_blocked"); ok {
			tmp := isOutOfCapacityBlocked.(bool)
			details.IsOutOfCapacityBlocked = &tmp
		}
		if isRevocationStatusTimeoutBlocked, ok := s.D.GetOkExists("is_revocation_status_timeout_blocked"); ok {
			tmp := isRevocationStatusTimeoutBlocked.(bool)
			details.IsRevocationStatusTimeoutBlocked = &tmp
		}
		if isUnknownRevocationStatusBlocked, ok := s.D.GetOkExists("is_unknown_revocation_status_blocked"); ok {
			tmp := isUnknownRevocationStatusBlocked.(bool)
			details.IsUnknownRevocationStatusBlocked = &tmp
		}
		if isUnsupportedCipherBlocked, ok := s.D.GetOkExists("is_unsupported_cipher_blocked"); ok {
			tmp := isUnsupportedCipherBlocked.(bool)
			details.IsUnsupportedCipherBlocked = &tmp
		}
		if isUnsupportedVersionBlocked, ok := s.D.GetOkExists("is_unsupported_version_blocked"); ok {
			tmp := isUnsupportedVersionBlocked.(bool)
			details.IsUnsupportedVersionBlocked = &tmp
		}
		if isUntrustedIssuerBlocked, ok := s.D.GetOkExists("is_untrusted_issuer_blocked"); ok {
			tmp := isUntrustedIssuerBlocked.(bool)
			details.IsUntrustedIssuerBlocked = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateDecryptionProfileDetails = details
	case strings.ToLower("SSL_INBOUND_INSPECTION"):
		details := oci_network_firewall.CreateSslInboundInspectionProfileDetails{}
		if isOutOfCapacityBlocked, ok := s.D.GetOkExists("is_out_of_capacity_blocked"); ok {
			tmp := isOutOfCapacityBlocked.(bool)
			details.IsOutOfCapacityBlocked = &tmp
		}
		if isUnsupportedCipherBlocked, ok := s.D.GetOkExists("is_unsupported_cipher_blocked"); ok {
			tmp := isUnsupportedCipherBlocked.(bool)
			details.IsUnsupportedCipherBlocked = &tmp
		}
		if isUnsupportedVersionBlocked, ok := s.D.GetOkExists("is_unsupported_version_blocked"); ok {
			tmp := isUnsupportedVersionBlocked.(bool)
			details.IsUnsupportedVersionBlocked = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateDecryptionProfileDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionProfileResourceCrud) populateTopLevelPolymorphicUpdateDecryptionProfileRequest(request *oci_network_firewall.UpdateDecryptionProfileRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SSL_FORWARD_PROXY"):
		details := oci_network_firewall.UpdateSslForwardProxyProfileDetails{}
		if areCertificateExtensionsRestricted, ok := s.D.GetOkExists("are_certificate_extensions_restricted"); ok {
			tmp := areCertificateExtensionsRestricted.(bool)
			details.AreCertificateExtensionsRestricted = &tmp
		}
		if isAutoIncludeAltName, ok := s.D.GetOkExists("is_auto_include_alt_name"); ok {
			tmp := isAutoIncludeAltName.(bool)
			details.IsAutoIncludeAltName = &tmp
		}
		if isExpiredCertificateBlocked, ok := s.D.GetOkExists("is_expired_certificate_blocked"); ok {
			tmp := isExpiredCertificateBlocked.(bool)
			details.IsExpiredCertificateBlocked = &tmp
		}
		if isOutOfCapacityBlocked, ok := s.D.GetOkExists("is_out_of_capacity_blocked"); ok {
			tmp := isOutOfCapacityBlocked.(bool)
			details.IsOutOfCapacityBlocked = &tmp
		}
		if isRevocationStatusTimeoutBlocked, ok := s.D.GetOkExists("is_revocation_status_timeout_blocked"); ok {
			tmp := isRevocationStatusTimeoutBlocked.(bool)
			details.IsRevocationStatusTimeoutBlocked = &tmp
		}
		if isUnknownRevocationStatusBlocked, ok := s.D.GetOkExists("is_unknown_revocation_status_blocked"); ok {
			tmp := isUnknownRevocationStatusBlocked.(bool)
			details.IsUnknownRevocationStatusBlocked = &tmp
		}
		if isUnsupportedCipherBlocked, ok := s.D.GetOkExists("is_unsupported_cipher_blocked"); ok {
			tmp := isUnsupportedCipherBlocked.(bool)
			details.IsUnsupportedCipherBlocked = &tmp
		}
		if isUnsupportedVersionBlocked, ok := s.D.GetOkExists("is_unsupported_version_blocked"); ok {
			tmp := isUnsupportedVersionBlocked.(bool)
			details.IsUnsupportedVersionBlocked = &tmp
		}
		if isUntrustedIssuerBlocked, ok := s.D.GetOkExists("is_untrusted_issuer_blocked"); ok {
			tmp := isUntrustedIssuerBlocked.(bool)
			details.IsUntrustedIssuerBlocked = &tmp
		}
		request.UpdateDecryptionProfileDetails = details
	case strings.ToLower("SSL_INBOUND_INSPECTION"):
		details := oci_network_firewall.UpdateSslInboundInspectionProfileDetails{}
		if isOutOfCapacityBlocked, ok := s.D.GetOkExists("is_out_of_capacity_blocked"); ok {
			tmp := isOutOfCapacityBlocked.(bool)
			details.IsOutOfCapacityBlocked = &tmp
		}
		if isUnsupportedCipherBlocked, ok := s.D.GetOkExists("is_unsupported_cipher_blocked"); ok {
			tmp := isUnsupportedCipherBlocked.(bool)
			details.IsUnsupportedCipherBlocked = &tmp
		}
		if isUnsupportedVersionBlocked, ok := s.D.GetOkExists("is_unsupported_version_blocked"); ok {
			tmp := isUnsupportedVersionBlocked.(bool)
			details.IsUnsupportedVersionBlocked = &tmp
		}
		request.UpdateDecryptionProfileDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
