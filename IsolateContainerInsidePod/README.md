Simple setup to validate that a container can expose a port to a sidecar only (not accessible outside of the pod).


# Deploy
```
ko apply -f config/
```

# Test
From inside the cluster:
```
curl <POD_IP>:8080
```
should fail.

```
curl <POD_IP>:9090
```
should succeed.
