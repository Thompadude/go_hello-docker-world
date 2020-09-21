# What's This?
Something hacked together to have something to play around with on AWS ECS.

## Build on Local Machine
### Build & Run
```bash
> docker build . --tag foo && docker run -d -p 8080:8080 --name bar foo
```
### Clean-up
```bash
> docker rm -f bar && docker rmi foo
```

## Usage
```bash
> curl http://localhost:8080                                               
hello Docker world!
```