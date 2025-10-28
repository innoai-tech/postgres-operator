postgres-operator := 'go tool postgres-operator'

clean:
    rm -rf ./target/postgres

serve-singlenode:
    {{ postgres-operator }} \
    	-c \
    	serve singlenode \
    	    --daemon-off \
    	    --daemon-exit-on-error

serve-singlenode-in-docker: dep build
    docker run --init -it --rm --name=postgres-operator \
      -p=35432:5432 \
      -p=8080:80 \
      -e=POSTGRES_DATA_DIR=/var/lib/postgresql/data \
      -e=POSTGRES_NAME=${POSTGRES_NAME} \
      -e=POSTGRES_USER=${POSTGRES_USER} \
      -e=POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
      -e=POSTGRES_CPU=${POSTGRES_CPU} \
      -e=POSTGRES_MEM=${POSTGRES_MEM} \
      -v=${PWD}/target/postgres:/var/lib/postgresql/data \
      -v=${PWD}/target/postgres-operator_linux_arm64/postgres-operator:/postgres-operator \
      --entrypoint=/postgres-operator \
      postgres:16.10 serve singlenode -c

test-integration:
    CGO_ENABLED=0 \
      go test -failfast -count=1 ./integration-testing/...

dump-k8s:
    {{ postgres-operator }} serve singlenode \
        --dump-k8s
