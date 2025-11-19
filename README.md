# End-to-End Observability for Kubernetes Using Prometheus, Loki & Grafana

This project provides a complete **observability setup** for a Kubernetes cluster using:

- **Prometheus** → Metrics collection  
- **Node Exporter** → Node-level metrics  
- **Application Instrumentation (Go App)** → Custom metrics  
- **Loki** → Log collection  
- **Grafana** → Dashboards and alerting  

The goal is to monitor a Kubernetes app end-to-end using real metrics, logs, dashboards, and alerts.

---

## 🚀 Overview

This setup includes:

### 1. **App Instrumentation**
A Go application exposing `/metrics` using Prometheus Go client library.

### 2. **Prometheus Stack**
Installed using `kube-prometheus-stack`, includes:
- Prometheus
- Alertmanager
- Node Exporter
- Grafana dashboards (optional)

### 3. **Loki Stack**
Used for log aggregation:
- Loki
- Promtail

### 4. **Grafana**
Used for:
- Dashboards  
- Querying logs and metrics  
- Alert management  

---

