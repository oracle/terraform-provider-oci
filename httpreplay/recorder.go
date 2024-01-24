// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package httpreplay

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"sync"
	"time"
)

// Mode represents recording/playback mode
type Mode int

// Recorder states
const (
	ModeRecording Mode = iota
	ModeReplaying
	ModeDisabled
)

// Transformer converts a request and a saved interaction into a result.  The Interaction is passed by value to suggest that it should not be modified.
type Transformer func(*Request, Interaction, *Response)

type jsonObj map[string]interface{}
type jsonArr []jsonObj
type jsonStr string

func defaultTransformer(req *Request, i Interaction, res *Response) {
}

type roundTripperProxy struct {
	recorder *Recorder
	chained  http.RoundTripper
}

func (rtp *roundTripperProxy) RoundTrip(r *http.Request) (*http.Response, error) {
	res, err := rtp.recorder.RoundTrip(r, rtp.chained)
	if err != nil && err.Error() == "Requested interaction not found" {
		debugLogf("stop RoundTrip for err: %v", err)
		panic(err)
	}

	return res, err
}

func (rtp *roundTripperProxy) CancelRequest(r *http.Request) {
	rtp.recorder.CancelRequest(r, rtp.chained)
}

// Recorder represents a type used to record and replay
// client and server interactions
type Recorder struct {
	// Operating mode of the recorder
	mode Mode

	// Scenario used by the recorder
	scenario *Scenario

	// transformer is used to adjust responses to match changes in requests
	transformer Transformer

	// count is for debug logging -- how many requests have been matched
	count int
}

// HookTransport makes a new transport and chains the one passed in with it, returning the new one
func (r *Recorder) HookTransport(client *http.Client) error {
	if r == nil {
		return errors.New("The test case missing calling SetScenerio() ")
	}
	if _, ok := client.Transport.(*roundTripperProxy); !ok {
		proxy := roundTripperProxy{
			recorder: r,
			chained:  client.Transport,
		}
		client.Transport = &proxy
	}
	return nil
}

// SetTransformer can be used to override the default (no-op) transformer
func (r *Recorder) SetTransformer(t Transformer) {
	r.transformer = t
}

var mut sync.RWMutex

func (r *Recorder) invokeTransformer(req *http.Request) (*Interaction, *Response, error) {
	mut.Lock()
	defer mut.Unlock()
	if err := req.ParseForm(); err != nil {
		debugLogf("\t-> Returning error from invokeTransformer: %v", err)
		//return nil, nil, err
	}

	reqBody := make([]byte, req.ContentLength)
	if _, err := io.ReadFull(req.Body, reqBody); err != nil {
		debugLogf("\t-> Returning error from invokeTransformer: %v", err)
		return nil, nil, err
	}

	var bodyParsed interface{}
	if len(reqBody) != 0 {
		bodyParsed, _ = unmarshal(reqBody)
	}

	request := Request{
		Body:       string(reqBody),
		BodyParsed: bodyParsed,
		Form:       req.PostForm,
		Headers:    req.Header,
		URL:        req.URL.String(),
		Method:     req.Method,
	}

	i, err := r.scenario.GetInteraction(request)
	if err != nil {
		if err.Error() == "Requested interaction not found" {
			debugLogf("\t-> Convert full path of request to find Interaction:")
			i, err = r.scenario.GetInteractionWithFullPath(request)
			if err != nil {
				debugLogf("\t-> Returning error from invokeTransformer: %v", err)
				return nil, nil, err
			}
		}
	}
	i.Request.BodyParsed, _ = unmarshal([]byte(i.Request.Body))
	i.Response.BodyParsed, _ = unmarshal([]byte(i.Response.Body))
	debugLogf("\t=> => Request %d matched interaction %d", r.count, i.Index)
	r.count++

	res := i.Response
	response := Response{
		Body:     res.Body,
		Headers:  res.Headers,
		Status:   res.Status,
		Code:     res.Code,
		Duration: res.Duration,
	}

	if len(res.Body) > 0 {
		if bodyParsed, err := unmarshal([]byte(res.Body)); err == nil {
			response.BodyParsed = bodyParsed
		}
	}

	r.transformer(&request, *i, &response)

	// Pick up changes from response.BodyParsed and put them into
	// response.Body to send back to the ultimate requestor.
	if response.BodyParsed != nil {
		resBody, err := json.Marshal(response.BodyParsed)
		if err != nil {
			debugLogf("\t-> Returning error from invokeTransformer: %v", err)
			return nil, nil, err
		}
		response.Body = string(resBody)
	}

	return i, &response, nil
}

func (r *Recorder) recordInteraction(req *http.Request, realTransport http.RoundTripper) (*Interaction, *Response, error) {
	// Copy the original request, so we can read the form values
	reqBytes, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		debugLogf("-=-=-=- Error from httputil.DumpRequestOut: %v", err)
		return nil, nil, err
	}

	reqBuffer := bytes.NewBuffer(reqBytes)
	copiedReq, err := http.ReadRequest(bufio.NewReader(reqBuffer))
	if err != nil {
		debugLogf("-=-=-=- Error from http.ReadRequest: %v", err)
		return nil, nil, err
	}

	err = copiedReq.ParseForm()
	if err != nil {
		debugLogf("-=-=-=- Error from copiedReq.ParseForm: %v", err)
		return nil, nil, err
	}

	reqBody := &bytes.Buffer{}
	if req.Body != nil {
		// Record the request body so we can add it to the scenario
		req.Body = ioutil.NopCloser(io.TeeReader(req.Body, reqBody))
	}

	// Perform client request to its original
	// destination and record interactions
	resp, err := realTransport.RoundTrip(req)
	if err != nil {
		debugLogf("-=-=-=- Error from realTransport.RoundTrip: %v", err)
		return nil, nil, err
	}

	req.Header.Del("Authorization")
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		debugLogf("-=-=-=- Error from ioutil.ReadAll: %v", err)
		return nil, nil, err
	}

	// Add interaction to scenario
	interaction := &Interaction{
		Request: Request{
			Body:    reqBody.String(),
			Form:    copiedReq.PostForm,
			Headers: req.Header,
			URL:     req.URL.String(),
			Method:  req.Method,
		},
		Response: Response{
			Body:    string(respBody),
			Headers: resp.Header,
			Status:  resp.Status,
			Code:    resp.StatusCode,
		},
	}
	r.scenario.AddInteraction(interaction)

	return interaction, &interaction.Response, nil
}

func (r *Recorder) requestHandler(req *http.Request, realTransport http.RoundTripper) (*Interaction, *Response, error) {
	// Return interaction from scenario if in replay mode
	if r.mode == ModeReplaying {
		return r.invokeTransformer(req)
	}
	return r.recordInteraction(req, realTransport)
}

func InstallRecorderForRecodReplay(client *http.Client, recorder *Recorder) (HTTPRecordingClient, error) {
	err := recorder.HookTransport(client)
	if err != nil {
		debugLogf("Fail install Recorder: %v", err)
		return nil, err
	}
	return client, nil
}

// unmarshal seems like it should not be necessary, but sometimes json.Unmarshal will choose a type of map[interface{}]interface{} which will not downcast into a map[string]interface{}.
func unmarshal(body []byte) (interface{}, error) {
	var bodyParsed interface{}

	var decode = func(result interface{}) error {
		jsonDecoder := json.NewDecoder(bytes.NewReader(body))
		jsonDecoder.UseNumber()
		return jsonDecoder.Decode(result)
	}

	if 0 < len(body) {
		var bodyObjParsed jsonObj
		//if err := json.Unmarshal(body, &bodyObjParsed); err == nil {
		if err := decode(&bodyObjParsed); err == nil {
			bodyParsed = bodyObjParsed
		} else {
			var bodyArrParsed jsonArr
			//if err := json.Unmarshal(body, &bodyArrParsed); err == nil {
			if err := decode(&bodyArrParsed); err == nil {
				bodyParsed = bodyArrParsed
			} else {
				var bodyStrParsed jsonStr
				//if err := json.Unmarshal(body, &bodyStrParsed); err == nil {
				if err := decode(&bodyStrParsed); err == nil {
					bodyParsed = bodyStrParsed
				} else {
					return nil, err
				}
			}
		}
	}
	return bodyParsed, nil
}

// NewRecorder creates a new recorder
func NewRecorder(scenarioName string) (*Recorder, error) {
	// Default mode is "replay" if file exists
	return NewRecorderAsMode(scenarioName, ModeReplaying)
}

// NewRecorderAsMode creates a new recorder in the specified mode
func NewRecorderAsMode(scenarioName string, mode Mode) (*Recorder, error) {
	var s *Scenario
	var err error

	if mode != ModeDisabled {
		// Depending on whether the scenario file exists or not we
		// either create a new empty scenario or load from file
		if mode == ModeRecording {
			// Create new scenario and enter in recording mode
			s = NewScenario(scenarioName)
		} else {
			// Load scenario from file and enter replay mode
			s, err = Load(scenarioName)
			if err != nil {
				return nil, err
			}
			mode = ModeReplaying
		}
	}

	r := &Recorder{
		mode:        mode,
		scenario:    s,
		transformer: defaultTransformer,
	}

	return r, nil
}

// Stop is used to stop the recorder and save any recorded interactions
func (r *Recorder) Stop() error {
	if r.mode == ModeRecording {
		if err := r.scenario.Save(); err != nil {
			return err
		}
	}

	return nil
}

// RoundTrip implements the http.RoundTripper interface
func (r *Recorder) RoundTrip(req *http.Request, realTransport http.RoundTripper) (*http.Response, error) {
	if r.mode == ModeDisabled {
		response, err := realTransport.RoundTrip(req)
		if err != nil {
			return response, err
		}
	}
	// Pass scenario and mode to handler, so that interactions can be
	// retrieved or recorded depending on the current recorder mode
	interaction, resp, err := r.requestHandler(req, realTransport)

	if err != nil {
		debugLogf("-==-==-==- Error return from RoundTrip: %v", err)
		return nil, err
	}

	select {
	case <-req.Context().Done():
		debugLogf("-==-==-==- Error return from RoundTrip: %v", req.Context().Err())
		return nil, req.Context().Err()
	default:
		buf := bytes.NewBuffer([]byte(resp.Body))
		// apply the duration defined in the interaction
		if false && resp.Duration != "" {
			d, err := time.ParseDuration(interaction.Duration)
			if err != nil {
				debugLogf("-==-==-==- Error return from RoundTrip: %v", err)
				return nil, err
			}
			// block for the configured 'duration' to simulate the network latency and server processing time.
			if _, ok := os.LookupEnv("TF_VAR_SLOW_REPLAY"); ok {
				<-time.After(d)
			}
		}

		theResp := http.Response{
			Status:        resp.Status,
			StatusCode:    resp.Code,
			Proto:         "HTTP/1.0",
			ProtoMajor:    1,
			ProtoMinor:    0,
			Request:       req,
			Header:        resp.Headers,
			Close:         true,
			ContentLength: int64(buf.Len()),
			Body:          ioutil.NopCloser(buf),
		}
		return &theResp, nil
	}
}

// CancelRequest implements the github.com/coreos/etcd/client.CancelableTransport interface
func (r *Recorder) CancelRequest(req *http.Request, realTransport http.RoundTripper) {
	type cancelableTransport interface {
		CancelRequest(req *http.Request)
	}
	if ct, ok := realTransport.(cancelableTransport); ok {
		ct.CancelRequest(req)
	}
}

// SetMatcher sets a function to match requests against recorded HTTP interactions.
func (r *Recorder) SetMatcher(matcher Matcher) {
	if r.scenario != nil {
		r.scenario.Matcher = matcher
	}
}
