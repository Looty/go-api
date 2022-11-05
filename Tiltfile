# version_settings() enforces a minimum Tilt version
# https://docs.tilt.dev/api.html#api.version_settings
version_settings(constraint='>=0.22.2')
allow_k8s_contexts('k3d-lab')

docker_build(
  'go-api',
  context='.',
  dockerfile='./Dockerfile',
  match_in_env_vars=True
)

local_resource(
  'go-test',
  cmd='gotestsum',
  labels=['go']
)

k8s_yaml(
  ['./manifests/deployment.yaml',
  './manifests/configmap.yaml']
)

# k8s_resource allows customization where necessary such as adding port forwards and labels
# https://docs.tilt.dev/api.html#api.k8s_resource
k8s_resource(
  'go-api',
  port_forwards='3000:8080'
)

# config.main_path is the absolute path to the Tiltfile being run
# there are many Tilt-specific built-ins for manipulating paths, environment variables, parsing JSON/YAML, and more!
# https://docs.tilt.dev/api.html#api.config.main_path
tiltfile_path = config.main_path

# print writes messages to the (Tiltfile) log in the Tilt UI
# the Tiltfile language is Starlark, a simplified Python dialect, which includes many useful built-ins
# config.tilt_subcommand makes it possible to only run logic during `tilt up` or `tilt down`
# https://github.com/bazelbuild/starlark/blob/master/spec.md#print
# https://docs.tilt.dev/api.html#api.config.tilt_subcommand
if config.tilt_subcommand == 'up':
    print("""
    \033[32m\033[32mHello world!\033[0m
    """.format(tiltfile=tiltfile_path))