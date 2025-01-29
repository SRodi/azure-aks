.PHONY: init plan apply destroy fmt test

init:
	cd live/aks && tofu init

plan:
	cd live/aks && tofu plan

apply:
	cd live/aks && tofu apply --auto-approve

destroy:
	cd live/aks && tofu destroy --auto-approve

fmt:
	tofu fmt -recursive

test:
	cd test && go test -v ./...