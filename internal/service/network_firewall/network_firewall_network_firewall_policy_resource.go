// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("10m"),
			Update: tfresource.GetTimeoutDuration("10m"),
			Delete: tfresource.GetTimeoutDuration("5m"),
		},
		Create: createNetworkFirewallNetworkFirewallPolicy,
		Read:   readNetworkFirewallNetworkFirewallPolicy,
		Update: updateNetworkFirewallNetworkFirewallPolicy,
		Delete: deleteNetworkFirewallNetworkFirewallPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//Required
						"application_list_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						//Optional
						"application_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									// Optional
									"icmp_type": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"icmp_code": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"minimum_port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"maximum_port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"decryption_profiles": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				//Set:      originsHashCodeForProfiles,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						// Optional
						"is_out_of_capacity_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_unsupported_cipher_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_unsupported_version_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"are_certificate_extensions_restricted": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_auto_include_alt_name": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_expired_certificate_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_revocation_status_timeout_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_untrusted_issuer_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_unknown_revocation_status_blocked": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						// Computed
					},
				},
			},
			"decryption_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"condition": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"destinations": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sources": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"decryption_profile": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ip_address_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ip_address_list_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address_list_value": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"mapped_secrets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				//Set:      mappedSecretsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"vault_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version_number": {
							Type:     schema.TypeInt,
							Default:  443,
							Optional: true,
						},
						// Computed
					},
				},
			},
			"security_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"condition": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"applications": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"destinations": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sources": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"urls": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"inspection": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"url_lists": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_list_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"url_list_values": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"pattern": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			// Computed
			"is_firewall_attached": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.NetworkFirewallPolicy
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateCreating),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateActive),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleting),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_network_firewall.LifecycleStateDeleted),
	}
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Create() error {
	request := oci_network_firewall.CreateNetworkFirewallPolicyRequest{}

	if applicationLists, ok := s.D.GetOkExists("application_lists"); ok {
		resultApplicationLists, err := s.objectMapToApplicationListsMap(applicationLists)
		if err != nil {
			return err
		}
		if len(resultApplicationLists) > 0 {
			request.ApplicationLists = resultApplicationLists
		}
	}

	if urlLists, ok := s.D.GetOkExists("url_lists"); ok {
		resultUrlLists, err := s.objectMapToUrlListsMap(urlLists)
		if err != nil {
			return err
		}
		if len(resultUrlLists) > 0 {
			request.UrlLists = resultUrlLists
		}
	}

	if ipAddressLists, ok := s.D.GetOk("ip_address_lists"); ok {
		request.IpAddressLists = s.objectMapToStringListMap(ipAddressLists)
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if decryptionProfiles, ok := s.D.GetOkExists("decryption_profiles"); ok {
		resultMap, err := s.objectMapToDecryptionProfileMap(decryptionProfiles)
		if err != nil {
			return err
		}
		if len(resultMap) > 0 {
			request.DecryptionProfiles = resultMap
		}
	}

	if mappedSecrets, ok := s.D.GetOkExists("mapped_secrets"); ok {
		resultSecrets, err := s.objectMapToMappedSecrets(mappedSecrets)
		if err != nil {
			return err
		}
		if len(resultSecrets) > 0 {
			request.MappedSecrets = resultSecrets
		}
	}

	if decryptionRules, ok := s.D.GetOkExists("decryption_rules"); ok {
		interfaces := decryptionRules.([]interface{})
		tmp := make([]oci_network_firewall.DecryptionRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "decryption_rules", stateDataIndex)
			converted, err := s.mapToDecryptionRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("decryption_rules") {
			request.DecryptionRules = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if securityRules, ok := s.D.GetOkExists("security_rules"); ok {
		interfaces := securityRules.([]interface{})
		tmp := make([]oci_network_firewall.SecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "security_rules", stateDataIndex)
			converted, err := s.mapToSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("security_rules") {
			request.SecurityRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) objectMapToDecryptionProfileMap(decryptionProfiles interface{}) (map[string]oci_network_firewall.DecryptionProfile, error) {

	resultMap := map[string]oci_network_firewall.DecryptionProfile{}
	tmpList := decryptionProfiles.([]interface{})
	i := 0
	for _, ifc := range tmpList {
		fmt.Printf("current index %q", ifc)
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "decryption_profiles", i)
		converted, err := s.mapToDecryptionProfile(fieldKeyFormat)
		if err != nil {
			return nil, err
		}

		if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
			tmp := type_.(string)
			switch tmp {
			case "SSL_INBOUND_INSPECTION":
				if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
					keyV := key.(string)
					resultMap[keyV] = converted
				}
			case "SSL_FORWARD_PROXY":
				if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
					keyV := key.(string)
					resultMap[keyV] = converted
				}
			}
		}
		i++
	}

	return resultMap, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) objectMapToMappedSecrets(mappedSecrets interface{}) (map[string]oci_network_firewall.MappedSecret, error) {

	resultMap := map[string]oci_network_firewall.MappedSecret{}
	tmpList := mappedSecrets.([]interface{})
	i := 0
	for _, ifc := range tmpList {
		fmt.Printf("Current secret %q", ifc)
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mapped_secrets", i)
		converted, err := s.mapToMappedSecret(fieldKeyFormat)
		if err != nil {
			return nil, err
		}

		if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
			tmp := type_.(string)
			switch tmp {
			case "SSL_INBOUND_INSPECTION":
				if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
					keyV := key.(string)
					resultMap[keyV] = converted
				}
			case "SSL_FORWARD_PROXY":
				if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
					keyV := key.(string)
					resultMap[keyV] = converted
				}
			}
		}
		i++
	}

	return resultMap, nil
}
func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToDecryptionProfile(fieldKeyFormat string) (oci_network_firewall.DecryptionProfile, error) {
	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		switch tmp {
		case "SSL_FORWARD_PROXY":
			result := oci_network_firewall.SslForwardProxyProfile{}
			if isExpiredCertificateBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_expired_certificate_blocked")); ok {
				tmp := isExpiredCertificateBlocked.(bool)
				result.IsExpiredCertificateBlocked = &tmp
			}
			if isUntrustedIssuerBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_untrusted_issuer_blocked")); ok {
				tmp := isUntrustedIssuerBlocked.(bool)
				result.IsUntrustedIssuerBlocked = &tmp
			}
			if isRevocationStatusTimeoutBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_revocation_status_timeout_blocked")); ok {
				tmp := isRevocationStatusTimeoutBlocked.(bool)
				result.IsRevocationStatusTimeoutBlocked = &tmp
			}
			if isUnsupportedVersionBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unsupported_version_blocked")); ok {
				tmp := isUnsupportedVersionBlocked.(bool)
				result.IsUnsupportedVersionBlocked = &tmp
			}
			if isUnsupportedCipherBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unsupported_cipher_blocked")); ok {
				tmp := isUnsupportedCipherBlocked.(bool)
				result.IsUnsupportedCipherBlocked = &tmp
			}
			if isUnknownRevocationStatusBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unknown_revocation_status_blocked")); ok {
				tmp := isUnknownRevocationStatusBlocked.(bool)
				result.IsUnknownRevocationStatusBlocked = &tmp
			}
			if areCertificateExtensionsRestricted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_certificate_extensions_restricted")); ok {
				tmp := areCertificateExtensionsRestricted.(bool)
				result.AreCertificateExtensionsRestricted = &tmp
			}
			if isAutoIncludeAltName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_auto_include_alt_name")); ok {
				tmp := isAutoIncludeAltName.(bool)
				result.IsAutoIncludeAltName = &tmp
			}
			if isOutOfCapacityBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_out_of_capacity_blocked")); ok {
				tmp := isOutOfCapacityBlocked.(bool)
				result.IsOutOfCapacityBlocked = &tmp
			}
			return result, nil
		case "SSL_INBOUND_INSPECTION":
			result := oci_network_firewall.SslInboundInspectionProfile{}
			if isUnsupportedVersionBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unsupported_version_blocked")); ok {
				tmp := isUnsupportedVersionBlocked.(bool)
				result.IsUnsupportedVersionBlocked = &tmp
			}
			if isUnsupportedCipherBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_unsupported_cipher_blocked")); ok {
				tmp := isUnsupportedCipherBlocked.(bool)
				result.IsUnsupportedCipherBlocked = &tmp
			}
			if isOutOfCapacityBlocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_out_of_capacity_blocked")); ok {
				tmp := isOutOfCapacityBlocked.(bool)
				result.IsOutOfCapacityBlocked = &tmp
			}
			return result, nil
		}
	}
	return nil, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToMappedSecret(fieldKeyFormat string) (oci_network_firewall.MappedSecret, error) {
	result := oci_network_firewall.VaultMappedSecret{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		switch tmp {
		case "SSL_INBOUND_INSPECTION":
			result.Type = "SSL_INBOUND_INSPECTION"
		case "SSL_FORWARD_PROXY":
			result.Type = "SSL_FORWARD_PROXY"
		}
	}
	if vaultSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_secret_id")); ok {
		tmp := vaultSecretId.(string)
		result.VaultSecretId = &tmp
	}
	if versionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_number")); ok {
		tmp := versionNumber.(int)
		result.VersionNumber = &tmp
	}
	return result, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) getNetworkFirewallPolicyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_firewall.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkFirewallPolicyId, err := networkFirewallPolicyWaitForWorkRequest(workId, "networkfirewallpolicy",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, networkFirewallPolicyId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_network_firewall.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*networkFirewallPolicyId)

	return s.Get()
}

func networkFirewallPolicyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "network_firewall", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_network_firewall.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func networkFirewallPolicyWaitForWorkRequest(wId *string, entityType string, action oci_network_firewall.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_firewall.NetworkFirewallClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_firewall")
	retryPolicy.ShouldRetryOperation = networkFirewallPolicyWorkRequestShouldRetryFunc(timeout)

	response := oci_network_firewall.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_network_firewall.OperationStatusInProgress),
			string(oci_network_firewall.OperationStatusAccepted),
			string(oci_network_firewall.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_network_firewall.OperationStatusSucceeded),
			string(oci_network_firewall.OperationStatusFailed),
			string(oci_network_firewall.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_network_firewall.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_network_firewall.OperationStatusFailed || response.Status == oci_network_firewall.OperationStatusCanceled {
		return nil, getErrorFromNetworkFirewallNetworkFirewallPolicyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromNetworkFirewallNetworkFirewallPolicyWorkRequest(client *oci_network_firewall.NetworkFirewallClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_network_firewall.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_network_firewall.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Get() error {
	request := oci_network_firewall.GetNetworkFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NetworkFirewallPolicy
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_network_firewall.UpdateNetworkFirewallPolicyRequest{}

	if applicationLists, ok := s.D.GetOkExists("application_lists"); ok {
		resultApplicationLists, err := s.objectMapToApplicationListsMap(applicationLists)
		if err != nil {
			return err
		}
		if len(resultApplicationLists) > 0 {
			request.ApplicationLists = resultApplicationLists
		}
	}

	if urlLists, ok := s.D.GetOkExists("url_lists"); ok {
		resultUrlLists, err := s.objectMapToUrlListsMap(urlLists)
		if err != nil {
			return err
		}
		if len(resultUrlLists) > 0 {
			request.UrlLists = resultUrlLists
		}
	}

	if decryptionProfiles, ok := s.D.GetOkExists("decryption_profiles"); ok {
		resultMap, err := s.objectMapToDecryptionProfileMap(decryptionProfiles)
		if err != nil {
			return err
		}
		if len(resultMap) > 0 {
			request.DecryptionProfiles = resultMap
		}
	}

	if decryptionRules, ok := s.D.GetOkExists("decryption_rules"); ok {
		interfaces := decryptionRules.([]interface{})
		tmp := make([]oci_network_firewall.DecryptionRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "decryption_rules", stateDataIndex)
			converted, err := s.mapToDecryptionRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("decryption_rules") {
			request.DecryptionRules = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if ipAddressLists, ok := s.D.GetOkExists("ip_address_lists"); ok {
		request.IpAddressLists = s.objectMapToStringListMap(ipAddressLists)
	}

	if mappedSecrets, ok := s.D.GetOkExists("mapped_secrets"); ok {
		resultSecrets, err := s.objectMapToMappedSecrets(mappedSecrets)
		if err != nil {
			return err
		}
		if len(resultSecrets) > 0 {
			request.MappedSecrets = resultSecrets
		}
	}

	tmp := s.D.Id()
	request.NetworkFirewallPolicyId = &tmp

	if securityRules, ok := s.D.GetOkExists("security_rules"); ok {
		interfaces := securityRules.([]interface{})
		tmp := make([]oci_network_firewall.SecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "security_rules", stateDataIndex)
			converted, err := s.mapToSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("security_rules") {
			request.SecurityRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNetworkFirewallPolicyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall"), oci_network_firewall.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteNetworkFirewallPolicyRequest{}

	tmp := s.D.Id()
	request.NetworkFirewallPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.DeleteNetworkFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkFirewallPolicyWaitForWorkRequest(workId, "networkfirewallpolicy",
		oci_network_firewall.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) SetData() error {
	if s.Res.ApplicationLists != nil {
		s.D.Set("application_lists", propertyApplicationListsToMap(s.Res.ApplicationLists))
	} else {
		s.D.Set("application_lists", nil)
	}

	if s.Res.UrlLists != nil {
		s.D.Set("url_lists", propertyUrlListsToMap(s.Res.UrlLists))
	} else {
		s.D.Set("url_lists", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DecryptionProfiles != nil {
		s.D.Set("decryption_profiles", DecryptionProfileMapToMap(s.Res.DecryptionProfiles))
	} else {
		s.D.Set("decryption_profiles", nil)
	}

	if s.Res.MappedSecrets != nil {
		s.D.Set("mapped_secrets", MappedSecretsToMap(s.Res.MappedSecrets))
	} else {
		s.D.Set("mapped_secrets", nil)
	}

	decryptionRules := []interface{}{}
	for _, item := range s.Res.DecryptionRules {
		decryptionRules = append(decryptionRules, DecryptionRuleToMap(item))
	}
	s.D.Set("decryption_rules", decryptionRules)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddressLists != nil {
		s.D.Set("ip_address_lists", ipAddressListsToMap(s.Res.IpAddressLists))
	} else {
		s.D.Set("ip_address_lists", nil)
	}

	if s.Res.IsFirewallAttached != nil {
		s.D.Set("is_firewall_attached", *s.Res.IsFirewallAttached)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	securityRules := []interface{}{}
	for _, item := range s.Res.SecurityRules {
		securityRules = append(securityRules, SecurityRuleToMap(item))
	}
	s.D.Set("security_rules", securityRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ipAddressListsToMap(ipAddressLists map[string][]string) []interface{} {
	var ipAddrLists []interface{}
	for listName, ipAddrs := range ipAddressLists {
		var ipAddrrObj = make(map[string]interface{})
		ipAddrrObj["ip_address_list_name"] = listName
		ipAddrrObj["ip_address_list_value"] = ipAddrs
		ipAddrLists = append(ipAddrLists, ipAddrrObj)
	}
	return ipAddrLists
}

func propertyUrlListsToMap(obj map[string][]oci_network_firewall.UrlPattern) []interface{} {

	var urls []interface{}
	for k, v := range obj {
		var urlListK = make(map[string]interface{})
		urlListK["url_list_name"] = k
		var urlLists []interface{}
		for _, item := range v {
			result := item.(map[string]interface{})
			var url = make(map[string]interface{})
			url["type"] = "SIMPLE"
			for key, value := range result {
				if key == "pattern" {
					url["pattern"] = value
				}
			}
			urlLists = append(urlLists, url)
		}
		urlListK["url_list_values"] = urlLists
		urls = append(urls, urlListK)
	}

	return urls
}

func propertyApplicationListsToMap(obj map[string][]oci_network_firewall.Application) []interface{} {

	var applications []interface{}
	for k, v := range obj {
		var applicationListK = make(map[string]interface{})
		applicationListK["application_list_name"] = k
		var applicationLists []interface{}
		for _, item := range v {
			var application = make(map[string]interface{})
			result := item.(map[string]interface{})
			for key, value := range result {
				if key == "type" {
					application["type"] = value
				} else if key == "icmpType" {
					application["icmp_type"] = value
				} else if key == "icmpCode" {
					application["icmp_code"] = value
				} else if key == "minimumPort" {
					application["minimum_port"] = value
				} else if key == "maximumPort" {
					application["maximum_port"] = value
				}
			}
			applicationLists = append(applicationLists, application)
		}
		applicationListK["application_values"] = applicationLists
		applications = append(applications, applicationListK)
	}
	return applications
}

func DecryptionProfileMapToMap(decryptionProfileMap map[string]oci_network_firewall.DecryptionProfile) []interface{} {
	profiles := []interface{}{}

	// This is because we model the API's map as a List for Terraform convenience
	for key, profile := range decryptionProfileMap {
		profileResultMap := DecryptionProfileToMap(profile)
		profileResultMap["key"] = key
		profiles = append(profiles, profileResultMap)
	}
	return profiles
}

func DecryptionProfileToMap(obj oci_network_firewall.DecryptionProfile) map[string]interface{} {
	result := map[string]interface{}{}

	switch v := (obj).(type) {
	case oci_network_firewall.SslForwardProxyProfile:
		result["is_expired_certificate_blocked"] = v.IsExpiredCertificateBlocked
		result["is_untrusted_issuer_blocked"] = v.IsUntrustedIssuerBlocked
		result["is_revocation_status_timeout_blocked"] = v.IsRevocationStatusTimeoutBlocked
		result["is_unsupported_version_blocked"] = v.IsUnsupportedVersionBlocked
		result["is_untrusted_issuer_blocked"] = v.IsUntrustedIssuerBlocked
		result["is_unknown_revocation_status_blocked"] = v.IsUnknownRevocationStatusBlocked
		result["are_certificate_extensions_restricted"] = v.AreCertificateExtensionsRestricted
		result["is_auto_include_alt_name"] = v.IsAutoIncludeAltName
		result["is_out_of_capacity_blocked"] = v.IsOutOfCapacityBlocked
		result["type"] = "SSL_FORWARD_PROXY"
	case oci_network_firewall.SslInboundInspectionProfile:
		result["is_unsupported_version_blocked"] = v.IsUnsupportedVersionBlocked
		result["is_out_of_capacity_blocked"] = v.IsOutOfCapacityBlocked
		result["is_unsupported_cipher_blocked"] = v.IsUnsupportedCipherBlocked
		result["type"] = "SSL_INBOUND_INSPECTION"
	}
	return result
}

func MappedSecretsToMap(mappedSecretMap map[string]oci_network_firewall.MappedSecret) []interface{} {
	secrets := []interface{}{}

	// This is because we model the API's map as a List for Terraform convenience
	for key, secret := range mappedSecretMap {
		profileResultMap := MappedSecretToMap(secret)
		profileResultMap["key"] = key
		secrets = append(secrets, profileResultMap)
	}

	return secrets
}

func MappedSecretToMap(obj oci_network_firewall.MappedSecret) map[string]interface{} {
	result := map[string]interface{}{}

	obj2 := obj.(oci_network_firewall.VaultMappedSecret)

	result["vault_secret_id"] = *obj2.VaultSecretId
	result["version_number"] = *obj2.VersionNumber
	if obj.GetType() == "SSL_INBOUND_INSPECTION" {
		result["type"] = "SSL_INBOUND_INSPECTION"
	} else {
		result["type"] = "SSL_FORWARD_PROXY"
	}
	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToDecryptionRule(fieldKeyFormat string) (oci_network_firewall.DecryptionRule, error) {
	result := oci_network_firewall.DecryptionRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_network_firewall.DecryptionRuleActionEnum(action.(string))
		if action == "DECRYPT" {
			if decryptionProfile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "decryption_profile")); ok {
				tmp := decryptionProfile.(string)
				result.DecryptionProfile = &tmp
			}
			if secret, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret")); ok {
				tmp := secret.(string)
				result.Secret = &tmp
			}
		}
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "condition"), 0)
			tmp, err := s.mapToDecryptionRuleMatchCriteria(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert condition, encountered error: %v", err)
			}
			result.Condition = &tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func DecryptionRuleToMap(obj oci_network_firewall.DecryptionRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Condition != nil {
		result["condition"] = []interface{}{DecryptionRuleMatchCriteriaToMap(obj.Condition)}
	}

	if obj.DecryptionProfile != nil {
		result["decryption_profile"] = string(*obj.DecryptionProfile)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Secret != nil {
		result["secret"] = string(*obj.Secret)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToDecryptionRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.DecryptionRuleMatchCriteria, error) {
	result := oci_network_firewall.DecryptionRuleMatchCriteria{}

	if destinations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destinations")); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destinations")) {
			result.Destinations = tmp
		}
	}

	if sources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
		interfaces := sources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
			result.Sources = tmp
		}
	}

	return result, nil
}

func DecryptionRuleMatchCriteriaToMap(obj *oci_network_firewall.DecryptionRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["destinations"] = obj.Destinations
	result["sources"] = obj.Sources

	return result
}

func NetworkFirewallPolicySummaryToMap(obj oci_network_firewall.NetworkFirewallPolicySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToSecurityRule(fieldKeyFormat string) (oci_network_firewall.SecurityRule, error) {
	result := oci_network_firewall.SecurityRule{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_network_firewall.SecurityRuleActionEnum(action.(string))
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "condition"), 0)
			tmp, err := s.mapToSecurityRuleMatchCriteria(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert condition, encountered error: %v", err)
			}
			result.Condition = &tmp
		}
	}

	if inspection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "inspection")); ok {
		result.Inspection = oci_network_firewall.SecurityRuleInspectionEnum(inspection.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func SecurityRuleToMap(obj oci_network_firewall.SecurityRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Condition != nil {
		result["condition"] = []interface{}{SecurityRuleMatchCriteriaToMap(obj.Condition)}
	}

	result["inspection"] = string(obj.Inspection)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) mapToSecurityRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.SecurityRuleMatchCriteria, error) {
	result := oci_network_firewall.SecurityRuleMatchCriteria{}

	if applications, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "applications")); ok {
		interfaces := applications.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "applications")) {
			result.Applications = tmp
		}
	}

	if destinations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destinations")); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destinations")) {
			result.Destinations = tmp
		}
	}

	if sources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
		interfaces := sources.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
			result.Sources = tmp
		}
	}

	if urls, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "urls")); ok {
		interfaces := urls.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "urls")) {
			result.Urls = tmp
		}
	}

	return result, nil
}

func SecurityRuleMatchCriteriaToMap(obj *oci_network_firewall.SecurityRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["applications"] = obj.Applications

	result["destinations"] = obj.Destinations

	result["sources"] = obj.Sources

	result["urls"] = obj.Urls

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_network_firewall.ChangeNetworkFirewallPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.NetworkFirewallPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.ChangeNetworkFirewallPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

type IpAddressObject struct {
	// Name of the IpAddr List
	IpAddressListName string `mandatory:"true" json:"ip_address_list_name"`

	// Version number of the secret to be used.
	IpAddressListValue []string `mandatory:"true" json:"ip_address_list_value"`
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) objectMapToStringListMap(obj interface{}) map[string][]string {
	resultMap := make(map[string][]string)
	IpAddrObjList := obj.([]interface{})

	for _, ipAddrObj := range IpAddrObjList {
		ipAddrActualObj := ipAddrObj.(map[string]interface{})
		ipAddrList := ipAddrActualObj["ip_address_list_value"].([]interface{})
		var parsedIpAddrList []string
		for ipAddr := range ipAddrList {
			if v, ok := ipAddrList[ipAddr].(string); ok {
				parsedIpAddrList = append(parsedIpAddrList, v)
			}
		}
		resultMap[ipAddrActualObj["ip_address_list_name"].(string)] = parsedIpAddrList
	}
	return resultMap
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) objectMapToUrlListsMap(urlLists interface{}) (map[string][]oci_network_firewall.UrlPattern, error) {
	resultMap := make(map[string][]oci_network_firewall.UrlPattern)
	urlList := urlLists.([]interface{})

	for _, list := range urlList {
		urlListK := list.(map[string]interface{})
		urlPatterns := urlListK["url_list_values"].([]interface{})
		listKey := urlListK["url_list_name"].(string)
		for _, v := range urlPatterns {
			urlPattern := oci_network_firewall.SimpleUrlPattern{}
			result := v.(map[string]interface{})
			tmp := result["pattern"].(string)
			urlPattern.Pattern = &tmp
			resultMap[listKey] = append(resultMap[listKey], urlPattern)
		}
	}
	return resultMap, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyResourceCrud) objectMapToApplicationListsMap(applicationLists interface{}) (map[string][]oci_network_firewall.Application, error) {
	resultMap := make(map[string][]oci_network_firewall.Application)
	applicationList := applicationLists.([]interface{})

	for _, list := range applicationList {
		applicationListK := list.(map[string]interface{})
		applications := applicationListK["application_values"].([]interface{})
		listKey := applicationListK["application_list_name"].(string)
		for _, v := range applications {
			application := v.(map[string]interface{})
			tmp := application["type"]
			switch tmp {
			case "ICMP":
				result := oci_network_firewall.IcmpApplication{}
				tmp := application["icmp_type"].(int)
				result.IcmpType = &tmp
				tmp1 := application["icmp_code"].(int)
				result.IcmpCode = &tmp1
				resultMap[listKey] = append(resultMap[listKey], result)
			case "ICMP6":
				result := oci_network_firewall.Icmp6Application{}
				tmp := application["icmp_type"].(int)
				result.IcmpType = &tmp
				tmp1 := application["icmp_code"].(int)
				result.IcmpCode = &tmp1
				resultMap[listKey] = append(resultMap[listKey], result)
			case "UDP":
				result := oci_network_firewall.UdpApplication{}
				tmp := application["minimum_port"].(int)
				result.MinimumPort = &tmp
				tmp1 := application["maximum_port"].(int)
				result.MaximumPort = &tmp1
				resultMap[listKey] = append(resultMap[listKey], result)
			case "TCP":
				result := oci_network_firewall.TcpApplication{}
				tmp := application["minimum_port"].(int)
				result.MinimumPort = &tmp
				tmp1 := application["maximum_port"].(int)
				result.MaximumPort = &tmp1
				resultMap[listKey] = append(resultMap[listKey], result)
			}
		}
	}
	return resultMap, nil
}
