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

func JmsJavaFamiliesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaFamilies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"family_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_supported_version": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"java_family_collection": {
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
									"release_date": {
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
					},
				},
			},
		},
	}
}

func readJmsJavaFamilies(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaFamiliesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaFamiliesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListJavaFamiliesResponse
}

func (s *JmsJavaFamiliesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaFamiliesDataSourceCrud) Get() error {
	request := oci_jms.ListJavaFamiliesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	if isSupportedVersion, ok := s.D.GetOkExists("is_supported_version"); ok {
		tmp := isSupportedVersion.(bool)
		request.IsSupportedVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListJavaFamilies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaFamilies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaFamiliesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaFamiliesDataSource-", JmsJavaFamiliesDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaFamily := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaFamilySummaryToMap(item))
	}
	javaFamily["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaFamiliesDataSource().Schema["java_family_collection"].Elem.(*schema.Resource).Schema)
		javaFamily["items"] = items
	}

	resources = append(resources, javaFamily)
	if err := s.D.Set("java_family_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaFamilySummaryToMap(obj oci_jms.JavaFamilySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DocUrl != nil {
		result["doc_url"] = string(*obj.DocUrl)
	}

	if obj.EndOfSupportLifeDate != nil {
		result["end_of_support_life_date"] = obj.EndOfSupportLifeDate.String()
	}

	if obj.FamilyVersion != nil {
		result["family_version"] = string(*obj.FamilyVersion)
	}

	if obj.IsSupportedVersion != nil {
		result["is_supported_version"] = bool(*obj.IsSupportedVersion)
	}

	if obj.LatestReleaseVersion != nil {
		result["latest_release_version"] = string(*obj.LatestReleaseVersion)
	}

	if obj.ReleaseDate != nil {
		result["release_date"] = obj.ReleaseDate.String()
	}

	result["support_type"] = string(obj.SupportType)

	return result
}
