load('ext://helm_resource', 'helm_resource')
load('ext://helm_resource', 'helm_repo')

helm_repo('minio-helm', 'https://charts.bitnami.com/bitnami', labels=["helm"])

helm_resource('minio',
            chart='minio-helm/minio',
            release_name='minio',
            resource_deps=['minio-helm'],
            labels=["S3"],
            flags=[
            '--set', 'statefulset.replicaCount=1',
            '--set', 'auth.rootUser=minio123',
            '--set', 'auth.rootPassword=minio123',
            '--set', 'defaultBuckets=fencelive:public test:public nats:public logs:public'
            ]
)

k8s_resource('minio', port_forwards=["9000:9000", "9001:9001"], labels=["S3"])
