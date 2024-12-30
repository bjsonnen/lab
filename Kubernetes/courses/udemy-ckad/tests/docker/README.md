# Udemy CKAD

## Chapter 3: Configuration

### Docker

1.) How many images are available on this host?

- `docker image ls`

2.) What is the size of the ubuntu image?

- `docker image ls`

3.) We just pulled a new image. What is the tag on the newly pulled NGINX image?

- `docker image ls`

4.) We just downloaded the code of an application. What is the base image used in the Dockerfile?

- `cat Dockerfile`

5.) To what location within the container is the application code copied to during a Docker build?

- `cat Dockerfile`

6.) When a container is created using the image built with this Dockerfile, what is the command used to RUN the application inside it. 

- `cat Dockerfile`

7.) What port is the web application run within the container?

- `cat Dockerfile`

8.) Build a docker image using the Dockerfile and name it webapp-color. No tag to be specified.

- `docker build . -t webapp-color`

9.) Run an instance of the image webapp-color and publish port 8080 on the container to 8282 on the host.

- `docker run -p 8282:8080 webapp-color`

10.) What is the base Operating System used by the python:3.6 image? 

- `docker run python:3.6 cat /etc/*release*`

11.) What is the approximate size of the webapp-color image?

- `docker image ls`

12.) Build a new smaller docker image by modifying the same Dockerfile and name it webapp-color and tag it lite.

- Change base image from `python:3.6` to `python:3.6-slim`
- Run `docker build . -t webapp-color:lite`
