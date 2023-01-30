// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaReleaseDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsJavaRelease,
		Schema: map[string]*schema.Schema{
			"release_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"artifact_content_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"artifacts": {
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
						"artifact_content_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_id": {
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
			"family_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"family_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"support_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"family_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"license_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_release_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_notes_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsJavaRelease(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaReleaseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaReleaseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetJavaReleaseResponse
}

func (s *JmsJavaReleaseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaReleaseDataSourceCrud) Get() error {
	request := oci_jms.GetJavaReleaseRequest{}

	if releaseVersion, ok := s.D.GetOkExists("release_version"); ok {
		tmp := releaseVersion.(string)
		request.ReleaseVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetJavaRelease(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaReleaseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaReleaseDataSource-", JmsJavaReleaseDataSource(), s.D))

	s.D.Set("artifact_content_types", s.Res.ArtifactContentTypes)
	s.D.Set("artifact_content_types", s.Res.ArtifactContentTypes)

	artifacts := []interface{}{}
	for _, item := range s.Res.Artifacts {
		artifacts = append(artifacts, JavaArtifactToMap(item))
	}
	s.D.Set("artifacts", artifacts)

	if s.Res.FamilyDetails != nil {
		s.D.Set("family_details", []interface{}{JavaFamilyToMap(s.Res.FamilyDetails)})
	} else {
		s.D.Set("family_details", nil)
	}

	if s.Res.FamilyVersion != nil {
		s.D.Set("family_version", *s.Res.FamilyVersion)
	}

	if s.Res.LicenseDetails != nil {
		s.D.Set("license_details", []interface{}{JavaLicenseToMap(s.Res.LicenseDetails)})
	} else {
		s.D.Set("license_details", nil)
	}

	s.D.Set("license_type", s.Res.LicenseType)

	if s.Res.ParentReleaseVersion != nil {
		s.D.Set("parent_release_version", *s.Res.ParentReleaseVersion)
	}

	if s.Res.ReleaseDate != nil {
		s.D.Set("release_date", s.Res.ReleaseDate.String())
	}

	if s.Res.ReleaseNotesUrl != nil {
		s.D.Set("release_notes_url", *s.Res.ReleaseNotesUrl)
	}

	s.D.Set("release_type", s.Res.ReleaseType)

	s.D.Set("security_status", s.Res.SecurityStatus)

	return nil
}
