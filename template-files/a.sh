#!/bin/bash - 
#===============================================================================
#
#          FILE: a.sh
# 
#         USAGE: ./a.sh 
# 
#   DESCRIPTION: 
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: YOUR NAME (), 
#  ORGANIZATION: 
#       CREATED: 06/04/20 09:57:06
#      REVISION:  ---
#===============================================================================

set -o nounset                              # Treat unset variables as an error
find . \( -name "*.go" -o -name "*.toml" -o -name "*.txt" \) -print |
	while read R
	do
		mv $R $R.gohtml
	done

