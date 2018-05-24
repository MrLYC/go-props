VERSION = 0.0.1

ROOTDIR = $(shell pwd)
APPNAME = go-props
GODIR = /tmp/gopath
SRCDIR = ${GODIR}/src/${APPNAME}
TARGET = bin/${APPNAME}
TESTDATA = ${ROOTDIR}/testdata

GO15VENDOREXPERIMENT = 1

GO = cd ${SRCDIR} && go
DEP = cd ${SRCDIR} && dep

LDFLAGS = 

.PHONY: .EXPORT_ALL_VARIABLES

.PHONY: build
build: ${SRCDIR} .EXPORT_ALL_VARIABLES
	${GO} build -i -ldflags="${LDFLAGS}" -o ${TARGET} ${APPNAME}

${SRCDIR}:
	mkdir -p bin
	mkdir -p `dirname "${SRCDIR}"`
	ln -s `dirname "${ROOTDIR}"`/${APPNAME} ${SRCDIR}

.PHONY: init
init: ${SRCDIR} update

.PHONY: update
update: ${SRCDIR} .EXPORT_ALL_VARIABLES
	${DEP} ensure || true

.PHONY: test
test: ${SRCDIR} build .EXPORT_ALL_VARIABLES
	python ${TESTDATA}/driven.py -v

.PHONY: lint
lint: .EXPORT_ALL_VARIABLES
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -exec golint {} \;

.PHONY: go-env
go-env: .EXPORT_ALL_VARIABLES
	@go env
