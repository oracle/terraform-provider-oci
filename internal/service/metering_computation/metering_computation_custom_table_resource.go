// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_metering_computation "github.com/oracle/oci-go-sdk/v58/usageapi"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationCustomTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMeteringComputationCustomTable,
		Read:     readMeteringComputationCustomTable,
		Update:   updateMeteringComputationCustomTable,
		Delete:   deleteMeteringComputationCustomTable,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"saved_custom_table": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"display_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"column_group_by": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compartment_depth": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"group_by_tag": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"row_group_by": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"version": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"saved_report_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createMeteringComputationCustomTable(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.CreateResource(d, sync)
}

func readMeteringComputationCustomTable(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

func updateMeteringComputationCustomTable(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMeteringComputationCustomTable(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationCustomTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MeteringComputationCustomTableResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_metering_computation.UsageapiClient
	Res                    *oci_metering_computation.CustomTable
	DisableNotFoundRetries bool
}

func (s *MeteringComputationCustomTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MeteringComputationCustomTableResourceCrud) Create() error {
	request := oci_metering_computation.CreateCustomTableRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if savedCustomTable, ok := s.D.GetOkExists("saved_custom_table"); ok {
		if tmpList := savedCustomTable.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "saved_custom_table", 0)
			tmp, err := s.mapToSavedCustomTable(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SavedCustomTable = &tmp
		}
	}

	if savedReportId, ok := s.D.GetOkExists("saved_report_id"); ok {
		tmp := savedReportId.(string)
		request.SavedReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.CreateCustomTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomTable
	return nil
}

func (s *MeteringComputationCustomTableResourceCrud) Get() error {
	request := oci_metering_computation.GetCustomTableRequest{}

	tmp := s.D.Id()
	request.CustomTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.GetCustomTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomTable
	return nil
}

func (s *MeteringComputationCustomTableResourceCrud) Update() error {
	request := oci_metering_computation.UpdateCustomTableRequest{}

	tmp := s.D.Id()
	request.CustomTableId = &tmp

	if savedCustomTable, ok := s.D.GetOkExists("saved_custom_table"); ok {
		if tmpList := savedCustomTable.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "saved_custom_table", 0)
			tmp, err := s.mapToSavedCustomTable(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SavedCustomTable = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	response, err := s.Client.UpdateCustomTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomTable
	return nil
}

func (s *MeteringComputationCustomTableResourceCrud) Delete() error {
	request := oci_metering_computation.DeleteCustomTableRequest{}

	tmp := s.D.Id()
	request.CustomTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "metering_computation")

	_, err := s.Client.DeleteCustomTable(context.Background(), request)
	return err
}

func (s *MeteringComputationCustomTableResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.SavedCustomTable != nil {
		s.D.Set("saved_custom_table", []interface{}{SavedCustomTableToMap(s.Res.SavedCustomTable)})
	} else {
		s.D.Set("saved_custom_table", nil)
	}

	if s.Res.SavedReportId != nil {
		s.D.Set("saved_report_id", *s.Res.SavedReportId)
	}

	return nil
}

func CustomTableSummaryToMap(obj oci_metering_computation.CustomTableSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SavedCustomTable != nil {
		result["saved_custom_table"] = []interface{}{SavedCustomTableToMap(obj.SavedCustomTable)}
	}

	return result
}

func (s *MeteringComputationCustomTableResourceCrud) mapToSavedCustomTable(fieldKeyFormat string) (oci_metering_computation.SavedCustomTable, error) {
	result := oci_metering_computation.SavedCustomTable{}

	if columnGroupBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "column_group_by")); ok {
		interfaces := columnGroupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "column_group_by")) {
			result.ColumnGroupBy = tmp
		}
	}

	if compartmentDepth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_depth")); ok {
		tmp := float32(compartmentDepth.(float64))
		result.CompartmentDepth = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if groupByTag, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_by_tag")); ok {
		interfaces := groupByTag.([]interface{})
		tmp := make([]oci_metering_computation.Tag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "group_by_tag"), stateDataIndex)
			converted, err := s.mapToTag(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_by_tag")) {
			result.GroupByTag = tmp
		}
	}

	if rowGroupBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "row_group_by")); ok {
		interfaces := rowGroupBy.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "row_group_by")) {
			result.RowGroupBy = tmp
		}
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := float32(version.(float64))
		result.Version = &tmp
	}

	return result, nil
}

func SavedCustomTableToMap(obj *oci_metering_computation.SavedCustomTable) map[string]interface{} {
	result := map[string]interface{}{}

	result["column_group_by"] = obj.ColumnGroupBy

	if obj.CompartmentDepth != nil {
		result["compartment_depth"] = float32(*obj.CompartmentDepth)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	groupByTag := []interface{}{}
	for _, item := range obj.GroupByTag {
		groupByTag = append(groupByTag, TagToMap(item))
	}
	result["group_by_tag"] = groupByTag

	result["row_group_by"] = obj.RowGroupBy

	if obj.Version != nil {
		result["version"] = float32(*obj.Version)
	}

	return result
}

func (s *MeteringComputationCustomTableResourceCrud) mapToTag(fieldKeyFormat string) (oci_metering_computation.Tag, error) {
	result := oci_metering_computation.Tag{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func TagToMap(obj oci_metering_computation.Tag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
