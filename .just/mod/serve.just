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
      postgres:16.11 serve singlenode -c

prepare-old-version-postgres-toolchain:
    docker run -it --rm --name=postgres-toolchains \
        -v=${PWD}/target/postgres-toolchain/16/usr-lib:/postgres-toolchain/usr/lib/postgresql/16 \
        -v=${PWD}/target/postgres-toolchain/16/usr-share:/postgres-toolchain/usr/share/postgresql/16 \
        postgres:16.10 sh -c "cp -r /usr/lib/postgresql/16/* /postgres-toolchain/usr/lib/postgresql/16/; cp -r /usr/share/postgresql/16/* /postgres-toolchain/usr/share/postgresql/16/";

serve-singlenode-in-docker-with-new: prepare-old-version-postgres-toolchain dep build
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
      -v=${PWD}/target/postgres-toolchain/16/usr-lib:/postgres-toolchain/usr/lib/postgresql/16 \
      -v=${PWD}/target/postgres-toolchain/16/usr-share:/postgres-toolchain/usr/share/postgresql/16 \
      --entrypoint=/postgres-operator \
      postgres:18.1 serve singlenode -c --daemon-exit-on-error

test-integration:
    CGO_ENABLED=0 \
      go test -failfast -count=1 ./integration-testing/...

dump-k8s:
    {{ postgres-operator }} serve singlenode \
        --dump-k8s
