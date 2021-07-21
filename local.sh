source ./provision/local/.env

# up
if [ "x$1" = "xup" ]; then
    cd provision/local
    docker-compose up -d

# restart
elif [ "x$1" = "xrestart" ]; then
    ./local.sh down
    ./local.sh up

    ./local.sh logs

# down
elif [ "x$1" = "xdown" ]; then
    cd provision/local
    docker-compose down

# deps
elif [ "x$1" = "xdeps" ]; then
    cd ./app-web
    go mod vendor

# logs
elif [ "x$1" = "xlogs" ]; then
    cd provision/local
    if [ -z "$2" ]; then
        docker-compose logs -f
    else
        docker-compose logs -f "$2"
    fi

# destroy
elif [ "x$1" = "xdestroy" ]; then
    rm_container_id=$(docker images -a -q ${PROJECT_NAME})
    docker rmi -f ${rm_container_id}

    rm_container_id=$(docker images -a -q ${PROJECT_NAME}/*)
    docker rmi -f ${rm_container_id}

else
    echo "You have to specify which action to be excuted. [ up / restart / down / deps / logs / destroy ]" 1>&2
    exit 1
fi
