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
		config: POSTGRES_SIGN_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBNTVmY2s3aUJUYzBEWm1wUFlsZ2VqUFVBU1MzQzVOVkpXbHg3aVdZSnFia1hhSmdoCktnVTRXSGlIZzNvaFVyRDdWK1VBV0Uyd0NXS3YvaEFzczQ4TFFvR1BSYmFoTkNYWlUwQk9TTGs1Nlp1TGRNOEgKMFc1WDZjTHp0U1pLNFRXZmlBSnBHRnpBejhPRXp5Y2RRYjl4aGQrQnlEdmwyMXA0UWQvVmNQZGgxb0dBUXdvLwpWMHQwdE9rTlMvYXdvQWplWDFyVmlQZTRPK25QWEE0ZnNZeXVvVDdlNVRKQWI4bW9xUzh1dG1BY2l4YW9remtKCkkrSmY1OVI2T3EzTThkK2hlTTM1NkJTczErQmJBaFlZb2dLVlpzdHkxV2Q1anQ2TE9Bb3lOanVMK3Qranc3Qk4KNXRkK2R1c0VZZnZYbWwwdENXeVJ0SlJwUnNDTDhmWlpNQ1c3bHdJREFRQUJBb0lCQUNSTG9ZOHBLR0tvRWJOWQpUKzZxdnNiNmtydmxSeGtUOERZUUxuQS9KSjZMMm9aUzZVZXJuOFFGeDc4c1FkS3kyQUZUYTNWclgzNXBFQW55ClFQODNHTjJvYm1yUGozNGJ2RzZXc200bFVWVXlRbTNoRUdtYk1IMzdSclNLTnN3SzJtQzNkLzhRY2t0ZGFoQUoKSWtXdW1qMFA1VWdiVXcvSW92WTMvVHlDR0x4RUxTSTJkTXlDTDVjenc4Sy9QbWF6OGREcjBUZWQva1IwWnBOUApJR2dXQ1JIeE85SmQ2WUhiTnd2dXFYOE9zbVFOSVNReUtHZ2owaktuMjA0UFErWUFWcFRQbjgwd1NpSlJBWWpxCjNOY045alVIMVNTZFVHbXlodWNiZ3FSMEdSdlIvSkFuWGw3ZVlQK2pWVUVSa09aV1d6ZkhNMlZrbWpoMEZGSm0KVG45ZUFEa0NnWUVBOVAyQmpWbThtNGNubmo1UGNFSmw4NFJPZTJ1aEVKZzdoOTdYVkpHbWlMNE5jUjFXbyt6TQpkbFMvN3FvNU1pck1wTjFaMWJKYzI5VjBCMlI1Z1hFKzBmSHdIVFh5MDdOditqM1JIK3hoV012RmREZlJzdHBKCnlYWGVlZWdnVW54WHczVnd2TVo5NEE0WldGNWpnTXBWdHVKZHl1dVdkR1phcFRaVVhqRGtESFVDZ1lFQThnQTYKb1oyc2RNeEVVRUkrZHdiUXltSGZ6dk0xOU90U2V6MWRZanlnUVhSZStNdDZnaXdWaG44UUU5ejliZzhNREJvVwpCUm51VnhJemI1TGVibEMrQWZocThkWWxXc0tCeDI3TEVsV0FOSWU3RjQzTDR2NlZZOU8xRmRrS3oycmM2TUJuCmlHV1l3WWNZRmlIN2lpSjVJT2hES2Q3TkVIRVQydlRreCtLT1Zsc0NnWUVBdXlnWEduYmRZU3RGRnR5Zkx6RnEKVlRoVUJIRmFvQ1RNQmFZMWRnTzI1MnZaTlBxbXY2QWRLcURmNTJIZlEvWHlWRmhOVXJWcHZ2ODRIcFdoUlVQUQpLKzdaOGxiT1pQQUZzWFFjR0hrcWZQMWVvTVFyektoNkNnK2pvQm0yNTR6YU54VzJ4R0FXdFYzUCt2UlFxNGpuCkprbVVRWHJzZTR2ZDM2eTdreUZpZGlrQ2dZRUEzSDdLN1RDZHpubkRrS0ViQ3haaHJOVEt1R1F2ZUczbFpEYkEKWEY3QzRZQ25lK2JpUTdMcEZmZTE3WE5BVWtSUmhNRkw2Ty84a3NjWnFJSllPb2xFNXNTeXBKQ3F6bXhGRlRKawp3dHEyaXFaVkdKdnc5bTFpTG1mYUtnTHM3NW45bC9DZkpNNzFCUGdUUVM5TFlrd1FzNlFNZGh6MEdSUCt0RlRJCkV6dmcxdHNDZ1lBOTBncUxnR1oyVGVMSHZEaU5mTC9XV3JZd3IyRWp6SkJUMVNKQTBDNGhLeGg1UjRDUnBNMFcKamJDZGorNlpLS3dRUVNTR3h5QmZmbTF6SGRoNFdZSjJxUFNWZzdpMDczMk1BRTBVMExPVlJJdHJRd0FnYzhCQwpxV3dnbWNxanVIRm96UHAydWM1TVJ6dWU2ZitSYWM3YlFKd2dETTk4M0R0c3hiMWsvWjVHbWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

		// 加密传输用私钥 
		config: POSTGRES_ENC_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBMlN2QXFNN0p0NWVFUlZ4UzVwZGVrVW1IaGVkWHZiZU9uc2Y2eEQ3cEVyQnBoVXBuCkFZK0tBT0E2QkpPeFV2WjJtSkFsR0F4SE9qWHpkMCtJOWZGMGx0Z3RVUkVMeW0yY1pEOThCNnk1bEN3OVpQWGcKd256TFdxREZBd2N0VXhtYlAxY2RjRzVDd0FoOXYrZElWMktDT3VKSktxdjVCSjZoS3V1enlwaVo4N1JCN0Z4Mwpwd2FSei9Wajh1UnU2ME9lNEhYMldNN3lEc3ZQckpETjNiYTlnZEw0UzlISFNGcG9TYk1qMzYycGZDci9qakMxCjNyWXF3VFN1UVlCUFlPbmJJR2RaRzlPYXV3SG5wdkhIOWxPOVExVTJCZVdlejFSV2d5eFRVQlQ4R3lmVUtYQmgKcFErRUZHSTE0UGVjdGJNaWhndlB1Qlg0RXBQSk1PeENvQWdRd3dJREFRQUJBb0lCQUJSMTVWQmlMeW9QYWVSWgpCdXgzdlc2bHhWRnQwdHZxYzlMczEyanUzbW1tMmtJM0dNNVNWamV2NkhkdTdNbDZ0QUM2cit2OG1DZFpWdWRhClJIYWJlWmNPcGZKWnowa1NtNlhzUDBTLzFyb3BrdDAwdHlqTHRpTURPNXlGN0JURXVGWTI0R3lyenJrbDg2Y0YKWk5EcTFJYzBzNTBFRjlHQ0dvQWZOZEMwSU9rUFdacmNkd041aVJmV0ltRWdiaDl6MkhReFhYSjNJa2NyQ05rWQo0YmFFaUsvakZhbmQzUHpJaTJIWUpnZW5aaVJ3Mi82UHBlQXprUVN5MUNaazVyNklKcmtSOHhQeGZhU2hYRSs0CjB6MVNhWmttNS96RjhXRVdaKzhYSklwb0xtQzY2ZmNLN1FBUExvNk4xU3d1TklFa0R1ekRrSFhwYnhIMnVIRm8KWU5aaGNNa0NnWUVBK0k0QzVZS0dyekFGdG0xbDhGR0NYRy9EdDZ6WnZsaHRmemVtVVh2dnlFVGR5K25wNUpuZQovY2h4cXFkOGI5VkNuemEwb0VJclZ0bFVoRTc0VGhNbUpLNlB1ZXU1QTkzNEFBMzdKbU1ieVFiUlNhVnBLb1FZCng3VUROdTV3V2FJY2ZJdWNLb0dEYzBuT1J0QjFaVWF0V1Q4eGJqVW5CaVpEa2Rmd0ZNT2ZUTWtDZ1lFQTM2MFUKeFpQSFJ5d3dXUWkvRHppTklGdVpweGtQQmFNSTRpK2xaalJlQ0dWMEpTWklzOTIvQXRZdjRPdllWM0svK2JWdAoyR0ZNRVZHcldWay9sb1l0SXhWclBLNXVwY2pJOEIxQ1liZ2lGalg4TEx6SDZCME4zZm1XY0F3ZTgvMFFkRjhWCkh0MDQwLy83NklOY094RmhkQlVFZnUzTmx5dG5tcDdLdE1qWlV5c0NnWUVBalo5dEt4VEtaVHU1cFk2RURRQ0UKaE9MeDQ5QkxhVmU5WEVWN01PYXJZN05KcFl5c3hxS2VHb0NCczdrbkFCbkZraTU3a096akFPTm9jdE1FVEloQQpyWm9CTHZDUFJSTE80a2tWRjNSVk9wLzExRDY1dzQzdENLMnRIVG1UTHA5ZUYrRDhwSU9UUUxlSEgzWmJ6YzhOCnF0S0UrY1N6YlorVDFKL3puZ0V3M1hFQ2dZQXZOdy90YlBaaDFiZ1c4enV4Y05TSmdneDdNMVR5Y2FuTVpSWmEKN3E3eXdzZXpsOU04OUkvL2YwcjRCWkRUVk11bFlHRGhqaGhLaDV6TjdZTDR1VFlKODltQkk2a0RvZENZcnZSMgpRREloMGg5N0toWmdydEZnaS9EdmtmOXVyWHF0dGV4MWFXazNodytiMHk3QzRUWmJGSnl3Vm01UmZMNFA2M2tLCmxHTWJwUUtCZ0RZVkZ1bTBZTUMwcjl2RzRrdktkcGNkaER3SFJrR0FlbFRrRE9vZFFvWmFtSzZ5b3J4VE0wTi8KR2ZZSHMxOUIvSG0yWUNremdqZ1V3aWpMY1RQbjI1a2RHNU1QL1ZwbWhLUmpaTWdKV0VwcnFZSnVKQ1RuN1lQYwpqc1RnenNUSUlaVG1sSnYyUXBUclZLN2ZlS2RhQmhXOGh0cWtHMEs0STltb01IMmJjMExMCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

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
					path:   _ | *"/"
					port:   _ | *ports."http"
					scheme: _ | *"HTTP"
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
