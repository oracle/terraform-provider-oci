// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaFamilyDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsJavaFamily,
		Schema: map[string]*schema.Schema{
			"family_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"doc_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_of_support_life_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_supported_version": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"latest_release_artifacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"approximate_file_size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"architecture": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_content_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_file_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"download_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_type_detail": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"script_checksum_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"script_download_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"latest_release_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsJavaFamily(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaFamilyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaFamilyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetJavaFamilyResponse
}

func (s *JmsJavaFamilyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaFamilyDataSourceCrud) Get() error {
	request := oci_jms.GetJavaFamilyRequest{}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetJavaFamily(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaFamilyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaFamilyDataSource-", JmsJavaFamilyDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DocUrl != nil {
		s.D.Set("doc_url", *s.Res.DocUrl)
	}

	if s.Res.EndOfSupportLifeDate != nil {
		s.D.Set("end_of_support_life_date", s.Res.EndOfSupportLifeDate.String())
	}

	if s.Res.IsSupportedVersion != nil {
		s.D.Set("is_supported_version", *s.Res.IsSupportedVersion)
	}

	latestReleaseArtifacts := []interface{}{}
	for _, item := range s.Res.LatestReleaseArtifacts {
		latestReleaseArtifacts = append(latestReleaseArtifacts, JavaArtifactToMap(item))
	}
	s.D.Set("latest_release_artifacts", latestReleaseArtifacts)

	if s.Res.LatestReleaseVersion != nil {
		s.D.Set("latest_release_version", *s.Res.LatestReleaseVersion)
	}

	s.D.Set("support_type", s.Res.SupportType)

	return nil
}
