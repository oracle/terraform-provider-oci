// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v34/keymanagement"
)

func init() {
	RegisterDatasource("oci_kms_vault", KmsVaultDataSource())
}

func KmsVaultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vault_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(KmsVaultResource(), fieldMap, readSingularKmsVault)
}

func readSingularKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient()

	return ReadResource(sync)
}

type KmsVaultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.GetVaultResponse
}

func (s *KmsVaultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsVaultDataSourceCrud) Get() error {
	request := oci_kms.GetVaultRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.GetVault(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsVaultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CryptoEndpoint != nil {
		s.D.Set("crypto_endpoint", *s.Res.CryptoEndpoint)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	if s.Res.RestoredFromVaultId != nil {
		s.D.Set("restored_from_vault_id", *s.Res.RestoredFromVaultId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	s.D.Set("vault_type", s.Res.VaultType)

	return nil
}
