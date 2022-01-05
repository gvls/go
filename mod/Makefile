# =================
# ===use method====
# =================
#
# if not add parameter in end of make, execute first "name"
# if add @ to front of command, terminal will not output command
# <name> : <condition>
# 	<command>
#
# ${variateName} get value of variate
#
# when "name" is same as name of file, also execute "name"
# .PHONY



# path and name of binary file
BINARY=./tifapi

# parameter
VERSION=`git log --pretty=format:"%H" -1`
BUILD=`date +%FT%T%z`

# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY:help
help:
	@echo "make build : \n\tgenerate binary file with parameters to specific directory \n\tparameters : (${LDFLAGS})\n\tposition : (${BINARY})"
	@echo "make install : \n\tcopy binary file to ${GOPATH}/bin/ "
	@echo "make clean: \n\tclean binary file"

.PHONY:build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}

.PHONY:install
install:
	go install ${LDFLAGS}

.PHONY:clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

