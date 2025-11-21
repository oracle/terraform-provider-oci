// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardAdhocQueryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardAdhocQuery,
		Read:     readCloudGuardAdhocQuery,
		Delete:   deleteCloudGuardAdhocQuery,
		Schema: map[string]*schema.Schema{
			// Required
			"adhoc_query_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"adhoc_query_resources": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"resource_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"resource_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"query": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"adhoc_query_regional_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"expected_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expired_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"failed_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"regional_error": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"regional_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"succeeded_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"error_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
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

func createCloudGuardAdhocQuery(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardAdhocQuery(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func deleteCloudGuardAdhocQuery(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardAdhocQueryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardAdhocQueryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.AdhocQuery
	DisableNotFoundRetries bool
}

func (s *CloudGuardAdhocQueryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardAdhocQueryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardAdhocQueryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardAdhocQueryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardAdhocQueryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardAdhocQueryResourceCrud) Create() error {
	request := oci_cloud_guard.CreateAdhocQueryRequest{}

	if adhocQueryDetails, ok := s.D.GetOkExists("adhoc_query_details"); ok {
		if tmpList := adhocQueryDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "adhoc_query_details", 0)
			tmp, err := s.mapToAdhocQueryDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdhocQueryDetails = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateAdhocQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AdhocQuery
	return nil
}

func (s *CloudGuardAdhocQueryResourceCrud) Get() error {
	request := oci_cloud_guard.GetAdhocQueryRequest{}

	tmp := s.D.Id()
	request.AdhocQueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetAdhocQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AdhocQuery
	return nil
}

func (s *CloudGuardAdhocQueryResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteAdhocQueryRequest{}

	tmp := s.D.Id()
	request.AdhocQueryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteAdhocQuery(context.Background(), request)
	return err
}

func (s *CloudGuardAdhocQueryResourceCrud) SetData() error {
	if s.Res.AdhocQueryDetails != nil {
		s.D.Set("adhoc_query_details", []interface{}{AdhocQueryDetailsToMap(s.Res.AdhocQueryDetails)})
	} else {
		s.D.Set("adhoc_query_details", nil)
	}

	adhocQueryRegionalDetails := []interface{}{}
	for _, item := range s.Res.AdhocQueryRegionalDetails {
		adhocQueryRegionalDetails = append(adhocQueryRegionalDetails, AdhocQueryRegionalDetailsToMap(item))
	}
	s.D.Set("adhoc_query_regional_details", adhocQueryRegionalDetails)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.ErrorMessage != nil {
		s.D.Set("error_message", *s.Res.ErrorMessage)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *CloudGuardAdhocQueryResourceCrud) mapToAdhocQueryDetails(fieldKeyFormat string) (oci_cloud_guard.AdhocQueryDetails, error) {
	result := oci_cloud_guard.AdhocQueryDetails{}

	if adhocQueryResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "adhoc_query_resources")); ok {
		interfaces := adhocQueryResources.([]interface{})
		tmp := make([]oci_cloud_guard.AdhocQueryResource, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "adhoc_query_resources"), stateDataIndex)
			converted, err := s.mapToAdhocQueryResource(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "adhoc_query_resources")) {
			result.AdhocQueryResources = tmp
		}
	}

	if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
		tmp := query.(string)
		result.Query = &tmp
	}

	return result, nil
}

func AdhocQueryDetailsToMap(obj *oci_cloud_guard.AdhocQueryDetails) map[string]interface{} {
	result := map[string]interface{}{}

	adhocQueryResources := []interface{}{}
	for _, item := range obj.AdhocQueryResources {
		adhocQueryResources = append(adhocQueryResources, AdhocQueryResourceToMap(item))
	}
	result["adhoc_query_resources"] = adhocQueryResources

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	return result
}

func AdhocQueryRegionalDetailsToMap(obj oci_cloud_guard.AdhocQueryRegionalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExpectedCount != nil {
		result["expected_count"] = string(*obj.ExpectedCount)
	}

	if obj.ExpiredCount != nil {
		result["expired_count"] = string(*obj.ExpiredCount)
	}

	if obj.FailedCount != nil {
		result["failed_count"] = string(*obj.FailedCount)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.RegionalError != nil {
		result["regional_error"] = string(*obj.RegionalError)
	}

	result["regional_status"] = string(obj.RegionalStatus)

	if obj.SucceededCount != nil {
		result["succeeded_count"] = string(*obj.SucceededCount)
	}

	return result
}

func (s *CloudGuardAdhocQueryResourceCrud) mapToAdhocQueryResource(fieldKeyFormat string) (oci_cloud_guard.AdhocQueryResource, error) {
	result := oci_cloud_guard.AdhocQueryResource{}

	if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
		tmp := region.(string)
		result.Region = &tmp
	}

	if resourceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_ids")); ok {
		interfaces := resourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "resource_ids")) {
			result.ResourceIds = tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	return result, nil
}

func AdhocQueryResourceToMap(obj oci_cloud_guard.AdhocQueryResource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["resource_ids"] = obj.ResourceIds

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}

func AdhocQuerySummaryToMap(obj oci_cloud_guard.AdhocQuerySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdhocQueryDetails != nil {
		result["adhoc_query_details"] = []interface{}{AdhocQueryDetailsToMap(obj.AdhocQueryDetails)}
	}

	adhocQueryRegionalDetails := []interface{}{}
	for _, item := range obj.AdhocQueryRegionalDetails {
		adhocQueryRegionalDetails = append(adhocQueryRegionalDetails, AdhocQueryRegionalDetailsToMap(item))
	}
	result["adhoc_query_regional_details"] = adhocQueryRegionalDetails

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
