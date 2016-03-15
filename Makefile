.PHONY: build-client


build-client:
	@echo Building Mattermost web client

	cd web/react/ && npm install
