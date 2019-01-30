// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KeyVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularKeyVersion,
		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_version_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
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

func readSingularKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KeyVersionDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

type KeyVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.GetKeyVersionResponse
}

func (s *KeyVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KeyVersionDataSourceCrud) Get() error {
	request := oci_kms.GetKeyVersionRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyVersionId, ok := s.D.GetOkExists("key_version_id"); ok {
		tmp := keyVersionId.(string)
		request.KeyVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.GetKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KeyVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(getKeyVersionCompositeId(*s.Res.KeyId, *s.Res.Id))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}
