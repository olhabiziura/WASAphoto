## WASAPhoto - a full stuck web-app developed for the course Web & Software Architecture in Sapienza University.

This repository contains the basic structure for [Web and Software Architecture](http://gamificationlab.uniroma1.it/en/wasa/) project.

# Project Description

 

Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you.

## Project specification - Simplified login

In real-world scenarios, new developments avoid implementing registration, login, and password-lost
flows as they are a security nightmare, cumbersome, error-prone, and outside the project scope. So,
why lose money and time on implementing those? The best practice is now to delegate those tasks to
a separate service (“identity provider”), either in-house (owned by the same company) or a third party
(like “Login with Apple/Facebook/Google” buttons).

In this project, we do not have an external service like this. Instead, we decided to provide you with a
specification for a login API so that you won’t spend time dealing with the design of the endpoint. 

The login endpoint accepts a username – like “Maria” – without any password. If the username already
exists, the user is logged in. If the username is new, the user is registered and logged in. The API will
return the user identifier you need to pass into the Authorization header in any other API.

![image](https://github.com/user-attachments/assets/13f61198-80a6-45fb-9938-b521916b0393)


![image](https://github.com/user-attachments/assets/fba9801b-d0de-4773-b5d1-7d4035bec035)

## Project functionality 

Each user is presented with a stream of photos (images) in reverse chronological order, with
information about when each photo was uploaded (date and time) and how many likes and comments
it has. The stream is composed by photos from “following” (other users that the user follows). Users
can place (and later remove) a “like” to photos from other users. Also, users can add comments to any
image (even those uploaded by themself). Only authors can remove their comments.

![image](https://github.com/user-attachments/assets/619f458b-dbe7-453e-8b3c-91e4b4c3e94f)

![image](https://github.com/user-attachments/assets/98a7d256-251e-4844-bdc5-a8e459dceef1)


Users can ban other users. If user Alice bans user Eve, Eve won’t be able to see any information about
Alice. Alice can decide to remove the ban at any moment.


![image](https://github.com/user-attachments/assets/2b91074a-ae77-4021-8195-f420c7e9d1fe)


![image](https://github.com/user-attachments/assets/05f118f0-fa34-4dff-bccb-c5c5edbfda37)




Users have their profiles. The personal profile page for the user shows: the user’s photos (in reverse
chronological order), how many photos have been uploaded, and the user’s followers and following.

![image](https://github.com/user-attachments/assets/4d8dbeb3-8907-4893-9c3c-7f2094874076)



Users can change their usernames, upload photos, remove photos, and follow/unfollow other users.
Removal of an image will also remove likes and comments.

A user can search other user profiles via username.

![image](https://github.com/user-attachments/assets/d5b66c7b-0b9d-42e5-a6dd-e5bf0c51ab36)


## Project structure

* `cmd/` contains all executables; Go programs here should only do "executable-stuff", like reading options from the CLI/env, etc.
	* `cmd/healthcheck` is an example of a daemon for checking the health of servers daemons; useful when the hypervisor is not providing HTTP readiness/liveness probes (e.g., Docker engine)
	* `cmd/webapi` contains an example of a web API server daemon
* `demo/` contains a demo config file
* `doc/` contains the documentation (usually, for APIs, this means an OpenAPI file)
* `service/` has all packages for implementing project-specific functionalities
	* `service/api` contains an example of an API server
	* `service/globaltime` contains a wrapper package for `time.Time` (useful in unit testing)
* `vendor/` is managed by Go, and contains a copy of all dependencies
* `webui/` is an example of a web frontend in Vue.js; it includes:
	* Bootstrap JavaScript framework
	* a customized version of "Bootstrap dashboard" template
	* feather icons as SVG
	* Go code for release embedding

Other project files include:
* `open-npm.sh` starts a new (temporary) container using `node:lts` image for safe web frontend development (you don't want to use `npm` in your system, do you?)

## Go vendoring

This project uses [Go Vendoring](https://go.dev/ref/mod#vendoring). You must use `go mod vendor` after changing some dependency (`go get` or `go mod tidy`) and add all files under `vendor/` directory in your commit.


## Node/NPM vendoring

This repository contains the `webui/node_modules` directory with all dependencies for Vue.JS. You should commit the content of that directory and both `package.json` and `package-lock.json`.

## How to set up a new project from this template

You need to:

* Change the Go module path to your module path in `go.mod`, `go.sum`, and in `*.go` files around the project
* If no cronjobs or health checks are needed, remove them from `cmd/`
* Update top/package comment inside `cmd/webapi/main.go` to reflect the actual project usage, goal, and general info

## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

