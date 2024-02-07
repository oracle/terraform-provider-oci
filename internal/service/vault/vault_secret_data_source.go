// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VaultSecretDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["secret_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(VaultSecretResource(), fieldMap, readSingularVaultSecret)
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

	if s.Res.LastRotationTime != nil {
		s.D.Set("last_rotation_time", s.Res.LastRotationTime.String())
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.NextRotationTime != nil {
		s.D.Set("next_rotation_time", s.Res.NextRotationTime.String())
	}

	if s.Res.RotationConfig != nil {
		s.D.Set("rotation_config", []interface{}{RotationConfigToMap(s.Res.RotationConfig)})
	} else {
		s.D.Set("rotation_config", nil)
	}

	s.D.Set("rotation_status", s.Res.RotationStatus)

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
