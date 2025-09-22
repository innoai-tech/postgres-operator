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

		// Enable debug mode 
		config: POSTGRES_ENABLE_DEBUG: string | *"false"

		// WithoutPg serve without postgres 
		config: POSTGRES_DAEMON_OFF: string | *"false"

		//  
		config: POSTGRES_DAEMON_EXIT_ON_ERROR: string | *"false"

		//  
		config: POSTGRES_AUTO_ARCHIVER_PERIOD: string | *"@never"

		//  
		config: POSTGRES_AUTO_ARCHIVER_CLEAN_PERIOD: string | *"@midnight"

		//  
		config: POSTGRES_AUTO_ARCHIVER_MAX_ARCHIVES_IN_SAME_DAY: string | *"1"

		//  
		config: POSTGRES_AUTO_ARCHIVER_KEEP_UNTIL_DAYS: string | *"7"

		services: "#": ports: containers."postgres-operator".ports

		containers: "postgres-operator": {

			ports: http: _ | *80

			env: POSTGRES_ADDR: _ | *":\(ports."http")"

			readinessProbe: {
				httpGet: {
					path:   "/"
					port:   ports."http"
					scheme: "HTTP"
				}
				initialDelaySeconds: _ | *5
				timeoutSeconds:      _ | *1
				periodSeconds:       _ | *10
				successThreshold:    _ | *1
				failureThreshold:    _ | *3
			}
			livenessProbe: readinessProbe
		}

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
