// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v29/keymanagement"
)

func init() {
	RegisterDatasource("oci_kms_encrypted_data", KmsEncryptedDataDataSource())
}

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
	client, err := m.(*OracleClients).KmsCryptoClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
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
		request.AssociatedData = objectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if plaintext, ok := s.D.GetOkExists("plaintext"); ok {
		tmp := plaintext.(string)
		request.Plaintext = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

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

	s.D.SetId(GenerateDataSourceID())

	if s.Res.Ciphertext != nil {
		s.D.Set("ciphertext", *s.Res.Ciphertext)
	}

	return nil
}
