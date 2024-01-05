// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package secrets

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_secrets "github.com/oracle/oci-go-sdk/v65/secrets"
)

func SecretsSecretbundleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularSecretsSecretbundle,
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_version_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stage": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Computed
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"secret_bundle_content": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"content_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_of_deletion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_expiry": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularSecretsSecretbundle(d *schema.ResourceData, m interface{}) error {
	sync := &SecretsSecretbundleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecretsClient()

	return tfresource.ReadResource(sync)
}

type SecretsSecretbundleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_secrets.SecretsClient
	Res    *oci_secrets.GetSecretBundleResponse
}

func (s *SecretsSecretbundleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecretsSecretbundleDataSourceCrud) Get() error {
	request := oci_secrets.GetSecretBundleRequest{}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	if secretVersionName, ok := s.D.GetOkExists("secret_version_name"); ok {
		tmp := secretVersionName.(string)
		request.SecretVersionName = &tmp
	}

	if stage, ok := s.D.GetOkExists("stage"); ok {
		request.Stage = oci_secrets.GetSecretBundleStageEnum(stage.(string))
	}

	if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
		tmp := versionNumber.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert versionNumber string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.VersionNumber = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "secrets")

	response, err := s.Client.GetSecretBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SecretsSecretbundleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SecretsSecretbundleDataSource-", SecretsSecretbundleDataSource(), s.D))

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.SecretBundleContent != nil {
		secretBundleContentArray := []interface{}{}
		if secretBundleContentMap := SecretBundleContentDetailsToMap(&s.Res.SecretBundleContent); secretBundleContentMap != nil {
			secretBundleContentArray = append(secretBundleContentArray, secretBundleContentMap)
		}
		s.D.Set("secret_bundle_content", secretBundleContentArray)
	} else {
		s.D.Set("secret_bundle_content", nil)
	}

	s.D.Set("stages", s.Res.Stages)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.TimeOfExpiry != nil {
		s.D.Set("time_of_expiry", s.Res.TimeOfExpiry.String())
	}

	if s.Res.VersionName != nil {
		s.D.Set("version_name", *s.Res.VersionName)
	}

	if s.Res.VersionNumber != nil {
		s.D.Set("version_number", strconv.FormatInt(*s.Res.VersionNumber, 10))
	}

	return nil
}

func SecretBundleContentDetailsToMap(obj *oci_secrets.SecretBundleContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_secrets.Base64SecretBundleContentDetails:
		result["content_type"] = "BASE64"

		if v.Content != nil {
			result["content"] = string(*v.Content)
		}
	default:
		log.Printf("[WARN] Received 'content_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
