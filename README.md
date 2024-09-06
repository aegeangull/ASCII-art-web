# ASCII-Art-Web-Dockerize

Program that takes user input as a text and displays it as an ASCII art, stylized and dockerized.

## How to run

- [Install Docker](https://www.docker.com/)
- Log out and log in
- Go to project directory
- Start Docker
- Build image
  ```
  docker build -t ascii-art .
  ```
- Start Container
  ```
  docker container run -p 8080:8080 ascii-art
  ```
- Go to http://localhost:8080
- Type the text in box, choose the font and press "Submit" button
- ctrl+C in terminal stops container

## Clean
- Delete image
```
docker rmi -f ascii-art
```
