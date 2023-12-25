FROM golang:buster

WORKDIR /usr/app

# Create the gitpod user. UID must be 33333.
RUN useradd -l -u 33333 -md /home/gitpod -s /bin/bash -p gitpod gitpod

# Install required packages
RUN apt-get update && \
    apt-get install -y -q --allow-unauthenticated \
    git \
    sudo \
    fish

# Switch to the gitpod user
USER gitpod

# Clone Homebrew repository
RUN git clone https://github.com/Homebrew/brew /home/gitpod/.linuxbrew/Homebrew

# Create necessary directories
RUN mkdir -p /home/gitpod/.linuxbrew/bin
RUN ln -s /home/gitpod/.linuxbrew/Homebrew/bin/brew /home/gitpod/.linuxbrew/bin/brew

# Set up environment variables
ENV PATH="/home/gitpod/.linuxbrew/bin:${PATH}"

# Perform Homebrew update and doctor checks
RUN brew update
RUN brew doctor

# Switch back to root user for final configuration
USER root

# Change ownership of Homebrew folders
RUN chown -R gitpod: /home/gitpod/.linuxbrew

# Set the final user
USER gitpod
