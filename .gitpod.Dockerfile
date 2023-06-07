from golang:buster

WORKDIR /usr/app

COPY . .

# Create the gitpod user. UID must be 33333.
RUN useradd -l -u 33333 -G sudo -md /home/gitpod -s /bin/bash -p gitpod gitpod

USER gitpod

# RUN sudo apt update && sudo apt upgrade -y
# RUN sudo apt install fish -y