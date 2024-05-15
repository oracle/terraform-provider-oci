// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubErratumDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubErratum,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"advisory_severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"advisory_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"classification_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_families": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_latest": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_families": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
									"is_mandatory_for_autonomous_linux": {
										Type:     schema.TypeBool,
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
				},
			},
			"references": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"related_cves": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"solution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"synopsis": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_issued": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsManagementHubErratum(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubErratumDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubErratumDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetErratumResponse
}

func (s *OsManagementHubErratumDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubErratumDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetErratumRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetErratum(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubErratumDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubErratumDataSource-", OsManagementHubErratumDataSource(), s.D))

	s.D.Set("advisory_severity", s.Res.AdvisorySeverity)

	s.D.Set("advisory_type", s.Res.AdvisoryType)

	s.D.Set("classification_type", s.Res.ClassificationType)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.From != nil {
		s.D.Set("from", *s.Res.From)
	}

	s.D.Set("os_families", s.Res.OsFamilies)

	packages := []interface{}{}
	for _, item := range s.Res.Packages {
		packages = append(packages, SoftwarePackageSummaryToMap(item))
	}
	s.D.Set("packages", packages)

	if s.Res.References != nil {
		s.D.Set("references", *s.Res.References)
	}

	s.D.Set("related_cves", s.Res.RelatedCves)

	s.D.Set("repositories", s.Res.Repositories)

	if s.Res.Solution != nil {
		s.D.Set("solution", *s.Res.Solution)
	}

	if s.Res.Synopsis != nil {
		s.D.Set("synopsis", *s.Res.Synopsis)
	}

	if s.Res.TimeIssued != nil {
		s.D.Set("time_issued", s.Res.TimeIssued.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
