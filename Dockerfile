FROM gcr.io/distroless/static:latest
COPY bin/terraform-registry_linux_amd64 /terraform-registry
ENTRYPOINT ["/terraform-registry"]
