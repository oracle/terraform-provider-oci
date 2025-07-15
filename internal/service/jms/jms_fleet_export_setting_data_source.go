// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetExportSettingDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetExportSetting,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"export_data_filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"application_name_contains": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"application_name_equal_to": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"java_major_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"java_vendors": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"java_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"security_statuses": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"export_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"export_frequency": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"export_resources": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"export_setting_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cross_region_acknowledged": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"target_bucket_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_bucket_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_bucket_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetExportSetting(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetExportSettingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetExportSettingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetExportSettingResponse
}

func (s *JmsFleetExportSettingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetExportSettingDataSourceCrud) Get() error {
	request := oci_jms.GetExportSettingRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetExportSetting(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetExportSettingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetExportSettingDataSource-", JmsFleetExportSettingDataSource(), s.D))

	if s.Res.ExportDataFilters != nil {
		s.D.Set("export_data_filters", []interface{}{ExportDataFiltersToMap(s.Res.ExportDataFilters)})
	} else {
		s.D.Set("export_data_filters", nil)
	}

	s.D.Set("export_duration", s.Res.ExportDuration)

	s.D.Set("export_frequency", s.Res.ExportFrequency)

	s.D.Set("export_resources", s.Res.ExportResources)

	if s.Res.ExportSettingKey != nil {
		s.D.Set("export_setting_key", *s.Res.ExportSettingKey)
	}

	if s.Res.IsCrossRegionAcknowledged != nil {
		s.D.Set("is_cross_region_acknowledged", *s.Res.IsCrossRegionAcknowledged)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.TargetBucketName != nil {
		s.D.Set("target_bucket_name", *s.Res.TargetBucketName)
	}

	if s.Res.TargetBucketNamespace != nil {
		s.D.Set("target_bucket_namespace", *s.Res.TargetBucketNamespace)
	}

	if s.Res.TargetBucketRegion != nil {
		s.D.Set("target_bucket_region", *s.Res.TargetBucketRegion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func (s *JmsFleetExportSettingDataSourceCrud) mapToExportDataFilters(fieldKeyFormat string) (oci_jms.ExportDataFilters, error) {
	result := oci_jms.ExportDataFilters{}

	if applicationNameContains, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_name_contains")); ok {
		interfaces := applicationNameContains.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "application_name_contains")) {
			result.ApplicationNameContains = tmp
		}
	}

	if applicationNameEqualTo, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_name_equal_to")); ok {
		interfaces := applicationNameEqualTo.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "application_name_equal_to")) {
			result.ApplicationNameEqualTo = tmp
		}
	}

	if javaMajorVersions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "java_major_versions")); ok {
		interfaces := javaMajorVersions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "java_major_versions")) {
			result.JavaMajorVersions = tmp
		}
	}

	if javaVendors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "java_vendors")); ok {
		interfaces := javaVendors.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "java_vendors")) {
			result.JavaVendors = tmp
		}
	}

	if javaVersions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "java_versions")); ok {
		interfaces := javaVersions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "java_versions")) {
			result.JavaVersions = tmp
		}
	}

	if securityStatuses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "security_statuses")); ok {
		interfaces := securityStatuses.([]interface{})
		tmp := make([]oci_jms.JreSecurityStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_jms.JreSecurityStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "security_statuses")) {
			result.SecurityStatuses = tmp
		}
	}

	return result, nil
}

func ExportDataFiltersToMap(obj *oci_jms.ExportDataFilters) map[string]interface{} {
	result := map[string]interface{}{}

	result["application_name_contains"] = obj.ApplicationNameContains

	result["application_name_equal_to"] = obj.ApplicationNameEqualTo

	result["java_major_versions"] = obj.JavaMajorVersions

	result["java_vendors"] = obj.JavaVendors

	result["java_versions"] = obj.JavaVersions

	result["security_statuses"] = obj.SecurityStatuses

	return result
}
