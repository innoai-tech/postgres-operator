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

		// jwt token 签发方 
		config: POSTGRES_SIGN_ISSUER: string | *"octohelm"

		// jwt 签发私钥 
		config: POSTGRES_SIGN_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBMVMvbFZwNzlaaG9lVWI3Y0FPRnVpYk82Rm1HSm55SGNVZTNvdjlseUVpUjc4NmRZClFTQlFRSlI2b3NvRmNDdTcwSUMya0J4QkVFVGZwME1yMnRVNmRyeUVnSHdMbVhSRkRmaDQxNjlVQTc2b2c1YVUKVGVNSmVhYUJrZFEvMGc2WXNnYXZyNnJhT0FWbTJLbUVOYUw1NDRBejJOazFiV1ZaYWFRM21pU2R0L0Z0Y1JGMApzYnR4Z214b0dKbGhmaUNTaW84R3kxeG1LUnV2L2lqanhMZ1F1d0pmVDg0TG5kS092WVlaeUJ1Tm9tVEN4MUkyCmRGLytYdHpKYlBocnFGWEd3cTVtK0hzZHdQWGhoWWRydmIxTlN3NUI3WGcrS0NZeUcrNkNJZjN5VWRWOFNwS3MKTkxEdkF6OCtyZmNBTTU4VUg4ZzgwT2g3TEd4YWlHU3ltd0hZRlFJREFRQUJBb0lCQUF6K2F0bVJQanVyYXQ2cgpwZWd0MVZseVNXelV6QmtWYk1MWG4rQ1ZwZFRDUVlYVDFOYS9XL1RidCsvVkpXaTFXYkMzTDZsdjkyMUE3V3JaCjlzSFRUa2x2YXhvVHRYZElkVzhKRG9DQzhMbDd0UFMwU0Z6SThrcG1ZaTVidy9vUEpySnZJdVV2b2paWTZmQloKd0xPdU45SHNmZnlCSExjS3Z0em9BL1dTdHNRL3g4dStWYU93WEVUajQ2MHBGOVdPeWxFUGdnNnBOc01SclV0QQpCdno3czZBV0FPS3dDZmdveFNCYStmdnA0VEhYL2NmYUtQanBQb3RCRmxhTW9ELzZaQVVlUVN3bnNGYUxPRjRjCm5McktQai9xSmEzWlJZZlRLbkRJN0RxWjVtWDRzRHFrOWx6YkFNQVZTUzQ0ZVVHaGw3THZyQTFQWXlJaG1RbXgKZVZPazA3RUNnWUVBL1hPWUVZNkNvbE1QakhPUHJ3MklQUFlKRzFldzRuR3UvZHlrWGttc3B1WStRbm15M0szVAowaWJoOERyZGhZdXJCdWZnVmZpK1c4WnRSUHBKMTk4NUE1ZHdhdmRxS2pIYUhHZUd0d0J1ZDd3R1lrcmhPc0I0CnlCdnU2M3NrNUp6anRqWEIvc0s3OHFYRDZYNUMxTUxDWGxxS01TK0VVRmwrd1lacTNJK2ppUTBDZ1lFQTExU28KWHhFRldXZ3k0RUY5NDU0bld1Q0N4WmpTM3B0OGNEQmRrcFlEZGFNT3Rvb2FQbmRWZGkzNmtEcGFZamJpeGp1cgo0MHhhUG0wQjY3ZWNhOUhxVzNFVW9RRThQazJLNTRzVkFBUnFXNFphZWliZTliOS9ZYnhjdFNmNko3UHAySlRaCjEzZHg4MnRhd2xQM3dJa0FRZjVTTTYvSWI5QkpJd0M1dHZSQk9Ta0NnWUE1SzV6dXdPV290Zkc0N2N4SVlWWU0KNGphMXF6Y2tMUjZhOXUxQXMzKzFlaFFyTElaekR1Ykw3YWRqWi9QV0R1WTZQWTIxOXRFQnBLVzdQSFh4c3RHSwpoTnQvMVdWbU5TNk0zN0Zja0VWYXp0Zmx6aHcyQTNwN01Rbllwa054S2c2WGFGTGxJNG4vdHZLVk5iemZmenNXCnkzZUdsc0JTMjQwakhDMzRxSkhyOVFLQmdDZ01jY2hFcFNjTXp6R1FYTGFoNnBYblhjc2NjbE8rdlhVc09hc1MKeFkveWhNRUVqSDhEdU54akR0QVdXa3NjQkM5MFY3TE50NWNXdFIyL0o1T1NGakZ0cGJXUVFrRWdNTzZCbXFWSApMUFRhMXljUzViTDFLOTdrcHFWMnl2cnBabHZHeTRGY2tOQUNMbjRvR2RNd1orRGVTVEdkeHZ5czAzMTBIQlpmCjlXV0pBb0dCQU1KdzdRUE9sT3A5QmdMVUY0dzV4Q0pCSHRQbHd4VFowU3JvQ0RRNDNZRmtFNURLdTlBbTFxcXgKNGxjZDZmK2xIY3UrRTI2ZGlVV0ErUi9nd3JBWTFocUpIYzI3SVhzU3VWbFZNaFljNTlKb3M2cjdlUHhvTDFaVgpONmJuc0NCNlZYZkpQelVRN3hoY0dNSkNZaFVVenZHNHFxUjdlbCtXTGV6YnJiaEN0MisrCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

		// 加密传输用私钥 
		config: POSTGRES_ENC_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBckpOZjdWdFZydmF1aU9KUnJqMlpxTDV4RWRuY2NWV1JlKzlzTnkrSXB5VzVMTmpJCjJLSlYwMjh2UXdzeWJWTDJoNFVCLzdnUVZpNFQwMWdQbHZYTktlWGV3Nm9HTzVXK2huZE1XY3h0K2NHU3ducjYKWVhJempFcU94bTA1K1A2ZXE5Vyt1NDBIdzkraHFIRUo0dGp0eEtuMlgwRjhCM2JJUXpIb3hrQW85MEdiQ3NRbwpEOXNwbHVOT053M3d5L25ISS9rN041M2VyLzNxSzdxVzJNckMzNWUyVmF1UGVkWnphTTcrdm9BMTBxQnh2QUFjClhCUVhMQ3R4QkZZUUEzQ2JUZzFmaHJGUHAwYkpsNExRRmV4RWgySmpCWHFVd0d0RVVhN1hMR3Y3bFBQYlU0NlgKNldBMmpmbmR2M3dwekNqQi96R1Z0VHcvbER5cVRtUFN1VjRsV3dJREFRQUJBb0lCQUJuZzJhUXFxNDBha2E2VwpIWUIrM1VGb0dXVi9ZV2FmV2JzWGVvZEYrZnh5bERPTTZJVlB4b1gzcU1NcGRTemVvWGhOOGpCS0JpMGVLTHZxCit4OThpWURzWnZ5NVRNajNtZ1BvOU1xTWRMdGNreWpWcnlFWXNuRXBwd0pMVThPcVR0bmVxZ1RPejZqSUtxSXIKSFR2bW9yVDRkYWp3Rld5N0c2bXozTCtKMUhyTEhQTlNZKzFSOXFrZUlNeTVJSnpWZXlUelJpSVZWSEcyK3RwKwovSjU1TkR0cFpSNjFzK01idjZEM3RoZUZXSUg3QWE0WnZVME5kRGsxRWhsL01FL2x2R0lPU0YrVzQ0Y2Urdmd1Clk0TG9ZcnQwaTQvamhlMnh6VENreEFRZ3E5bjhNdmxTOW5BQUJZdmtnTDVNaHQrZEJua3NLSDlGRmFzTE5FaCsKbkFWcWRmMENnWUVBd1BGV0ExSEJ2ZHd0KzlLWTdWWG5GS2hQSE16TTRhUXhodElMcno2N05jMDhycDhBVDdIUwp6TVVaSXprM2tkdmdxWjZ1SkxrZzdLWHdVYVllS000VklZRjczbDhQQklNK3AwM2s5THBVRGZOaml3a1JLZGxSCkZhU3BZQ3pQRTVqY2hML2ZvTzhRTE9oY2dsSUYyMEdpKzNjYWplbWIydExZNEFybURjVDBNQTBDZ1lFQTVQb0UKMVNGWnBWL2FpUlkrN1NOUDBwUGhLSzgxdTgzb09GU1lRSFVKUkk5dzluNEEwWmp6MjkyZlRCWThxNGtsYUxnUgpoaGZDOUdCbVRlcWQ2dHBrU1VMMVZ5cW1HMUJhd3BkUUFOeHpUREN3S1VMOWcwNHBHVDBaK1h1REE0QjFxZmtwCkRSN21hQWFsc0h2Zm1BSjRMTzByeVpDOVpyRll3MEJOaytDSjZRY0NnWUVBajFkOHNIQ2dBRjdBNXZLWjNORlMKSVQyOXNNYlNlOXlSVXZsZjV1MHpCcENZd1o2dEM0Z3Y5U09GUG03Mnd1MVk1b3RXRTBCYW5wWFZpY05oYXExWQpjNUVRSnEvMnAwS2VYSXQ1U3Z2WEVKbyszUDk2ZWQzUzZNSnhkMXN5NlB0SzhYRGZRbC81WTNPcHJzUWpSN1ptCjBHMjNFN0Yzc2NXdGpCMXN0dFFaR2swQ2dZRUFpNDczTXcvSW04cjRYMlYzcFFGSXZZZjBTOSsrV0dEL2tKVysKMWtwL0E1S0p2eks5UFFLRVh4V002Y3NEMzJrUHErdkVnbjRwRE5sVWdWam1OeVkweVpKT0JucXdFeVcrcTZ1ZAp6MmlOdlhwUFpGYTRQVGQrUlN2QWtSWittN3ZIKzNrcFZCM3BRSzRNZnF5QmN4ek9NbE83eEhhN2VjUE4zZk5yClZSNGQ4REVDZ1lFQXVCTXVWcEFHWm5oR01zTXN4cEtoWC9mRUJaREhpcGRRcXJlS01sZ2t4L0UrWnY2VnBwTFcKazU3cys5K0lrdVVrSFVoMkRmSDZCT3J4WG5TTG4wNnFjWW9VQUg2N0VqaE82dmhXS0xMOHJ3S1pxcjh2MitqSQpvcTdrbmxmd3M1eUFMdjhFdys2eGVLd1ZESzlqeFBIa3dIemxQUmkzdTZncjRHMi9HajVaem5vPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

		// db data-dir
		config: POSTGRES_DATA_DIR: string

		// pg bin version, don't set this unless you know what will be happen 
		config: POSTGRES_PG_VERSION: string | *""

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
		config: POSTGRES_MAX_CONNECTIONS: string | *"0"

		// db which application 
		config: POSTGRES_APPLICATION_TYPE: string | *"MIXED"

		// disk type 
		config: POSTGRES_DISK_TYPE: string | *"SSD"

		// [前端] api 请求调用需要跟随 bash href 
		config: POSTGRES_WEB_UI_ALL_API_PREFIX_WITH_BASE_HREF: string | *"true"

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
