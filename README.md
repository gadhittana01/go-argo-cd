# Go ArgoCD Application

A simple Go HTTP server with health check endpoint, containerized and ready for deployment with ArgoCD.

## Application Structure

```
├── main.go                 # Go application source
├── go.mod                  # Go module definition
├── Dockerfile              # Multi-stage Docker build
├── .github/workflows/      # CI/CD Pipeline
│   └── cd.yml              # Simple deployment workflow
├── k8s/                    # Kubernetes manifests
│   ├── configmap.yaml      # Application configuration
│   ├── deployment.yaml     # Application deployment
│   ├── service.yaml        # Service to expose the app
│   └── kustomization.yaml  # Kustomize configuration
└── README.md              # This file
```

## Quick Start

### 1. Build and Run Locally

```bash
# Run locally
go run main.go

# Test health endpoint
curl http://localhost:8080/health
```

### 2. Build Docker Image

```bash
# Build the Docker image
docker build -t go-app:latest .

# Run the container
docker run -p 8080:8080 go-app:latest

# Test health endpoint
curl http://localhost:8080/health
```

### 3. Deploy with kubectl

```bash
# Apply all Kubernetes manifests (with kustomize)
kubectl apply -k k8s/

# OR apply individual files
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml

# Check deployment status
kubectl get pods -l app=go-app
kubectl get svc go-app-service

# Test the service (using port-forward)
kubectl port-forward svc/go-app-service 8080:80
curl http://localhost:8080/health
```

### 4. Deploy with ArgoCD

Since you have ArgoCD configured elsewhere, create an ArgoCD Application that points to this repository:

**Via ArgoCD CLI:**
```bash
argocd app create go-app \
  --repo https://github.com/your-username/go-argo-cd.git \
  --path k8s \
  --dest-server https://kubernetes.default.svc \
  --dest-namespace default \
  --sync-policy automated \
  --auto-prune \
  --self-heal
```

**Via ArgoCD UI:**
1. Click "New App"
2. Set Application Name: `go-app`
3. Set Repository URL: `https://github.com/your-username/go-argo-cd.git`
4. Set Path: `k8s`
5. Set Cluster URL: `https://kubernetes.default.svc`
6. Set Namespace: `default`
7. Enable "Auto-Sync" and "Auto-Prune"

## Simple Deployment Pipeline

The repository includes a minimal GitHub Actions workflow (`.github/workflows/cd.yml`) that automatically deploys your application to Docker Hub.

### 🚀 **How it works:**

```bash
# 1. Make changes to your code
vim main.go

# 2. Commit and push to main branch
git add .
git commit -m "Update feature"
git push origin main

# 3. Automatic deployment happens:
# - Builds Docker image
# - Pushes to Docker Hub
# - Updates k8s/deployment.yaml
# - ArgoCD syncs the changes
```

### 📋 **What the pipeline does:**

1. **Build**: Creates Docker image from your code
2. **Push**: Uploads image to Docker Hub (`docker.io/your-username/go-argo-cd`)
3. **Update**: Modifies `k8s/deployment.yaml` with new image tag
4. **Commit**: Pushes the update back to Git
5. **Deploy**: ArgoCD detects the change and deploys

### ⚙️ **Setup:**

**You need to configure Docker Hub credentials:**

1. **Create Docker Hub Access Token:**
   - Go to [Docker Hub](https://hub.docker.com/)
   - Login → Account Settings → Security → New Access Token
   - Copy the token

2. **Add GitHub Secrets:**
   - Go to your GitHub repository
   - Settings → Secrets and variables → Actions
   - Add these secrets:
     ```
     DOCKER_USERNAME = your-docker-hub-username
     DOCKER_PASSWORD = your-docker-hub-access-token
     ```

3. **Create Docker Hub Repository (Optional):**
   - The pipeline will create `your-username/go-argo-cd` automatically
   - Or you can pre-create it on Docker Hub

### 🏷️ **Image tags:**
- `latest` - Always points to the most recent build
- `<commit-sha>` - Specific commit for traceability

## Configuration

The application uses a ConfigMap for configuration with the following environment variables:

- `PORT`: Application port (default: 8080)
- `LOG_LEVEL`: Logging level (default: info)
- `APP_ENV`: Application environment (default: production)

## Health Check

The application provides a health check endpoint at `/health` that returns:

```json
{"status": "okay"}
```

## Resource Requirements

- **Requests**: 50m CPU, 64Mi memory
- **Limits**: 100m CPU, 128Mi memory
- **Replicas**: 3 (for high availability)

## Monitoring

The deployment includes:

- **Liveness probe**: Checks if the application is running
- **Readiness probe**: Checks if the application is ready to serve traffic
- Both probes use the `/health` endpoint

## Notes

- 🚀 **One-push deployment**: Just push to main branch, everything else is automatic
- 🐳 **Docker Hub**: Container images stored at `docker.io/your-username/go-argo-cd`
- 🎯 **GitOps ready**: ArgoCD automatically detects and syncs changes
- 🔑 **Requires setup**: Need Docker Hub username and access token in GitHub secrets
- 🏷️ **Commit-based tagging**: Each deployment is traceable to a specific commit 