// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowRunStatementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataflowRunStatement,
		Read:     readDataflowRunStatement,
		Delete:   deleteDataflowRunStatement,
		Schema: map[string]*schema.Schema{
			// Required
			"code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"run_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"output": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"data": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"type": {
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
						"error_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"traceback": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"progress": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDataflowRunStatement(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.CreateResource(d, sync)
}

func readDataflowRunStatement(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

func deleteDataflowRunStatement(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunStatementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataflowRunStatementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataflow.DataFlowClient
	Res                    *oci_dataflow.Statement
	DisableNotFoundRetries bool
}

func (s *DataflowRunStatementResourceCrud) ID() string {
	return GetRunStatementCompositeId(s.D.Get("run_id").(string), *s.Res.Id)
}

func (s *DataflowRunStatementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dataflow.StatementLifecycleStateAccepted),
		string(oci_dataflow.StatementLifecycleStateInProgress),
	}
}

func (s *DataflowRunStatementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataflow.StatementLifecycleStateFailed),
		string(oci_dataflow.StatementLifecycleStateSucceeded),
	}
}

func (s *DataflowRunStatementResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataflow.StatementLifecycleStateAccepted),
		string(oci_dataflow.StatementLifecycleStateInProgress),
		string(oci_dataflow.StatementLifecycleStateCancelling),
	}
}

func (s *DataflowRunStatementResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataflow.StatementLifecycleStateFailed),
		string(oci_dataflow.StatementLifecycleStateSucceeded),
		string(oci_dataflow.StatementLifecycleStateCancelled),
	}
}

func (s *DataflowRunStatementResourceCrud) Create() error {
	request := oci_dataflow.CreateStatementRequest{}

	if code, ok := s.D.GetOkExists("code"); ok {
		tmp := code.(string)
		request.Code = &tmp
	}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.CreateStatement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Statement
	s.D.SetId(GetRunStatementCompositeId(*response.RunId, *response.Id))

	return nil
}

func (s *DataflowRunStatementResourceCrud) Get() error {
	request := oci_dataflow.GetStatementRequest{}

	runId, statementId, err := parseRunStatementCompositeId(s.D.Id())

	if err == nil {
		request.RunId = &runId
		tmp := strconv.FormatInt(statementId, 10)
		request.StatementId = &tmp
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.GetStatement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Statement
	return nil
}

func (s *DataflowRunStatementResourceCrud) Delete() error {
	request := oci_dataflow.DeleteStatementRequest{}

	runId, statementId, err := parseRunStatementCompositeId(s.D.Id())

	if err == nil {
		request.RunId = &runId
		tmp := strconv.FormatInt(statementId, 10)
		request.StatementId = &tmp
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, error := s.Client.DeleteStatement(context.Background(), request)
	return error
}

func (s *DataflowRunStatementResourceCrud) SetData() error {

	runId, _, err := parseRunStatementCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("run_id", &runId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Code != nil {
		s.D.Set("code", *s.Res.Code)
	}

	if s.Res.Output != nil {
		s.D.Set("output", []interface{}{StatementOutputToMap(s.Res.Output)})
	} else {
		s.D.Set("output", nil)
	}

	if s.Res.Progress != nil {
		s.D.Set("progress", *s.Res.Progress)
	}

	if s.Res.RunId != nil {
		s.D.Set("run_id", *s.Res.RunId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetRunStatementCompositeId(runId string, statementId int64) string {
	runId = url.PathEscape(runId)
	compositeId := "runs/" + runId + "/statements/" + strconv.FormatInt(statementId, 10)
	return compositeId
}

func parseRunStatementCompositeId(compositeId string) (runId string, statementId int64, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("runs/.*/statements/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	runId, _ = url.PathUnescape(parts[1])
	statementId, _ = strconv.ParseInt(parts[3], 10, 64)

	return
}

func StatementOutputToMap(obj *oci_dataflow.StatementOutput) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Data != nil {
		dataArray := []interface{}{}
		if dataMap := StatementOutputDataToMap(&obj.Data); dataMap != nil {
			dataArray = append(dataArray, dataMap)
		}
		result["data"] = dataArray
	}

	if obj.ErrorName != nil {
		result["error_name"] = string(*obj.ErrorName)
	}

	if obj.ErrorValue != nil {
		result["error_value"] = string(*obj.ErrorValue)
	}

	result["status"] = string(obj.Status)

	result["traceback"] = obj.Traceback
	result["traceback"] = obj.Traceback

	return result
}

func StatementOutputDataToMap(obj *oci_dataflow.StatementOutputData) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_dataflow.ImagePngStatementOutputData:
		result["type"] = "IMAGE_PNG"
		if v.Value != nil {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, v.Value)
			if err != nil {
				result["value"] = buf.String()
			} else {
				result["value"] = ""
			}
		}
	case oci_dataflow.TextHtmlStatementOutputData:
		result["type"] = "TEXT_HTML"

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_dataflow.TextPlainStatementOutputData:
		result["type"] = "TEXT_PLAIN"

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func StatementSummaryToMap(obj oci_dataflow.StatementSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = strconv.FormatInt(*obj.Id, 10)
	}

	if obj.RunId != nil {
		result["run_id"] = string(*obj.RunId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCompleted != nil {
		result["time_completed"] = obj.TimeCompleted.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
