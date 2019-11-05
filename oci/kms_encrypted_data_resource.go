// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KmsEncryptedDataResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createKmsEncryptedData,
		Read:     readKmsEncryptedData,
		Delete:   deleteKmsEncryptedData,
		Schema: map[string]*schema.Schema{
			// Required
			"crypto_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"plaintext": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"associated_data": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"logging_context": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"ciphertext": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEncryptedDataResourceCrud{}
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

	return CreateResource(d, sync)
}

func readKmsEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEncryptedDataResourceCrud{}
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

func deleteKmsEncryptedData(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KmsEncryptedDataResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.EncryptedData
	DisableNotFoundRetries bool
}

func (s *KmsEncryptedDataResourceCrud) ID() string {
	return string(hashcode.String(*s.Res.Ciphertext))
}

func (s *KmsEncryptedDataResourceCrud) Create() error {
	request := oci_kms.EncryptRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = objectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if loggingContext, ok := s.D.GetOkExists("logging_context"); ok {
		request.LoggingContext = objectMapToStringMap(loggingContext.(map[string]interface{}))
	}

	if plaintext, ok := s.D.GetOkExists("plaintext"); ok {
		tmp := plaintext.(string)
		request.Plaintext = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.Encrypt(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EncryptedData
	return nil
}

func (s *KmsEncryptedDataResourceCrud) Get() error {

	if cipherText, ok := s.D.GetOkExists("ciphertext"); ok {
		tmp := cipherText.(string)
		encryptedData := oci_kms.EncryptedData{Ciphertext: &tmp}
		s.Res = &encryptedData
	} else {
		return s.Create()
	}

	return nil
}

func (s *KmsEncryptedDataResourceCrud) SetData() error {
	if s.Res.Ciphertext != nil {
		s.D.Set("ciphertext", *s.Res.Ciphertext)
	}

	return nil
}
