// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementExadataInfrastructureFleetMetricDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementExadataInfrastructureFleetMetric,
		Schema: map[string]*schema.Schema{
			"compare_baseline_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compare_target_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compare_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_by_exadata_infrastructure_deployment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter_by_exadata_infrastructure_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"exadata_infrastructure_fleet_summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregated_metrics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"baseline_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"dimension_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"dimension_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"percentage_change": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"target_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"inventory": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"deployment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"inventory_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rack_size": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"fleet_exadata_infrastructures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"infrastructure_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"infrastructure_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metrics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"baseline_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"dimension_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"dimension_value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"percentage_change": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"target_value": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"timestamp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"number_of_db_systems": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rack_size": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_server_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDatabaseManagementExadataInfrastructureFleetMetric(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementExadataInfrastructureFleetMetricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementExadataInfrastructureFleetMetricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetExadataInfrastructureFleetHealthMetricsResponse
}

func (s *DatabaseManagementExadataInfrastructureFleetMetricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementExadataInfrastructureFleetMetricDataSourceCrud) Get() error {
	request := oci_database_management.GetExadataInfrastructureFleetHealthMetricsRequest{}

	if compareBaselineTime, ok := s.D.GetOkExists("compare_baseline_time"); ok {
		tmp := compareBaselineTime.(string)
		request.CompareBaselineTime = &tmp
	}

	if compareTargetTime, ok := s.D.GetOkExists("compare_target_time"); ok {
		tmp := compareTargetTime.(string)
		request.CompareTargetTime = &tmp
	}

	if compareType, ok := s.D.GetOkExists("compare_type"); ok {
		request.CompareType = oci_database_management.GetExadataInfrastructureFleetHealthMetricsCompareTypeEnum(compareType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if filterByExadataInfrastructureDeploymentType, ok := s.D.GetOkExists("filter_by_exadata_infrastructure_deployment_type"); ok {
		request.FilterByExadataInfrastructureDeploymentType = oci_database_management.GetExadataInfrastructureFleetHealthMetricsFilterByExadataInfrastructureDeploymentTypeEnum(filterByExadataInfrastructureDeploymentType.(string))
	}

	if filterByExadataInfrastructureLifecycleState, ok := s.D.GetOkExists("filter_by_exadata_infrastructure_lifecycle_state"); ok {
		request.FilterByExadataInfrastructureLifecycleState = oci_database_management.ExadataInfrastructureLifecycleStateValuesStateEnum(filterByExadataInfrastructureLifecycleState.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetExadataInfrastructureFleetHealthMetrics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementExadataInfrastructureFleetMetricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementExadataInfrastructureFleetMetricDataSource-", DatabaseManagementExadataInfrastructureFleetMetricDataSource(), s.D))

	s.D.Set("compare_type", s.Res.CompareType)

	if s.Res.ExadataInfrastructureFleetSummary != nil {
		s.D.Set("exadata_infrastructure_fleet_summary", []interface{}{ExadataInfrastructureFleetSummaryToMap(s.Res.ExadataInfrastructureFleetSummary)})
	} else {
		s.D.Set("exadata_infrastructure_fleet_summary", nil)
	}

	fleetExadataInfrastructures := []interface{}{}
	for _, item := range s.Res.FleetExadataInfrastructures {
		fleetExadataInfrastructures = append(fleetExadataInfrastructures, ExadataInfrastructureUsageMetricsToMap(item))
	}
	s.D.Set("fleet_exadata_infrastructures", fleetExadataInfrastructures)

	return nil
}

func ExadataFleetMetricDefinitionToMap(obj oci_database_management.ExadataFleetMetricDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaselineValue != nil {
		result["baseline_value"] = float64(*obj.BaselineValue)
	}

	dimensions := []interface{}{}
	for _, item := range obj.Dimensions {
		dimensions = append(dimensions, ExadataMetricDimensionDefinitionToMap(item))
	}
	result["dimensions"] = dimensions

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.PercentageChange != nil {
		result["percentage_change"] = float64(*obj.PercentageChange)
	}

	if obj.TargetValue != nil {
		result["target_value"] = float64(*obj.TargetValue)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func ExadataFleetMetricSummaryDefinitionToMap(obj oci_database_management.ExadataFleetMetricSummaryDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaselineValue != nil {
		result["baseline_value"] = float64(*obj.BaselineValue)
	}

	dimensions := []interface{}{}
	for _, item := range obj.Dimensions {
		dimensions = append(dimensions, ExadataMetricDimensionDefinitionToMap(item))
	}
	result["dimensions"] = dimensions

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.PercentageChange != nil {
		result["percentage_change"] = float64(*obj.PercentageChange)
	}

	if obj.TargetValue != nil {
		result["target_value"] = float64(*obj.TargetValue)
	}

	if obj.Unit != nil {
		result["unit"] = string(*obj.Unit)
	}

	return result
}

func ExadataInfrastructureFleetStatusByCategoryToMap(obj oci_database_management.ExadataInfrastructureFleetStatusByCategory) map[string]interface{} {
	result := map[string]interface{}{}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.InventoryCount != nil {
		result["inventory_count"] = int(*obj.InventoryCount)
	}

	result["rack_size"] = string(obj.RackSize)

	return result
}

func ExadataInfrastructureFleetSummaryToMap(obj *oci_database_management.ExadataInfrastructureFleetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	aggregatedMetrics := []interface{}{}
	for _, item := range obj.AggregatedMetrics {
		aggregatedMetrics = append(aggregatedMetrics, ExadataFleetMetricSummaryDefinitionToMap(item))
	}
	result["aggregated_metrics"] = aggregatedMetrics

	inventory := []interface{}{}
	for _, item := range obj.Inventory {
		inventory = append(inventory, ExadataInfrastructureFleetStatusByCategoryToMap(item))
	}
	result["inventory"] = inventory

	return result
}

func ExadataInfrastructureUsageMetricsToMap(obj oci_database_management.ExadataInfrastructureUsageMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["deployment_type"] = string(obj.DeploymentType)

	if obj.InfrastructureId != nil {
		result["infrastructure_id"] = string(*obj.InfrastructureId)
	}

	if obj.InfrastructureName != nil {
		result["infrastructure_name"] = string(*obj.InfrastructureName)
	}

	metrics := []interface{}{}
	for _, item := range obj.Metrics {
		metrics = append(metrics, ExadataFleetMetricDefinitionToMap(item))
	}
	result["metrics"] = metrics

	if obj.NumberOfDbSystems != nil {
		result["number_of_db_systems"] = int(*obj.NumberOfDbSystems)
	}

	result["rack_size"] = string(obj.RackSize)

	result["state"] = string(obj.State)

	if obj.StorageServerCount != nil {
		result["storage_server_count"] = int(*obj.StorageServerCount)
	}

	return result
}

func ExadataMetricDimensionDefinitionToMap(obj oci_database_management.ExadataMetricDimensionDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DimensionName != nil {
		result["dimension_name"] = string(*obj.DimensionName)
	}

	if obj.DimensionValue != nil {
		result["dimension_value"] = string(*obj.DimensionValue)
	}

	return result
}
