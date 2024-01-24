// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsEncryptedDataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularEncryptedData,
		Schema: map[string]*schema.Schema{
			"associated_data": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plaintext": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ciphertext": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &EncryptedDataDataSourceCrud{}
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

type EncryptedDataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsCryptoClient
	Res    *oci_kms.EncryptResponse
}

func (s *EncryptedDataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EncryptedDataDataSourceCrud) Get() error {
	request := oci_kms.EncryptRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = tfresource.ObjectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if plaintext, ok := s.D.GetOkExists("plaintext"); ok {
		tmp := plaintext.(string)
		request.Plaintext = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.Encrypt(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EncryptedDataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	if s.Res.Ciphertext != nil {
		s.D.Set("ciphertext", *s.Res.Ciphertext)
	}

	return nil
}
