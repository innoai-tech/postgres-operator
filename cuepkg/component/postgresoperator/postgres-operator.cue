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
		config: POSTGRES_SIGN_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBc0MvTzE5S3R4QWhmQ0lpNGRVcnRqaHVPRytaOVVvcFBQV3c4TkdQd0h0LzA4QVJsCkxwK0o0dWZVaVUxeG9FWGFWRE5zM2Z2ZStQQUFhR0txZUc4aGRHS2MwbmdDY1RPdmxRd2Vpc0ZmdWhaaUpVamoKZXlPRlh4UzkxaFF5WmpXVnd4RDhGSU05UU9uOWVDLzN0Q25ZRWE3d3loa0tPZHNSc1FJSitaOC9GUWJaNHlpOQpGaG1UK0praENlSXBab2xld3dGME9Xbk9PQlZpQ3IrZFRNVHpBUG8zZ0IrdHlJK0hYK05MN1Z3WUM1M28zR3pnCjdvaHFML1NQTFVkUVJ2ZVlvMlZudXIzb1Bka2tvS3ZlM29HWS8rRlVpak9uZ1NkZWdQWTFzYjJVZnFOdXlUT0sKMmVlYTYwQ2dxaEp2REZadHU0WWdleU5vNlFrWUFUVG1oa1hMM1FJREFRQUJBb0lCQUJJQWZJVzBUdEZzR291MApib1o3aHF6QnQ0STF5WUhNeEg1dDUvUGhiemcwWVdKQnpMWFlFcjVtdWlKaVVaVVM5aFVxNFBtSWdmYVZEVEdkClBUWkxncE5SL1J5TEt4SzJlWDhFNnp6NXQ2WjdvWUVXUXJvQy9kNWliK2duRGFSQVh1STh3c2oweE9ZMFVsUDgKc0xaTm5UaG5VODZPZ3VzNDdVU1ZCUUJtcTBzaVNkdWp1NDkrNUJSSU5hM0xXWTJHUEJzNjlreTV3aDhqeG96SwoyMG03MUxSKzdpSENuN0hsYk5jTVNlYUtqN3VrWjB5SkwrdDExU2JUWjEvNndNWFNQVW5DS2dFN1pzMDc1cmlvCkM0N1k0M0VTWW13UUpPQjhWYXZkNzh4R3UwVDJ4bnV1eUJYVGhRV05GQjhhWUlqTk9iU1RoNGVVNzU0dVIrRmMKU2JDU3N3RUNnWUVBNEVpN2pxQjd1TElwdkUreXNsUDl3V3kyaHIzNFpxcTNSMjFMbWVobDZVL2FuU3Zqcm5ScQovWkZtUWQrRGhnL29QNm1mOHF1VlpCOEtOblpQcS85Wi9YdkZsaVUzeUVVaHZUZ2NsNGxCNDQ0M0NpQng1bFhlCjA4Y1RUYi96N2F6T2JSSlpBTmVrSUVHK3NIS0hob29KQVdMUmgxeVNZanVUMWdldS9VL0N1dDBDZ1lFQXlSbnAKWk5CMExFdzV2Umk2WEtHajRGWk10RWhNVVQwNHVmNnF2eG1DL0gxcjJURUxDUnIyOHZxMDI0akVmNVl2Qm9MeApBTE80SGJJb3pEYUw5TThrVEFjUVBBcGNCYjJITUFlUmt1WS93NzVoakZDUGZsWmp1aXY5Z2cvZnRYUlVGaDhQCk9ybDZnZ3dseVh1UVk1Z21sTGM5ZWxTRk9YRnBYYUVEYmNxanhRRUNnWUFLWGJJMWZGdGJoUGlDMkpna3Y5Y0oKbXBHeEZwU2xnaHhvYzdlN3pFN3hncHhUQ1ZWRG1lUGNCbDFZakJFVElDY1cwaXN2VnhqWGdNdkRDcUxTQzBKSQpnZkQyNkk5MGRTV0REbFhiOXg4UmVtQTIycHNKRDB6Mk1zeEVtcXVXZjVjbTJXTTlzN25GTitTdFdRM0VmUnEvCmNxYkdmOVBRTUhxN3VLMHd6Zi90RFFLQmdFa2JlMmFrQldmSk9rQk5TZ3E4MUllTXVBdmVNS2hqK2toYmxaVXQKWUJvTU9uWFZ4MjVDK0QyeStLYktuS3pKVzBVaHV4MkhPRXJMWnR0K2hQaTFpVHQrWWQrQjRKeFJuMkROajVWNAowUHVITFkxR0NEUmtrWkt2eFZSUUV3S3pUSkpTTGtZcUFhaGZaQ2xRRzFpcGluVUwvSFNKWDNsWGluZDRQbmZlCmNCVUJBb0dBZDVhOElxYmVEUWJkcmZ3RmllY0RJZVBPVUFUdU0xY3NTVXE5dXZlNjhwaGU3TldGUTNtM2dsQWkKMlZnMTJpWGN5RmJXN0RXRThyc1pnQ2c2YjQ2QTFiUFkxRjNpMlFQMXp0Y2ZPS0ZqZXN2R1R6UHdQNFcwejAxagpqVjUvak5XT3NMVkx0TDEwSStrT3QrdUtNTU9mRFFseFdkT3JsTExwY0FFWUUxWlBadmM9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

		// 加密传输用私钥 
		config: POSTGRES_ENC_PRIVATE_KEY: string | *"data:;base64,LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBeEdUM2hDM3Y0UDUweHFjZWNCMXp2dm9NWTdBNUJVOVhnMmViclpnNVFJaE1Dc0xiCjJGdXhESk1kcjZzenBDckZGc3lmRWtvVUhMUUF1c00wY0xQN29hczVJRXpLSmU0TW1YaWFqcDRRZlBFSGF4a2EKMTBxYlE4UG5tOWFzUGcybEtTcHZsRURJb1RGNFZRV2dZVVZTTnJYaG5rZVNSeGNFU1lKQkpZWFlKcDQvRnNGdQp4T0gxRUVOY3V0SUZxcW4yOU00dHk5QXdBbjVsdEZQOU16U2JsMDNzbFBwSHJIUWIvOHRpd0tRVG44amJuWnNyClcwN3ovcWhMS0FqSXRFczd2SGJuZVBUelhYU2w1TktKTUlPOThoZkgwVmFwSDB1S0l6WXU5cXRSZ0tzYll0K1UKNnZJMlR4MHJjQ3pJU1EySlZ5TURSU1FnUTRLaitGNVl5Z3JIdHdJREFRQUJBb0lCQUFQcXlYTytJMnhmVFNobQpJZXV5MmFSd1dFZlVUTHhxdnhPN2c1K0tUaTRIQUN6QXhicmJtSFNxNTB4a3ZobHc2SzhNaVFXRkNmMnl2akM0CnBmYW9qSVVPOHljUTdOcE1zRzhVVkdlaGwybUN1SmNCWGtwUVdXUjhWcW44M0tVdkVvaGJtTjJubEliQ3RkdTgKR0doUW5IS1hIL29NWjRleHc0eEhXNXo2OVlDbjByZCtXb0tlR1l5VDJ0YUFramY0aFZFUnFLTWI1RStWczhURApOYmdXOHgyRUVKdVpuRk81N29wNDZya25xUkFPejF4eU1sMWtmeXpEa1FSK01EWTBsblYxVzBqeTRQVW9hRDRECnJlYUovVzhMejBJTGVBbmZXaEZ3cjVWRUpwMWQ2ckpaYUdIYWNaWmhLdGFSNmhwNXROUmVkZnU0dCtJVTJHZloKSkJ3K1ZYMENnWUVBMmRRNEg4bkVkZE05ZktEZ2h1V3dmdWtwNGVJMGFQbXJyQzRlbWRaTVNxTVBoM3RrSSs4dQpBU2JTVFF3SXBrdjdwVGJlemtHbVFUbzdxeE5CWWk4L0NEOHU2ajhxdlBvSFp4bGZGQTV4cTV3UkNUdXRpajV3ClJGV2FqOWVyU09GdXM2cUVQWkdabVBHbFFWdTJlaUt0UUE1akxCVGpvTjh2aHBHdWo0ejIzSFVDZ1lFQTVzOHgKOEhsYXZ6eVQ0M2FFL004a01TN1JZc1A1bFBtV1VvRHJzYlhybWVpM3FLSm14YU0rYWlEaDAzVVJ5MXpsSVVzaQp3b1NhNGRyOEJNRmwxRXVDSWNLMDBJcisxUldUSk5vZTZFQjErOWM5dEdxWUVlcGRIZjVhZjVQN0IyUnpXQ3pKCnNGMTBaZWhFejdhb2NBWGdvUDNBY3RNUDlXQ0RGQWoyRFR0OS9mc0NnWUJ1ZkNhdmNPNlYrTHdTTDZOU3FNUDUKeXhmME80RHIwZDlTYU40YWwwaUEvdTRNZ1BpTkJXN25KS0s4YzZNYmZpUzRhdmxkMG95YVB3Z3V3SWlFWFlSeQpFV0loYkVLb01ZVjV0TE8xLzVHR1FwV2Rna2lHZXg4RWVncjRkS0tyUjdTWHFxQ0NmZ3hUT0JYaTdickRmajB6CjVWaEY3cDU2WlZtOHZyMjBrQUpTVlFLQmdFMncrTEpsWnZKd1JhZVBRWHlIalRzdzh6STNuVTlVSkJGcGErekoKdTZCM0FUczJUem0vbFViTUFyZlc5RUpyNW9TcWNlemdEZkp3Yjl4NTdQams0Y3pUYVdHNUo4WTZHT21Tc2t5agpSaE9iaEIyeXM3VjBHaHY2ZmlQcmY5Z2hLK0pHVVMyWWg1RzErVk9odkZqWTRaL1BTblJjTDBiOVVhSHcwa0hxCkFLTDVBb0dBUEpwV3FQM2dyeFB6RzdReTRvT0NtNE9Qa2h0eGlxc2xleWduOXh3K1RrK0tqOHVGKzk2L1hyT3IKV1c2bFR1dnRWejRuNDE0R3VueHB2cERUMTJUYjd1VWFZYkpSQ1JqRzJ6NmNLdlpXbU96cklyRWVxSUNqSkJsRwo4ZVBvamZ5dXlpYmdSZ2RkMWRIZjRHQ05QV29vVE9hUndFK1ZJanFIR2JTR0lYTUJVQ2M9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

		// db data-dir
		config: POSTGRES_DATA_DIR: string

		// archive data-dir 
		config: POSTGRES_ARCHIVE_DATA_DIR: string | *""

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
