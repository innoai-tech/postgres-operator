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
		config: POSTGRES_SIGN_ISSUER: string | *"algo.industai.com"

		// jwt 签发私钥 (base64 std encoding 格式) 
		config: POSTGRES_SIGN_PRIVATE_KEY: string | *"LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBd1pMOEM5M1hkam40K0ZtN096cVFGblpMZXpzYVZva2h0bWk3bHRsOHZnM1JvYkFMCmlFb0dzekx2SUNaQ1JxYzdaWGxCR1h2TTFyMURoWVZZb3ZvcGFjQzY1ZUNuU2NOTkcraWpveU4xbndYK3l3am4KV3RRREZ2NXVwYnhCWmdiWUxVaTZVV2syUUFGODY3Z1l5NW0wVkRxMU5Zb256QzVtdHR4c0R2d2RFdUo2Vzg1YgpnMysraENZTGxhK1VDT0VQMjVZWXdIZXZ1R3Fma2lVTm94MFl4OVllT25YUTZjUkIzb3dYallNNTNzVTJqaVoyClpzRWVyQWk5OUowSitQL3NYN0pTRVFMR1R2N2hSYXFSVGJ4d0p5WmgwMThOQTM1V1lZYzYvTzFTS3FWUVhZUzUKSmZxMGw4SFZicW9sYkEwSUt5SGorUDE1Q3dQcnpXT1ZiMTd1WlFJREFRQUJBb0lCQUZiZmhFWS9GVWQxMVFxQQpHc3NHQ1V4TFlxeTNaYWFZZkl6RWpBdXpKNHlUM3hUVlVSZWxRVVNEQTFaR1ZWQW9GVHdCSXZvdzJVazJQMnRZCjRPS2pxcHBQYUpGcU5TbWhnS1daem5hVW4yRHk2OXBxOEltT3hLUTVJNmMxeVpQa1ZuaGNQMkh4K2VyWHdCeDUKem9ieFFFRGM2QjVURm5GSUVjalhPNWV1QmM1SEZnN2F2U3dOQ0t1bXlFdnhrTWF1U2RjOTRVTHhXWlFIOWdyego0SEF0ekVWRVduT1dEREEvbUNraGdQdjJyekpveFc2OHNPcnM4ODdCVlY5YUhUSzY0TDUzbGp4MHZ5MUhaWlJlClVwWVM2Z0pyMXNMVlZIZFF1Y1FSdStyUy9uUm1mbkptZ05La0lZdEVNUDRIUnNqOUoxbFFUMUFWcHRVR2JkeSsKbVcxTTk3RUNnWUVBMmNGdGVQaHNoTGpYTUxTZ2ZRWHhwNlYyU2Nid0tRMmlGZzJwSEo3TmZnbmplVEtlVUM1VQp4anlEQUtUSDQrNW45K1hlYjJUeEVmb0l0V1VNMkVETmY5Tm1RNmpiZlplbUxTTkYrSmNIZHFRYTdUeCtxdnpsCkZTS0lLeEtkZFJsZ0FKcytpNGEwWlpzamF1K0ZlZGI5UzJlWURtYlNVaDJHNWFabXRUQXAySjhDZ1lFQTQ1SlUKUHlENDd3RUlrQUJwWlVacWk3WEx3a3R0bFhPb0d3NkRBNjEyQ1grNkNXZW1lL29zNEh6ZlpZdUljVUorcTFKTgppQUN1ckI1NTJqLzZuditzU0p0RDFXbkRpa3FZTjZ1UldBWExrYi9hQ3FnY0IwRnRSZkIwMkFmOEJDTFBRdkJjCnBCLzY1V1ZkUFJ6SlYyNHNQNTFaT2Q4alFpRmFVb3VRLzNvNzVuc0NnWUJ5czJHUHZBT2xjWnZnU3ZUU2hrUUIKVTdYWUxOTXFQTVVwS3E3dXBYT0d0WHk5eFgwQXJUYjRhTDEwcEZlcVoxVWFqbG01Z1lrK045OUlkVjdydGs3SQp5emp0NE04ays3R0x5eTI3UVFxc1lzclFOZkplN1BGYVRhTStWUUxkd0swQ0JNUlFTSGRUb2dGby9adm9UdWVEClJhRWh6T1Z5WFoxRjhIM2ZhT3hSN3dLQmdEaVBCYUVxOExOMTdrcmJ6MVg4U2o2dTBCVlo2Y0piSEV5ZExyS2IKU2RyU2c4b1NtSzMzWDIvcm1Sb0RzOFZ5WVVqekM1SGhtWWZ0aXh0VEMxYTQ5SGlYSGlUTVJHejZYUnA0NllhVwo2aUV2TFdHNHBqNm5aQ0Vkd2V4dkQ1TjNrMHR2c1phWTNDSm5MdVVoWW1qNFdNYjJ0RlpOdnRXUXRJSVBiZG9RCjV4SkJBb0dCQUsxWE90VUtjZ1FCVW0zajd2d29RV3FFZVBVdjVVRWdiOFFUcVBwNHU2Sk9OYWdPWDFyUkYzbUoKeVVrTnhLZDlML0k2aVNlRE03ZzhYcmxkNVRwU3I3S05HOFkvM2dIMXFQQkJIeE1td3pJeGUrQXFzVWNPM2ZabgpRNDRxN0xweTFQOC9FNXVvYTBoanVhM1RDNVlHbENtWGNqWS9lYzgwdlQwODhub3R2ajVRCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

		// 加密传输用私钥 
		config: POSTGRES_ENC_PRIVATE_KEY: string | *"LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBaytOU1JxaXJQUmpSNFFqclNRUFdJSE84b3lIOCsrZjRFcFE3cEp0S2V2MC9yVlphCjM0dWdrN05KY3FjWkRrQ1ljYmFzWTRQYTIxQk9YRVZMb1JPbThlRnExZlNiKzhtZTZaQmdCeTQrc0lGclNNLy8KYjJtNE1nYzN4QWVBWXdJd1k5M2ZRWGFvU25lU051aWYxYmxVeUdQQ3oxZnV4WXVUTEtSQ0xNQytsK0twUGZuSwp2UUQxemlhNDlDM25SQlNUN3VpTkI1YzAxWkJOSER4aW1FVFNVdHlHVU5pbEZXcW9QemhJZTQ4TzU3OElsNFZGCk5LVjQ4a3JQN0ZsVmlnTm14NnJ4M1ZkZ1ZidjRsZk1UVFRmdDV5YWk1YkZMOGU2RisyRzZDVmZkVUgrdVZCM1EKZllNelUrTEtmajR0UDJGMXZtZkIwWlpMZm5pY0FHZjMxY1hPQndJREFRQUJBb0lCQUJkM1I1UGNxYWM5Sy9oWQpNUFZaL1gwU3dFam1XTDBqOXJJRUFWN0NJVlJZb09VQlFkUWMxTEw3RE5sN2VEZFJ1cDY0akt3UnZRVjF2YWxZCmVZS3FtcC9nMDlIMVNWYnZqQjZsZ1FkQ3FFSGJsMGxyYjYyZkNkSXVZUVBpNk5WTkFKNm9Fd2ZSS3l6OXdoNmkKVk5KVjRRVVhCMDlkTlk3MDhicHdxZXR0L2hBZGNsckgrL29tL2crRnVVYWNaVGlCZnd3UlUyM0M4OHZpbDAvbwo3ekFxYUFCendyeC8rSkV4elRVck9KbVRlWjlCaE5idThXOG1pN0tJRG5xTHpQUEZqc0JUMkQ3TkloSE1saytzCjRHbTZJcDdYV29UcWRYbWlGNnROLzhRZ2dwVFVxSmo1R1k2WjU5RjZETTI1UWZXZ3ozQ1hYaHgrVVhjLzFTWXoKVHdhb3BUa0NnWUVBd0lQVU94RjJoSXJQQ0ZtdExVTDBIMnlTbVVQWE01MnJtckIwa0lMSElYaVBKZjVDbU5nKwp6dXVHM3BBcUZhV3VSdjlaOGY4c3RhTTBPdjV5bDV0UWVCRTdZTjA2MEJveXZQVW10L29zMUQ5b1AvVk0ydHBLCmlKc1Y2YWxhaml1RENlN3lSZ3pldFJUYSt0OEIrTytDTDBpM2hmZDJhdmtNNWhZVEZPdUplMnNDZ1lFQXhLZ1IKcHNJWFpnQUVyd3RiMUNhZ1o1SW40Vkxra3JCS0N4cy84KzVQTzZhdlJzTlJhWEFNRHJrN0RxcWZqeHpwN29LNwpodHJIeWlhWkd0WEU0VmZneUUvUWVPN1lZQnRGU1FPQ3hhaUFzTThQRXhqWEFNNzZ5TlVFWHhzSFF3eDhmd2pXCkdxaDlKTWFnMmtsOWN2M1oreEg5d3kzalhodGlmbWJKcEk4YjJ0VUNnWUFlbHBPcEV2dyswUU5XTTFGMXlKYWIKUzVmN2JERU1UWGdQcXd2S1RrMHZmMFZYWncvVDAwQWZob0syYURlWG11eVc4VW1zVHJ3ZTNDQ1hZd2g0R0VCdgo0MlVJM2YzVFJPWmM2YUxPUnB6SzJJeEs1VUhoNEI5SmwwS2pEcnFKcmxZeXhObVAwY210QWZSTk9oUEpKdDBMCmdFVlFydUlNMnlkMmczbGlzSTMyb3dLQmdIaVoxWkZkMERtbVl6anlEMU93aUloYnNvZFZiWVdrQlJTQWxwekMKbGVhd0Z5ZWZXb3I5d3ZjNGswWXdUdi9XTElRdnVrOCtWbTNiYytOb0c0QlNnekIwK2hWZjdHUXI5VGFNcDZNTgo3allBRlcrUnVURHV2Zzl4eTJlRHpOVktrQUxiNldtWjBIMUFMcDhQbUpadnBVRjZ6QlBwVFZtR0U4WU94VFJjCjZ5a3RBb0dBUGlnTnZadWlTZy94TG5QQlFkZTQ4MmtpSGJVbTdiSDBVcGVBeDkrS2R2L01UakhETjNYcXpWUDkKcTE0ZmdlTzJWYUpjNUo3eW5iVzdjaVh2eS83OXN3SXUrU2VWMlExbzE4OTZvN2xjMll3TG1QckpEa1FQemgyagpvclRtTG5DRXRMVy9hbmhsUS91bzMxWUhMcG9UQXhaSDNCY1JuNTN1cDdZUk9hSGliTkk9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

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
