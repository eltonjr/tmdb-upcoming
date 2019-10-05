## Upcoming movies Webapp

#### About

This Webapp displays a list of upcoming movies and their details.
It is made in HTML+JS+CSS and it access the backend api from `../server`.

#### Building and running

A `Makefile` is provided to perform most tasks.

You can run the project using your local tools (node and npm needed) or using docker.

Make targets:

    Usage:
      make <target>

    Targets:
      help                  Display this help
      deps                  Install node dependencies based on npm
      build                 Compiles the system's dist
      run                   Run the system locally
      lint                  Check the system syntax
      ops-build             Compiles the system using docker
      ops-run               Runs the system using docker


#### Architectural decisions

The project was built using a [VueJS template for webpack](http://vuejs-templates.github.io/webpack/).
Application files are organized following the Vue style-guide for single file components.

The project depends on:
- Vue: the Single Page Application framework;
- Bulma: a css framework best known for being clean and easy to use;
- Axios: a small lib for http requests, more complete and popular than just `fetch`;
- Moment: a small lib for parsing dates;
- Vue-infinite-scroll: a simple vue directive to simplify the tricky parts about infinite scrolling.

This frontend webapp is served by an nginx, which is also used to reverse-proxy the requests to the backend.
