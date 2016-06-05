
# lists all available targets
list:
	@echo "List of available targets:"
	@sh -c "$(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | grep -v 'make\[1\]' | grep -v 'Makefile' | sort"
# required for list
no_targets__:

build:
	@go install \
		# && cp "$GOPATH/bin/$GOOS_$GOARCH/measurement-protocol-service" . \
		# && docker build --rm=true -t zollie/measurement-protocol-service . \

push:
	@docker push zollie/measurement-protocol-service \

start:
	@docker run -d --restart=always -p 2020:2020 zollie/measurement-protocol-service "$@" \