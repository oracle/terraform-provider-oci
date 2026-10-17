// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerStackResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createResourcemanagerStack,
		Read:     readResourcemanagerStack,
		Update:   updateResourcemanagerStack,
		Delete:   deleteResourcemanagerStack,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The OCID of the compartment containing the stack.",
			},
			"config_source": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				MinItems:    1,
				Description: "The Terraform configuration source for the stack. Only ZIP_UPLOAD is currently supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_source_type": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Configuration source type. Only ZIP_UPLOAD is currently supported.",
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ZIP_UPLOAD",
							}, true),
						},
						"zip_file_base64encoded": {
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							Description: "Base64-encoded ZIP archive containing the Terraform configuration for the stack.",
						},
						"working_directory": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "File path to the directory inside the ZIP archive from which Terraform runs.",
						},
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"terraform_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"variables": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"stack_drift_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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
			"time_drift_last_checked": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createResourcemanagerStack(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.CreateResource(d, sync)
}

func readResourcemanagerStack(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

func updateResourcemanagerStack(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteResourcemanagerStack(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerStackResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ResourcemanagerStackResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resourcemanager.ResourceManagerClient
	Res                    *oci_resourcemanager.Stack
	DisableNotFoundRetries bool
}

func (s *ResourcemanagerStackResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ResourcemanagerStackResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resourcemanager.StackLifecycleStateCreating),
	}
}

func (s *ResourcemanagerStackResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resourcemanager.StackLifecycleStateActive),
	}
}

func (s *ResourcemanagerStackResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resourcemanager.StackLifecycleStateDeleting),
	}
}

func (s *ResourcemanagerStackResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resourcemanager.StackLifecycleStateDeleted),
	}
}

func (s *ResourcemanagerStackResourceCrud) Create() error {
	request := oci_resourcemanager.CreateStackRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	configSource, err := s.mapToCreateConfigSourceDetails("config_source.0.%s")
	if err != nil {
		return err
	}
	request.ConfigSource = configSource

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if terraformVersion, ok := s.D.GetOkExists("terraform_version"); ok {
		tmp := terraformVersion.(string)
		request.TerraformVersion = &tmp
	}

	if variables, ok := s.D.GetOkExists("variables"); ok {
		request.Variables = tfresource.ObjectMapToStringMap(variables.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.CreateStack(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stack
	return nil
}

func (s *ResourcemanagerStackResourceCrud) Get() error {
	request := oci_resourcemanager.GetStackRequest{}

	tmp := s.D.Id()
	request.StackId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.GetStack(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stack
	return nil
}

func (s *ResourcemanagerStackResourceCrud) Update() error {
	request := oci_resourcemanager.UpdateStackRequest{}

	tmp := s.D.Id()
	request.StackId = &tmp

	if s.D.HasChange("config_source") {
		configSource, err := s.mapToUpdateConfigSourceDetails("config_source.0.%s")
		if err != nil {
			return err
		}
		request.ConfigSource = configSource
	}

	if s.D.HasChange("defined_tags") {
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			request.DefinedTags = convertedDefinedTags
		} else {
			request.DefinedTags = map[string]map[string]interface{}{}
		}
	}

	if s.D.HasChange("description") {
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			request.Description = &tmp
		} else {
			request.Description = nil
		}
	}

	if s.D.HasChange("display_name") {
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			request.DisplayName = &tmp
		} else {
			request.DisplayName = nil
		}
	}

	if s.D.HasChange("freeform_tags") {
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		} else {
			request.FreeformTags = map[string]string{}
		}
	}

	if s.D.HasChange("terraform_version") {
		if terraformVersion, ok := s.D.GetOkExists("terraform_version"); ok {
			tmp := terraformVersion.(string)
			request.TerraformVersion = &tmp
		} else {
			request.TerraformVersion = nil
		}
	}

	if s.D.HasChange("variables") {
		if variables, ok := s.D.GetOkExists("variables"); ok {
			request.Variables = tfresource.ObjectMapToStringMap(variables.(map[string]interface{}))
		} else {
			request.Variables = map[string]string{}
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	response, err := s.Client.UpdateStack(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stack
	return nil
}

func (s *ResourcemanagerStackResourceCrud) Delete() error {
	request := oci_resourcemanager.DeleteStackRequest{}

	tmp := s.D.Id()
	request.StackId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager")

	_, err := s.Client.DeleteStack(context.Background(), request)
	return err
}

func (s *ResourcemanagerStackResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigSource != nil {
		configSourceArray := []interface{}{}
		downloadedConfigSource, err := s.getDownloadedStackConfigSource()
		if err != nil {
			return err
		}
		if configSourceMap := StackConfigSourceToMap(&s.Res.ConfigSource, downloadedConfigSource); configSourceMap != nil {
			configSourceArray = append(configSourceArray, configSourceMap)
		}
		s.D.Set("config_source", configSourceArray)
	} else {
		s.D.Set("config_source", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("stack_drift_status", s.Res.StackDriftStatus)
	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TerraformVersion != nil {
		s.D.Set("terraform_version", *s.Res.TerraformVersion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDriftLastChecked != nil {
		s.D.Set("time_drift_last_checked", s.Res.TimeDriftLastChecked.String())
	}

	s.D.Set("variables", s.Res.Variables)

	return nil
}

func (s *ResourcemanagerStackResourceCrud) getDownloadedStackConfigSource() ([]interface{}, error) {
	return getDownloadedStackConfigSource(s.Client, s.Res, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resourcemanager"), getExistingStackConfigSource(s.D))
}

func getDownloadedStackConfigSource(client *oci_resourcemanager.ResourceManagerClient, stack *oci_resourcemanager.Stack, retryPolicy *oci_common.RetryPolicy, existingConfigSource []interface{}) ([]interface{}, error) {
	if stack == nil || stack.ConfigSource == nil {
		return nil, nil
	}

	switch v := stack.ConfigSource.(type) {
	case oci_resourcemanager.ZipUploadConfigSource:
		request := oci_resourcemanager.GetStackTfConfigRequest{
			StackId: stack.Id,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		}

		response, err := client.GetStackTfConfig(context.Background(), request)
		if err != nil {
			return nil, err
		}
		defer response.Content.Close()

		content, err := io.ReadAll(response.Content)
		if err != nil {
			return nil, err
		}

		result := map[string]interface{}{
			"config_source_type":     "ZIP_UPLOAD",
			"zip_file_base64encoded": base64.StdEncoding.EncodeToString(content),
		}
		if v.WorkingDirectory != nil {
			result["working_directory"] = *v.WorkingDirectory
		}

		return []interface{}{result}, nil
	default:
		return existingConfigSource, nil
	}
}

func (s *ResourcemanagerStackResourceCrud) mapToCreateConfigSourceDetails(fieldKeyFormat string) (oci_resourcemanager.CreateConfigSourceDetails, error) {
	configSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_source_type"))
	configSourceType := ""
	if ok {
		configSourceType = configSourceTypeRaw.(string)
	}

	switch strings.ToLower(configSourceType) {
	case strings.ToLower("ZIP_UPLOAD"):
		details := oci_resourcemanager.CreateZipUploadConfigSourceDetails{}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		if zipFileBase64Encoded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "zip_file_base64encoded")); ok {
			tmp := zipFileBase64Encoded.(string)
			details.ZipFileBase64Encoded = &tmp
		}
		return details, nil
	default:
		return nil, fmt.Errorf("unknown config_source_type '%v' was specified", configSourceType)
	}
}

func (s *ResourcemanagerStackResourceCrud) mapToUpdateConfigSourceDetails(fieldKeyFormat string) (oci_resourcemanager.UpdateConfigSourceDetails, error) {
	configSourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_source_type"))
	configSourceType := ""
	if ok {
		configSourceType = configSourceTypeRaw.(string)
	}

	switch strings.ToLower(configSourceType) {
	case strings.ToLower("ZIP_UPLOAD"):
		details := oci_resourcemanager.UpdateZipUploadConfigSourceDetails{}
		if workingDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "working_directory")); ok {
			tmp := workingDirectory.(string)
			details.WorkingDirectory = &tmp
		}
		if zipFileBase64Encoded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "zip_file_base64encoded")); ok {
			tmp := zipFileBase64Encoded.(string)
			details.ZipFileBase64Encoded = &tmp
		}
		return details, nil
	default:
		return nil, fmt.Errorf("unknown config_source_type '%v' was specified", configSourceType)
	}
}

func getExistingStackConfigSource(d *schema.ResourceData) []interface{} {
	if existingConfigSource, ok := d.GetOkExists("config_source"); ok {
		if interfaces, ok := existingConfigSource.([]interface{}); ok {
			return interfaces
		}
	}

	return nil
}

func StackConfigSourceToMap(obj *oci_resourcemanager.ConfigSource, existingConfigSource []interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	switch v := (*obj).(type) {
	case oci_resourcemanager.ZipUploadConfigSource:
		result["config_source_type"] = "ZIP_UPLOAD"

		if v.WorkingDirectory != nil {
			result["working_directory"] = *v.WorkingDirectory
		}

		if len(existingConfigSource) > 0 {
			if existingConfigSourceMap, ok := existingConfigSource[0].(map[string]interface{}); ok {
				if zipFileBase64Encoded, ok := existingConfigSourceMap["zip_file_base64encoded"]; ok && zipFileBase64Encoded != nil {
					result["zip_file_base64encoded"] = zipFileBase64Encoded.(string)
				}
			}
		}
	default:
		log.Printf("[WARN] Received 'config_source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
