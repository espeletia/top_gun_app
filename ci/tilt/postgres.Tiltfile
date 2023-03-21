load('ext://helm_resource', 'helm_resource')
load('ext://helm_resource', 'helm_repo')

helm_repo('postgresql-helm', 'https://charts.bitnami.com/bitnami', labels=["helm"])

helm_resource('postgresql',
            chart='postgresql-helm/postgresql',
            release_name='postgresql',
            resource_deps=['postgresql-helm'],
            labels=["DB"],
            flags=[
            '--set',  'image.tag=13-debian-10',
            '--set',  'nameOverride=fencelive',
            '--set',  'auth.enablePostgresUser=true',
            '--set',  'auth.postgresPassword=postgres',
            '--set',  'auth.database=fencelive',
            ]
)

k8s_resource('postgresql', port_forwards="5434:5432", labels=["DB"])

# This is for experimenting with DB pool
# k8s_yaml("ci/postgres/pgbouncer/deployment.yaml")
# k8s_resource(
#     workload="pgbouncer",
#     labels=["DB"]
# )
