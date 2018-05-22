VERSION = 0.0.1

ROOTDIR = $(shell pwd)
APPNAME = go-props
GODIR = /tmp/gopath
SRCDIR = ${GODIR}/src/${APPNAME}
TARGET = bin/${APPNAME}
TESTDATA = ${ROOTDIR}/testdata

GOENV = GOPATH=${GODIR}:${GOPATH} GO15VENDOREXPERIMENT=1

GO = ${GOENV} go
DEP = ${GOENV} dep

LDFLAGS = 

.PHONY: build
build: ${SRCDIR}
	${GO} build -i -ldflags="${LDFLAGS}" -o ${TARGET} ${APPNAME}

${SRCDIR}:
	mkdir -p bin
	mkdir -p `dirname "${SRCDIR}"`
	ln -s `dirname "${ROOTDIR}"`/${APPNAME} ${SRCDIR}

.PHONY: init
init: ${SRCDIR} update

.PHONY: update
update: ${SRCDIR}
	cd ${SRCDIR} && ${DEP} ensure || true

.PHONY: test
test: ${SRCDIR} build
	env PROPS_TARGET=${TARGET} PROPS_TESTDATA=${TESTDATA} python ${TESTDATA}/driven.py -v

.PHONY: lint
lint:
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -exec golint {} \;

.PHONY: go-env
go-env:
	@go env
