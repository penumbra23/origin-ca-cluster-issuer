# Origin CA Issuer

Origin CA Issuer is a [cert-manager](https://github.com/cert-manager/cert-manager) CertificateRequest controller for Cloudflare's [Origin CA](https://developers.cloudflare.com/ssl/origin-configuration/origin-ca) feature.

## Getting Started
We assume you have a Kubernetes cluster (1.16 or newer) with cert-manager (1.0 or newer) installed. We also assume you have permissions to create Custom Resource Definitions.

### Installing Origin CA Issuer
First, we need to install the Custom Resource Definitions for the Origin CA Issuer.

```sh
kubectl apply -f deploy/crds
```

Then install the RBAC rules, which will allow the Origin CA Issuer to operate with OriginClusterIssuer and CertificateRequest resources

```sh
kubectl apply -f deploy/rbac
```

Then install the controller, which will process Certificate Requests created by cert-manager.

```sh
kubectl apply -f deploy/manifests
```

```
$ kubectl get -n origin-ca-issuer pod
NAME                                READY   STATUS      RESTARTS    AGE
pod/origin-ca-issuer-1234568-abcdw  1/1     Running     0           1m
```

### Adding an OriginClusterIssuer
With running the controller out of the way, we can now setup an issuer that's connected to our Cloudflare account via the Cloudflare API.

We need to fetch our API service key for Origin CA. This key can be found by navigating to the [API Tokens](https://dash.cloudflare.com/profile/api-tokens) section of the Cloudflare Dashboard and viewing the "Origin CA Key" API key. This key will begin with "v1.0-" and is different than your normal API key. It is not currently possible to use an API Token with the Origin CA API at this time.

Once you've copied your Origin CA Key, you can use this to create the Secret used by the OriginClusterIssuer.

```sh
kubectl create secret generic \
    --dry-run \
    -n default service-key \
    --from-literal key=v1.0-FFFFFFF-FFFFFFFF -oyaml
```

Then create an OriginClusterIssuer referencing the secret created above.

```yaml
apiVersion: cert-manager.k8s.cloudflare.com/v1
kind: OriginClusterIssuer
metadata:
  name: prod-issuer
  namespace: default
spec:
  requestType: OriginECC
  auth:
    serviceKeyRef:
      name: service-key
      key: key
      namespace: default
```

**NOTE**: The ServiceKey secret doesn't have to be in the same namespace as the OriginClusterIssuer, because it's being referenced under `serviceKeyRef`.

```
$ kubectl apply -f service-key.yaml -f issuer.yaml
originclusterissuer.cert-manager.k8s.cloudflare.com/prod-issuer created
secret/service-key created
```

The status conditions of the OriginClusterIssuer resource will be updated once the Origin CA Issuer is ready.

```
$ kubectl get originclusterissuer.cert-manager.k8s.cloudflare.com prod-issuer -o json | jq .status.conditions
[
  {
    "lastTransitionTime": "2020-10-07T00:05:00Z",
    "message": "OriginClusterIssuer verified an ready to sign certificates",
    "reason": "Verified",
    "status": "True",
    "type": "Ready"
  }
]
```

### Creating our first certificate

We can create a cert-manager managed certificate, which will be automatically rotated by cert-manager before expiration.

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: example-com
  namespace: default
spec:
  # The secret name where cert-manager should store the signed certificate
  secretName: example-com-tls
  dnsNames:
    - example.com
  # Duration of the certificate
  duration: 168h
  # Renew a day before the certificate expiration
  renewBefore: 24h
  # Reference the Origin CA Issuer you created above.
  issuerRef:
    group: cert-manager.k8s.cloudflare.com
    kind: OriginClusterIssuer
    name: prod-issuer
```

Note that the Origin CA API has stricter limitations than the Certificate object. For example, DNS SANs must be used, IP addresses are not allowed, and further restrictions on wildcards. See the Origin CA documentation for further details.

## Ingress Certificate
You can use cert-manager's support for [Securing Ingress Resources](https://cert-manager.io/docs/usage/ingress/) along with the Origin CA Issuer to automatically create and renew certificates for Ingress resources, without needing to create a Certificate resource manually.
As this is a cluster-wide resource, any ingress from any namespace can use it, but there's a bit more to it.

**IMPORTANT**: To have the cert-manager reference the correct issuer, please set all of the annotation mentioned in the below example (that is, `issuer` is the name of your issuer, `issuer-kind: OriginClusterIssuer` and `issuer-group: cert-manager.k8s.cloudflare.com`). 

```yaml
apiVersion: networking/v1
kind: Ingress
metadata:
  annotations:
    # Reference the Origin CA Issuer you created above.
    # NOTE: set all three annotations
    cert-manager.io/issuer: prod-issuer
    cert-manager.io/issuer-kind: OriginClusterIssuer
    cert-manager.io/issuer-group: cert-manager.k8s.cloudflare.com
  name: example
  namespace: default
spec:
  rules:
    - host: example.com
      http:
        paths:
         - pathType: Prefix
           path: /
           backend:
              service:
                name: examplesvc
                port:
                  number: 80
  tls:
    # specifying a host in the TLS section will tell cert-manager what
    # DNS SANs should be on the created certificate.
    - hosts:
        - example.com
      # cert-manager will create this secret
      secretName: example-tls
```

You may need additional annotations or `spec` fields for your specific Ingress controller.

## Disable Approval Check
The Origin Issuer will wait for CertificateRequests to have an [approved condition set](https://cert-manager.io/docs/concepts/certificaterequest/#approval) before signing. If using an older version of cert-manager (pre-v1.3), you can disable this check by supplying the command line flag `--disable-approved-check` to the Issuer Deployment.
