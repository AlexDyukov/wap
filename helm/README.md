Для работы helm chart нужен рабочий kubernetes кластер.

Для тестирования и проверки можно использовать k3s на одной ноде, к примеру:
* `curl -sfL https://get.k3s.io | sh -`
* `curl -O https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3`
* `bash ./get-helm-3`
* `export KUBECONFIG=/etc/rancher/k3s/k3s.yaml`
