// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext,
		ReadContext:   readDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext,
		UpdateContext: updateDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext,
		DeleteContext: deleteDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata_source": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},

			// Computed
			"lifecycle_details": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseToolsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_tools.DatabaseToolsClient
	Res                    *oci_database_tools.DatabaseToolsDatabaseApiGatewayConfig
	DisableNotFoundRetries bool
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) ID() string {
	databaseToolsDatabaseApiGatewayConfig := *s.Res
	return *databaseToolsDatabaseApiGatewayConfig.GetId()
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive),
	}
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted),
	}
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database_tools.CreateDatabaseToolsDatabaseApiGatewayConfigRequest{}
	err := s.populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.CreateDatabaseToolsDatabaseApiGatewayConfig(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfig
	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database_tools.GetDatabaseToolsDatabaseApiGatewayConfigRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.GetDatabaseToolsDatabaseApiGatewayConfig(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfig
	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_tools.UpdateDatabaseToolsDatabaseApiGatewayConfigRequest{}
	err := s.populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	response, err := s.Client.UpdateDatabaseToolsDatabaseApiGatewayConfig(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DatabaseToolsDatabaseApiGatewayConfig
	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_database_tools.DeleteDatabaseToolsDatabaseApiGatewayConfigRequest{}

	tmp := s.D.Id()
	request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	_, err := s.Client.DeleteDatabaseToolsDatabaseApiGatewayConfig(ctx, request)
	return err
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigDefault:
		s.D.Set("type", "DEFAULT")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, DbtoolsDbApiGatewayConfigResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		s.D.Set("metadata_source", v.MetadataSource)

		s.D.Set("state", v.LifecycleState)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func DatabaseToolsDatabaseApiGatewayConfigSummaryToMap(obj oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigSummary) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.GetId() != nil {
		result["id"] = *obj.GetId()
	}
	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = *obj.GetCompartmentId()
	}
	if obj.GetDisplayName() != nil {
		result["display_name"] = *obj.GetDisplayName()
	}
	result["state"] = string(obj.GetLifecycleState())

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	switch (obj).(type) {
	case oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigDefaultSummary:
		result["type"] = "DEFAULT"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}
	return result
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_database_tools.ResourceLock, error) {
	result := oci_database_tools.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database_tools.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func DbtoolsDbApiGatewayConfigResourceLockToMap(obj oci_database_tools.ResourceLock) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.RelatedResourceId != nil {
		result["related_resource_id"] = string(*obj.RelatedResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) populateTopLevelPolymorphicCreateDatabaseToolsDatabaseApiGatewayConfigRequest(request *oci_database_tools.CreateDatabaseToolsDatabaseApiGatewayConfigRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools.CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_database_tools.ResourceLock, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToResourceLock(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if metadataSource, ok := s.D.GetOkExists("metadata_source"); ok {
			details.MetadataSource = oci_database_tools.DatabaseApiGatewayConfigMetadataSourceEnum(metadataSource.(string))
		}
		request.CreateDatabaseToolsDatabaseApiGatewayConfigDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) populateTopLevelPolymorphicUpdateDatabaseToolsDatabaseApiGatewayConfigRequest(request *oci_database_tools.UpdateDatabaseToolsDatabaseApiGatewayConfigRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_database_tools.UpdateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails{}
		tmp := s.D.Id()
		request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDatabaseToolsDatabaseApiGatewayConfigDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_database_tools.ChangeDatabaseToolsDatabaseApiGatewayConfigCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DatabaseToolsDatabaseApiGatewayConfigId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_tools")

	_, err := s.Client.ChangeDatabaseToolsDatabaseApiGatewayConfigCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
