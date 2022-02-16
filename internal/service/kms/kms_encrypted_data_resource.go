// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"
)

func KmsEncryptedDataResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
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
			"encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
	client, err := m.(*client.OracleClients).KmsCryptoClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.CreateResource(d, sync)
}

func readKmsEncryptedData(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEncryptedDataResourceCrud{}
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

func deleteKmsEncryptedData(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KmsEncryptedDataResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.EncryptedData
	DisableNotFoundRetries bool
}

func (s *KmsEncryptedDataResourceCrud) ID() string {
	return fmt.Sprint(hashcode.String(*s.Res.Ciphertext))
}

func (s *KmsEncryptedDataResourceCrud) Create() error {
	request := oci_kms.EncryptRequest{}

	if associatedData, ok := s.D.GetOkExists("associated_data"); ok {
		request.AssociatedData = utils.ObjectMapToStringMap(associatedData.(map[string]interface{}))
	}

	if encryptionAlgorithm, ok := s.D.GetOkExists("encryption_algorithm"); ok {
		request.EncryptionAlgorithm = oci_kms.EncryptDataDetailsEncryptionAlgorithmEnum(encryptionAlgorithm.(string))
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyVersionId, ok := s.D.GetOkExists("key_version_id"); ok {
		tmp := keyVersionId.(string)
		request.KeyVersionId = &tmp
	}

	if loggingContext, ok := s.D.GetOkExists("logging_context"); ok {
		request.LoggingContext = utils.ObjectMapToStringMap(loggingContext.(map[string]interface{}))
	}

	if plaintext, ok := s.D.GetOkExists("plaintext"); ok {
		tmp := plaintext.(string)
		request.Plaintext = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

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

	s.D.Set("encryption_algorithm", s.Res.EncryptionAlgorithm)

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.KeyVersionId != nil {
		s.D.Set("key_version_id", *s.Res.KeyVersionId)
	}

	return nil
}
