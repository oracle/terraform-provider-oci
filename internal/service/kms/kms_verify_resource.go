// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"
)

func KmsVerifyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsVerify,
		Read:     readKmsVerify,
		Delete:   deleteKmsVerify,
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
			"key_version_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"message": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signing_algorithm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"message_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"is_signature_valid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func createKmsVerify(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVerifyResourceCrud{}
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

func readKmsVerify(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteKmsVerify(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KmsVerifyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.VerifiedData
	DisableNotFoundRetries bool
}

func (s *KmsVerifyResourceCrud) ID() string {
	return s.D.Get("signature").(string)
}

func (s *KmsVerifyResourceCrud) Create() error {
	request := oci_kms.VerifyRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyVersionId, ok := s.D.GetOkExists("key_version_id"); ok {
		tmp := keyVersionId.(string)
		request.KeyVersionId = &tmp
	}

	if message, ok := s.D.GetOkExists("message"); ok {
		tmp := message.(string)
		request.Message = &tmp
	}

	if messageType, ok := s.D.GetOkExists("message_type"); ok {
		request.MessageType = oci_kms.VerifyDataDetailsMessageTypeEnum(messageType.(string))
	}

	if signature, ok := s.D.GetOkExists("signature"); ok {
		tmp := signature.(string)
		request.Signature = &tmp
	}

	if signingAlgorithm, ok := s.D.GetOkExists("signing_algorithm"); ok {
		request.SigningAlgorithm = oci_kms.VerifyDataDetailsSigningAlgorithmEnum(signingAlgorithm.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.Verify(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.VerifiedData
	return nil
}

func (s *KmsVerifyResourceCrud) SetData() error {
	if s.Res.IsSignatureValid != nil {
		s.D.Set("is_signature_valid", *s.Res.IsSignatureValid)
	}

	return nil
}
