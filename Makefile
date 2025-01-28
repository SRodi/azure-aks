.PHONY: init plan apply destroy fmt test

init:
	cd live/aks && tofu init

plan:
	cd live/aks && tofu plan

apply:
	cd live/aks && tofu apply

destroy:
	cd live/aks && tofu destroy

fmt:
	tofu fmt -recursive

test:
	cd test && go test -v ./...