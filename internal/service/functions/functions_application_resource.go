// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"
)

func FunctionsApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFunctionsApplication,
		Read:     readFunctionsApplication,
		Update:   updateFunctionsApplication,
		Delete:   deleteFunctionsApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"image_policy_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_policy_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"key_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"kms_key_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"network_security_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"syslog_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFunctionsApplication(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FunctionsApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_functions.FunctionsManagementClient
	Res                    *oci_functions.Application
	DisableNotFoundRetries bool
}

func (s *FunctionsApplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FunctionsApplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateCreating),
	}
}

func (s *FunctionsApplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateActive),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleting),
	}
}

func (s *FunctionsApplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_functions.ApplicationLifecycleStateDeleted),
	}
}

func (s *FunctionsApplicationResourceCrud) Create() error {
	request := oci_functions.CreateApplicationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToImagePolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		request.Shape = oci_functions.CreateApplicationDetailsShapeEnum(shape.(string))
	}

	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subnet_ids") {
			request.SubnetIds = tmp
		}
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Get() error {
	request := oci_functions.GetApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_functions.UpdateApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	if config, ok := s.D.GetOkExists("config"); ok {
		request.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
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

	if imagePolicyConfig, ok := s.D.GetOkExists("image_policy_config"); ok {
		if tmpList := imagePolicyConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "image_policy_config", 0)
			tmp, err := s.mapToImagePolicyConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImagePolicyConfig = &tmp
		}
	}

	if networkSecurityGroupIds, ok := s.D.GetOkExists("network_security_group_ids"); ok {
		set := networkSecurityGroupIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("network_security_group_ids") {
			request.NetworkSecurityGroupIds = tmp
		}
	}

	if syslogUrl, ok := s.D.GetOkExists("syslog_url"); ok {
		tmp := syslogUrl.(string)
		request.SyslogUrl = &tmp
	}

	if traceConfig, ok := s.D.GetOkExists("trace_config"); ok && s.D.HasChange("trace_config") {
		if tmpList := traceConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trace_config", 0)
			tmp, err := s.mapToApplicationTraceConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TraceConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *FunctionsApplicationResourceCrud) Delete() error {
	request := oci_functions.DeleteApplicationRequest{}

	tmp := s.D.Id()
	request.ApplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	_, err := s.Client.DeleteApplication(context.Background(), request)
	return err
}

func (s *FunctionsApplicationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config", s.Res.Config)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImagePolicyConfig != nil {
		s.D.Set("image_policy_config", []interface{}{ImagePolicyConfigToMapFunctions(s.Res.ImagePolicyConfig)})
	} else {
		s.D.Set("image_policy_config", nil)
	}

	networkSecurityGroupIds := []interface{}{}
	for _, item := range s.Res.NetworkSecurityGroupIds {
		networkSecurityGroupIds = append(networkSecurityGroupIds, item)
	}
	s.D.Set("network_security_group_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds))

	s.D.Set("shape", s.Res.Shape)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.SyslogUrl != nil {
		s.D.Set("syslog_url", *s.Res.SyslogUrl)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TraceConfig != nil {
		s.D.Set("trace_config", []interface{}{ApplicationTraceConfigToMap(s.Res.TraceConfig)})
	} else {
		s.D.Set("trace_config", nil)
	}

	return nil
}

func (s *FunctionsApplicationResourceCrud) mapToApplicationTraceConfig(fieldKeyFormat string) (oci_functions.ApplicationTraceConfig, error) {
	result := oci_functions.ApplicationTraceConfig{}

	if domainId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_id")); ok {
		tmp := domainId.(string)
		result.DomainId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func ApplicationTraceConfigToMap(obj *oci_functions.ApplicationTraceConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainId != nil {
		result["domain_id"] = string(*obj.DomainId)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *FunctionsApplicationResourceCrud) mapToImagePolicyConfig(fieldKeyFormat string) (oci_functions.ImagePolicyConfig, error) {
	result := oci_functions.ImagePolicyConfig{}

	if isPolicyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_policy_enabled")); ok {
		tmp := isPolicyEnabled.(bool)
		result.IsPolicyEnabled = &tmp
	}

	if keyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_details")); ok {
		interfaces := keyDetails.([]interface{})
		tmp := make([]oci_functions.KeyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "key_details"), stateDataIndex)
			converted, err := s.mapToKeyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "key_details")) {
			result.KeyDetails = tmp
		}
	}

	return result, nil
}

func ImagePolicyConfigToMapFunctions(obj *oci_functions.ImagePolicyConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsPolicyEnabled != nil {
		result["is_policy_enabled"] = bool(*obj.IsPolicyEnabled)
	}

	keyDetails := []interface{}{}
	for _, item := range obj.KeyDetails {
		keyDetails = append(keyDetails, KeyDetailsToMapFunctions(item))
	}
	result["key_details"] = keyDetails

	return result
}

func (s *FunctionsApplicationResourceCrud) mapToKeyDetails(fieldKeyFormat string) (oci_functions.KeyDetails, error) {
	result := oci_functions.KeyDetails{}

	if kmsKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kms_key_id")); ok {
		tmp := kmsKeyId.(string)
		result.KmsKeyId = &tmp
	}

	return result, nil
}

func KeyDetailsToMapFunctions(obj oci_functions.KeyDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	return result
}

func (s *FunctionsApplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_functions.ChangeApplicationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ApplicationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "functions")

	_, err := s.Client.ChangeApplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FunctionsApplicationResourceCrud) ExtraWaitPostDelete() time.Duration {
	if httpreplay.ShouldRetryImmediately() {
		return time.Duration(1 * time.Second)
	}
	log.Printf("[DEBUG] Waiting for 5 minutes post destroy of application resource due to known service issue")
	return time.Duration(5 * time.Minute)
}
