// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"
)

func KmsVaultUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularKmsVaultUsage,
		Schema: map[string]*schema.Schema{
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"key_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key_version_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_key_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_key_version_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularKmsVaultUsage(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

type KmsVaultUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.GetVaultUsageResponse
}

func (s *KmsVaultUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsVaultUsageDataSourceCrud) Get() error {
	request := oci_kms.GetVaultUsageRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.GetVaultUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsVaultUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsVaultUsageDataSource-", KmsVaultUsageDataSource(), s.D))

	if s.Res.KeyCount != nil {
		s.D.Set("key_count", *s.Res.KeyCount)
	}

	if s.Res.KeyVersionCount != nil {
		s.D.Set("key_version_count", *s.Res.KeyVersionCount)
	}

	if s.Res.SoftwareKeyCount != nil {
		s.D.Set("software_key_count", *s.Res.SoftwareKeyCount)
	}

	if s.Res.SoftwareKeyVersionCount != nil {
		s.D.Set("software_key_version_count", *s.Res.SoftwareKeyVersionCount)
	}

	return nil
}
