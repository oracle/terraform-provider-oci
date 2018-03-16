// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

// Package common provides supporting functions and structs used by service packages
package common

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

const (
	// DefaultHostURLTemplate The default url template for service hosts
	DefaultHostURLTemplate        = "%s.%s.oraclecloud.com"
	defaultScheme                 = "https"
	defaultSDKMarker              = "Oracle-GoSDK"
	defaultUserAgentTemplate      = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultRequestTimeout         = 0
	defaultConnectionTimeout      = 10 * time.Second
	defaultTLSHandshakeTimeout    = 5 * time.Second
	defaultConfigFileName         = "config"
	defaultConfigDirName          = ".oci"
	secondaryConfigDirName        = ".oraclebmc"
	maxBodyLenForDebug            = 1024 * 1000
	generatedRetryTokenLength     = 30
	retryTokenKey                 = "opc-retry-token"
	absoluteMaximumRequestTimeout = 7 * 24 * time.Hour
)

// RequestInterceptor function used to customize the request before calling the underlying service
type RequestInterceptor func(*http.Request) error

// HTTPRequestDispatcher wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HTTPRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

// BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	//HTTPClient performs the http network operations
	HTTPClient HTTPRequestDispatcher

	//Signer performs auth operation
	Signer HTTPRequestSigner

	//Provides an on-behalf-of token
	Obo OboTokenProvider

	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor

	//The host of the service
	Host string

	//The user agent
	UserAgent string

	//Base path for all operations of this client
	BasePath string

	//Random number generator for this client
	generator *rand.Rand
}

func defaultUserAgent() string {
	userAgent := fmt.Sprintf(defaultUserAgentTemplate, defaultSDKMarker, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
	return userAgent
}

var clientCounter int64

func getNextSeed() int64 {
	newCounterValue := atomic.AddInt64(&clientCounter, 1)
	return newCounterValue + time.Now().UnixNano()
}

func newBaseClient(signer HTTPRequestSigner, dispatcher HTTPRequestDispatcher) BaseClient {
	return BaseClient{
		UserAgent:   defaultUserAgent(),
		Interceptor: nil,
		Signer:      signer,
		Obo:         NewEmptyOboTokenProvider(),
		HTTPClient:  dispatcher,
		generator:   rand.New(rand.NewSource(getNextSeed())),
	}
}

func defaultHTTPDispatcher() http.Client {
	httpClient := http.Client{
		Timeout: defaultRequestTimeout,
	}
	return httpClient
}

func defaultBaseClient(provider KeyProvider) BaseClient {
	dispatcher := defaultHTTPDispatcher()
	signer := DefaultRequestSigner(provider)
	return newBaseClient(signer, &dispatcher)
}

//DefaultBaseClientWithSigner creates a default base client with a given signer
func DefaultBaseClientWithSigner(signer HTTPRequestSigner) BaseClient {
	dispatcher := defaultHTTPDispatcher()
	return newBaseClient(signer, &dispatcher)
}

// NewClientWithConfig Create a new client with a configuration provider, the configuration provider
// will be used for the default signer as well as reading the region
// This function does not check for valid regions to implement forward compatibility
func NewClientWithConfig(configProvider ConfigurationProvider) (client BaseClient, err error) {
	var ok bool
	if ok, err = IsConfigurationProviderValid(configProvider); !ok {
		err = fmt.Errorf("can not create client, bad configuration: %s", err.Error())
		return
	}

	client = defaultBaseClient(configProvider)
	return
}

func getHomeFolder() string {
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}

// DefaultConfigProvider returns the default config provider. The default config provider
// will look for configurations in 3 places: file in $HOME/.oci/config, HOME/.obmcs/config and
// variables names starting with the string TF_VAR. If the same configuration is found in multiple
// places the provider will prefer the first one.
func DefaultConfigProvider() ConfigurationProvider {
	homeFolder := getHomeFolder()
	defaultConfigFile := path.Join(homeFolder, defaultConfigDirName, defaultConfigFileName)
	secondaryConfigFile := path.Join(homeFolder, secondaryConfigDirName, defaultConfigFileName)

	defaultFileProvider, _ := ConfigurationProviderFromFile(defaultConfigFile, "")
	secondaryFileProvider, _ := ConfigurationProviderFromFile(secondaryConfigFile, "")
	environmentProvider := environmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}

	provider, _ := ComposingConfigurationProvider([]ConfigurationProvider{defaultFileProvider, secondaryFileProvider, environmentProvider})
	Debugf("Configuration provided by: %s", provider)
	return provider
}

func (client *BaseClient) prepareRequest(request *http.Request) (err error) {
	if client.UserAgent == "" {
		return fmt.Errorf("user agent can not be blank")
	}

	if request.Header == nil {
		request.Header = http.Header{}
	}
	request.Header.Set("User-Agent", client.UserAgent)

	if !strings.Contains(client.Host, "http") &&
		!strings.Contains(client.Host, "https") {
		client.Host = fmt.Sprintf("%s://%s", defaultScheme, client.Host)
	}

	clientURL, err := url.Parse(client.Host)
	if err != nil {
		return fmt.Errorf("host is invalid. %s", err.Error())
	}
	request.URL.Host = clientURL.Host
	request.URL.Scheme = clientURL.Scheme
	currentPath := request.URL.Path
	request.URL.Path = path.Clean(fmt.Sprintf("/%s/%s", client.BasePath, currentPath))
	return
}

func (client BaseClient) intercept(request *http.Request) (err error) {
	if client.Interceptor != nil {
		err = client.Interceptor(request)
	}
	return
}

func checkForSuccessfulResponse(res *http.Response) error {
	familyStatusCode := res.StatusCode / 100
	if familyStatusCode == 4 || familyStatusCode == 5 {
		return newServiceFailureFromResponse(res)
	}
	return nil

}

func generateRetryToken(randGen *rand.Rand) string {
	alphanumericChars := []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	retryToken := make([]rune, generatedRetryTokenLength)
	for i := range retryToken {
		retryToken[i] = alphanumericChars[randGen.Intn(len(alphanumericChars))]
	}
	return string(retryToken)
}

func addRetryTokenToRequestIfNeeded(request *http.Request, generator *rand.Rand) {
	// ensure request.Header exists
	if request.Header == nil {
		request.Header = http.Header{}
	}

	if _, present := request.Header[retryTokenKey]; !present {
		generatedRetryToken := generateRetryToken(generator)
		request.Header[retryTokenKey] = []string{generatedRetryToken}
	}
}

// getRetryPolicy assembles a retry policy using the specified retry policy options, and information about the request.
// The retry policy options decorate the default retry policy for any given request (defined by services). If no options
// are present, then the default behavior will be No Retry. This behavior will change in the future, as teams evolve
// the default retry behavior per supported operation. If any options are present, then the default behavior is assumed
// to be an Exponential Backoff retry policy, and the specified options override that policy.
func getRetryPolicy(request *http.Request, options ...RetryPolicyOption) RetryPolicy {
	if len(options) == 0 {
		return NoRetryPolicy()
	}

	return BuildRetryPolicy(options...)
}

// GetMaximumTimeout ensures that the policy MaximumTimeout (which can be set to 'unlimited') is still bounded by
// the OCI absolute maximum (currently set to 1 week).
func GetMaximumTimeout(policy RetryPolicy) time.Duration {
	// even if a user says poll forever by specifying zero for the maximum timeout, we're going to stop them
	if policy.MaximumTimeout == 0 {
		return absoluteMaximumRequestTimeout
	}
	return policy.MaximumTimeout
}

type durationExceedsDeadlineError struct{}

func (durationExceedsDeadlineError) Error() string { return "now() + duration exceeds context deadline" }

// DurationExceedsDeadline is the error returned by Call() when GetNextDuration() returns a time.Duration that would
// force the user to wait past the request deadline before re-issuing a request. This enables us to exit early, since
// we cannot succeed based on the configured retry policy.
var DurationExceedsDeadline error = durationExceedsDeadlineError{}

// CallConfig is an encapsulation of the arguments needed to configure the Call function.
type CallConfig struct {
	KeepResponseBodyOpen bool                              // Whether not to close the response body. Defaults to false (always close).
	ResponseCallback     func(*http.Response, error) error // Callback to process the response & error.
	RetryPolicyOptions   []RetryPolicyOption               // Retry policy options.
}

// Call executes the http request with the given context according to the specified retry policy (if present)
func (client BaseClient) Call(ctx context.Context, request *http.Request, config CallConfig) error {
	responseCallback := config.ResponseCallback
	// Define a no-op/passthrough callback if a callback wasn't supplied.
	if responseCallback == nil {
		responseCallback = func(response *http.Response, err error) error {
			return err
		}
	}

	err := client.prepareRequest(request)
	if err != nil {
		return err
	}

	if len(config.RetryPolicyOptions) == 0 {
		response, err := client.doRequest(ctx, request)
		if !config.KeepResponseBodyOpen {
			defer closeBodyIfValid(response)
		}
		return responseCallback(response, err)
	}

	// the request Body is closed by the underlying http.Transport
	// => store off the request body as a byte array for reconstituting the body on each retry attempt
	requestBodyAsByteSlice := []byte{}
	if request.Body != nil {
		var err error
		requestBodyAsByteSlice, err = ioutil.ReadAll(request.Body)
		if err != nil {
			// error reading the body of the request
			return err
		}
	}

	policy := getRetryPolicy(request, config.RetryPolicyOptions...)
	addRetryTokenToRequestIfNeeded(request, client.generator)

	deadlineContext, deadlineCancel := context.WithTimeout(ctx, GetMaximumTimeout(policy))
	defer deadlineCancel()

	select {
	case <-deadlineContext.Done():
		// return why the request was aborted (could be user interrupted or deadline exceeded)
		return responseCallback(nil, deadlineContext.Err())
	default:
		// non-blocking select
	}

	for currentOperationAttempt := uint(1); ShouldContinueIssuingRequests(currentOperationAttempt, policy.MaximumNumberAttempts); currentOperationAttempt++ {
		// reset the request body on each operation attempt
		request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyAsByteSlice))

		Debugln(fmt.Sprintf("operation attempt #%v", currentOperationAttempt))

		response, err := client.doRequest(deadlineContext, request)
		if !config.KeepResponseBodyOpen {
			defer closeBodyIfValid(response)
		}

		if policy.ShouldRetryOperation(response, err, currentOperationAttempt) {
			// this conditional is explicitly not added to the encompassing if condition to retry based on response
			// => it is only to determine if, on the last round of this loop, we still skip sleeping (if we're the
			//    last attempt, then there's no point sleeping before we round the loop again and fall out to the
			//    Maximum Number Attempts exceeded error)
			if currentOperationAttempt != policy.MaximumNumberAttempts {
				// sleep before retrying the operation
				duration := policy.GetNextDuration(currentOperationAttempt)
				if deadline, ok := ctx.Deadline(); ok && time.Now().Add(duration).After(deadline) {
					// we want to retry the operation, but the policy is telling us to wait for a duration that exceeds
					// the specified overall deadline for the operation => instead of waiting for however long that
					// time period is and then aborting, abort now and save the cycles
					return responseCallback(response, DurationExceedsDeadline)
				}
				Debugln(fmt.Sprintf("waiting %v before retrying operation", duration))
				time.Sleep(duration)
			}
		} else {
			// we should NOT retry operation based on response and/or error => return
			return responseCallback(response, err)
		}
	}
	return responseCallback(nil, fmt.Errorf("maximum number of attempts exceeded (%v)", policy.MaximumNumberAttempts))
}

//doRequest executes the http request with the given context
func (client BaseClient) doRequest(ctx context.Context, request *http.Request) (response *http.Response, err error) {
	Debugln("Atempting to call downstream service")
	request = request.WithContext(ctx)

	//Fetch an obo token from the provider, and if one is returned put it into the request headers
	oboToken, err := client.Obo.OboToken()
	if err != nil {
		return
	}
	if oboToken != "" {
		request.Header.Set("Opc-Obo-Token", oboToken)
	}

	//Intercept
	err = client.intercept(request)
	if err != nil {
		return
	}

	//Sign the request
	err = client.Signer.Sign(request)
	if err != nil {
		return
	}

	IfDebug(func() {
		dumpBody := true
		if request.ContentLength > maxBodyLenForDebug {
			Logln("not dumping body too big")
			dumpBody = false
		}
		if dump, e := httputil.DumpRequest(request, dumpBody); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	//Execute the http request
	response, err = client.HTTPClient.Do(request)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		dumpBody := true
		if response.ContentLength > maxBodyLenForDebug {
			Logln("not dumping body too big")
			dumpBody = false
		}

		if dump, e := httputil.DumpResponse(response, dumpBody); e == nil {
			Logf("Dump Response %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	if err != nil {
		return
	}

	err = checkForSuccessfulResponse(response)
	return
}

// closeBodyIfValid closes the body of an http response if the response and the body are valid
func closeBodyIfValid(httpResponse *http.Response) {
	if httpResponse != nil && httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
}
