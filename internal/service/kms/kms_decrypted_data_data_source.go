// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"
)

func KmsDecryptedDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDecryptedData,
		Schema: map[string]*schema.Schema{
			"associated_data": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"ciphertext": {
				Type:     schema.TypeString,
				Required: true,
			},
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"plaintext": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plaintext_checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDecryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &DecryptedDataDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("crypto_endpoint")
	if !ok {
		return fmt.Errorf("crypto_endpoint missing")
	}
	client, err := m.(*client.OracleClients).KmsCryptoClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type DecryptedDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsCryptoClient
	Res    *oci_kms.DecryptResponse
}

func (s *DecryptedDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DecryptedDataDataSourceCrud) Get() error {
	request := oci_kms.DecryptRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = utils.ObjectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if ciphertext, ok := s.D.GetOkExists("ciphertext"); ok {
		tmp := ciphertext.(string)
		request.Ciphertext = &tmp
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.Decrypt(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DecryptedDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	if s.Res.Plaintext != nil {
		s.D.Set("plaintext", *s.Res.Plaintext)
	}

	if s.Res.PlaintextChecksum != nil {
		s.D.Set("plaintext_checksum", *s.Res.PlaintextChecksum)
	}

	return nil
}
