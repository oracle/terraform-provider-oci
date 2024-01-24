// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vault

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"
)

func VaultSecretVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularVaultSecretVersion,
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_version_number": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"version_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularVaultSecretVersion(d *schema.ResourceData, m interface{}) error {
	sync := &VaultSecretVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VaultsClient()

	return tfresource.ReadResource(sync)
}

type VaultSecretVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_vault.VaultsClient
	Res    *oci_vault.GetSecretVersionResponse
}

func (s *VaultSecretVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VaultSecretVersionDataSourceCrud) Get() error {
	request := oci_vault.GetSecretVersionRequest{}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
		tmp := secretVersionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert secretVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.SecretVersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "vault")

	response, err := s.Client.GetSecretVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VaultSecretVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("VaultSecretVersionDataSource-", VaultSecretVersionDataSource(), s.D))

	s.D.Set("content_type", s.Res.ContentType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("stages", s.Res.Stages)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfCurrentVersionExpiry != nil {
		s.D.Set("time_of_current_version_expiry", s.Res.TimeOfCurrentVersionExpiry.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.VersionNumber != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	return nil
}
