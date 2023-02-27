load('ext://restart_process', 'docker_build_with_restart')
k8s_yaml("busybox.yaml")


local_resource(
      'compile coolProject',
      'cd coolProject && ./build.sh',
      deps=[
      './coolProject/',
      ],
      ignore=[
      'tilt_modules',
      'Tiltfile',
      'coolProject/graph/schema.graphqls',
      'coolProject/build',
      'coolProject/dep',
      'coolProject/ci/docker-compose.yaml',
      'coolProject/swagger.yaml',
      'coolProject/internal/handlers/swagger.yaml',
      'coolProject/internal/handlers/generated.go',
      'coolProject/**/testdata'
      ],
  )

docker_build_with_restart('coolproject',
        '.',
        dockerfile='coolProject/Dockerfile',
        entrypoint='/app/start_server',
        only=[
            './coolProject/build'
        ],
        live_update=[
            sync('./coolProject/build', '/app'),
        ])

k8s_resource("coolproject", port_forwards=["0.0.0.0:8098:8080"])
