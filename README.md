# LQIP-API

A microservice for generating LQIP images from a POST, in a customizable and configurable way.


## Example

Note the `--data-binary` flag

```curl localhost:9980 --data-binary @testimg.jpg -o out.svg```

## Dockerized

Example startup:

Totally basic/default:
```docker run --rm -p 9980:80 lqip-api:0.0.0```

Pass args:
```docker run --rm -p 9980:80 --entrypoint="/go/bin/lqip-api" lqip-api:0.0.0 -ar```

Host networking:
```docker run --rm --entrypoint="/go/bin/lqip-api" --network=host lqip-api:0.0.0 -ar -p 9980```