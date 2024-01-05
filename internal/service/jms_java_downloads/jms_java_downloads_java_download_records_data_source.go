// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaDownloadsJavaDownloadRecords,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"architecture": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"family_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_type_detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_download_record_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
									"download_source_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"download_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"family_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"family_version": {
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
									"release_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_downloaded": {
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

func readJmsJavaDownloadsJavaDownloadRecords(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.ListJavaDownloadRecordsResponse
}

func (s *JmsJavaDownloadsJavaDownloadRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadRecordsDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.ListJavaDownloadRecordsRequest{}

	if architecture, ok := s.D.GetOkExists("architecture"); ok {
		tmp := architecture.(string)
		request.Architecture = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if familyVersion, ok := s.D.GetOkExists("family_version"); ok {
		tmp := familyVersion.(string)
		request.FamilyVersion = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		tmp := osFamily.(string)
		request.OsFamily = &tmp
	}

	if packageTypeDetail, ok := s.D.GetOkExists("package_type_detail"); ok {
		tmp := packageTypeDetail.(string)
		request.PackageTypeDetail = &tmp
	}

	if releaseVersion, ok := s.D.GetOkExists("release_version"); ok {
		tmp := releaseVersion.(string)
		request.ReleaseVersion = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.ListJavaDownloadRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaDownloadRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaDownloadsJavaDownloadRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaDownloadRecordsDataSource-", JmsJavaDownloadsJavaDownloadRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaDownloadRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaDownloadRecordSummaryToMap(item))
	}
	javaDownloadRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaDownloadsJavaDownloadRecordsDataSource().Schema["java_download_record_collection"].Elem.(*schema.Resource).Schema)
		javaDownloadRecord["items"] = items
	}

	resources = append(resources, javaDownloadRecord)
	if err := s.D.Set("java_download_record_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaDownloadRecordSummaryToMap(obj oci_jms_java_downloads.JavaDownloadRecordSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Architecture != nil {
		result["architecture"] = string(*obj.Architecture)
	}

	if obj.DownloadSourceId != nil {
		result["download_source_id"] = string(*obj.DownloadSourceId)
	}

	if obj.DownloadType != nil {
		result["download_type"] = string(*obj.DownloadType)
	}

	if obj.FamilyDisplayName != nil {
		result["family_display_name"] = string(*obj.FamilyDisplayName)
	}

	if obj.FamilyVersion != nil {
		result["family_version"] = string(*obj.FamilyVersion)
	}

	if obj.OsFamily != nil {
		result["os_family"] = string(*obj.OsFamily)
	}

	if obj.PackageType != nil {
		result["package_type"] = string(*obj.PackageType)
	}

	if obj.PackageTypeDetail != nil {
		result["package_type_detail"] = string(*obj.PackageTypeDetail)
	}

	if obj.ReleaseVersion != nil {
		result["release_version"] = string(*obj.ReleaseVersion)
	}

	if obj.TimeDownloaded != nil {
		result["time_downloaded"] = obj.TimeDownloaded.String()
	}

	return result
}
