load('ext://restart_process', 'docker_build_with_restart')
load_dynamic('./ci/tilt/postgres.Tiltfile')
load_dynamic('./ci/tilt/minio.Tiltfile')
k8s_yaml("./ci/fencelive.yaml")

local_resource(
    'regenerate-fencelive',
    'cd fencelive && go generate cmd/main.go',
    deps=[
    './fencelive/graph/'
    ],
    ignore=[
    './fencelive/graph/*.go',
    './fencelive/graph/generated',
    './fencelive/graph/model',
    ],
    resource_deps=['postgresql'],
    labels=["compile"],
)

local_resource(
      'compile fencelive',
      'cd fencelive && bash ./ci/build.sh',
      deps=[
      './fencelive/',
      ],
      ignore=[
      'tilt_modules',
      'Tiltfile',
      'fencelive/graph/schema.graphqls',
      'fencelive/build',
      'fencelive/dep',
      'fencelive/ci/docker-compose.yaml',
      'fencelive/swagger.yaml',
      'fencelive/internal/handlers/swagger.yaml',
      'fencelive/internal/handlers/generated.go',
      'fencelive/**/testdata'
      ],
      labels=["compile"],
  )

docker_build_with_restart('fencelive',
        '.',
        dockerfile='fencelive/ci/Dockerfile',
        entrypoint='/app/start_server',
        only=[
            './fencelive/build',
            './fencelive/dep',
            './fencelive/ci',
            './fencelive/configurations',
            './fencelive/certs',
            './fencelive/migrations',
            './fencelive/cmd/migrationsgo'
        ],
        live_update=[
            sync('./fencelive/build', '/app'),
        ])

docker_build_with_restart('fencelive-migrations',
    '.',
    dockerfile='fencelive/ci/Dockerfile',
    entrypoint='/app/run_migrations',
    only=[
        './fencelive/build',
        './fencelive/ci',
        './fencelive/configurations',
        './fencelive/certs',
        './fencelive/migrations'
    ],
    live_update=[
        sync('./fencelive/build' , '/app'),
        sync('./fencelive/configurations' , '/app/configurations')
    ])

k8s_resource("fencelive", port_forwards=["0.0.0.0:8080:8080"], resource_deps=['minio', 'postgresql'], labels=["BE"])
