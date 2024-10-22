# Azure Kubernetes Service (AKS) Deployment with OpenTofu

This project sets up an Azure Kubernetes Service (AKS) cluster using OpenTofu. OpenTofu is an open-source infrastructure as code tool that allows you to define and provision infrastructure using a high-level configuration language. For more information on OpenTofu and installation instructions, visit the [official documentation](https://opentofu.org/docs).

## Prerequisites

- OpenTofu installed on your machine. Follow the [installation guide](https://opentofu.org/docs/installation) to set it up.
- Azure CLI installed and authenticated.

## Project Structure

The project directory is organized as follows:

```
├── main.tf
├── variables.tf
├── outputs.tf
├── README.md
└── .terraform.lock.hcl
```

- `main.tf`: Contains the main configuration for the AKS cluster.
- `variables.tf`: Defines the input variables for the project.
- `outputs.tf`: Defines the outputs of the project.
- `README.md`: Project documentation.
- `.terraform.lock.hcl`: Automatically managed by OpenTofu, should not be edited manually.

## Instructions

### Create `terraform.tfvars`

Create a `terraform.tfvars` file in the root directory of the project. This file should contain the values for the variables defined in `variables.tf`. Below is an example of what this file might look like:

```hcl
subscription_id     = "your-subscription-id"
tenant_id           = "your-tenant-id"
location            = "uksouth"
resource_group_name = "your-resource-group-name"
prefix              = "your-prefix"
labels = {
  environment = "test"
  team        = "devops"
}
```

### Initialize the Project
Run the following command to initialize the project. This will download the necessary providers and set up the backend.

```sh
tofu init
```

### Plan the Deployment
Run the following command to create an execution plan. This will show you what changes will be made without actually applying them.

```sh
tofu plan
```

### Apply the Deployment
Run the following command to apply the changes and create the resources.

```sh
tofu apply
```

### Destroy the Deployment
If you need to destroy the resources created by this project, run the following command.

```sh
tofu destroy
```

## Notes

For more detailed information on each command, refer to the [OpenTofu CLI documentation](https://opentofu.org/docs/cli).