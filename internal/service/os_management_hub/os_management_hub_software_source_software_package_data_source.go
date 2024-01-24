// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceSoftwarePackageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubSoftwareSourceSoftwarePackage,
		Schema: map[string]*schema.Schema{
			"software_package_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checksum_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dependencies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dependency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dependency_modifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dependency_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"checksum": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"checksum_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_latest": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_modified_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsManagementHubSoftwareSourceSoftwarePackage(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceSoftwarePackageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceSoftwarePackageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetSoftwarePackageResponse
}

func (s *OsManagementHubSoftwareSourceSoftwarePackageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceSoftwarePackageDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetSoftwarePackageRequest{}

	if softwarePackageName, ok := s.D.GetOkExists("software_package_name"); ok {
		tmp := softwarePackageName.(string)
		request.SoftwarePackageName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetSoftwarePackage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubSoftwareSourceSoftwarePackageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceSoftwarePackageDataSource-", OsManagementHubSoftwareSourceSoftwarePackageDataSource(), s.D))

	if s.Res.Architecture != nil {
		s.D.Set("architecture", *s.Res.Architecture)
	}

	if s.Res.Checksum != nil {
		s.D.Set("checksum", *s.Res.Checksum)
	}

	if s.Res.ChecksumType != nil {
		s.D.Set("checksum_type", *s.Res.ChecksumType)
	}

	dependencies := []interface{}{}
	for _, item := range s.Res.Dependencies {
		dependencies = append(dependencies, SoftwarePackageDependencyToMap(item))
	}
	s.D.Set("dependencies", dependencies)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	files := []interface{}{}
	for _, item := range s.Res.Files {
		files = append(files, SoftwarePackageFileToMap(item))
	}
	s.D.Set("files", files)

	if s.Res.IsLatest != nil {
		s.D.Set("is_latest", *s.Res.IsLatest)
	}

	if s.Res.LastModifiedDate != nil {
		s.D.Set("last_modified_date", *s.Res.LastModifiedDate)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	softwareSources := []interface{}{}
	for _, item := range s.Res.SoftwareSources {
		softwareSources = append(softwareSources, SoftwareSourceDetailsToMap(item))
	}
	s.D.Set("software_sources", softwareSources)

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
