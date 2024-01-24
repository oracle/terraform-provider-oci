// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package secrets

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_secrets "github.com/oracle/oci-go-sdk/v65/secrets"
)

func SecretsSecretbundleVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSecretsSecretbundleVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"secret_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_bundle_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"secret_id": {
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
						"version_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSecretsSecretbundleVersions(d *schema.ResourceData, m interface{}) error {
	sync := &SecretsSecretbundleVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecretsClient()

	return tfresource.ReadResource(sync)
}

type SecretsSecretbundleVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_secrets.SecretsClient
	Res    *oci_secrets.ListSecretBundleVersionsResponse
}

func (s *SecretsSecretbundleVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecretsSecretbundleVersionsDataSourceCrud) Get() error {
	request := oci_secrets.ListSecretBundleVersionsRequest{}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "secrets")

	response, err := s.Client.ListSecretBundleVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecretBundleVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SecretsSecretbundleVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SecretsSecretbundleVersionsDataSource-", SecretsSecretbundleVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		secretbundleVersion := map[string]interface{}{
			"secret_id": *r.SecretId,
		}

		secretbundleVersion["stages"] = r.Stages

		if r.TimeCreated != nil {
			secretbundleVersion["time_created"] = r.TimeCreated.String()
		}

		if r.TimeOfDeletion != nil {
			secretbundleVersion["time_of_deletion"] = r.TimeOfDeletion.String()
		}

		if r.TimeOfExpiry != nil {
			secretbundleVersion["time_of_expiry"] = r.TimeOfExpiry.String()
		}

		if r.VersionName != nil {
			secretbundleVersion["version_name"] = *r.VersionName
		}

		if r.VersionNumber != nil {
			secretbundleVersion["version_number"] = strconv.FormatInt(*r.VersionNumber, 10)
		}

		resources = append(resources, secretbundleVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, SecretsSecretbundleVersionsDataSource().Schema["secret_bundle_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("secret_bundle_versions", resources); err != nil {
		return err
	}

	return nil
}
