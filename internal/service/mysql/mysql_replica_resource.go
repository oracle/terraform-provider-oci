// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlReplicaResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createMysqlReplica,
		Read:   readMysqlReplica,
		Update: updateMysqlReplica,
		Delete: deleteMysqlReplica,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"is_delete_protected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"replica_overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"configuration_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mysql_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"security_attributes": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"telemetry_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 0,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"logs": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: false,
										MinItems: 0,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"destination": {
													Type:     schema.TypeString,
													Required: true,
												},
												"destination_configurations": {
													Type:     schema.TypeSet,
													Required: true,
													Set:      destinationConfigurationsHashCodeForSets,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"key": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},
												"log_types": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
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

						// Computed
					},
				},
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encrypt_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key_generation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_x": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secure_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"certificate_generation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"telemetry_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"logs": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"destination_configurations": {
										Type:     schema.TypeSet,
										Computed: true,
										Set:      destinationConfigurationsHashCodeForSets,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"log_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
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

func createMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.CreateResource(d, sync)
}

func readMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlReplicaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.ReplicasClient
	Res                    *oci_mysql.Replica
	DisableNotFoundRetries bool
}

func (s *MysqlReplicaResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlReplicaResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateCreating),
		string(oci_mysql.ReplicaLifecycleStateUpdating),
	}
}

func (s *MysqlReplicaResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateActive),
		string(oci_mysql.ReplicaLifecycleStateNeedsAttention),
	}
}

func (s *MysqlReplicaResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateDeleting),
	}
}

func (s *MysqlReplicaResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateDeleted),
	}
}

func (s *MysqlReplicaResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateUpdating),
	}
}

func (s *MysqlReplicaResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateActive),
		string(oci_mysql.ReplicaLifecycleStateNeedsAttention),
	}
}

func (s *MysqlReplicaResourceCrud) Create() error {
	request := oci_mysql.CreateReplicaRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDeleteProtected, ok := s.D.GetOkExists("is_delete_protected"); ok {
		tmp := isDeleteProtected.(bool)
		request.IsDeleteProtected = &tmp
	}

	if replicaOverrides, ok := s.D.GetOkExists("replica_overrides"); ok {
		if tmpList := replicaOverrides.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_overrides", 0)
			tmp, err := s.mapToReplicaOverrides(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicaOverrides = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replica
	return nil
}

func (s *MysqlReplicaResourceCrud) Get() error {
	request := oci_mysql.GetReplicaRequest{}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replica
	return nil
}

func (s *MysqlReplicaResourceCrud) Update() error {
	request := oci_mysql.UpdateReplicaRequest{}

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

	if isDeleteProtected, ok := s.D.GetOkExists("is_delete_protected"); ok {
		tmp := isDeleteProtected.(bool)
		request.IsDeleteProtected = &tmp
	}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	if replicaOverrides, ok := s.D.GetOkExists("replica_overrides"); ok {
		if tmpList := replicaOverrides.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_overrides", 0)
			tmp, err := s.mapToReplicaOverrides(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicaOverrides = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateReplica(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlReplicaResourceCrud) Delete() error {
	request := oci_mysql.DeleteReplicaRequest{}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteReplica(context.Background(), request)
	return err
}

func (s *MysqlReplicaResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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

	if s.Res.EncryptData != nil {
		s.D.Set("encrypt_data", []interface{}{EncryptDataDetailsToMap(s.Res.EncryptData)})
	} else {
		s.D.Set("encrypt_data", nil)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsDeleteProtected != nil {
		s.D.Set("is_delete_protected", *s.Res.IsDeleteProtected)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	if s.Res.ReplicaOverrides != nil {
		s.D.Set("replica_overrides", []interface{}{ReplicaOverridesToMap(s.Res.ReplicaOverrides, false)})
	} else {
		s.D.Set("replica_overrides", nil)
	}

	if s.Res.SecureConnections != nil {
		s.D.Set("secure_connections", []interface{}{SecureConnectionDetailsToMap(s.Res.SecureConnections)})
	} else {
		s.D.Set("secure_connections", nil)
	}

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TelemetryConfiguration != nil {
		s.D.Set("telemetry_configuration", []interface{}{TelemetryConfigurationDetailsToMap(s.Res.TelemetryConfiguration, false)})
	} else {
		s.D.Set("telemetry_configuration", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *MysqlReplicaResourceCrud) mapToDestinationConfiguration(fieldKeyFormat string) (oci_mysql.DestinationConfiguration, error) {
	result := oci_mysql.DestinationConfiguration{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *MysqlReplicaResourceCrud) mapToLoggingDestinationConfiguration(fieldKeyFormat string) (oci_mysql.LoggingDestinationConfiguration, error) {
	result := oci_mysql.LoggingDestinationConfiguration{}

	if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
		result.Destination = oci_mysql.LoggingDestinationConfigurationDestinationEnum(destination.(string))
	}

	if destinationConfigurations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_configurations")); ok {
		set := destinationConfigurations.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_mysql.DestinationConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := destinationConfigurationsHashCodeForSets(interfaces[i])
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination_configurations"), stateDataIndex)
			converted, err := s.mapToDestinationConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination_configurations")) {
			result.DestinationConfigurations = tmp
		}
	}

	if logTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_types")); ok {
		interfaces := logTypes.([]interface{})
		tmp := make([]oci_mysql.LoggingDestinationConfigurationLogTypesEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_mysql.LoggingDestinationConfigurationLogTypesEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "log_types")) {
			result.LogTypes = tmp
		}
	}

	return result, nil
}

func (s *MysqlReplicaResourceCrud) mapToReplicaOverrides(fieldKeyFormat string) (oci_mysql.ReplicaOverrides, error) {
	result := oci_mysql.ReplicaOverrides{}

	configurationIdField := fmt.Sprintf(fieldKeyFormat, "configuration_id")
	if configurationId, ok := s.D.GetOkExists(configurationIdField); ok && s.D.HasChange(configurationIdField) {
		tmp := configurationId.(string)
		result.ConfigurationId = &tmp
	}

	mysqlVersionField := fmt.Sprintf(fieldKeyFormat, "mysql_version")
	if mysqlVersion, ok := s.D.GetOkExists(mysqlVersionField); ok && s.D.HasChange(mysqlVersionField) {
		tmp := mysqlVersion.(string)
		result.MysqlVersion = &tmp
	}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	securityAttributesField := fmt.Sprintf(fieldKeyFormat, "security_attributes")
	if securityAttributes, ok := s.D.GetOkExists(securityAttributesField); ok && s.D.HasChange(securityAttributesField) {
		result.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
	}

	shapeNameField := fmt.Sprintf(fieldKeyFormat, "shape_name")
	if shapeName, ok := s.D.GetOkExists(shapeNameField); ok && s.D.HasChange(shapeNameField) {
		tmp := shapeName.(string)
		result.ShapeName = &tmp
	}

	telemetryConfigurationField := fmt.Sprintf(fieldKeyFormat, "telemetry_configuration")
	if telemetryConfiguration, ok := s.D.GetOkExists(telemetryConfigurationField); ok && s.D.HasChange(telemetryConfigurationField) {
		if tmpList := telemetryConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", telemetryConfigurationField, 0)
			tmp, err := s.mapToTelemetryConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert telemetry_configuration, encountered error: %v", err)
			}
			result.TelemetryConfiguration = &tmp
		}
	}

	return result, nil
}

func ReplicaOverridesToMap(obj *oci_mysql.ReplicaOverrides, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigurationId != nil {
		result["configuration_id"] = string(*obj.ConfigurationId)
	}

	if obj.MysqlVersion != nil {
		result["mysql_version"] = string(*obj.MysqlVersion)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SecurityAttributes != nil {
		result["security_attributes"] = tfresource.SecurityAttributesToMap(obj.SecurityAttributes)
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	if obj.TelemetryConfiguration != nil {
		result["telemetry_configuration"] = []interface{}{TelemetryConfigurationDetailsToMap(obj.TelemetryConfiguration, datasource)}
	}

	return result
}

func (s *MysqlReplicaResourceCrud) mapToTelemetryConfigurationDetails(fieldKeyFormat string) (oci_mysql.TelemetryConfigurationDetails, error) {
	result := oci_mysql.TelemetryConfigurationDetails{}

	if logs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logs")); ok {
		interfaces := logs.([]interface{})
		tmp := make([]oci_mysql.LoggingDestinationConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "logs"), stateDataIndex)
			converted, err := s.mapToLoggingDestinationConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "logs")) {
			result.Logs = tmp
		}
		if len(tmp) == 0 {
			result.Logs = nil
		}
	}

	return result, nil
}
