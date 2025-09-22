package postgresoperator

import (
	kubepkg "github.com/octohelm/kubepkgspec/cuepkg/kubepkg"
)

#PostgresOperator: kubepkg.#KubePkg & {
	metadata: name: string | *"postgres-operator"
	spec: {
		version: _

		deploy: kind: "StatefulSet"

		// Log level 
		config: POSTGRES_LOG_LEVEL: string | *"info"

		// Log format 
		config: POSTGRES_LOG_FORMAT: string | *"json"

		// When set, will collect traces 
		config: POSTGRES_TRACE_COLLECTOR_ENDPOINT: string | *""

		//  
		config: POSTGRES_METRIC_COLLECTOR_ENDPOINT: string | *""

		//  
		config: POSTGRES_METRIC_COLLECT_INTERVAL_SECONDS: string | *"0"

		// db data-dir
		config: POSTGRES_DATA_DIR: string

		// db name
		config: POSTGRES_NAME: string

		// db user
		config: POSTGRES_USER: string

		// db password
		config: POSTGRES_PASSWORD: string

		// db listen port 
		config: POSTGRES_PORT: string | *"5432"

		// db cpu requests 
		config: POSTGRES_CPU: string | *"2"

		// db mem requests 
		config: POSTGRES_MEM: string | *"4Gi"

		// db max connections 
		config: POSTGRES_MAX_CONNECTIONS: string | *"200"

		//  
		config: POSTGRES_EXIT_ON_ERROR: string | *"false"

		containers: "postgres-operator": {
			image: {
				name: _ | *"ghcr.io/innoai-tech/postgres-operator"
				tag:  _ | *"\(version)"
			}

			args: [
				"serve", "singlenode",
			]
		}
	}
}
