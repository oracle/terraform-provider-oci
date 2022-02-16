// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"golang.org/x/mod/semver"

	"github.com/hashicorp/terraform-exec/tfinstall"

	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	tf_provider "github.com/terraform-providers/terraform-provider-oci/internal/provider"
	utils "github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

type ResourceDiscoveryStage int

const (
	Discovery       ResourceDiscoveryStage = 1
	GeneratingState                        = 2
)

var referenceMap map[string]string //	stores references to replace the ocids in config
var refMapLock sync.Mutex
var referenceResourceNameSet map[string]bool   // this set contains terraform resource names for the references in referenceMap
var failedResourceReferenceSet map[string]bool // stores the terraform reference name for failed resources, used to remove InterpolationString type values if a resource failed to import
var vars map[string]string
var resourceNameCount map[string]int
var resourceNameCountLock sync.Mutex
var resourcesMap map[string]*schema.Resource
var datasourcesMap map[string]*schema.Resource
var compartmentScopeServices []string
var tenancyScopeServices []string
var isMissingRequiredAttributes bool
var missingAttributesPerResourceLock sync.Mutex
var sem chan struct{}
var exportConfigProvider oci_common.ConfigurationProvider
var tfHclVersion TfHclVersion

func elapsed(what string, step *resourceDiscoveryBaseStep, stage ResourceDiscoveryStage) func() {
	start := time.Now()
	return func() {
		totalTime := time.Since(start)
		utils.Debugf("[DEBUG] %s took %v\n", what, totalTime)
		if step != nil {
			switch stage {
			case Discovery:
				step.updateTimeTakenForDiscovery(totalTime)
			case GeneratingState:
				step.updateTimeTakenForGeneratingState(totalTime)
			}
		}
	}
}

func init() {
	resourceNameCount = map[string]int{}
	vars = map[string]string{}
	referenceMap = map[string]string{}

	compartmentScopeServices = make([]string, len(compartmentResourceGraphs))
	idx := 0
	for mode := range compartmentResourceGraphs {
		compartmentScopeServices[idx] = mode
		idx++
	}

	tenancyScopeServices = make([]string, len(tenancyResourceGraphs))
	idx = 0
	for mode := range tenancyResourceGraphs {
		tenancyScopeServices[idx] = mode
		idx++
	}

	isMissingRequiredAttributes = false
}

func printResourceGraphResources(resourceGraphs map[string]TerraformResourceGraph, scope string) error {
	for graphName, resourceGraph := range resourceGraphs {
		// Need a set here because the same resource type may have multiple associations in the same graph
		// This avoids adding duplicates of those resource types
		resourceSet := map[string]bool{}
		for _, association := range resourceGraph {
			for _, hint := range association {
				if _, isResource := resourcesMap[hint.resourceClass]; isResource {
					resourceSet[hint.resourceClass] = true
				}
			}
		}

		if len(resourceSet) > 0 {
			supportedResources := make([]string, len(resourceSet))
			idx := 0
			for resourceClass := range resourceSet {
				supportedResources[idx] = resourceClass
				idx++
			}

			sort.Strings(supportedResources)
			utils.Logf("%s (%s-scope resources)", graphName, scope)
			utils.Log("===========")
			for _, resourceClass := range supportedResources {
				utils.Logf("- %s", resourceClass)
			}
			utils.Logln("")
		}
	}
	return nil
}

func RunListExportableResourcesCommand() error {
	resourcesMap = tf_provider.ResourcesMap()
	datasourcesMap = tf_provider.DataSourcesMap()

	utils.Logln("List of Discoverable Oracle Cloud Infrastructure Resources")

	if err := printResourceGraphResources(tenancyResourceGraphs, "tenancy"); err != nil {
		return err
	}

	if err := printResourceGraphResources(compartmentResourceGraphs, "compartment"); err != nil {
		return err
	}
	return nil
}

type ExportService struct {
	Name  string
	Scope string
}

const (
	TenancyScope     = "tenancy"
	CompartmentScope = "compartment"
)

func RunListExportableServicesCommand(listExportServicesPath string) error {

	utils.Logln("List Discoverable Oracle Cloud Infrastructure Services")

	services := make([]*ExportService, 0)
	for _, service := range tenancyScopeServices {
		services = append(services, &ExportService{
			Name:  service,
			Scope: TenancyScope,
		})
	}

	for _, service := range compartmentScopeServices {
		services = append(services, &ExportService{
			Name:  service,
			Scope: CompartmentScope,
		})
	}

	servicesJson, err := json.MarshalIndent(services, "", "")
	if err != nil {
		return fmt.Errorf("[ERROR] Error marshalling services to JSON: %v", err)
	}

	if listExportServicesPath != "" {
		if err := ioutil.WriteFile(listExportServicesPath, servicesJson, 0644); err != nil {
			return err
		} else {
			utils.Logf("[INFO] Services written to json file at: %s", listExportServicesPath)
		}
	}
	utils.Logf(string(servicesJson))
	return nil
}

type ExportCommandArgs struct {
	CompartmentId                *string
	CompartmentName              *string
	IDs                          []string
	Services                     []string
	OutputDir                    *string
	GenerateState                bool
	TFVersion                    *TfHclVersion
	RetryTimeout                 *string
	ExcludeServices              []string
	IsExportWithRelatedResources bool
	Parallelism                  int
}

func RunExportCommand(args *ExportCommandArgs) (err error, status Status) {
	defer func() {
		if r := recover(); r != nil {
			utils.Logf("[ERROR] panic in RunExportCommand, exiting with status %v", StatusFail)
			debug.PrintStack()
			err = errors.New("[ERROR] panic in RunExportCommand: unknown error occurred in export")
			status = StatusFail
		}
	}()
	resourcesMap = tf_provider.ResourcesMap()
	datasourcesMap = tf_provider.DataSourcesMap()

	if err := args.validate(); err != nil {
		return err, StatusFail
	}

	tfHclVersion = *args.TFVersion

	r := &schema.Resource{
		Schema: tf_provider.SchemaMap(),
	}
	d := r.Data(nil)

	err = readEnvironmentVars(d)
	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}

	clients, err := getExportConfig(d)
	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}

	if args.CompartmentName != nil && *args.CompartmentName != "" {
		var err error
		args.CompartmentId, err = resolveCompartmentId(clients.(*tf_client.OracleClients), args.CompartmentName)
		if err != nil {
			utils.Logln(err.Error())
			return err, StatusFail
		}
	}

	/* Getting the tenancy ocid from env var export_tenancy_id, if not specified get customer tenancy ocid from configuration provider */
	tenancyOcid := utils.GetEnvSettingWithBlankDefault("export_tenancy_id")

	if tenancyOcid == "" {
		/* Keep the tenancy lookup for backward compatibility */
		if args.CompartmentId != nil && *args.CompartmentId != "" {
			tenancyOcid, err = getTenancyOcidFromCompartment(clients.(*tf_client.OracleClients), *args.CompartmentId)
			if err != nil {
				utils.Logln(err.Error())
				return err, StatusFail
			}
		} else {
			// If compartment ocid not provided in arguments, get it from configuration provider
			tenancyId, exists := clients.(*tf_client.OracleClients).Configuration["tenancy_ocid"]
			if !exists {
				return fmt.Errorf("[ERROR] could not get a tenancy OCID during initialization"), StatusFail
			}
			tenancyOcid = tenancyId
		}

	}
	if args.Parallelism < 1 {
		return fmt.Errorf("[ERROR] invalid value for arument parallelism, specify a value >= 1"), StatusFail
	}

	sem = make(chan struct{}, args.Parallelism)

	ctx, err := createResourceDiscoveryContext(clients.(*tf_client.OracleClients), args, tenancyOcid)
	if err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}
	args.finalizeServices(ctx)

	/*
		Setting retry timeout to a lower value for resource discovery
		This is done to handle the 404 and 500 errors in case
		any resource is unavailable in a region or in case the service is down
		The time out value is configurable from export command
	*/
	if args.RetryTimeout != nil && *args.RetryTimeout != "" {
		tfresource.LongRetryTime = *tfresource.GetTimeoutDuration(*args.RetryTimeout)
		tfresource.ShortRetryTime = tfresource.LongRetryTime
	}

	utils.Logf("[INFO] resource discovery retry timeout duration set to %v", tfresource.ShortRetryTime)

	if err := runExportCommand(ctx); err != nil {
		utils.Logln(err.Error())
		return err, StatusFail
	}
	if len(ctx.errorList.errors) > 0 {
		// If the errors were from discovery of resources return partial success status
		ctx.printErrors()
		return nil, StatusPartialSuccess
	}
	return nil, StatusSuccess
}

func convertStringSliceToSet(slice []string, omitEmptyStrings bool) map[string]bool {
	result := map[string]bool{}
	for _, item := range slice {
		if omitEmptyStrings && item == "" {
			continue
		}
		result[item] = false
	}
	return result
}

func (args *ExportCommandArgs) finalizeServices(ctx *resourceDiscoveryContext) {
	if len(args.Services) == 0 {
		args.Services = compartmentScopeServices

		/*
			If compartmentId provided is not provided or is a root compartment then discover tenancy scope resources too
		*/
		if args.CompartmentId != nil && (*args.CompartmentId == "" || *args.CompartmentId == ctx.tenancyOcid) {
			args.Services = append(args.Services, tenancyScopeServices...)
		}
	}

	// Dedupes possible repeating services from command line and sorts them
	finalServices := []string{}
	serviceSet := convertStringSliceToSet(args.Services, true)
	excludeServicesSet := convertStringSliceToSet(args.ExcludeServices, true)
	for service := range serviceSet {
		if _, exists := excludeServicesSet[service]; !exists {
			finalServices = append(finalServices, service)
		}
	}
	args.Services = finalServices
	sort.Strings(args.Services)
}

// Validate export command arguments and returns nil if there are no issues
func (args *ExportCommandArgs) validate() error {
	if args.OutputDir == nil || *args.OutputDir == "" {
		return fmt.Errorf("[ERROR] no output directory specified")
	}

	path, err := os.Stat(*args.OutputDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("[ERROR] output_path does not exist: %s", err)
	}

	if !path.IsDir() {
		return fmt.Errorf("[ERROR] output_path %s should be a directory", *args.OutputDir)
	}

	return nil
}

func getExportConfig(d *schema.ResourceData) (interface{}, error) {
	clients := &tf_client.OracleClients{
		SdkClientMap:  make(map[string]interface{}, len(tf_client.OracleClientRegistrationsVar.RegisteredClients)),
		Configuration: make(map[string]string),
	}

	userAgentString := fmt.Sprintf(globalvar.ExportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, globalvar.Version)
	httpClient := tf_provider.BuildHttpClient()

	sdkConfigProvider, err := tf_provider.GetSdkConfigProvider(d, clients)
	if err != nil {
		return nil, err
	}
	exportConfigProvider = sdkConfigProvider

	// Note: In case of Instance Principal auth, the TenancyOCID will return
	// the ocid for the tenancy for the compute instance and not the one for the customer
	clients.Configuration["tenancy_ocid"], err = sdkConfigProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClientLocal, err := tf_provider.BuildConfigureClientFn(sdkConfigProvider, httpClient)
	if err != nil {
		return nil, err
	}

	configureClientWithUserAgent := func(client *oci_common.BaseClient) error {
		if err := configureClientLocal(client); err != nil {
			return err
		}
		client.UserAgent = userAgentString
		return nil
	}
	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	tf_client.ConfigureClientVar = configureClientWithUserAgent
	err = tf_client.CreateSDKClients(clients, sdkConfigProvider, configureClientWithUserAgent)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func runExportCommand(ctx *resourceDiscoveryContext) error {
	utils.Logf("[INFO] Running export command\n")
	utils.Logf("[INFO] parallelism: %d", ctx.Parallelism)
	defer ctx.printSummary()
	exportStart := time.Now()
	defer elapsed("entire export command", nil, 0)()
	steps, err := getDiscoverResourceSteps(ctx)
	if err != nil {
		return err
	}
	discoveryStart := time.Now()
	var discoverWg sync.WaitGroup
	discoverWg.Add(len(steps))
	for i, step := range steps {

		sem <- struct{}{}

		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] discover: Running step %d", i)
			defer elapsed(fmt.Sprintf("time taken in discovering resources for step %d", i), step.getBaseStep(), Discovery)()
			defer func() {
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in discover goroutine")
					utils.Logf("[ERROR] panic in discover goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] discoverWg done: step %d", i)
				discoverWg.Done()
			}()

			err := step.discover()
			if err != nil {
				// All errors in discover are added to the ctx.errorList
				utils.Debugf("[ERROR] error occurred while discovering resources for step %d", i)
				utils.Logf("[ERROR] error occurred while discovering resources: %s", err.Error())
				return
			}
			// Cull any references from the ref map that contain omitted resources
			// This is to avoid omitted resources from being referenced in generated configs
			for _, omittedResource := range step.getOmittedResources() {
				for key, reference := range referenceMap {
					if strings.Contains(reference, omittedResource.getTerraformReference()) {
						// refactor referenceMap to data structure with lock and methods to modify
						refMapLock.Lock()
						delete(referenceMap, key)
						refMapLock.Unlock()
					}
				}
			}

			utils.Debugf("[DEBUG] discover: Completed step %d", i)
			utils.Debugf("[DEBUG] discovered %d resources for step %d", len(step.getDiscoveredResources()), i)
			<-sem
		}(i, step)

	}

	// Wait for all steps to complete discovery
	discoverWg.Wait()
	totalDiscoveryTime := time.Since(discoveryStart)
	utils.Debugf("discovering resources for all services took %v\n", totalDiscoveryTime)
	ctx.timeTakenToDiscover = totalDiscoveryTime
	utils.Debug("[DEBUG] ~~~~~~ discover steps completed ~~~~~~")

	if ctx.GenerateState {
		stateStart := time.Now()
		// Run import commands
		if ctx.Parallelism > 1 {
			utils.Debug("[DEBUG] Generating state in parallel")
			if err := generateStateParallel(ctx, steps); err != nil {
				return err
			}
		} else {
			utils.Debug("[DEBUG] Generating state sequentially")
			if err := generateState(ctx, steps); err != nil {
				return err
			}
		}

		// remove invalid references from referenceMap for the resources with import error
		if ctx.isImportError {
			// lock not required for referenceMap as only 1 thread is running at this point
			deleteInvalidReferences(referenceMap, ctx.discoveredResources)
		}
		timeForStateGeneration := time.Since(stateStart)
		utils.Debugf("[DEBUG] state generation took %v\n", timeForStateGeneration)
		ctx.timeTakenToGenerateState = timeForStateGeneration
	}

	// Reset discovered resources if already set by writeTmpConfigurationForImport
	ctx.discoveredResources = make([]*OCIResource, 0)

	errorChannel := make(chan error)
	var configWg sync.WaitGroup
	configWg.Add(len(steps))

	wgDone := make(chan bool)

	// Write configuration for imported resources
	configStart := time.Now()
	for i, step := range steps {

		sem <- struct{}{}
		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] writeConfiguration: Running step %d", i)
			defer func() {
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in writeConfiguration goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] configWg done: step %d", i)
				configWg.Done()
			}()
			if err := step.writeConfiguration(); err != nil {
				errorChannel <- fmt.Errorf("[ERROR] error writing final configuration for resources found: %s", err.Error())
			}

			utils.Debugf("[DEBUG] writeConfiguration: Completed step %d", i)
			<-sem
		}(i, step)
	}

	// goroutine to wait until configWg is done
	go func() {
		utils.Debugf("[DEBUG] waiting for all configWg goroutines to finish...")
		configWg.Wait()
		close(wgDone)
	}()

	// Wait until either configWg is done or an error is received through the errorChannel
	select {
	case <-wgDone:
		utils.Debugf("[DEBUG] ~~~~~~ writeConfiguration steps completed ~~~~~~")
		utils.Debugf("[DEBUG] writing config took %v\n", time.Since(configStart))
		break
	case err := <-errorChannel:
		close(errorChannel)
		utils.Logf("[ERROR] error writing final configuration for resources found: %s", err.Error())
		return err
	}

	region, err := exportConfigProvider.Region()
	if err != nil {
		return err
	}
	vars["region"] = fmt.Sprintf("\"%s\"", region)

	if err := generateProviderFile(ctx.OutputDir); err != nil {
		return err
	}

	if err := generateVarsFile(vars, ctx.OutputDir); err != nil {
		return err
	}

	if isMissingRequiredAttributes {
		ctx.summaryStatements = append(ctx.summaryStatements, "")
		ctx.summaryStatements = append(ctx.summaryStatements, globalvar.MissingRequiredAttributeWarning)
		ctx.summaryStatements = append(ctx.summaryStatements, "Missing required attributes:")
		for key, value := range ctx.missingAttributesPerResource {
			ctx.summaryStatements = append(ctx.summaryStatements, fmt.Sprintf("%s: %s", key, strings.Join(value, ",")))
		}
	}
	ctx.timeTakenForEntireExport = time.Since(exportStart)
	ctx.postValidate()
	return nil
}

/*
generateStateParallel is used if value of parallelism arg > 1
- writes temp config for the discovered resources e.g. `resource_type resource_name {}` in order to run import
- writes temp state file for each service in parallel by running import for each of the found resources
- finally it merges all the state files generated into one state file using json merge
*/
func generateStateParallel(ctx *resourceDiscoveryContext, steps []resourceDiscoveryStep) error {

	// isInitDone is to make sure that multiple threads do not call terraform init
	utils.Debugf("[DEBUG] Reset isInitDone")
	isInitDone = false
	// Cleanup the temporary state files created for each input service
	defer cleanupTempStateFiles(ctx)
	defer elapsed("generating state in parallel", nil, 0)()

	errorChannel := make(chan error)
	var stateWg sync.WaitGroup
	wgDone := make(chan bool)

	stateWg.Add(len(steps))

	for i, step := range steps {
		if len(step.getDiscoveredResources()) == 0 {
			stateWg.Done()
			continue
		}

		sem <- struct{}{}

		go func(i int, step resourceDiscoveryStep) {
			utils.Debugf("[DEBUG] writing temp config and state: Running step %d", i)
			defer elapsed(fmt.Sprintf("time taken by step %s to generate state", fmt.Sprint(i)), step.getBaseStep(), GeneratingState)()
			defer func() {
				if r := recover(); r != nil {
					utils.Logf("[ERROR] panic in writing temp config and state goroutine")
					debug.PrintStack()
				}
				utils.Debugf("[DEBUG] stateWg done: step %d", i)
				stateWg.Done()
			}()

			/* Generate temporary HCL configs from all discovered resources to run import
			   Final configuration will be generated after import so that we can exclude the resources for which import failed
			   and also remove the references to failed resources from the referenceMap */
			if err := step.writeTmpConfigurationForImport(); err != nil {
				errorChannel <- fmt.Errorf("[ERROR] error writing temp config for resources found: %s", err.Error())
			}

			// Write temp state file for each service, this step will import resources into a separate state file for each service in parallel
			if err := step.writeTmpState(); err != nil {
				errorChannel <- fmt.Errorf("[ERROR] error writing temp state for resources found: %s", err.Error())
			}

			utils.Debugf("writing temp config and state: Completed step %d", i)
			<-sem
		}(i, step)
	}

	// goroutine to wait until stateWg is done
	go func() {
		utils.Debugf("[DEBUG] waiting for all stateWg threads to finish...")
		stateWg.Wait()
		close(wgDone)
	}()

	// Wait until either stateWg is done or an error is received through the errorChannel
	select {
	case <-wgDone:
		utils.Debugf("[DEBUG] ~~~~~~ writing temp config and state steps completed ~~~~~~")
		break
	case err := <-errorChannel:
		close(errorChannel)
		return err
	}

	// Generate final state by merging state json generated for all services
	for _, step := range steps {
		if err := step.mergeGeneratedStateFile(); err != nil {
			return fmt.Errorf("[ERROR] Resource discovery failed. Error merging generated states: %s", err.Error())
		}
	}

	if ctx.state == nil {
		utils.Logf("[INFO] ~~~~~~ no resources were imported to the state file ~~~~~~")
		return nil
	}

	// Create final state file to write state
	stateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultStateFilename)

	f, err := os.OpenFile(stateOutputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to Create state file %s: %s", stateOutputFile, err.Error())
	}

	// write generate state to file
	stateBytes, _ := json.MarshalIndent(ctx.state, "", "\t")
	if err := ioutil.WriteFile(stateOutputFile, stateBytes, 0644); err != nil {
		return fmt.Errorf("[ERROR] error writing state file at %s: %s", stateOutputFile, err.Error())
	} else {
		utils.Logf("[INFO] state written to file at: %s", stateOutputFile)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("[WARN] unable to close state file: %s", err.Error())
	}

	// Run init in output_path to validate generated state
	backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption

	if ctx.terraformProviderBinaryPath != "" {
		utils.Logf("[INFO] plugin dir set to: '%s'", ctx.terraformProviderBinaryPath)
		initArgs = append(initArgs, tfexec.PluginDir(ctx.terraformProviderBinaryPath))
	}

	if err := ctx.terraform.Init(backgroundCtx, initArgs...); err != nil {
		return err
	}
	return nil
}

/*
generateState is used if value of parallelism arg == 1
- writes temp config for the discovered resources e.g. `resource_type resource_name {}` in order to run import
- writes state file by running import for each of the found resources
*/
func generateState(ctx *resourceDiscoveryContext, steps []resourceDiscoveryStep) error {

	utils.Debugln("[DEBUG] writing temp config files for import")
	for _, step := range steps {

		/* Generate temporary HCL configs from all discovered resources to run import
		   Final configuration will be generated after import so that we can exclude the resources for which import failed
		   and also remove the references to failed resources from the referenceMap */
		if err := step.writeTmpConfigurationForImport(); err != nil {
			return err
		}
	}

	// Run init command
	backgroundCtx := context.Background()

	var initArgs []tfexec.InitOption

	if ctx.terraformProviderBinaryPath != "" {
		utils.Logf("[INFO] plugin dir set to: '%s'", ctx.terraformProviderBinaryPath)
		initArgs = append(initArgs, tfexec.PluginDir(ctx.terraformProviderBinaryPath))
	}
	if err := ctx.terraform.Init(backgroundCtx, initArgs...); err != nil {
		return err
	}

	stateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultStateFilename)
	tmpStateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), globalvar.DefaultTmpStateFile)
	if err := os.RemoveAll(tmpStateOutputFile); err != nil {
		utils.Logf("[WARN] unable to delete existing tmp state file %s", tmpStateOutputFile)
		return err
	}

	// Run import for all resources
	for _, resource := range ctx.discoveredResources {
		importResource(ctx, resource, tmpStateOutputFile)
	}

	if _, err := os.Stat(tmpStateOutputFile); !os.IsNotExist(err) {
		if err := os.Rename(tmpStateOutputFile, stateOutputFile); err != nil {
			return err
		}
	}

	return nil
}

/*
importResource runs terraform import for a given resource using Terraform exec and writes to the given state file
*/
func importResource(ctx *resourceDiscoveryContext, resource *OCIResource, tmpStateOutputFile string) {
	utils.Logf("[INFO] ===> Importing resource '%s'", resource.getTerraformReference())
	utils.Debugf("[DEBUG] ===> Importing resource '%s'", resource.getTerraformReference())

	resourceDefinition, exists := resourcesMap[resource.terraformClass]
	if !exists {
		utils.Logf("[INFO] skip importing '%s' since it is not a Terraform OCI resource", resource.getTerraformReference())
		utils.Debugf("[DEBUG] skip importing '%s' since it is not a Terraform OCI resource", resource.getTerraformReference())
		return
	}

	if resourceDefinition.Importer == nil {
		utils.Logf("[WARN] unable to import '%s' because import is not supported for '%s'", resource.getTerraformReference(), resource.terraformClass)
		return
	}

	importId := resource.importId
	if len(importId) == 0 {
		importId = resource.id
	}

	importArgs := []tfexec.ImportOption{
		tfexec.Config(*ctx.OutputDir),
		tfexec.State(tmpStateOutputFile),
	}
	if importErr := ctx.terraform.Import(context.Background(), resource.getTerraformReference(), importId, importArgs...); importErr != nil {
		utils.Logf("[ERROR] terraform import command failed for resource '%s' at id '%s': %s", resource.getTerraformReference(), importId, importErr.Error())

		// mark resource as errored so that it can be skipped while writing configurations
		resource.isErrorResource = true

		ctx.ctxLock.Lock()
		ctx.isImportError = true
		ctx.ctxLock.Unlock()

		err := fmt.Errorf("[ERROR] terraform import command failed for resource '%s' at id '%s': %s Any references to this resource have been replaced with hard coded values in generated configurations", resource.getTerraformReference(), importId, importErr.Error())

		var rdError *ResourceDiscoveryError
		if ctx.targetSpecificResources {
			rdError = &ResourceDiscoveryError{
				resource.terraformClass,
				"",
				err,
				nil}
		} else {
			rdError = &ResourceDiscoveryError{
				resource.terraformClass,
				resource.parent.terraformName,
				err,
				nil}
		}
		ctx.addErrorToList(rdError)

	}
}

func getDiscoverResourceSteps(ctx *resourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	if !ctx.targetSpecificResources {
		return getDiscoverResourceWithGraphSteps(ctx)
	}

	result := make([]resourceDiscoveryStep, 1)
	result[0] = &resourceDiscoveryWithTargetIds{
		resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
			ctx:                 ctx,
			name:                "resources",
			discoveredResources: []*OCIResource{},
			omittedResources:    []*OCIResource{},
		},
	}
	return result, nil
}

func getDiscoverResourceWithGraphSteps(ctx *resourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	defer elapsed("Building resource discovery graph", nil, 0)()
	if ctx.CompartmentId == nil || *ctx.CompartmentId == "" {
		*ctx.CompartmentId = ctx.tenancyOcid
	}
	var result []resourceDiscoveryStep

	// Discover tenancy scope resources only if compartmentId is tenancy ocid
	if *ctx.CompartmentId == ctx.tenancyOcid {
		tenancyResource := &OCIResource{
			compartmentId: ctx.tenancyOcid,
			TerraformResource: TerraformResource{
				id:             ctx.tenancyOcid,
				terraformClass: "oci_identity_tenancy",
				terraformName:  "export",
			},
		}

		for _, mode := range ctx.Services {
			if resourceGraph, exists := tenancyResourceGraphs[mode]; exists {
				result = append(result, &resourceDiscoveryWithGraph{
					root:                      tenancyResource,
					resourceGraph:             resourceGraph,
					resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
				})

				vars["tenancy_ocid"] = fmt.Sprintf("\"%s\"", ctx.tenancyOcid)
				referenceMap[ctx.tenancyOcid] = tfHclVersion.getVarHclString("tenancy_ocid")
			}
		}
	}

	compartmentResource := &OCIResource{
		compartmentId: *ctx.CompartmentId,
		TerraformResource: TerraformResource{
			id:             *ctx.CompartmentId,
			terraformClass: "oci_identity_compartment",
			terraformName:  "export",
		},
	}

	for _, mode := range ctx.Services {
		if resourceGraph, exists := compartmentResourceGraphs[mode]; exists {
			result = append(result, &resourceDiscoveryWithGraph{
				root:                      compartmentResource,
				resourceGraph:             resourceGraph,
				resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
			})

			vars["compartment_ocid"] = fmt.Sprintf("\"%s\"", *ctx.CompartmentId)
			referenceMap[*ctx.CompartmentId] = tfHclVersion.getVarHclString("compartment_ocid")
		}
	}

	return result, nil
}

func findResources(ctx *resourceDiscoveryContext, root *OCIResource, resourceGraph TerraformResourceGraph) (foundResources []*OCIResource, err error) {
	// findResources will never return error, it will add the errors encountered to the errorList and print those after the discovery finishes
	// If find resources needs to fail in some scenario, this func needs to be modified to return error instead of continuing discovery
	// Errors so far are API errors or the errors when service/feature is not available
	foundResources = []*OCIResource{}

	childResourceTypes, exists := resourceGraph[root.terraformClass]
	if !exists {
		return foundResources, nil
	}

	utils.Logf("[INFO] resource discovery: visiting %s\n", root.getTerraformReference())
	utils.Debugf("[DEBUG] resource discovery: visiting %s\n", root.getTerraformReference())

	for _, childType := range childResourceTypes {
		func() {
			defer handlePanicFindResources(&childType, &err)

			findResourceFn := findResourcesGeneric
			if childType.findResourcesOverrideFn != nil {
				findResourceFn = childType.findResourcesOverrideFn
			}
			results, err := findResourceFn(ctx, &childType, root, &resourceGraph)
			if err != nil {
				// add error to the errorList and continue discovering rest of the resources
				rdError := &ResourceDiscoveryError{
					childType.resourceClass,
					root.terraformName,
					err,
					&resourceGraph}
				ctx.addErrorToList(rdError)
				return
			}

			if childType.processDiscoveredResourcesFn != nil {
				results, err = childType.processDiscoveredResourcesFn(ctx, results)
				if err != nil {
					// add error to the errorList and continue discovering rest of the resources
					rdError := &ResourceDiscoveryError{
						childType.resourceClass,
						root.terraformName,
						err,
						&resourceGraph}
					ctx.addErrorToList(rdError)
					return
				}
			}
			foundResources = append(foundResources, results...)

			for _, resource := range results {
				//referenceMap[resource.id] = resource.getHclReferenceIdString()
				if ctx.expectedResourceIds != nil && len(ctx.expectedResourceIds) > 0 {
					if _, shouldExport := ctx.expectedResourceIds[resource.id]; shouldExport {
						resource.omitFromExport = false
						ctx.ctxLock.Lock()
						ctx.expectedResourceIds[resource.id] = true
						ctx.ctxLock.Unlock()
					} else {
						resource.omitFromExport = !childType.alwaysExportable
					}
				}

				subResources, err := findResources(ctx, resource, resourceGraph)
				if err != nil {
					continue
				}
				foundResources = append(foundResources, subResources...)
			}
		}()
	}

	return foundResources, nil
}

func generateVarsFile(vars map[string]string, outputDir *string) error {
	varsTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), globalvar.VarsFile)
	varsOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), globalvar.VarsFile)
	file, err := os.OpenFile(varsTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	for variable, defaultVal := range vars {
		if defaultVal != "" {
			_, _ = file.WriteString(fmt.Sprintf("variable %s { default = %s }\n", variable, defaultVal))
		} else {
			_, _ = file.WriteString(fmt.Sprintf("variable %s {}\n", variable))
		}
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := os.Rename(varsTmpFile, varsOutputFile); err != nil {
		return err
	}

	return nil
}

func generateProviderFile(outputDir *string) error {
	providerTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), globalvar.ProviderFile)
	providerOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), globalvar.ProviderFile)
	file, err := os.OpenFile(providerTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("provider oci {\n\tregion = %s\n}\n", tfHclVersion.getVarHclString("region")))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(providerTmpFile, providerOutputFile); err != nil {
		return err
	}

	return nil
}

type OCIResource struct {
	TerraformResource
	compartmentId    string
	rawResource      interface{}
	sourceAttributes map[string]interface{}
	getHclStringFn   func(*strings.Builder, *OCIResource, map[string]string) error
	parent           *OCIResource
	isErrorResource  bool
}

type TerraformResource struct {
	id                         string
	importId                   string
	terraformClass             string
	terraformName              string
	terraformReferenceIdString string // syntax independent interpolation- `resource_type.resource_name.id`
	terraformTypeInfo          *TerraformResourceHints
	omitFromExport             bool
}

func (tr *TerraformResource) getHclReferenceIdString() string {
	if tr.terraformReferenceIdString != "" {
		return tfHclVersion.getSingleExpHclString(tr.terraformReferenceIdString)
	}
	return tfHclVersion.getDoubleExpHclString(tr.getTerraformReference(), "id")
}

func (tr *TerraformResource) getTerraformReference() string {
	return fmt.Sprintf("%s.%s", tr.terraformClass, tr.terraformName)
}

func getHCLStringFromMap(builder *strings.Builder, sourceAttributes map[string]interface{}, resourceSchema *schema.Resource, interpolationMap map[string]string, ociRes *OCIResource, attributePrefix string) error {
	sortedKeys := make([]string, len(resourceSchema.Schema))
	cnt := 0
	for k := range resourceSchema.Schema {
		sortedKeys[cnt] = k
		cnt++
	}
	sort.Strings(sortedKeys)

	for _, tfAttribute := range sortedKeys {
		tfSchema := resourceSchema.Schema[tfAttribute]
		if tfSchema.Deprecated != "" || tfSchema.Removed != "" || (!tfSchema.Required && !tfSchema.Optional) {
			continue
		}

		if attributeVal, exists := sourceAttributes[tfAttribute]; exists {
			switch v := attributeVal.(type) {
			case InterpolationString:
				if ok := failedResourceReferenceSet[v.resourceReference]; ok {
					builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, v.value))
				} else {
					builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, v.interpolation))
				}
				continue
			case string:
				if varOverride, exists := interpolationMap[fmt.Sprintf("%v", v)]; exists {
					v = varOverride
					builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, v))
				} else {
					builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, escapeTFStrings(v)))
				}
				continue
			case int, bool, float64:
				builder.WriteString(fmt.Sprintf("%s = \"%v\"\n", tfAttribute, v))
				continue
			case []interface{}:
				switch tfSchema.Type {
				case schema.TypeString:
					if tfAttribute == "delivery_policy" {
						builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, parseDeliveryPolicy(v[0].(interface{}))))
						continue
					}
				case schema.TypeList, schema.TypeSet:
					switch elem := tfSchema.Elem.(type) {
					case *schema.Resource:
						for i, item := range v {
							if val := item.(map[string]interface{}); val != nil {
								builder.WriteString(fmt.Sprintf("%s {\n", tfAttribute))
								attributePrefixForRecursiveCall := attributePrefix
								if attributePrefix == "" {
									attributePrefixForRecursiveCall = fmt.Sprintf("%s[%d]", tfAttribute, i)
								} else {
									attributePrefixForRecursiveCall = fmt.Sprintf("%s.%s[%d]", attributePrefix, tfAttribute, i)
								}
								if err := getHCLStringFromMap(builder, val, elem, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
									return err
								}
								builder.WriteString("}\n")
							}
						}
						continue
					case *schema.Schema, schema.ValueType, InterpolationString:
						builder.WriteString(fmt.Sprintf("%s = [\n", tfAttribute))
						for _, item := range v {
							switch trueListVal := item.(type) {
							case InterpolationString:
								if ok := failedResourceReferenceSet[trueListVal.resourceReference]; ok {
									builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, trueListVal.value))
								} else {
									builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, trueListVal.interpolation))
								}
							case string:
								if varOverride, exists := interpolationMap[fmt.Sprintf("%v", trueListVal)]; exists {
									trueListVal = varOverride
									builder.WriteString(fmt.Sprintf("%v,\n", trueListVal))
								} else {
									builder.WriteString(fmt.Sprintf("%q,\n", escapeTFStrings(trueListVal)))
								}
							case int, bool, float64:
								builder.WriteString(fmt.Sprintf("\"%v\",\n", trueListVal))
							default:
								return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': List element type mismatch", tfAttribute, tfAttribute)
							}
						}
						builder.WriteString("]\n")
						continue
					}

					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': List element is neither schema.Resource or schema.Schema", tfAttribute, tfAttribute)
				}
			case map[string]interface{}:
				switch tfSchema.Type {
				case schema.TypeList:
					if nestedResource := tfSchema.Elem.(*schema.Resource); nestedResource != nil {
						builder.WriteString(fmt.Sprintf("%s {\n", tfAttribute))
						attributePrefixForRecursiveCall := attributePrefix
						if attributePrefix == "" {
							attributePrefixForRecursiveCall = tfAttribute
						} else {
							attributePrefixForRecursiveCall = attributePrefix + "." + tfAttribute
						}
						if err := getHCLStringFromMap(builder, v, nestedResource, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
							return err
						}
						builder.WriteString("}\n")
						continue
					}
					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': Nested resource type mismatch", tfAttribute, tfAttribute)
				case schema.TypeMap:
					builder.WriteString(fmt.Sprintf("%s = {\n", tfAttribute))

					keys := utils.GetSortedKeys(v)
					for _, mapKey := range keys {
						switch mapVal := v[mapKey].(type) {
						case InterpolationString:
							if ok := failedResourceReferenceSet[mapVal.resourceReference]; ok {
								builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, mapVal.value))
							} else {
								builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, mapVal.interpolation))
							}
						case string:
							if varOverride, exists := interpolationMap[fmt.Sprintf("%v", mapVal)]; exists {
								mapVal = varOverride
								builder.WriteString(fmt.Sprintf("\"%s\" = %v\n", mapKey, mapVal))
							} else {
								builder.WriteString(fmt.Sprintf("\"%s\" = %q\n", mapKey, escapeTFStrings(mapVal)))
							}
						case int, bool, float64:
							builder.WriteString(fmt.Sprintf("\"%s\" = \"%v\"\n", mapKey, mapVal))
						default:
							builder.WriteString(fmt.Sprintf("#%s = <<Placeholder due to complex map value>>\n", mapKey))
						}
					}
					builder.WriteString("}\n")
					continue
				default:
					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': Source attribute is nested object but TF attribute is not", tfAttribute, tfAttribute)
				}
			case nil:
				utils.Logf("[INFO] TF attribute '%s' is nil in source\n", tfAttribute)
				if !tfSchema.Required {
					continue
				}
			default:
				utils.Logf("[WARN] TF attribute '%s' is unknown type in source\n", tfAttribute)
			}
		}

		if tfSchema.Required {
			utils.Logf("[WARN] Required TF attribute '%s' not found in source\n", tfAttribute)
			/* Set missing value if specified in resource hints. This is to avoid plan failure for existing infrastructure.
			This is only done for required attributes as the Optional attributes will not cause plan failure
			We can extend this in future to provide this option to customer to add default values for attributes
			and add this logic to Optional attributes too */

			if ociRes.terraformTypeInfo == nil {
				ociRes.terraformTypeInfo = &TerraformResourceHints{}
			}

			if ociRes.terraformTypeInfo.defaultValuesForMissingAttributes == nil {
				ociRes.terraformTypeInfo.defaultValuesForMissingAttributes = make(map[string]interface{})
			}
			if tfAttributeVal, exists := ociRes.terraformTypeInfo.defaultValuesForMissingAttributes[tfAttribute]; exists {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, tfAttributeVal))
			} else {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, globalvar.PlaceholderValueForMissingAttribute))
			}
			builder.WriteString("\t#Required attribute not found in discovery, placeholder value set to avoid plan failure\n")
			isMissingRequiredAttributes = true

			/* Add missing required attribute to ignorableRequiredMissingAttributes to be generated in lifecycle ignore_changes */
			if ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes == nil {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes = make(map[string]bool)
			}
			if attributePrefix == "" {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes[tfAttribute] = true
			} else {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes[attributePrefix+"."+tfAttribute] = true
			}

		} else if tfSchema.Optional {
			utils.Logf("[INFO] Optional TF attribute '%s' not found in source\n", tfAttribute)
			builder.WriteString(fmt.Sprintf("#%s = <<Optional value not found in discovery>>\n", tfAttribute))
		}
	}
	return nil
}

func (resource *OCIResource) hasFreeformTag(tagKey string) bool {
	if freeformTags, exists := resource.sourceAttributes["freeform_tags"]; exists {
		if freeformTagMap, ok := freeformTags.(map[string]interface{}); ok {
			if _, hasFreeFormTag := freeformTagMap[tagKey]; hasFreeFormTag {
				return true
			}
		}
	}

	return false
}

func (resource *OCIResource) hasDefinedTag(tagKey string, tagValue string) bool {
	if definedTags, exists := resource.sourceAttributes["defined_tags"]; exists {
		if definedTagMap, ok := definedTags.(map[string]interface{}); ok {
			if definedTagValue, hasDefinedTag := definedTagMap[tagKey]; hasDefinedTag {
				return definedTagValue == tagValue
			}
		}
	}

	return false
}

func (ociRes *OCIResource) getHCLString(builder *strings.Builder, interpolationMap map[string]string) error {
	// Remove any potential cyclical references from the interpolation map
	selfReference := ociRes.getTerraformReference()
	resourceInterpolationMap := map[string]string{}
	for value, interpolation := range interpolationMap {
		if !strings.Contains(interpolation, selfReference) {
			resourceInterpolationMap[value] = interpolation
		}
	}

	if ociRes.getHclStringFn != nil {
		return ociRes.getHclStringFn(builder, ociRes, resourceInterpolationMap)
	}
	return getHclStringFromGenericMap(builder, ociRes, resourceInterpolationMap)
}

func getHclStringFromGenericMap(builder *strings.Builder, ociRes *OCIResource, interpolationMap map[string]string) error {
	resourceSchema := resourcesMap[ociRes.terraformClass]

	builder.WriteString(fmt.Sprintf("resource %s %s {\n", ociRes.terraformClass, ociRes.terraformName))
	if err := getHCLStringFromMap(builder, ociRes.sourceAttributes, resourceSchema, interpolationMap, ociRes, ""); err != nil {
		return err
	}

	if ociRes.terraformTypeInfo != nil && len(ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes) > 0 {
		builder.WriteString("\n# Required attributes that were not found in discovery have been added to " +
			"lifecycle ignore_changes")
		builder.WriteString("\n# This is done to avoid terraform plan failure for the existing infrastructure")
		builder.WriteString("\nlifecycle {\n" +
			"ignore_changes = [")

		missingAttributes := make([]string, 0, len(ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes))

		for attribute := range ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes {
			missingAttributes = append(missingAttributes, tfHclVersion.getReference(attribute))
		}
		builder.WriteString(strings.Join(missingAttributes, ","))

		builder.WriteString("]\n" +
			"}\n")
	}
	builder.WriteString("}\n\n")

	return nil
}

// This function attempts to convert resource data items to a map representation that omits attributes where no value was set.
func convertDatasourceItemToMap(d *schema.ResourceData, itemPrefix string, itemSchema map[string]*schema.Schema) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	for attributeKey, attributeSchema := range itemSchema {
		var key string
		if itemPrefix != "" {
			key = fmt.Sprintf("%s.%s", itemPrefix, attributeKey)
		} else {
			key = attributeKey
		}

		switch attributeSchema.Type {
		case schema.TypeBool, schema.TypeInt, schema.TypeFloat, schema.TypeString:
			if val, exists := d.GetOkExists(key); exists {
				result[attributeKey] = val
			}
		case schema.TypeList:
			switch v := attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					result[attributeKey] = val
				}
			case *schema.Resource:
				if val, exists := d.GetOkExists(key); exists {
					list := val.([]interface{})
					resourceList := make([]interface{}, len(list))
					for idx := range list {
						resourceList[idx], _ = convertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, idx), v.Schema)
					}
					result[attributeKey] = resourceList
				}
			}
		case schema.TypeMap:
			switch attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					result[attributeKey] = val
				}
			default:
				return result, fmt.Errorf("[ERROR] found a non-primitive element in schema for TypeMap attribute '%s'", attributeKey)
			}
		case schema.TypeSet:
			switch v := attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					setVal := val.(*schema.Set)
					result[attributeKey] = setVal.List()
				}
			case *schema.Resource:
				if val, exists := d.GetOkExists(key); exists {
					setVal := val.(*schema.Set)
					list := setVal.List()
					resourceList := make([]interface{}, len(list))
					for idx, item := range list {
						itemHashCode := setVal.F(item)
						resourceList[idx], _ = convertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, itemHashCode), v.Schema)
					}
					result[attributeKey] = resourceList
				}
			}
		}
	}

	return result, nil
}

func findResourcesGeneric(ctx *resourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	results := []*OCIResource{}
	clients := ctx.clients

	utils.Logf("[INFO] discovering resources with data source '%s'\n", tfMeta.datasourceClass)
	utils.Debugf("[DEBUG] discovering resources with data source '%s'\n", tfMeta.datasourceClass)
	datasource := datasourcesMap[tfMeta.datasourceClass]
	d := datasource.TestResourceData()
	d.Set("compartment_id", parent.compartmentId)

	for queryAttributeName, queryValue := range tfMeta.datasourceQueryParams {
		utils.Logf("[INFO] adding datasource query attribute '%s' from parent attribute '%s'\n", queryAttributeName, queryValue)
		utils.Debugf("[DEBUG] adding datasource query attribute '%s' from parent attribute '%s'\n", queryAttributeName, queryValue)

		if queryValue == "" || queryValue == "id" {
			d.Set(queryAttributeName, parent.id)
		} else if strings.HasPrefix(queryValue, "'") && strings.HasSuffix(queryValue, "'") { // Anything encapsulated in ' ' means to use the literal value
			d.Set(queryAttributeName, queryValue[1:len(queryValue)-1])
		} else if val, ok := parent.sourceAttributes[queryValue]; ok {
			d.Set(queryAttributeName, val)
		} else {
			utils.Logf("[WARN] no attribute '%s' found in parent '%s', returning no results for this resource\n", queryValue, parent.getTerraformReference())
			return results, nil
		}
	}

	if err := datasource.Read(d, clients); err != nil {
		return results, err
	}

	if !tfMeta.DiscoversWithSingularDatasource() {
		// Results are from a plural datasource
		itemSchema := datasource.Schema[tfMeta.datasourceItemsAttr]
		elemResource, ok := itemSchema.Elem.(*schema.Resource)
		if !ok {
			return results, fmt.Errorf("[ERROR] element schema is not of a resource")
		}
		datasourceItemsAttribute := tfMeta.datasourceItemsAttr

		if tfMeta.isDatasourceCollection {
			collectionItemSchema := elemResource.Schema["items"]

			elemResource, ok = collectionItemSchema.Elem.(*schema.Resource)
			if !ok {
				return results, fmt.Errorf("[ERROR] collection element schema is not of a resource")
			}
			datasourceItemsAttribute = tfMeta.datasourceItemsAttr + ".0.items"
		}

		foundItems, _ := d.GetOkExists(datasourceItemsAttribute)
		for idx, item := range foundItems.([]interface{}) {
			if itemMap, ok := item.(map[string]interface{}); ok {
				if state, exists := itemMap["state"].(string); exists && len(tfMeta.discoverableLifecycleStates) > 0 {
					discoverable := false
					for _, val := range tfMeta.discoverableLifecycleStates {
						if strings.EqualFold(state, val) {
							discoverable = true
							break
						}
					}

					if !discoverable {
						continue
					}
				}
			}
			var resource *OCIResource
			var err error
			if tfMeta.requireResourceRefresh {
				resourceSchema := resourcesMap[tfMeta.resourceClass]
				r := resourceSchema.TestResourceData()

				// Use resource to fill in all attributes (likely because the datasource doesn't return complete info)
				if tfMeta.getIdFn != nil {
					tmpResource, err := generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", datasourceItemsAttribute, idx), tfMeta, parent)
					if err != nil {
						rdError := &ResourceDiscoveryError{
							tfMeta.resourceClass,
							parent.terraformName,
							fmt.Errorf("[ERROR] error generating temporary resource from resource data returned in list datasource read: %v ", err),
							resourceGraph}
						ctx.addErrorToList(rdError)
						continue
					}

					itemId, err := tfMeta.getIdFn(tmpResource)
					if err != nil {
						rdError := &ResourceDiscoveryError{
							tfMeta.resourceClass,
							parent.terraformName,
							fmt.Errorf("[ERROR] failed to get a composite ID for the resource: %v ", err),
							resourceGraph}
						ctx.addErrorToList(rdError)
						continue
					}
					r.SetId(itemId)
				} else if idSchema, exists := elemResource.Schema["id"]; exists && idSchema.Type == schema.TypeString {
					itemId := item.(map[string]interface{})["id"]
					r.SetId(itemId.(string))
				} else {
					rdError := &ResourceDiscoveryError{
						tfMeta.resourceClass,
						parent.terraformName,
						fmt.Errorf("[ERROR] elements in datasource '%s' are missing an 'id' field and is unable to generate an id",
							tfMeta.datasourceClass),
						resourceGraph}
					ctx.addErrorToList(rdError)
					continue
				}

				if err = resourceSchema.Read(r, clients); err != nil {
					rdError := &ResourceDiscoveryError{
						tfMeta.resourceClass,
						parent.terraformName,
						fmt.Errorf("[ERROR] error refreshing resource using resource read: %v ", err),
						resourceGraph}
					ctx.addErrorToList(rdError)
					continue
				}
				// If state was voided because of error in Read (r.Id() is empty)
				if r.Id() == "" {
					rdError := &ResourceDiscoveryError{
						tfMeta.resourceClass,
						parent.terraformName,
						fmt.Errorf("[ERROR] error refreshing resource using resource read, state voided"),
						resourceGraph}
					ctx.addErrorToList(rdError)
					continue
				}
				resource, err = generateOciResourceFromResourceData(r, r, resourceSchema.Schema, "", tfMeta, parent)
				if err != nil {
					rdError := &ResourceDiscoveryError{
						tfMeta.resourceClass,
						parent.terraformName,
						fmt.Errorf("[ERROR] error generating resource from resource data returned in resource read: %v ", err),
						resourceGraph}
					ctx.addErrorToList(rdError)
					continue
				}
			} else {
				resource, err = generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", datasourceItemsAttribute, idx), tfMeta, parent)
				if err != nil {
					rdError := &ResourceDiscoveryError{
						tfMeta.resourceClass,
						parent.terraformName,
						fmt.Errorf("[ERROR] error generating resource from resource data returned in list datasource read: %v ", err),
						resourceGraph}
					ctx.addErrorToList(rdError)
					continue
				}
			}

			if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, elemResource.Schema); err != nil {
				resource.terraformName = fmt.Sprintf("%s_%s_%d", parent.terraformName, tfMeta.resourceAbbreviation, idx+1)
			}

			results = append(results, resource)
		}
	} else if d.Id() != "" {
		// Result is from a singular datasource that hasn't had its state voided (hence d.Id() is non-empty)
		resource, err := generateOciResourceFromResourceData(d, d, datasource.Schema, "", tfMeta, parent)
		if err != nil {
			return results, err
		}

		if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, datasource.Schema); err != nil {
			resource.terraformName = fmt.Sprintf("%s_%s", parent.terraformName, tfMeta.resourceAbbreviation)
		}

		discoverable := true
		if state, ok := resource.sourceAttributes["state"]; ok && len(tfMeta.discoverableLifecycleStates) > 0 {
			discoverable = false
			for _, val := range tfMeta.discoverableLifecycleStates {
				if strings.EqualFold(state.(string), val) {
					discoverable = true
					break
				}
			}
		}

		if discoverable {
			results = append(results, resource)
		}
	} else {
		utils.Debugf("[DEBUG] singular data source not able to find resource")
	}

	return results, nil
}

func getNormalizedTerraformName(source string) string {
	// Only alphanumeric, underscore, and hyphens are allowed. Strip out anything else.
	reg, err := regexp.Compile(`[^a-zA-Z0-9\-\_]+`)
	if err != nil {
		log.Fatal(err)
	}

	result := reg.ReplaceAllString(source, "-")
	result = fmt.Sprintf("export_%s", result)
	return result
}

func convertResourceDataToMap(schemaMap map[string]*schema.Schema, d *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	for key := range schemaMap {
		if val, ok := d.GetOkExists(key); ok {
			result[key] = val
		}
	}

	return result
}

// This function should only be used to escape TF-characters in strings
func escapeTFStrings(val string) string {
	val = strings.ReplaceAll(val, "%{", "%%{")
	val = strings.ReplaceAll(val, "${", "$${")
	return val
}

func generateTerraformNameFromResource(resourceAttributes map[string]interface{}, resourceSchema map[string]*schema.Schema) (string, error) {
	possibleNameAttributes := []string{
		"display_name",
		"name",
	}

	for _, nameAttribute := range possibleNameAttributes {
		if nameSchema, hasNameAttr := resourceSchema[nameAttribute]; hasNameAttr && nameSchema.Type == schema.TypeString {
			if value, exists := resourceAttributes[nameAttribute]; exists {
				terraformName := getNormalizedTerraformName(value.(string))
				resourceNameCountLock.Lock()
				if count, resourceNameExists := resourceNameCount[terraformName]; resourceNameExists {
					// Update count for existing name
					resourceNameCount[terraformName] = count + 1
					terraformName = fmt.Sprintf("%s_%d", terraformName, count)
				}
				// add the newly generated name to map
				resourceNameCount[terraformName] = 1
				resourceNameCountLock.Unlock()
				return terraformName, nil
			}
		}
	}

	return "", fmt.Errorf("unable to find a suitable name from the resource attributes")
}

func generateOciResourceFromResourceData(d *schema.ResourceData, rawResource interface{}, resourceSchema map[string]*schema.Schema, itemPrefix string, tfMeta *TerraformResourceAssociation, parent *OCIResource) (*OCIResource, error) {
	// The following conversion takes a ResourceData and converts it to a map where null values are preserved (and omitted).
	// Note that we don't use the raw map that Terraform gives us, because it will set zero-values even though the datasource didn't set one.
	//
	// TODO: An improvement on this logic would be to load the resource schema's Read function and invoke that to get the full
	// resource representation for attributes that a datasource might not give. The reasons for not doing this yet are:
	// 1) Adding an extra Read invocation could result in unnecessary traffic overhead against services for every resource we've discovered
	// 2) The result of resource Reads may return TypeSets with nested resources, which are hard to check for non-existent sub-attributes
	//
	// For now, assume that the datasource Read is good enough and add custom logic (via the process functions) if you need to add more info
	resourceMap, err := convertDatasourceItemToMap(d, itemPrefix, resourceSchema)
	if err != nil {
		return nil, err
	}

	resource := &OCIResource{
		compartmentId:    parent.compartmentId,
		sourceAttributes: resourceMap,
		rawResource:      rawResource,
		TerraformResource: TerraformResource{
			terraformClass:    tfMeta.resourceClass,
			terraformTypeInfo: tfMeta.TerraformResourceHints,
		},
		getHclStringFn: getHclStringFromGenericMap,
		parent:         parent,
	}

	if tfMeta.getIdFn != nil {
		if customId, err := tfMeta.getIdFn(resource); err == nil {
			resource.id = customId
		}
	} else if resourceId, resourceIdExists := resourceMap["id"]; resourceIdExists {
		resource.id = resourceId.(string)
	}

	if resource.id == "" {
		resource.id = d.Id()
	}

	if tfMeta.getHCLStringOverrideFn != nil {
		resource.getHclStringFn = tfMeta.getHCLStringOverrideFn
	}

	return resource, nil
}

func getOciResource(d *schema.ResourceData, resourceSchema map[string]*schema.Schema, compartmentId string, resourceHint *TerraformResourceHints, resourceId string) (*OCIResource, error) {
	resourceMap, err := convertDatasourceItemToMap(d, "", resourceSchema)
	if err != nil {
		return nil, err
	}

	resource := &OCIResource{
		compartmentId:    compartmentId,
		sourceAttributes: resourceMap,
		rawResource:      d,
		TerraformResource: TerraformResource{
			terraformClass:    resourceHint.resourceClass,
			terraformTypeInfo: resourceHint,
		},
		getHclStringFn: getHclStringFromGenericMap,
	}

	if resourceId != "" {
		resource.id = resourceId
	}

	if resource.id == "" {
		resource.id = d.Id()
	}

	return resource, nil
}

func resolveCompartmentId(clients *tf_client.OracleClients, compartmentName *string) (*string, error) {
	req := oci_identity.ListCompartmentsRequest{}

	rootCompartment, err := exportConfigProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}
	req.CompartmentId = &rootCompartment

	recursiveSearch := true
	req.CompartmentIdInSubtree = &recursiveSearch

	for {
		resp, err := clients.IdentityClient().ListCompartments(context.Background(), req)
		if err != nil {
			return nil, err
		}

		for _, compartment := range resp.Items {
			if compartment.Name != nil && *compartment.Name == *compartmentName {
				utils.Logf("[INFO] resolved compartment name '%s' to compartment id '%s'", *compartmentName, *compartment.Id)
				return compartment.Id, nil
			}
		}

		if resp.OpcNextPage == nil {
			break
		}
		req.Page = resp.OpcNextPage
	}

	return nil, fmt.Errorf("[ERROR] Could not find a compartment named '%s' in your tenancy", *compartmentName)
}

func readEnvironmentVars(d *schema.ResourceData) error {

	if err := d.Set(globalvar.AuthAttrName, utils.GetProviderEnvSettingWithDefault(globalvar.AuthAttrName, globalvar.AuthAPIKeySetting)); err != nil {
		return err
	}
	if err := d.Set(globalvar.ConfigFileProfileAttrName, utils.GetProviderEnvSettingWithDefault(globalvar.ConfigFileProfileAttrName, "")); err != nil {
		return err
	}
	if region := utils.GetProviderEnvSettingWithDefault(globalvar.RegionAttrName, ""); region != "" {
		if err := d.Set(globalvar.RegionAttrName, region); err != nil {
			return err
		}
	}

	if tenancyOcid := utils.GetProviderEnvSettingWithDefault(globalvar.TenancyOcidAttrName, ""); tenancyOcid != "" {
		if err := d.Set(globalvar.TenancyOcidAttrName, tenancyOcid); err != nil {
			return err
		}
	}

	if userOcid := utils.GetProviderEnvSettingWithDefault(globalvar.UserOcidAttrName, ""); userOcid != "" {
		if err := d.Set(globalvar.UserOcidAttrName, userOcid); err != nil {
			return err
		}
	}
	if fingerprint := utils.GetProviderEnvSettingWithDefault(globalvar.FingerprintAttrName, ""); fingerprint != "" {
		if err := d.Set(globalvar.FingerprintAttrName, fingerprint); err != nil {
			return err
		}
	}
	if privateKey := utils.GetProviderEnvSettingWithDefault(globalvar.PrivateKeyAttrName, ""); privateKey != "" {
		if err := d.Set(globalvar.PrivateKeyAttrName, privateKey); err != nil {
			return err
		}
	}
	if privateKeyPath := utils.GetProviderEnvSettingWithDefault(globalvar.PrivateKeyPathAttrName, ""); privateKeyPath != "" {
		if err := d.Set(globalvar.PrivateKeyPathAttrName, privateKeyPath); err != nil {
			return err
		}
	}
	if privateKeyPassword := utils.GetProviderEnvSettingWithDefault(globalvar.PrivateKeyPasswordAttrName, ""); privateKeyPassword != "" {
		if err := d.Set(globalvar.PrivateKeyPasswordAttrName, privateKeyPassword); err != nil {
			return err
		}
	}
	return nil
}

func getTenancyOcidFromCompartment(clients *tf_client.OracleClients, compartmentId string) (string, error) {

	for true {
		response, err := clients.IdentityClient().GetCompartment(context.Background(), oci_identity.GetCompartmentRequest{
			CompartmentId: &compartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(true, "identity"),
			},
		})
		if err != nil {
			return "", fmt.Errorf("[ERROR] could not get tenancy ocid from compartment ocid %v", err)
		}
		if response.CompartmentId == nil {
			utils.Logf("[INFO] root compartment found %v", compartmentId)
			return *response.Id, nil
		}
		compartmentId = *response.CompartmentId
	}

	return "", fmt.Errorf("[ERROR] could not get tenancy ocid from compartment ocid")
}

func deleteInvalidReferences(referenceMap map[string]string, discoveredResources []*OCIResource) {
	// intialize referenceResourceNameSet
	// This set contains unique terraform names for resource references
	if referenceResourceNameSet == nil {
		referenceResourceNameSet = make(map[string]bool)
		for _, value := range referenceMap {
			valueParts := strings.Split(value, ".")
			if len(valueParts) < 3 {
				continue
			}
			referenceResourceNameSet[valueParts[1]] = true
		}
	}
	if failedResourceReferenceSet == nil {
		failedResourceReferenceSet = make(map[string]bool)
	}

	for _, resource := range discoveredResources {

		// delete the entry if key is an OCID for a failed resource
		if resource.isErrorResource {
			// store failed resource reference, will be used later to remove InterpolationString type values when generating config
			failedResourceReferenceSet[resource.getTerraformReference()] = true
			if _, ok := referenceMap[resource.id]; ok {
				delete(referenceMap, resource.id)
			}

			// delete any entries that have references to a failed resource
			// e.g. oci_core_instance.instance1.volume_id should be replaced by volume ocid if instance import failed
			if ok := referenceResourceNameSet[resource.terraformName]; ok {
				for key, value := range referenceMap {
					valueParts := strings.Split(value, ".")
					if len(valueParts) < 3 {
						continue
					}
					if valueParts[1] == resource.terraformName {
						delete(referenceMap, key)
					}
				}
			}
		}
	}
}

/*
	Initialize Terraform struct from executable provided
	Terraform struct will later be copied to each resource discovery step for parallel runs
*/
func createTerraformStruct(args *ExportCommandArgs) (*tfexec.Terraform, string, error) {

	utils.Logln("[INFO] validating Terraform CLI")
	var err error
	terraformBinPath := utils.GetEnvSettingWithBlankDefault(globalvar.TerraformBinPathName)
	if terraformBinPath == "" {
		terraformBinPath, err = tfinstall.Find(context.Background(), tfinstall.LookPath())
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] error finding terraform CLI, either specify the path to terraform CLI "+
				"including name using env var 'terraform_bin_path' or add terraform CLI to your system path: %s", err.Error())
		}
	} else {
		// Validate the path provided
		file, err := os.Stat(terraformBinPath)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR]: error verifying the terraform binary path provided %s: %s", terraformBinPath, err)
		}
		if file.IsDir() {
			return nil, "", fmt.Errorf("[ERROR]: terraform CLI path provided is a directory: %s", terraformBinPath)
		}

	}

	// Initialize Terraform struct from executable provided
	// Setting the global var 'tf' here, will be later using while running terraform init and import
	tf, err := tfexec.NewTerraform(*args.OutputDir, terraformBinPath)
	if err != nil {
		return nil, terraformBinPath, err
	}

	// Set log path for TF Exec
	logPath := os.Getenv(globalvar.EnvLogFile)
	if logPath != "" {
		utils.Logf("[INFO] setting log path for Terraform exec to '%s'", logPath)
		if err := tf.SetLogPath(logPath); err != nil {
			return nil, terraformBinPath, err
		}
	}

	backgroundCtx := context.Background()

	// discard stdout to avoid showing tf version results in logs
	tf.SetStdout(ioutil.Discard)

	// validate Terraform CLI
	if tfVersion, _, err := tf.Version(backgroundCtx, true); err != nil {
		return nil, terraformBinPath, fmt.Errorf("[ERROR] error verifying the terraform binary provided: %s", err)
	} else {
		utils.Debugf("[DEBUG] version %v", tfVersion)

		// check for tf_version and terraform CLI version so as to avoid
		// scenarios where config is not compatible with TF version of state file
		// version should be >= 0.12.0
		inputTfVersion := "v" + tfVersion.String()

		if semver.MajorMinor(inputTfVersion) == semver.MajorMinor("v"+string(TfVersion11)) {
			return nil, terraformBinPath, fmt.Errorf("[ERROR] resource discovery does not support v0.11.* CLI, "+
				"please specify terraform CLI with version v0.12.*, terraform version provided: %s", tfVersion.String())
		}

		executableVersion := semver.MajorMinor(inputTfVersion)
		configVersion := semver.MajorMinor("v" + tfHclVersion.toString())

		if executableVersion < configVersion {
			return nil, terraformBinPath, fmt.Errorf("[ERROR] major and minor version of terraform CLI provided is not same as the generated configuration version, "+
				"configuration version: %s, terraform CLI version: %s, please provide CLI version >= %s ", tfHclVersion.toString(), tfVersion.String(), tfHclVersion.toString())
		}
	}
	// enable stdout again to show init and import output in logs
	tf.SetStdout(os.Stdout)

	return tf, terraformBinPath, nil
}

func handlePanicFindResources(tfMeta *TerraformResourceAssociation, err *error) {
	if r := recover(); r != nil {
		utils.Logf("[WARN] recovered from panic in findResourcesGeneric for resource: %s \n continuing discovery...", tfMeta.resourceClass)
		returnErr := fmt.Errorf("panic in findResourcesGeneric for resource %s", tfMeta.resourceClass)
		*err = returnErr
		debug.PrintStack()
	}
}

func cleanupTempStateFiles(ctx *resourceDiscoveryContext) {

	/* Clean up temp state files for individual services */
	if err := os.RemoveAll(fmt.Sprintf("%s%stmp%s", *ctx.OutputDir, string(os.PathSeparator), string(os.PathSeparator))); err != nil {
		utils.Logf("[ERROR] Error removing tmp state files: %s", err.Error())
	}
}

func parseDeliveryPolicy(policy interface{}) string {
	backoffRetryPolicy := policy.(map[string]interface{})["backoff_retry_policy"].([]interface{})
	maxRetryDuration := backoffRetryPolicy[0].(map[string]interface{})["max_retry_duration"]
	policyType := backoffRetryPolicy[0].(map[string]interface{})["policy_type"]
	return fmt.Sprintf("{\"backoffRetryPolicy\":{\"maxRetryDuration\":%v,\"policyType\":\"%v\"}}", maxRetryDuration, policyType)
}
