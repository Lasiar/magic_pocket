#!/bin/bash - 
#===============================================================================
#
#          FILE: won.sh
# 
#         USAGE: ./won.sh 
# 
#   DESCRIPTION: 
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: YOUR NAME (), 
#  ORGANIZATION: 
#       CREATED: 07/11/2017 12:40
#      REVISION:  ---
#===============================================================================

set -o nounset                              # Treat unset variables as an error
day=$(date '+%u')
if [ "$day" -lt 5  ]  ; then
	echo "in first if"
	time=$(date '+%H%M')
	if  [ "$time" -ge 0845 -a "$time" -lt 2100   ]  ; then
		echo "in second if"
		for mac  in $(./magic_packet -ls) ; do
			echo $mac | grep  -i '[0-9A-F]\{2\}\(:[0-9A-F]\{2\}\)\{5\}' 2> /dev/null
		done
	fi
fi



