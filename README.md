# LQIP-API

A microservice for generating LQIP images from a POST, in a customizable and configurable way.


### Example

Note the `--data-binary` flag

```curl localhost:9980 --data-binary @testimg.jpg -o out.svg```

## Server Flags

| Flag                     | Type          | Default     | Description                                                                             |
| ------------------------ | ------------- | ----------- | --------------------------------------------------------------------------------------- |
| --port (-p)              | int           | 80          | Port to listen on                                                                       |
| --async (-a)             | bool          | false       | Allow drop-off/pickup async requests                                                    |
| --cacheLife              | int           | 600         | Time (in seconds) an async svg should be stored (When async enabled)                    |
| --cacheGC                | int           | 900         | Cadence (in seconds) in between scraping cache for expired content (When async enabled) |
| --postback (-r)          | bool          | false       | Allow drop-off/send-back async requests                                                 |
| --defaultShapeCount (-s) | int           | 16          | Default number of shapes in an LQIP                                                     |
| --allowShapeCountQP      | bool          | true        | Allow user to specify non-default shape count                                           |
| --maxShapeCountQP        | int           | 32          | Maximum user shape count specifiable                                                    |
| --defaultMode (-m)       | int           | 1           | Default type of shape to generate for the LQIP (See Modes)                              |
| --allowedModeQPs         | string        | "012345678" | Allowable modes specifiable by user (See Modes)                                         |
| --blur (-b)              | int           | 12          | Default level of Gaussian blur filter                                                   |
| --allowBlurQP            | bool          | true        | Allow user to specify different blur level                                              |


## Request Query Parameters

| Name        | Type    | Description                                             |
| ----------- | ------- | ------------------------------------------------------- |
| shapecount  | int     | Number of shapes to generate in LQIP                    |
| mode        | int     | Type of shapes to generate (See Modes)                  |
| blur        | int     | Strength of Gaussian blur filter                        |
| postback    | string  | Url to post finished product back to (in postback mode) |

## Modes

| Mode | Shape                                 |
| ---- | --------------------------------------|
| 0    | Mixed Set (1, 2, 6)                   |
| 1    | Triangles                             |
| 2    | Rectangles                            |
| 3    | Ellipses                              |
| 4    | Circles                               |
| 5    | Rotated Rectangles                    |
| 6    | Rotated Ellipses                      |
| 7    | Beziers (Use with higher shape count) |
| 8    | Freeform Polygons                     |

## Dockerized

Example startup:

Totally basic/default:
```docker run --rm -p 9980:80 lqip-api:0.0.0```

Pass args:
```docker run --rm -p 9980:80 lqip-api:0.0.0 -ar```

Host networking:
```docker run --rm --network=host lqip-api:0.0.0 -ar -p 9980```