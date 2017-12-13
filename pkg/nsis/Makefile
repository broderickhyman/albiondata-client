TOP_SRCDIR = ../..
PACKAGE = albiondata-client
PACKAGE_NAME = Albion Data Client
PACKAGE_BUGREPORT = https://github.com/broderickhyman/albiondata-client/issues
INSTALLER_COMPRESSION = lzma
PACKAGE_EXE = albiondata-client.exe

all: nsis

nsis:
	set -e; if [ "x${CIRCLE_TAG}" = "x" ]; then export VERSION="0.0.1"; else export VERSION="${CIRCLE_TAG}"; fi; \
	makensis -DPACKAGE_VERSION="$$VERSION" -DPACKAGE="$(PACKAGE)" -DPACKAGE_NAME="$(PACKAGE_NAME)" -DPACKAGE_EXE="$(PACKAGE_EXE)" -DPACKAGE_BUGREPORT="$(PACKAGE_BUGREPORT)" -DTOP_SRCDIR="$(TOP_SRCDIR)" -DOUTFILE=../../albiondata-client-amd64-installer.exe -X"SetCompressor /FINAL /SOLID $(INSTALLER_COMPRESSION)" albiondata-client.nsi $<
