// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v56/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func KmsVaultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vault_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(KmsVaultResource(), fieldMap, readSingularKmsVault)
}

func readSingularKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.ManagementEndpoint != nil {
		s.D.Set("management_endpoint", *s.Res.ManagementEndpoint)
	}

	if s.Res.ReplicaDetails != nil {
		s.D.Set("replica_details", []interface{}{VaultReplicaDetailsToMap(s.Res.ReplicaDetails)})
	} else {
		s.D.Set("replica_details", nil)
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
