# Helm Template

If you know `kubectl`, you know about the `dry-run` flag. 

Dry-run works differently on Helm. Dry-run is used for inspecting and understanding how Helm installes or updates software. It combines YAML content with non-YAML content. 

To create a YAML file of an app, run: 
- `helm template <name> <image>`

Here is an example:
- `helm template mysite bitnami/drupal`

When using a template command, Helm will not use any custom resources. This makes the YAML almost independant of the Kubernetes version. If you install it directly, Helm uses custom Helm resources.
