#!/bin/bash
yum update -y
cpupower frequency-set -g performance
yum install mongodb
