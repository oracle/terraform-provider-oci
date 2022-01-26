// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_vault "github.com/oracle/oci-go-sdk/v56/vault"
)

func VaultSecretDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularVaultSecret,
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"secret_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"is_enforced_on_deleted_secret_versions": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_secret_content_retrieval_blocked_on_expiry": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"rule_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret_version_expiry_interval": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_absolute_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_current_version_expiry": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularVaultSecret(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.ReadResource(sync)
}

type VaultSecretDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vault.VaultsClient
	Res    *oci_vault.GetSecretResponse
}

func (s *VaultSecretDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VaultSecretDataSourceCrud) Get() error {
	request := oci_vault.GetSecretRequest{}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vault")

	response, err := s.Client.GetSecret(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VaultSecretDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentVersionNumber != nil {
		s.D.Set("current_version_number", strconv.FormatInt(*s.Res.CurrentVersionNumber, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.SecretName != nil {
		s.D.Set("secret_name", *s.Res.SecretName)
	}

	secretRules := []interface{}{}
	for _, item := range s.Res.SecretRules {
		secretRules = append(secretRules, SecretRuleToMap(item))
	}
	s.D.Set("secret_rules", secretRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfCurrentVersionExpiry != nil {
		s.D.Set("time_of_current_version_expiry", s.Res.TimeOfCurrentVersionExpiry.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func SecretRuleToMap(obj oci_vault.SecretRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_vault.SecretExpiryRule:
		result["rule_type"] = "SECRET_EXPIRY_RULE"

		if v.IsSecretContentRetrievalBlockedOnExpiry != nil {
			result["is_secret_content_retrieval_blocked_on_expiry"] = bool(*v.IsSecretContentRetrievalBlockedOnExpiry)
		}

		if v.SecretVersionExpiryInterval != nil {
			result["secret_version_expiry_interval"] = string(*v.SecretVersionExpiryInterval)
		}

		if v.TimeOfAbsoluteExpiry != nil {
			result["time_of_absolute_expiry"] = v.TimeOfAbsoluteExpiry.Format(time.RFC3339Nano)
		}
	case oci_vault.SecretReuseRule:
		result["rule_type"] = "SECRET_REUSE_RULE"

		if v.IsEnforcedOnDeletedSecretVersions != nil {
			result["is_enforced_on_deleted_secret_versions"] = bool(*v.IsEnforcedOnDeletedSecretVersions)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}
