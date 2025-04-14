// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dblm

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dblm "github.com/oracle/oci-go-sdk/v65/dblm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DblmPatchManagementDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDblmPatchManagement,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_release": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"images_patch_recommendation_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"total_images_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"up_to_date_images_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"image_patch_recommendations_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_operations_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"scheduled_patch_ops_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"running_patch_ops_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"successful_patch_ops_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"warnings_patch_ops_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"failed_patch_ops_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"agent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connector_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_platform_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_info": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"host_cores": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_cluster_db": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"license_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"resources_patch_compliance_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"total_resources_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"up_to_date_resources_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"non_compliant_resources_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"not_subscribed_resources_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"not_dblm_registered_resources_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDblmPatchManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DblmPatchManagementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbLifeCycleManagementClient()

	return tfresource.ReadResource(sync)
}

type DblmPatchManagementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dblm.DbLifeCycleManagementClient
	Res    *oci_dblm.GetPatchManagementResponse
}

func (s *DblmPatchManagementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DblmPatchManagementDataSourceCrud) Get() error {
	request := oci_dblm.GetPatchManagementRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseRelease, ok := s.D.GetOkExists("database_release"); ok {
		tmp := databaseRelease.(string)
		request.DatabaseRelease = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dblm.DblmVulnerabilityLifecycleStateEnum(state.(string))
	}

	if timeStartedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_started_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeStartedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeStartedLessThan, ok := s.D.GetOkExists("time_started_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStartedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeStartedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dblm")

	response, err := s.Client.GetPatchManagement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DblmPatchManagementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DblmPatchManagementDataSource-", DblmPatchManagementDataSource(), s.D))

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	resources_images_patch := []interface{}{}
	for _, item := range s.Res.Resources {
		resources_images_patch = append(resources_images_patch, DblmResourceInfoToMap(item))
	}
	s.D.Set("images_patch_recommendation_summary", resources_images_patch)

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	resources_patch_operations := []interface{}{}
	for _, item := range s.Res.Resources {
		resources_patch_operations = append(resources_patch_operations, DblmResourceInfoToMap(item))
	}
	s.D.Set("patch_operations_summary", resources_patch_operations)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, DblmResourceInfoToMap(item))
	}
	s.D.Set("resources", resources)

	resources_patch_compliance := []interface{}{}
	for _, item := range s.Res.Resources {
		resources_patch_compliance = append(resources_patch_compliance, DblmResourceInfoToMap(item))
	}
	s.D.Set("resources_patch_compliance_summary", resources_patch_compliance)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeEnabled != nil {
		s.D.Set("time_enabled", s.Res.TimeEnabled.String())
	}

	return nil
}

func DblmHostInfoToMap(obj oci_dblm.HostInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HostCores != nil {
		result["host_cores"] = int(*obj.HostCores)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	return result
}

func DblmResourceInfoToMap(obj oci_dblm.ResourceInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	if obj.DbPlatformType != nil {
		result["db_platform_type"] = string(*obj.DbPlatformType)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DeploymentType != nil {
		result["deployment_type"] = string(*obj.DeploymentType)
	}

	hostInfo := []interface{}{}
	for _, item := range obj.HostInfo {
		hostInfo = append(hostInfo, HostInfoToMap(item))
	}
	result["host_info"] = hostInfo

	if obj.IsClusterDb != nil {
		result["is_cluster_db"] = bool(*obj.IsClusterDb)
	}

	if obj.LicenseType != nil {
		result["license_type"] = string(*obj.LicenseType)
	}

	if obj.ResourceCompartmentId != nil {
		result["resource_compartment_id"] = string(*obj.ResourceCompartmentId)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}
