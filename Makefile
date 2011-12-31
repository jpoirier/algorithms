
PACKAGES := comp_geometry graph network_flow path_finding searching sorting

all: clean install test
.PHONY : all

install:
	$(foreach pkg,$(PACKAGES),$(MAKE) -C $(pkg) install ;)
	
build:
	$(foreach pkg,$(PACKAGES),$(MAKE) -C $(pkg) ;)

clean:
	$(foreach pkg,$(PACKAGES),$(MAKE) -C $(pkg) clean ;)

test:
	$(foreach pkg,$(PACKAGES),$(MAKE) -C $(pkg) test ;)
