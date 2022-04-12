hello-go-safeweb
---

This is a sample of [go-safeweb](https://github.com/google/go-safeweb), a collection of libraries for writing secure-by-default HTTP servers in Go.

## Usage (HTTP)

You can see the removal of DOM XSS vulnerabilities.

### Run safe

1. Run the application.

```shell
./local.sh up
```

2. Access the following URL and confirm that the string `hoge` is not displayed.

[http://localhost:8080/#hoge](http://localhost:8080/#hoge)

![safe](https://user-images.githubusercontent.com/52403055/163034831-cb1a64cf-eebb-4f98-a604-dce06fa5fd75.png)

3. Down the application

```shell
./local.sh down
```

### Run unsafe

1. Rewrite `RUN_SAFE` in the .env file with `false`.

```
RUN_SAFE=false
```

2. Run the application.

```shell
./local.sh up
```

3. Access the following URL and confirm that the string `hoge` is displayed.

[http://localhost:8080/#hoge](http://localhost:8080/#hoge)

![unsafe](https://user-images.githubusercontent.com/52403055/163035041-452a7962-374d-4a0f-b058-c16a9fc5bec8.png)

4. Down the application

```shell
./local.sh down
```

### Usage (SQL)

See [here](https://github.com/hyorimitsu/hello-go-safeweb/blob/master/app-web/main.go#L61-L72).
