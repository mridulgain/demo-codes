apiVersion: config.projectsveltos.io/v1beta1
kind: ClusterProfile
metadata:
  name: kyverno
spec:
  clusterSelector:
    matchExpressions:
      - key: env
        operator: In
        values:
        - dev
        - qa
  syncMode: ContinuousWithDriftDetection
  helmCharts:
  - repositoryURL:    https://kyverno.github.io/kyverno/
    repositoryName:   kyverno
    chartName:        kyverno/kyverno
    chartVersion:     v3.3.3
    releaseName:      kyverno-latest
    releaseNamespace: kyverno
    helmChartAction:  Install
  patches:
  - target:
      group: apps
      version: v1
      kind: Deployment
      name: "kyverno-admission-controller"
    patch: |
      - op: add
        path: /metadata/annotations/projectsveltos.io~1driftDetectionIgnore
        value: "ok"
