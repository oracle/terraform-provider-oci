// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"
)

func MysqlChannelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("30m"),
			Update: tfresource.GetTimeoutDuration("30m"),
			Delete: tfresource.GetTimeoutDuration("30m"),
		},
		Create: createMysqlChannel,
		Read:   readMysqlChannel,
		Update: updateMysqlChannel,
		Delete: deleteMysqlChannel,
		Schema: map[string]*schema.Schema{
			// Required
			"source": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"hostname": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"MYSQL",
							}, true),
						},
						"ssl_mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ssl_ca_certificate": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"certificate_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PEM",
										}, true),
									},
									"contents": {
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
			"target": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"db_system_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"target_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DBSYSTEM",
							}, true),
						},

						// Optional
						"applier_username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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

func createMysqlChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()

	return tfresource.CreateResource(d, sync)
}

func readMysqlChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMysqlChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ChannelsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlChannelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.ChannelsClient
	Res                    *oci_mysql.Channel
	DisableNotFoundRetries bool
}

func (s *MysqlChannelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlChannelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateCreating),
	}
}

func (s *MysqlChannelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateActive),
		string(oci_mysql.ChannelLifecycleStateNeedsAttention),
		string(oci_mysql.ChannelLifecycleStateInactive), // when is_enabled if false
	}
}

func (s *MysqlChannelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateDeleting),
	}
}

func (s *MysqlChannelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateDeleted),
	}
}

func (s *MysqlChannelResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateUpdating),
	}
}

func (s *MysqlChannelResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.ChannelLifecycleStateActive),
		string(oci_mysql.ChannelLifecycleStateNeedsAttention),
		string(oci_mysql.ChannelLifecycleStateInactive), // when is_enabled if false
	}
}

func (s *MysqlChannelResourceCrud) Create() error {
	request := oci_mysql.CreateChannelRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToCreateChannelSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToCreateChannelTargetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Channel
	return nil
}

func (s *MysqlChannelResourceCrud) Get() error {
	request := oci_mysql.GetChannelRequest{}

	tmp := s.D.Id()
	request.ChannelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Channel
	return nil
}

func (s *MysqlChannelResourceCrud) Update() error {
	request := oci_mysql.UpdateChannelRequest{}

	tmp := s.D.Id()
	request.ChannelId = &tmp

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		if tmpList := source.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source", 0)
			tmp, err := s.mapToUpdateChannelSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Source = tmp
		}
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		if tmpList := target.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target", 0)
			tmp, err := s.mapToUpdateChannelTargetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Target = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateChannel(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlChannelResourceCrud) Delete() error {
	request := oci_mysql.DeleteChannelRequest{}

	tmp := s.D.Id()
	request.ChannelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteChannel(context.Background(), request)
	return err
}

func (s *MysqlChannelResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.SetId(*s.Res.Id)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Source != nil {
		sourceArray := []interface{}{}
		if sourceMap := s.ChannelSourceToMap(&s.Res.Source); sourceMap != nil {
			sourceArray = append(sourceArray, sourceMap)
		}
		s.D.Set("source", sourceArray)
	} else {
		s.D.Set("source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Target != nil {
		targetArray := []interface{}{}
		if targetMap := ChannelTargetToMap(&s.Res.Target); targetMap != nil {
			targetArray = append(targetArray, targetMap)
		}
		s.D.Set("target", targetArray)
	} else {
		s.D.Set("target", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MysqlChannelResourceCrud) mapToCreateChannelSourceDetails(fieldKeyFormat string) (oci_mysql.CreateChannelSourceDetails, error) {
	var baseObject oci_mysql.CreateChannelSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("MYSQL"):
		details := oci_mysql.CreateChannelSourceFromMysqlDetails{}
		if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if sslMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_mode")); ok {
			details.SslMode = oci_mysql.ChannelSourceMysqlSslModeEnum(sslMode.(string))
		}
		if sslCaCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate")); ok {
			if tmpList := sslCaCertificate.([]interface{}); len(tmpList) > 0 {
				certificateFieldKeyFormat := fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate.0.%s")
				tmp, err := s.mapToCaCertificate(certificateFieldKeyFormat)
				if err != nil {
					return nil, err
				}
				details.SslCaCertificate = tmp
			}
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func (s *MysqlChannelResourceCrud) mapToCaCertificate(fieldKeyFormat string) (oci_mysql.CaCertificate, error) {
	var baseObject oci_mysql.CaCertificate
	//discriminator
	certificateTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_type"))
	var certificateType string
	if ok {
		certificateType = certificateTypeRaw.(string)
	} else {
		certificateType = "" // default value
	}
	switch strings.ToLower(certificateType) {
	case strings.ToLower("PEM"):
		details := oci_mysql.PemCaCertificate{}
		if contents, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "contents")); ok {
			tmp := contents.(string)
			details.Contents = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown certificate_type '%v' was specified", certificateType)
	}
	return baseObject, nil
}

func (s *MysqlChannelResourceCrud) CaCertificateToMap(obj *oci_mysql.CaCertificate) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.PemCaCertificate:
		result["certificate_type"] = "PEM"

		if v.Contents != nil {
			result["contents"] = string(*v.Contents)
		}
	default:
		log.Printf("[WARN] Received 'certificate_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlChannelResourceCrud) mapToUpdateChannelSourceDetails(fieldKeyFormat string) (oci_mysql.UpdateChannelSourceDetails, error) {
	var baseObject oci_mysql.UpdateChannelSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("MYSQL"):
		details := oci_mysql.UpdateChannelSourceFromMysqlDetails{}
		if hostname, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hostname")); ok {
			tmp := hostname.(string)
			details.Hostname = &tmp
		}
		if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if sslMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_mode")); ok {
			details.SslMode = oci_mysql.ChannelSourceMysqlSslModeEnum(sslMode.(string))
		}
		if sslCaCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate")); ok {
			if tmpList := sslCaCertificate.([]interface{}); len(tmpList) > 0 {
				certificateFieldKeyFormat := fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate.0.%s")
				tmp, err := s.mapToCaCertificate(certificateFieldKeyFormat)
				if err != nil {
					return nil, err
				}
				details.SslCaCertificate = tmp
			}
		}
		if sslMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_mode")); ok {
			details.SslMode = oci_mysql.ChannelSourceMysqlSslModeEnum(sslMode.(string))
		}
		if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func (s *MysqlChannelResourceCrud) ChannelSourceToMap(obj *oci_mysql.ChannelSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_mysql.ChannelSourceMysql:
		result["source_type"] = "MYSQL"

		if v.Hostname != nil {
			result["hostname"] = string(*v.Hostname)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.SslCaCertificate != nil {
			sslCaCertificateArray := []interface{}{}
			if sslCaCertificateMap := s.CaCertificateToMap(&v.SslCaCertificate); sslCaCertificateMap != nil {
				sslCaCertificateArray = append(sslCaCertificateArray, sslCaCertificateMap)
			}
			result["ssl_ca_certificate"] = sslCaCertificateArray
		}

		result["ssl_mode"] = string(v.SslMode)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}

		if password, ok := s.D.GetOkExists("source.0.password"); ok && password != nil {
			result["password"] = password.(string)
		}

	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *MysqlChannelResourceCrud) mapToCreateChannelTargetDetails(fieldKeyFormat string) (oci_mysql.CreateChannelTargetDetails, error) {
	var baseObject oci_mysql.CreateChannelTargetDetails
	//discriminator
	targetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_type"))
	var targetType string
	if ok {
		targetType = targetTypeRaw.(string)
	} else {
		targetType = "" // default value
	}
	switch strings.ToLower(targetType) {
	case strings.ToLower("DBSYSTEM"):
		details := oci_mysql.CreateChannelTargetFromDbSystemDetails{}
		if applierUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "applier_username")); ok {
			tmp := applierUsername.(string)
			details.ApplierUsername = &tmp
		}
		if channelName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channel_name")); ok {
			tmp := channelName.(string)
			details.ChannelName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "db_system_id")); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_type '%v' was specified", targetType)
	}
	return baseObject, nil
}

func (s *MysqlChannelResourceCrud) mapToUpdateChannelTargetDetails(fieldKeyFormat string) (oci_mysql.UpdateChannelTargetDetails, error) {
	var baseObject oci_mysql.UpdateChannelTargetDetails
	//discriminator
	targetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_type"))
	var targetType string
	if ok {
		targetType = targetTypeRaw.(string)
	} else {
		targetType = "" // default value
	}
	switch strings.ToLower(targetType) {
	case strings.ToLower("DBSYSTEM"):
		details := oci_mysql.UpdateChannelTargetFromDbSystemDetails{}
		if applierUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "applier_username")); ok {
			tmp := applierUsername.(string)
			details.ApplierUsername = &tmp
		}
		if channelName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channel_name")); ok {
			tmp := channelName.(string)
			details.ChannelName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown target_type '%v' was specified", targetType)
	}
	return baseObject, nil
}
