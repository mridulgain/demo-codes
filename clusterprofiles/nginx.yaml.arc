---
apiVersion: config.projectsveltos.io/v1beta1
kind: ClusterProfile
metadata:
  name: nginx-4-9
spec:
  clusterSelector:
    matchLabels:
      nginx: v4-9
  syncMode: ContinuousWithDriftDetection
  helmCharts:
  - repositoryURL: https://kubernetes.github.io/ingress-nginx
    repositoryName: ingress-nginx
    chartName: ingress-nginx/ingress-nginx
    chartVersion: 4.9.0
    releaseName: ingress-nginx
    releaseNamespace: ingress-nginx
---
apiVersion: config.projectsveltos.io/v1beta1
kind: ClusterProfile
metadata:
  name: nginx-4-10
spec:
  clusterSelector:
    matchLabels:
      nginx: v4-10
  syncMode: ContinuousWithDriftDetection
  helmCharts:
  - repositoryURL: https://kubernetes.github.io/ingress-nginx
    repositoryName: ingress-nginx
    chartName: ingress-nginx/ingress-nginx
    chartVersion: 4.10.0
    releaseName: ingress-nginx
    releaseNamespace: ingress-nginx
