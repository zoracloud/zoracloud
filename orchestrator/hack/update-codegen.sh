# The set command is mainly used to display the shell variables that already exist in the system 
# and to set the new variable values of shell variables When changing a shell feature with a set, 
# the symbols “+” and “-” turn the specified mode on and off, respectively. 
# The set command cannot define new shell variables. 
# If you want to define a new variable, you can use the declare command to define it in the form of the variable name = value.
# https://developpaper.com/examples-of-the-meaning-and-use-of-set-and-shopt-command-options-in-the-shell/
# errexit -e	Exit when the command returns a non-zero exit status (failed). Not set when reading initialization file

############################
set -o errexit
set -o nounset
set -o pipefail
############################

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})
echo "SCRIPT_ROOT is $SCRIPT_ROOT"