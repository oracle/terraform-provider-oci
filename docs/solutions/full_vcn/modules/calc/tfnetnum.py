"""
tfnetnum.py - Calculate the net number to be used in the Terraform cidrsubnet
              call for a sequence of subnet masks
@author: Steven B. Nelson, Principal Solutions Architect, Oracle Cloud Infrastructure
@date: 14 June 2018

Net number is a Terraform specific parameter to the cidrsubnet interpolation function.  The net number parameter is the
relative index position from the begining of the supernet.  So, if you are starting with a /24 subnet, and want to add
three /27 in succession, you would have three calls with netnumber 0, 1, 2.  However, if you have a /30, /27, /30, you
would end up with a sequence of 0, 1, 9 - first subnet break at the beginning (0 - /30), second after at the next break
for a /27 (1 - /27), then the next break at the ninth subnet for a /30 (9 - /30).

This works by calculating the successive netnumbers by accumulating the number of addresses consumed out of the
supernet, then determining which breakpoint in the supernet that corresponds to the required one for the subnet. """

import json
import sys

# Get input from Terraform datasource call.
inputDict = json.loads(str(sys.stdin.readline()))

# Instantiate the list that will hold the sequence of net numbers and accumulator
nnList = []
accum = 0

# Convert the input from Terraform into a simple list
maskList = [int(mask) for mask in inputDict.get(inputDict.keys()[0]).split(",")]

# Loop through each mask and calculate the number of addresses consumed. 
# Insert the result in to the accumlator
for maskPos in range(0, len(maskList)):
    # The first net number is always zero
    if maskPos == 0:
        accum = 0
    else:
        '''
        If the previous mask is larger (smaller subnet) then just accumlate the lower boundary.  But if the 
        previous mask is greater (larger subnet) then we have to accumulate to the upper boundary so that the larger 
        mask is calculated at the right spot.  Otherwise, you will try to calculate on the lower boundary on the 
        larger mask and end up in the middle of the smaller mask subnet. 
        '''
        if maskList[maskPos - 1] >= maskList[maskPos]:
            accum = accum + 2 ** (32 - maskList[maskPos])
        else:
            accum = accum + 2 ** (32 - maskList[maskPos - 1])
    # The net number is just the index of the number of addresses when compared to the number required by the mask
    netnum = accum / 2 ** (32 - maskList[maskPos])
    nnList.append(str(netnum))
print(json.dumps({"nnStrList": ",".join(nnList)}))
