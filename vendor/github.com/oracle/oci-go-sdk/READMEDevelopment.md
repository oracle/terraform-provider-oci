# Running the Code Generator
## Pre-requisites
- Maven
- Python (and tools, see below)
- Go(make sure you set the GOPATH enviroment correctly)
- oci-go-sdk commons go package
  - You can install the lastest version by installing pulling the current version of the go sdk
- Make

## Python setup
Use ``virtualenv`` to create a virtual environment that runs Python 2.x (of course, you have to have Python 2.x installed somewhere):

    # Install the virtual environment in the current directory
    virtualenv --python=<path to Python 2.x executable> temp/python2
    # Activate virtual environment
    source temp/python2/bin/activate
    # Install packages
    pip install PyYAML
    pip install six
    


## Start here!
The build functionality is driven by 2 make files

- Makefile: Is public and exposed, builds the sdk and runs unittest
- MakefileDevelopment: Private. generates new sdk, runs private integtests
 


## Generating
The generation makefile is: ***MakefileDevelopment.mk***
You run the code generator by executing. This will generate the code as well as build it

    make -f MakefileDevelopment.mk build
    
After executing this command the source code will be placed under the canonical repository `$GOPATH/src/$PROJECT_NAME` where $PROJECT_NAME is the fully qualified project name: `github.com/oracle/oci-go-sdk`

## Release
Instead of the `build` target. Execute the release target like so:

    make -f MakefileDevelopment.mk release

Do not forget to setup major, minor versions by updating the variables in: `MakefileDevelopment.mk`

    VER_MAJOR = x
    VER_MINOR = y
    
## Testing

Execute the `test-` targets, like so:
    
    make -f MakefileDevelopment.mk test-audit  ## Will execute the integ tests for audit
    
    make -f MakefileDevelopment.mk test-all ## Will execute all integtest
    



