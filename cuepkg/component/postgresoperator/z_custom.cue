package postgresoperator

import (
	kubepkgspec "github.com/octohelm/kubepkgspec/cuepkg/kubepkg"
)

#PostgresOperator: {
	metadata: _

	spec: {
		deploy: {
			spec: serviceName: metadata.name
		}

		config: {
			POSTGRES_DATA_DIR: _ | *"/var/lib/postgresql/data"
		}

		containers: "postgres-operator": {
			ports: tcp: _ | *5432

			resources: {
				requests: {
					cpu:    _ | *"\(config.POSTGRES_CPU)"
					memory: _ | *"\(config.POSTGRES_MEM)"
				}
				limits: {
					cpu:    _ | *"\(requests.cpu)"
					memory: _ | *"\(requests.memory)"
				}
			}

			readinessProbe: {
				httpGet: {
					path: "/api/postgres-operator/v1/status/readiness"
				}
				initialDelaySeconds: 5
				timeoutSeconds:      5
				periodSeconds:       15
			}
		}

		volumes: {
			"pgdata": kubepkgspec.#Volume & {
				mountPath: "/var/lib/postgresql/data"
				type:      "PersistentVolumeClaim"
				spec: {
					accessModes: ["ReadWriteOnce"]
					storageClassName: _ | *"local-path"
					resources: requests: storage: _ | *"500Gi"
				}
			}
		}
	}
}
