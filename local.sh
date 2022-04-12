source .env

# up
if [ "x$1" = "xup" ]; then
    docker-compose up -d

# down
elif [ "x$1" = "xdown" ]; then
    docker-compose down

# deps
elif [ "x$1" = "xdeps" ]; then
    docker-compose -f docker-compose-tools.yml run --rm go-mod

# logs
elif [ "x$1" = "xlogs" ]; then
    if [ -z "$2" ]; then
        docker-compose logs -f
    else
        docker-compose logs -f "$2"
    fi

else
    echo "You have to specify which action to be excuted. [ up / down / deps / logs ]" 1>&2
    exit 1
fi
