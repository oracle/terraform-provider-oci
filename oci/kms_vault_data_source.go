// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KmsVaultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularKmsVault,
		Schema: map[string]*schema.Schema{
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularKmsVault(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", *s.Res.TimeOfDeletion)
	}

	s.D.Set("vault_type", s.Res.VaultType)

	return nil
}
