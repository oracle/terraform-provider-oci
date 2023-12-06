#!/usr/bin/env bash

help()
{
   echo "Usage: The script returns untagged tests if any in the terraform-provider-oci repository. The script only takes one option at a time."
   echo "Checks for untagged tests in all test files if no option passed."
   echo
   echo "Syntax: ./scripts/check-untagged-tests.sh [-h|s|f]"
   echo "options:"
   echo "h     print help"
   echo "s     check for untagged tests in specific service e.g. ./check-untagged-tests -s ocvp"
   echo "f     check for untagged tests in specific test file e.g. ./check-untagged-tests -f ocvp_esxi_host_test.go"
   echo
   echo "make command syntax: make SERVICE=ocvp check-untagged-tests OR make check-untagged-tests"
   echo
}

checkFn () {
    SEARCH_PATTERN=$1
    echo "SEARCH_PATTERN: $SEARCH_PATTERN"
    # create temp files
    tmpfile=$(mktemp /tmp/check-untagged-tests.XXXXXX)
    tmpfile2=$(mktemp /tmp/check-untagged-tests.XXXXXX)
    trap 'rm -f $tmpfile $tmpfile2' 0 2 3 15

    echo "==> Checking for untagged test functions"
    # grep test function header with tagging information
    grep -B 1 -h '^func Test.*{$' $SEARCH_PATTERN > $tmpfile
    if [[ $? -ne 0 ]]
    then
        echo "grep command failed for SEARCH_PATTERN" $SEARCH_PATTERN
        exit 1
    fi
    # remove the tagged tests from the grep output
    vi -es '+g/issue-routing-tag/-1,+1d' '+w !tee' '+wq!' $tmpfile
    # parse the untagged test function names
    grep -oE '^func Test[^(]+\(' $tmpfile | sed 's/(//; s/func //' > $tmpfile2

    if [ -s "$tmpfile2" ]
        then
        echo "==> Following untagged test functions found:"
        echo "________________________________________________________________________"
        cat $tmpfile2
        echo "________________________________________________________________________"
        echo "==> [Action Required] Please provide correct tagging information for above test functions"
        exit 1
    else
        echo "untagged test functions not found!"
        exit 0
    fi
}

if [ $# -eq 0 ]
    then
        echo "Checking all test files"
        checkFn "internal/integrationtest/*_test.go"
fi

# Get the options
while getopts "hs:f:" option; do
    case $option in
        h) # check for untagged tests in specific service
            help
            exit 0;;
        s) # check for untagged tests in specific service
            checkFn "internal/integrationtest/$OPTARG*_test.go";;
        f) # check for untagged tests in specific test file
            checkFn "internal/integrationtest/$OPTARG";;
        *) # invalid option
            echo "Error: Invalid option"
            help
            exit 1;;
    esac
done
