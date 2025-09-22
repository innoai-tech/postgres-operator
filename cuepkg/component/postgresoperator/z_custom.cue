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
					cpu:    _ | *"4"
					memory: _ | *"4Gi"
				}
				limits: {
					cpu:    _ | *"4"
					memory: _ | *"4Gi"
				}
			}

			env: {
				POSTGRES_CPU: "@resource/limits.cpu"
				POSTGRES_MEM: "@resource/limits.memory"
			}
		}

		services: "#": containers["postgres-operator"].ports

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
