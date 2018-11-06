**HTTP Replay**
=====

Synopsis
-----

```
	// install the hook for HTTP replaying:
	
		if h, ok := client.HTTPClient.(*http.Client); ok {
			_, err := httpreplay.InstallRecorder(h)
			if err != nil {
				return err
			}
		}

func TestMyServiceResource_basic(t *testing.T) {
    // In a unit test, tell the recorder what test we are running
    httpreplay.SetScenario("TestMyServiceResource_basic")
    defer httpreplay.SaveScenario()
    ... testing happens ...
}
```

Description
-----

This library provides a recording mechanism for the OCI-GO-SDK. It hooks into
the 'Transport' layer of the HTTP request, calling through to the real
network as needed.

* bypass (default): Do nothing
* record: Store the Interaction
* replay: Load the Interaction file and send back the response
        
Select the record or replay mode by specifying a build tag to go: `-tags <mode>`


Functions
-----

InstallRecorder
* Install hooks into the `http.Client` to allow the record/replay library to 
  intercept HTTP calls.  `InstallRecorder` tries to be safe if called multiple
  times, but it is possible to fool it.  Best is to only call it once per 
  `http.Client`.

SetScenario
* Name the scenario that is about to run.  

  Currently, this specifies 
  the filename that the scenario data will be saved into, 
  with a `.yaml` extension.  
  - If `-tags record` is specified, the 
  requests are written to this file in the `SaveScenario` call. The Interaction file
   will be stored in directory  `~<prj_path>/record/` with the name passed into. 
   - If `-tags replay` is specified, then a file by that name is immediately 
  read and used for generating replies to network requests.

SaveScenario
* Save the scenario data.

  Currently, if `-tags record` is specified, this writes all the 
  recorded requests to the file named in `SetScenario`.
  

Record Storage 
-----
   
* In record mode: After running the test case, the record file will be stored under "oci/record/".
* In replay mode: Look for the record file under "oci/record/" and throw error if it is not found.


Example usage 
-----
* To run normally: `go test`
* Or run 1 specific test case: `go test -run <testname>`
----
* To record interactions: `go test -tags record`
* Or to record 1 specific test case: `go test -run <testname> -tags record`
----
* To replay interactions: `go test -tags replay`
* Or to replay 1 specific test case: `go test -run <testname> -tags replay`

### Example Output

Run with recording turned on, the test portion takes 2411 seconds:
    
    > go test -v -timeout 120m -run TestResourceCoreImageTestSuite -tags record
    === RUN   TestResourceCoreImageTestSuite
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_basic
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageTuple
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageUri
    --- PASS: TestResourceCoreImageTestSuite (2411.14s)
        --- PASS: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_basic (2410.58s)
        --- SKIP: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageTuple (0.00s)
            core_image_resource_test.go:191: Long running test, requires per tenancy namespace + bucket + image export object to run
        --- SKIP: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageUri (0.00s)
            core_image_resource_test.go:155: Long running test, requires exported image available via public url
    PASS


Now that we have a recording, run in replay mode, note that it is only 4.09 seconds:

    > go test -v -run TestResourceCoreImageTestSuite -tags replay
    === RUN   TestResourceCoreImageTestSuite
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_basic
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageTuple
    === RUN   TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageUri
    --- PASS: TestResourceCoreImageTestSuite (4.09s)
        --- PASS: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_basic (3.60s)
        --- SKIP: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageTuple (0.00s)
            core_image_resource_test.go:191: Long running test, requires per tenancy namespace + bucket + image export object to run
        --- SKIP: TestResourceCoreImageTestSuite/TestAccResourceCoreImage_createFromExport_objectStorageUri (0.00s)
            core_image_resource_test.go:155: Long running test, requires exported image available via public url
    PASS
