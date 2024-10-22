output "kubeconfig" {
    value = "az aks get-credentials --resource-group ${azurerm_resource_group.aks_rg.name} --name ${azurerm_kubernetes_cluster.aks.name} --admin"
}