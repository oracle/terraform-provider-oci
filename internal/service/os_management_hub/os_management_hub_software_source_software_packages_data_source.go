// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwareSourceSoftwarePackagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSourceSoftwarePackages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_latest": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_package_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
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
									"os_families": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
					},
				},
			},
		},
	}
}

func readOsManagementHubSoftwareSourceSoftwarePackages(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceSoftwarePackagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceSoftwarePackagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListSoftwarePackagesResponse
}

func (s *OsManagementHubSoftwareSourceSoftwarePackagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceSoftwarePackagesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListSoftwarePackagesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if isLatest, ok := s.D.GetOkExists("is_latest"); ok {
		tmp := isLatest.(bool)
		request.IsLatest = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListSoftwarePackages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSoftwarePackages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubSoftwareSourceSoftwarePackagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourceSoftwarePackagesDataSource-", OsManagementHubSoftwareSourceSoftwarePackagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSourceSoftwarePackage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SoftwarePackageSummaryToMap(item))
	}
	softwareSourceSoftwarePackage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourceSoftwarePackagesDataSource().Schema["software_package_collection"].Elem.(*schema.Resource).Schema)
		softwareSourceSoftwarePackage["items"] = items
	}

	resources = append(resources, softwareSourceSoftwarePackage)
	if err := s.D.Set("software_package_collection", resources); err != nil {
		return err
	}

	return nil
}
