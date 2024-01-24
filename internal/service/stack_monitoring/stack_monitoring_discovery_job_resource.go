// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"
)

func StackMonitoringDiscoveryJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringDiscoveryJob,
		Read:     readStackMonitoringDiscoveryJob,
		Delete:   deleteStackMonitoringDiscoveryJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"discovery_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"agent_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"properties": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"properties_map": {
										Type:     schema.TypeMap,
										Optional: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},

									// Computed
								},
							},
						},
						"resource_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"credentials": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"items": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"credential_name": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"credential_type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"properties": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"properties_map": {
																Type:     schema.TypeMap,
																Optional: true,
																//	Computed: true,
																ForceNew: true,
																Elem:     schema.TypeString,
															},

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional

									// Computed
								},
							},
						},
						"license": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"tags": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"properties_map": {
										Type:     schema.TypeMap,
										Optional: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
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
			"discovery_client": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"discovery_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"should_propagate_tags_to_discovered_resources": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createStackMonitoringDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func deleteStackMonitoringDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringDiscoveryJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.DiscoveryJob
	DisableNotFoundRetries bool
}

func (s *StackMonitoringDiscoveryJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringDiscoveryJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateCreating),
	}
}

func (s *StackMonitoringDiscoveryJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateActive),
	}
}

func (s *StackMonitoringDiscoveryJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateDeleting),
	}
}

func (s *StackMonitoringDiscoveryJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.LifecycleStateDeleted),
	}
}

func (s *StackMonitoringDiscoveryJobResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateDiscoveryJobRequest{}

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

	if discoveryClient, ok := s.D.GetOkExists("discovery_client"); ok {
		tmp := discoveryClient.(string)
		request.DiscoveryClient = &tmp
	}

	if discoveryDetails, ok := s.D.GetOkExists("discovery_details"); ok {
		if tmpList := discoveryDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "discovery_details", 0)
			tmp, err := s.mapToDiscoveryDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DiscoveryDetails = &tmp
		}
	}

	if discoveryType, ok := s.D.GetOkExists("discovery_type"); ok {
		request.DiscoveryType = oci_stack_monitoring.CreateDiscoveryJobDetailsDiscoveryTypeEnum(discoveryType.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if shouldPropagateTagsToDiscoveredResources, ok := s.D.GetOkExists("should_propagate_tags_to_discovered_resources"); ok {
		tmp := shouldPropagateTagsToDiscoveredResources.(bool)
		request.ShouldPropagateTagsToDiscoveredResources = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoveryJob
	return nil
}

func (s *StackMonitoringDiscoveryJobResourceCrud) Get() error {
	request := oci_stack_monitoring.GetDiscoveryJobRequest{}

	tmp := s.D.Id()
	request.DiscoveryJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DiscoveryJob
	return nil
}

func (s *StackMonitoringDiscoveryJobResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteDiscoveryJobRequest{}

	tmp := s.D.Id()
	request.DiscoveryJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteDiscoveryJob(context.Background(), request)
	return err
}

func (s *StackMonitoringDiscoveryJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DiscoveryClient != nil {
		s.D.Set("discovery_client", *s.Res.DiscoveryClient)
	}

	if s.Res.DiscoveryDetails != nil {
		s.D.Set("discovery_details", []interface{}{DiscoveryDetailsToMap(s.Res.DiscoveryDetails)})
	} else {
		s.D.Set("discovery_details", nil)
	}

	s.D.Set("discovery_type", s.Res.DiscoveryType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusMessage != nil {
		s.D.Set("status_message", *s.Res.StatusMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}

func (s *StackMonitoringDiscoveryJobResourceCrud) mapToCredentialCollection(fieldKeyFormat string) (oci_stack_monitoring.CredentialCollection, error) {
	result := oci_stack_monitoring.CredentialCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_stack_monitoring.CredentialDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToCredentialDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func CredentialCollectionToMap(obj *oci_stack_monitoring.CredentialCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, CredentialDetailsToMap(item))
	}
	result["items"] = items

	return result
}

func (s *StackMonitoringDiscoveryJobResourceCrud) mapToCredentialDetails(fieldKeyFormat string) (oci_stack_monitoring.CredentialDetails, error) {
	result := oci_stack_monitoring.CredentialDetails{}

	if credentialName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_name")); ok {
		tmp := credentialName.(string)
		result.CredentialName = &tmp
	}

	if credentialType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credential_type")); ok {
		tmp := credentialType.(string)
		result.CredentialType = &tmp
	}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		if tmpList := properties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), 0)
			tmp, err := s.mapToPropertyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert properties, encountered error: %v", err)
			}
			result.Properties = &tmp
		}
	}

	return result, nil
}

func CredentialDetailsToMap(obj oci_stack_monitoring.CredentialDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CredentialName != nil {
		result["credential_name"] = string(*obj.CredentialName)
	}

	if obj.CredentialType != nil {
		result["credential_type"] = string(*obj.CredentialType)
	}

	if obj.Properties != nil {
		result["properties"] = []interface{}{PropertyDetailsToMap(obj.Properties)}
	}

	return result
}

func (s *StackMonitoringDiscoveryJobResourceCrud) mapToDiscoveryDetails(fieldKeyFormat string) (oci_stack_monitoring.DiscoveryDetails, error) {
	result := oci_stack_monitoring.DiscoveryDetails{}

	if agentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_id")); ok {
		tmp := agentId.(string)
		result.AgentId = &tmp
	}

	if credentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials")); ok {
		if tmpList := credentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credentials"), 0)
			tmp, err := s.mapToCredentialCollection(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert credentials, encountered error: %v", err)
			}
			result.Credentials = &tmp
		}
	}

	if license, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license")); ok {
		result.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
	}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		if tmpList := properties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), 0)
			tmp, err := s.mapToPropertyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert properties, encountered error: %v", err)
			}
			result.Properties = &tmp
		}
	}

	if resourceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name")); ok {
		tmp := resourceName.(string)
		result.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		result.ResourceType = oci_stack_monitoring.DiscoveryDetailsResourceTypeEnum(resourceType.(string))
	}

	if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
		if tmpList := tags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), 0)
			tmp, err := s.mapToPropertyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tags, encountered error: %v", err)
			}
			result.Tags = &tmp
		}
	}

	return result, nil
}

func DiscoveryDetailsToMap(obj *oci_stack_monitoring.DiscoveryDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.Credentials != nil {
		result["credentials"] = []interface{}{CredentialCollectionToMap(obj.Credentials)}
	}

	result["license"] = string(obj.License)

	if obj.Properties != nil {
		result["properties"] = []interface{}{PropertyDetailsToMap(obj.Properties)}
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_type"] = string(obj.ResourceType)

	if obj.Tags != nil {
		result["tags"] = []interface{}{PropertyDetailsToMap(obj.Tags)}
	}

	return result
}

func DiscoveryJobSummaryToMap(obj oci_stack_monitoring.DiscoveryJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["discovery_type"] = string(obj.DiscoveryType)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.StatusMessage != nil {
		result["status_message"] = string(*obj.StatusMessage)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UserId != nil {
		result["user_id"] = string(*obj.UserId)
	}

	return result
}

func (s *StackMonitoringDiscoveryJobResourceCrud) mapToPropertyDetails(fieldKeyFormat string) (oci_stack_monitoring.PropertyDetails, error) {
	result := oci_stack_monitoring.PropertyDetails{}

	if propertiesMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties_map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(propertiesMap.(map[string]interface{}))
	}

	return result, nil
}

func PropertyDetailsToMap(obj *oci_stack_monitoring.PropertyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["properties_map"] = obj.PropertiesMap

	return result
}
