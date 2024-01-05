// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsSignResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsSign,
		Read:     readKmsSign,
		Delete:   deleteKmsSign,
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
			"message": {
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
			"key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"message_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsSign(d *schema.ResourceData, m interface{}) error {
	sync := &KmsSignResourceCrud{}
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

func readKmsSign(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteKmsSign(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KmsSignResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsCryptoClient
	Res                    *oci_kms.SignedData
	DisableNotFoundRetries bool
}

func (s *KmsSignResourceCrud) ID() string {
	return *s.Res.KeyId
}

func (s *KmsSignResourceCrud) Create() error {
	request := oci_kms.SignRequest{}

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
		request.MessageType = oci_kms.SignDataDetailsMessageTypeEnum(messageType.(string))
	}

	if signingAlgorithm, ok := s.D.GetOkExists("signing_algorithm"); ok {
		request.SigningAlgorithm = oci_kms.SignDataDetailsSigningAlgorithmEnum(signingAlgorithm.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.Sign(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SignedData
	return nil
}

func (s *KmsSignResourceCrud) SetData() error {
	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.KeyVersionId != nil {
		s.D.Set("key_version_id", *s.Res.KeyVersionId)
	}

	if s.Res.Signature != nil {
		s.D.Set("signature", *s.Res.Signature)
	}

	s.D.Set("signing_algorithm", s.Res.SigningAlgorithm)

	return nil
}
