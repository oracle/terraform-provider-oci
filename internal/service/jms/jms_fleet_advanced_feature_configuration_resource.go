// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetAdvancedFeatureConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsFleetAdvancedFeatureConfiguration,
		Read:     readJmsFleetAdvancedFeatureConfiguration,
		Update:   updateJmsFleetAdvancedFeatureConfiguration,
		Delete:   deleteJmsFleetAdvancedFeatureConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"advanced_usage_tracking": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"analytic_bucket_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analytic_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crypto_event_analysis": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"summarized_events_log": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"log_group_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"log_id": {
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
			"java_migration_analysis": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"jfr_recording": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"lcm": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"post_installation_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"add_logging_handler": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"disabled_tls_versions": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"global_logging_level": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"minimum_key_size_settings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"certpath": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"key_size": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"jar": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"key_size": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"tls": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"key_size": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"proxies": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"ftp_proxy_host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ftp_proxy_port": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"http_proxy_host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"http_proxy_port": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"https_proxy_host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"https_proxy_port": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"socks_proxy_host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"socks_proxy_port": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"use_system_proxies": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"should_replace_certificates_operating_system": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"performance_tuning_analysis": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
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
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAdvancedFeatureConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAdvancedFeatureConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

func updateJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetAdvancedFeatureConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsFleetAdvancedFeatureConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type JmsFleetAdvancedFeatureConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms.JavaManagementServiceClient
	Res                    *oci_jms.FleetAdvancedFeatureConfiguration
	DisableNotFoundRetries bool
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) ID() string {
	return GetFleetAdvancedFeatureConfigurationCompositeId(s.D.Get("fleet_id").(string))
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) Create() error {
	request := oci_jms.UpdateFleetAdvancedFeatureConfigurationRequest{}

	if advancedUsageTracking, ok := s.D.GetOkExists("advanced_usage_tracking"); ok {
		if tmpList := advancedUsageTracking.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advanced_usage_tracking", 0)
			tmp, err := s.mapToAdvancedUsageTracking(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdvancedUsageTracking = &tmp
		}
	}

	if analyticBucketName, ok := s.D.GetOkExists("analytic_bucket_name"); ok {
		tmp := analyticBucketName.(string)
		request.AnalyticBucketName = &tmp
	}

	if analyticNamespace, ok := s.D.GetOkExists("analytic_namespace"); ok {
		tmp := analyticNamespace.(string)
		request.AnalyticNamespace = &tmp
	}

	if cryptoEventAnalysis, ok := s.D.GetOkExists("crypto_event_analysis"); ok {
		if tmpList := cryptoEventAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "crypto_event_analysis", 0)
			tmp, err := s.mapToCryptoEventAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CryptoEventAnalysis = &tmp
		}
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if javaMigrationAnalysis, ok := s.D.GetOkExists("java_migration_analysis"); ok {
		if tmpList := javaMigrationAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "java_migration_analysis", 0)
			tmp, err := s.mapToJavaMigrationAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JavaMigrationAnalysis = &tmp
		}
	}

	if jfrRecording, ok := s.D.GetOkExists("jfr_recording"); ok {
		if tmpList := jfrRecording.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jfr_recording", 0)
			tmp, err := s.mapToJfrRecording(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JfrRecording = &tmp
		}
	}

	if lcm, ok := s.D.GetOkExists("lcm"); ok {
		if tmpList := lcm.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lcm", 0)
			tmp, err := s.mapToLcm(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Lcm = &tmp
		}
	}

	if performanceTuningAnalysis, ok := s.D.GetOkExists("performance_tuning_analysis"); ok {
		if tmpList := performanceTuningAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "performance_tuning_analysis", 0)
			tmp, err := s.mapToPerformanceTuningAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PerformanceTuningAnalysis = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateFleetAdvancedFeatureConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetAdvancedFeatureConfiguration
	return nil
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) Get() error {
	request := oci_jms.GetFleetAdvancedFeatureConfigurationRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	fleetId, err := parseFleetAdvancedFeatureConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.FleetId = &fleetId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.GetFleetAdvancedFeatureConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetAdvancedFeatureConfiguration
	return nil
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) Update() error {
	request := oci_jms.UpdateFleetAdvancedFeatureConfigurationRequest{}

	if advancedUsageTracking, ok := s.D.GetOkExists("advanced_usage_tracking"); ok {
		if tmpList := advancedUsageTracking.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "advanced_usage_tracking", 0)
			tmp, err := s.mapToAdvancedUsageTracking(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdvancedUsageTracking = &tmp
		}
	}

	if analyticBucketName, ok := s.D.GetOkExists("analytic_bucket_name"); ok {
		tmp := analyticBucketName.(string)
		request.AnalyticBucketName = &tmp
	}

	if analyticNamespace, ok := s.D.GetOkExists("analytic_namespace"); ok {
		tmp := analyticNamespace.(string)
		request.AnalyticNamespace = &tmp
	}

	if cryptoEventAnalysis, ok := s.D.GetOkExists("crypto_event_analysis"); ok {
		if tmpList := cryptoEventAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "crypto_event_analysis", 0)
			tmp, err := s.mapToCryptoEventAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CryptoEventAnalysis = &tmp
		}
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if javaMigrationAnalysis, ok := s.D.GetOkExists("java_migration_analysis"); ok {
		if tmpList := javaMigrationAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "java_migration_analysis", 0)
			tmp, err := s.mapToJavaMigrationAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JavaMigrationAnalysis = &tmp
		}
	}

	if jfrRecording, ok := s.D.GetOkExists("jfr_recording"); ok {
		if tmpList := jfrRecording.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "jfr_recording", 0)
			tmp, err := s.mapToJfrRecording(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JfrRecording = &tmp
		}
	}

	if lcm, ok := s.D.GetOkExists("lcm"); ok {
		if tmpList := lcm.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lcm", 0)
			tmp, err := s.mapToLcm(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Lcm = &tmp
		}
	}

	if performanceTuningAnalysis, ok := s.D.GetOkExists("performance_tuning_analysis"); ok {
		if tmpList := performanceTuningAnalysis.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "performance_tuning_analysis", 0)
			tmp, err := s.mapToPerformanceTuningAnalysis(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PerformanceTuningAnalysis = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateFleetAdvancedFeatureConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FleetAdvancedFeatureConfiguration
	return nil
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) SetData() error {

	fleetId, err := parseFleetAdvancedFeatureConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("fleet_id", &fleetId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AdvancedUsageTracking != nil {
		s.D.Set("advanced_usage_tracking", []interface{}{AdvancedUsageTrackingToMap(s.Res.AdvancedUsageTracking)})
	} else {
		s.D.Set("advanced_usage_tracking", nil)
	}

	if s.Res.AnalyticBucketName != nil {
		s.D.Set("analytic_bucket_name", *s.Res.AnalyticBucketName)
	}

	if s.Res.AnalyticNamespace != nil {
		s.D.Set("analytic_namespace", *s.Res.AnalyticNamespace)
	}

	if s.Res.CryptoEventAnalysis != nil {
		s.D.Set("crypto_event_analysis", []interface{}{CryptoEventAnalysisToMap(s.Res.CryptoEventAnalysis)})
	} else {
		s.D.Set("crypto_event_analysis", nil)
	}

	if s.Res.JavaMigrationAnalysis != nil {
		s.D.Set("java_migration_analysis", []interface{}{JavaMigrationAnalysisToMap(s.Res.JavaMigrationAnalysis)})
	} else {
		s.D.Set("java_migration_analysis", nil)
	}

	if s.Res.JfrRecording != nil {
		s.D.Set("jfr_recording", []interface{}{JfrRecordingToMap(s.Res.JfrRecording)})
	} else {
		s.D.Set("jfr_recording", nil)
	}

	if s.Res.Lcm != nil {
		s.D.Set("lcm", []interface{}{LcmToMap(s.Res.Lcm)})
	} else {
		s.D.Set("lcm", nil)
	}

	if s.Res.PerformanceTuningAnalysis != nil {
		s.D.Set("performance_tuning_analysis", []interface{}{PerformanceTuningAnalysisToMap(s.Res.PerformanceTuningAnalysis)})
	} else {
		s.D.Set("performance_tuning_analysis", nil)
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func GetFleetAdvancedFeatureConfigurationCompositeId(fleetId string) string {
	fleetId = url.PathEscape(fleetId)
	compositeId := "fleets/" + fleetId + "/advancedFeatureConfiguration"
	return compositeId
}

func parseFleetAdvancedFeatureConfigurationCompositeId(compositeId string) (fleetId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fleets/.*/advancedFeatureConfiguration", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fleetId, _ = url.PathUnescape(parts[1])

	return
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToAdvancedUsageTracking(fieldKeyFormat string) (oci_jms.AdvancedUsageTracking, error) {
	result := oci_jms.AdvancedUsageTracking{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func AdvancedUsageTrackingToMap(obj *oci_jms.AdvancedUsageTracking) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToCryptoEventAnalysis(fieldKeyFormat string) (oci_jms.CryptoEventAnalysis, error) {
	result := oci_jms.CryptoEventAnalysis{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if summarizedEventsLog, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "summarized_events_log")); ok {
		if tmpList := summarizedEventsLog.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "summarized_events_log"), 0)
			tmp, err := s.mapToSummarizedEventsLog(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert summarized_events_log, encountered error: %v", err)
			}
			result.SummarizedEventsLog = &tmp
		}
	}

	return result, nil
}

func CryptoEventAnalysisToMap(obj *oci_jms.CryptoEventAnalysis) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.SummarizedEventsLog != nil {
		result["summarized_events_log"] = []interface{}{SummarizedEventsLogToMap(obj.SummarizedEventsLog)}
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToJavaMigrationAnalysis(fieldKeyFormat string) (oci_jms.JavaMigrationAnalysis, error) {
	result := oci_jms.JavaMigrationAnalysis{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func JavaMigrationAnalysisToMap(obj *oci_jms.JavaMigrationAnalysis) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToJfrRecording(fieldKeyFormat string) (oci_jms.JfrRecording, error) {
	result := oci_jms.JfrRecording{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func JfrRecordingToMap(obj *oci_jms.JfrRecording) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToKeySizeAlgorithm(fieldKeyFormat string) (oci_jms.KeySizeAlgorithm, error) {
	result := oci_jms.KeySizeAlgorithm{}

	if keySize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_size")); ok {
		tmp := keySize.(int)
		result.KeySize = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		result.Name = oci_jms.AlgorithmsEnum(name.(string))
	}

	return result, nil
}

func KeySizeAlgorithmToMap(obj oci_jms.KeySizeAlgorithm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.KeySize != nil {
		result["key_size"] = int(*obj.KeySize)
	}

	result["name"] = string(obj.Name)

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToLcm(fieldKeyFormat string) (oci_jms.Lcm, error) {
	result := oci_jms.Lcm{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if postInstallationActions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "post_installation_actions")); ok {
		if tmpList := postInstallationActions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "post_installation_actions"), 0)
			tmp, err := s.mapToPostInstallationActionSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert post_installation_actions, encountered error: %v", err)
			}
			result.PostInstallationActions = &tmp
		}
	}

	return result, nil
}

func LcmToMap(obj *oci_jms.Lcm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.PostInstallationActions != nil {
		result["post_installation_actions"] = []interface{}{PostInstallationActionSettingsToMap(obj.PostInstallationActions)}
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToMinimumKeySizeSettings(fieldKeyFormat string) (oci_jms.MinimumKeySizeSettings, error) {
	result := oci_jms.MinimumKeySizeSettings{}

	if certpath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certpath")); ok {
		interfaces := certpath.([]interface{})
		tmp := make([]oci_jms.KeySizeAlgorithm, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "certpath"), stateDataIndex)
			converted, err := s.mapToKeySizeAlgorithm(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "certpath")) {
			result.Certpath = tmp
		}
	}

	if jar, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jar")); ok {
		interfaces := jar.([]interface{})
		tmp := make([]oci_jms.KeySizeAlgorithm, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "jar"), stateDataIndex)
			converted, err := s.mapToKeySizeAlgorithm(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "jar")) {
			result.Jar = tmp
		}
	}

	if tls, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tls")); ok {
		interfaces := tls.([]interface{})
		tmp := make([]oci_jms.KeySizeAlgorithm, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tls"), stateDataIndex)
			converted, err := s.mapToKeySizeAlgorithm(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tls")) {
			result.Tls = tmp
		}
	}

	return result, nil
}

func MinimumKeySizeSettingsToMap(obj *oci_jms.MinimumKeySizeSettings) map[string]interface{} {
	result := map[string]interface{}{}

	certpath := []interface{}{}
	for _, item := range obj.Certpath {
		certpath = append(certpath, KeySizeAlgorithmToMap(item))
	}
	result["certpath"] = certpath

	jar := []interface{}{}
	for _, item := range obj.Jar {
		jar = append(jar, KeySizeAlgorithmToMap(item))
	}
	result["jar"] = jar

	tls := []interface{}{}
	for _, item := range obj.Tls {
		tls = append(tls, KeySizeAlgorithmToMap(item))
	}
	result["tls"] = tls

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToPerformanceTuningAnalysis(fieldKeyFormat string) (oci_jms.PerformanceTuningAnalysis, error) {
	result := oci_jms.PerformanceTuningAnalysis{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func PerformanceTuningAnalysisToMap(obj *oci_jms.PerformanceTuningAnalysis) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToPostInstallationActionSettings(fieldKeyFormat string) (oci_jms.PostInstallationActionSettings, error) {
	result := oci_jms.PostInstallationActionSettings{}

	if addLoggingHandler, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "add_logging_handler")); ok {
		tmp := addLoggingHandler.(bool)
		result.AddLoggingHandler = &tmp
	}

	if disabledTlsVersions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disabled_tls_versions")); ok {
		interfaces := disabledTlsVersions.([]interface{})
		tmp := make([]oci_jms.TlsVersionsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_jms.TlsVersionsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "disabled_tls_versions")) {
			result.DisabledTlsVersions = tmp
		}
	}

	if globalLoggingLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "global_logging_level")); ok {
		result.GlobalLoggingLevel = oci_jms.GlobalLoggingLevelEnum(globalLoggingLevel.(string))
	}

	if minimumKeySizeSettings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum_key_size_settings")); ok {
		if tmpList := minimumKeySizeSettings.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "minimum_key_size_settings"), 0)
			tmp, err := s.mapToMinimumKeySizeSettings(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert minimum_key_size_settings, encountered error: %v", err)
			}
			result.MinimumKeySizeSettings = &tmp
		}
	}

	if proxies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "proxies")); ok {
		if tmpList := proxies.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "proxies"), 0)
			tmp, err := s.mapToProxies(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert proxies, encountered error: %v", err)
			}
			result.Proxies = &tmp
		}
	}

	if shouldReplaceCertificatesOperatingSystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_replace_certificates_operating_system")); ok {
		tmp := shouldReplaceCertificatesOperatingSystem.(bool)
		result.ShouldReplaceCertificatesOperatingSystem = &tmp
	}

	return result, nil
}

func PostInstallationActionSettingsToMap(obj *oci_jms.PostInstallationActionSettings) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddLoggingHandler != nil {
		result["add_logging_handler"] = bool(*obj.AddLoggingHandler)
	}

	result["disabled_tls_versions"] = obj.DisabledTlsVersions

	result["global_logging_level"] = string(obj.GlobalLoggingLevel)

	if obj.MinimumKeySizeSettings != nil {
		result["minimum_key_size_settings"] = []interface{}{MinimumKeySizeSettingsToMap(obj.MinimumKeySizeSettings)}
	}

	if obj.Proxies != nil {
		result["proxies"] = []interface{}{ProxiesToMap(obj.Proxies)}
	}

	if obj.ShouldReplaceCertificatesOperatingSystem != nil {
		result["should_replace_certificates_operating_system"] = bool(*obj.ShouldReplaceCertificatesOperatingSystem)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToProxies(fieldKeyFormat string) (oci_jms.Proxies, error) {
	result := oci_jms.Proxies{}

	if ftpProxyHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ftp_proxy_host")); ok {
		tmp := ftpProxyHost.(string)
		result.FtpProxyHost = &tmp
	}

	if ftpProxyPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ftp_proxy_port")); ok {
		tmp := ftpProxyPort.(int)
		result.FtpProxyPort = &tmp
	}

	if httpProxyHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_proxy_host")); ok {
		tmp := httpProxyHost.(string)
		result.HttpProxyHost = &tmp
	}

	if httpProxyPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_proxy_port")); ok {
		tmp := httpProxyPort.(int)
		result.HttpProxyPort = &tmp
	}

	if httpsProxyHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "https_proxy_host")); ok {
		tmp := httpsProxyHost.(string)
		result.HttpsProxyHost = &tmp
	}

	if httpsProxyPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "https_proxy_port")); ok {
		tmp := httpsProxyPort.(int)
		result.HttpsProxyPort = &tmp
	}

	if socksProxyHost, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "socks_proxy_host")); ok {
		tmp := socksProxyHost.(string)
		result.SocksProxyHost = &tmp
	}

	if socksProxyPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "socks_proxy_port")); ok {
		tmp := socksProxyPort.(int)
		result.SocksProxyPort = &tmp
	}

	if useSystemProxies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "use_system_proxies")); ok {
		tmp := useSystemProxies.(bool)
		result.UseSystemProxies = &tmp
	}

	return result, nil
}

func ProxiesToMap(obj *oci_jms.Proxies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FtpProxyHost != nil {
		result["ftp_proxy_host"] = string(*obj.FtpProxyHost)
	}

	if obj.FtpProxyPort != nil {
		result["ftp_proxy_port"] = int(*obj.FtpProxyPort)
	}

	if obj.HttpProxyHost != nil {
		result["http_proxy_host"] = string(*obj.HttpProxyHost)
	}

	if obj.HttpProxyPort != nil {
		result["http_proxy_port"] = int(*obj.HttpProxyPort)
	}

	if obj.HttpsProxyHost != nil {
		result["https_proxy_host"] = string(*obj.HttpsProxyHost)
	}

	if obj.HttpsProxyPort != nil {
		result["https_proxy_port"] = int(*obj.HttpsProxyPort)
	}

	if obj.SocksProxyHost != nil {
		result["socks_proxy_host"] = string(*obj.SocksProxyHost)
	}

	if obj.SocksProxyPort != nil {
		result["socks_proxy_port"] = int(*obj.SocksProxyPort)
	}

	if obj.UseSystemProxies != nil {
		result["use_system_proxies"] = bool(*obj.UseSystemProxies)
	}

	return result
}

func (s *JmsFleetAdvancedFeatureConfigurationResourceCrud) mapToSummarizedEventsLog(fieldKeyFormat string) (oci_jms.SummarizedEventsLog, error) {
	result := oci_jms.SummarizedEventsLog{}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func SummarizedEventsLogToMap(obj *oci_jms.SummarizedEventsLog) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}
