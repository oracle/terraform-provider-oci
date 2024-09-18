package devops

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Struct to extract deployment Parameters from the resourceDiff construct
type DeploymentParams struct {
	deployPipelineId    string
	deploymentType      string
	deployStageId       string
	deploymentArguments []interface{}

	debug bool
}

var enableFeature bool = true
var helmDiffDeploymentExecutionProgress oci_devops.DeploymentExecutionProgress

// Deployment Resource Custom Diff Implementation for invoking Helm Diff
func resourceOkeClusterHelmReleaseDiff(c context.Context, d *schema.ResourceDiff, m interface{}) error {
	deployParams := extractDeploymentParameters(d)
	devopsClient := m.(*client.OracleClients).DevopsClient()

	// Add DRY_RUN = true if there is a Helm Artifact in pipeline
	// and trigger a deployment
	if isHelmArtifactInPipeline(devopsClient, deployParams.deployPipelineId) {
		checkPipelineForHydrationWorkRequest(deployParams.deployPipelineId, m)
		deploymentArguments := deployParams.deploymentArguments

		if len(deploymentArguments) > 0 {
			arguments := deploymentArguments[0].(map[string]interface{})
			items := arguments["items"].([]interface{})

			// Loop through all the deployment arguments
			for _, item := range items {
				name := item.(map[string]interface{})["name"].(string)
				value := item.(map[string]interface{})["value"].(string)

				// Check each of the deployment argument to see if PLAN_DRY_RUN is set to true
				if name == "PLAN_DRY_RUN" && value == "true" {
					if enableFeature == true {
						enableFeature = false
						var deploymentDetails oci_devops.CreateDeploymentDetails
						deploymentName := fmt.Sprintf("HelmDryRunDeployment%s", time.Now().Format(time.RFC3339))
						switch deployParams.deploymentType {
						case "PIPELINE_DEPLOYMENT":
							deploymentDetails = oci_devops.CreateDeployPipelineDeploymentDetails{
								DeployPipelineId: &deployParams.deployPipelineId,
								DisplayName:      &deploymentName,
								DeploymentArguments: &oci_devops.DeploymentArgumentCollection{Items: []oci_devops.DeploymentArgument{oci_devops.DeploymentArgument{Name: common.String("DRY_RUN"),
									Value: common.String("true"),
								}}},
							}
						case "SINGLE_STAGE_DEPLOYMENT":
							log.Printf("[WARN] Creating a single stage deployment with deploy stage id %s", deployParams.deployStageId)
							deploymentDetails = oci_devops.CreateSingleDeployStageDeploymentDetails{
								DeployPipelineId: &deployParams.deployPipelineId,
								DeployStageId:    &deployParams.deployStageId,
								DisplayName:      &deploymentName,
								DeploymentArguments: &oci_devops.DeploymentArgumentCollection{Items: []oci_devops.DeploymentArgument{oci_devops.DeploymentArgument{Name: common.String("DRY_RUN"),
									Value: common.String("true"),
								}}},
							}
						default:
							log.Printf("[WARN] Deployment type not supported for rendering Helm Diff")
							return nil
						}
						if deployParams.deploymentType == "SINGLE_STAGE_DEPLOYMENT" || deployParams.deploymentType == "PIPELINE_DEPLOYMENT" {
							// API call for deployment
							log.Printf("[WARN] Deployment Details %s", deploymentDetails)
							req := oci_devops.CreateDeploymentRequest{CreateDeploymentDetails: deploymentDetails}
							resp, _ := devopsClient.CreateDeployment(context.Background(), req)

							// Get API Deployment call for querying status
							request := oci_devops.GetDeploymentRequest{}
							log.Printf("[WARN] Deployment request %s", request)
							request.DeploymentId = resp.GetId()
							deploymentResponse, _ := devopsClient.GetDeployment(context.Background(), request)
							status := deploymentResponse.GetLifecycleState()

							// Wait till deployment reaches a terminal state
							for status != oci_devops.DeploymentLifecycleStateSucceeded && status != oci_devops.DeploymentLifecycleStateFailed {
								deploymentResponse, _ = devopsClient.GetDeployment(context.Background(), request)
								time.Sleep(2 * time.Second)
								status = deploymentResponse.GetLifecycleState()
								log.Printf("[INFO] Waiting for the Helm Diff deployment to reach a terminal state - %s", status)
							}
							// Extract the deployment execution progress that contains Helm Diff
							progress := deploymentResponse.GetDeploymentExecutionProgress()
							helmDiffDeploymentExecutionProgress = *progress
						}
					}
					setHelmExecutionProgress(d, helmDiffDeploymentExecutionProgress)
				}
			}
		}
	}
	return nil
}

// Editing the Deployment execution progress to display only Helm stage progress during DRY_RUN plan
func setHelmExecutionProgress(d *schema.ResourceDiff, resultProgress oci_devops.DeploymentExecutionProgress) {
	// Post-process the Helm Diff Deployment result to retrieve progress for only Helm Stages
	var newStageProgress []oci_devops.DeployStageExecutionProgress
	depStageProgress := resultProgress.DeployStageExecutionProgress

	for _, eachStageProgress := range depStageProgress {
		switch v := eachStageProgress.(type) {
		case oci_devops.OkeHelmChartDeploymentStageExecutionProgress:
			diffString := v.HelmDiff
			formattedDiff := formatHelmDiff(*diffString)
			v.HelmDiff = &formattedDiff
			v.TimeStarted = nil
			v.TimeFinished = nil
			v.DeployStageExecutionProgressDetails = nil
			newStageProgress = append(newStageProgress, v)
		default:
			newStageProgress = append(newStageProgress, v)
		}
	}

	var helmDiffDeploymentExecutionProgress oci_devops.DeploymentExecutionProgress
	helmDiffDeploymentExecutionProgress.TimeStarted = nil
	helmDiffDeploymentExecutionProgress.TimeFinished = nil

	helmDiffDeploymentExecutionProgress.DeployStageExecutionProgress = make(map[string]oci_devops.DeployStageExecutionProgress)
	for _, deployStageExecutionProgress := range newStageProgress {
		helmDiffDeploymentExecutionProgress.DeployStageExecutionProgress[*deployStageExecutionProgress.GetDeployStageId()] = deployStageExecutionProgress
	}

	// Display the post-processed Helm Diff Execution Progress to the user in tf plan
	err := d.SetNew("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(&helmDiffDeploymentExecutionProgress)})
	if err != nil {
		log.Printf("[INFO] [HELMDIFF] DeployStageExecutionProgress SetNew Error: %s", err)
	}
}

// Formatting the Diff to show a message when no changes are detected
func formatHelmDiff(helmDiff string) string {

	result := make(map[string][]string)
	var name string
	var values []string
	var retval []string

	// remove any trailing newline as it will cause Problems(tm)
	helmDiff = strings.TrimRight(helmDiff, "\n")

	inside := false
	if helmDiff != "" {
		for _, line := range strings.Split(helmDiff, "\n") {
			// Group the output into the categories listed by Helm Diff plugin
			if strings.HasSuffix(line, "has changed:") || strings.HasSuffix(line, "has been added:") || strings.HasSuffix(line, "has been removed:") {
				if name != "" {
					result[name] = values
				}
				name = line
				values = []string{}
				inside = true
				continue
			}
			if inside {
				values = append(values, line)
			}
		}
		// get the last one
		result[name] = values
	}
	var result_keys []string
	for k := range result {
		result_keys = append(result_keys, k)
	}
	sort.Strings(result_keys)

	for _, k := range result_keys {
		retval = append(retval, k)
		for _, v := range result[k] {
			retval = append(retval, v)
		}
	}
	return strings.Join(retval, string('\n'))
}

// Check all the artifacts in the pipeline to determine if Helm stage exists
func isHelmArtifactInPipeline(client *oci_devops.DevopsClient, pipelineId string) bool {
	request := oci_devops.GetDeployPipelineRequest{}
	request.DeployPipelineId = &pipelineId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
	response, _ := client.GetDeployPipeline(context.Background(), request)

	result := response.DeployPipelineArtifacts
	if result != nil {
		for _, item := range result.Items {
			artifactRequest := oci_devops.GetDeployArtifactRequest{}
			artifactRequest.DeployArtifactId = item.DeployArtifactId
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "devops")
			artifactResponse, _ := client.GetDeployArtifact(context.Background(), artifactRequest)

			if artifactResponse.DeployArtifactType == oci_devops.DeployArtifactDeployArtifactTypeHelmChart {
				return true
			}
		}
	}

	return false
}

// Extract the deployment parameters from the diff construct for populating API calls
func extractDeploymentParameters(d *schema.ResourceDiff) *DeploymentParams {
	deployPipelineId := d.Get("deploy_pipeline_id").(string)
	deploymentType := d.Get("deployment_type").(string)
	deployStageId := d.Get("deploy_stage_id").(string)
	deploymentArguments := d.Get("deployment_arguments").([]interface{})

	param := &DeploymentParams{}
	param.deploymentArguments = deploymentArguments
	param.deployPipelineId = deployPipelineId
	param.deploymentType = deploymentType
	param.deployStageId = deployStageId
	return param

}

func checkPipelineForHydrationWorkRequest(pipelineId string, m interface{}) {
	common.Debugf("checkPipelineForHydrationWorkRequest: enter, pipelineId=%v\n", pipelineId)

	getDeployPipelineRequest := oci_devops.GetDeployPipelineRequest{}
	listWorkRequestsRequest := oci_devops.ListWorkRequestsRequest{}
	devOpsClient := m.(*client.OracleClients).DevopsClient()

	listWorkRequestsRequest.ResourceId = &pipelineId
	getDeployPipelineRequest.DeployPipelineId = &pipelineId
	// Get the compartment id from the Pipeline
	common.Debugf("checkPipelineForHydrationWorkRequest: getDeployPipelineRequest= %v\n", getDeployPipelineRequest)
	getDeployPipelineResponse, err := devOpsClient.GetDeployPipeline(context.Background(), getDeployPipelineRequest)
	if err != nil {
		common.Debugf("checkPipelineForHydrationWorkRequest: getDeployPipelineResponse err= %v\n", err)
	} else {
		common.Debugf("checkPipelineForHydrationWorkRequest: getDeployPipelineResponse.OpcRequestId= %v\n", *getDeployPipelineResponse.OpcRequestId)
		deployPipeline := getDeployPipelineResponse.DeployPipeline
		listWorkRequestsRequest.CompartmentId = deployPipeline.CompartmentId
	}

	waitForHydrationWorkRequest(listWorkRequestsRequest, devOpsClient)
	common.Debugf("checkPipelineForHydrationWorkRequest: exit")
}

func checkForHydrationWorkRequest(d *schema.ResourceData, m interface{}) {
	common.Debugf("checkForHydrationWorkRequest: enter, ResourceData= %v\n", d)

	getDeployPipelineRequest := oci_devops.GetDeployPipelineRequest{}
	listWorkRequestsRequest := oci_devops.ListWorkRequestsRequest{}
	devOpsClient := m.(*client.OracleClients).DevopsClient()

	if pipelineId, ok := d.GetOkExists("deploy_pipeline_id"); ok {
		common.Debugf("checkForHydrationWorkRequest: pipelineId= %v\n", pipelineId)
		tmp := pipelineId.(string)
		listWorkRequestsRequest.ResourceId = &tmp
		getDeployPipelineRequest.DeployPipelineId = &tmp
	}
	if compartmentId, ok := d.GetOkExists("compartment_id"); ok {
		common.Debugf("checkForHydrationWorkRequest: compartmentId= %v\n", compartmentId)
		tmp := compartmentId.(string)
		listWorkRequestsRequest.CompartmentId = &tmp
	} else {
		// Get the compartment id from the Pipeline
		common.Debugf("checkForHydrationWorkRequest: getDeployPipelineRequest= %v\n", getDeployPipelineRequest)
		getDeployPipelineResponse, err := devOpsClient.GetDeployPipeline(context.Background(), getDeployPipelineRequest)
		if err != nil {
			common.Debugf("checkForHydrationWorkRequest: getDeployPipelineResponse err= %v\n", err)
		} else {
			common.Debugf("checkForHydrationWorkRequest: getDeployPipelineResponse.OpcRequestId= %v\n", *getDeployPipelineResponse.OpcRequestId)
			deployPipeline := getDeployPipelineResponse.DeployPipeline
			listWorkRequestsRequest.CompartmentId = deployPipeline.CompartmentId
		}
	}
	waitForHydrationWorkRequest(listWorkRequestsRequest, devOpsClient)
	common.Debugf("checkForHydrationWorkRequest: exit")
}
func waitForHydrationWorkRequest(listWorkRequestsRequest oci_devops.ListWorkRequestsRequest, devOpsClient *oci_devops.DevopsClient) {
	common.Debugf("checkForHydrationWorkRequest: listWorkRequestsRequest= %v\n", listWorkRequestsRequest)
	if listWorkRequestsRequest.CompartmentId != nil {
		workRequestInProgress := true
		// Wait until all hydration work requests for the pipeline are complete.
		for workRequestInProgress {
			workRequestInProgress = false
		InProgress:
			log.Printf(" InProgress Block Execution")
			listWorkRequestsResponse, err := devOpsClient.ListWorkRequests(context.Background(), listWorkRequestsRequest)
			common.Debugf("checkForHydrationWorkRequest: listWorkRequestsResponse= %v\n", listWorkRequestsResponse)
			if err != nil {
				// If we can't list the work requests, we'll just continue with the deployment.
				common.Debugf("checkForHydrationWorkRequest: listWorkRequestResponse err= %v\n", err)
			} else {
				hasMorePages := true
				for hasMorePages {
					common.Debugf("checkForHydrationWorkRequest: listWorkRequestsResponse.OpcRequestId= %v\n", *listWorkRequestsResponse.OpcRequestId)
					workRequestCollection := listWorkRequestsResponse.WorkRequestCollection
					if len(workRequestCollection.Items) > 0 {
						common.Debugf("checkForHydrationWorkRequest: workRequestCollection.Items= %i\n", len(workRequestCollection.Items))
						for i, summary := range workRequestCollection.Items {
							if !(summary.Status == "SUCCEEDED" || summary.Status == "FAILED" ||
								summary.Status == "CANCELED" || summary.Status == "NEEDS_ATTENTION") {
								workRequestInProgress = true
								common.Debugf("checkForHydrationWorkRequest: WorkRequestSummary found in progress= %i %v\n", i, summary)
								goto InProgress
							}
						}
					}
					// Retrieve the next page of data
					if listWorkRequestsResponse.OpcNextPage != nil && len(*listWorkRequestsResponse.OpcNextPage) > 0 {
						common.Debugf("checkForHydrationWorkRequest: listWorkRequestResponse page = %v\n", *listWorkRequestsResponse.OpcNextPage)
						listWorkRequestsRequest.Page = listWorkRequestsResponse.OpcNextPage
						listWorkRequestsResponse, err = devOpsClient.ListWorkRequests(context.Background(), listWorkRequestsRequest)
						if err != nil {
							common.Debugf("checkForHydrationWorkRequest: listWorkRequestResponse page err= %v\n", err)
							hasMorePages = false
						}
						if listWorkRequestsResponse.OpcNextPage != nil && len(*listWorkRequestsResponse.OpcNextPage) > 0 {
							common.Debugf("checkForHydrationWorkRequest: listWorkRequestResponse next page = %v\n", *listWorkRequestsResponse.OpcNextPage)
						}
					} else {
						hasMorePages = false
					}
				}

			}
		}
	}
}
