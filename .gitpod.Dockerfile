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

# Set Homebrew prefix to user's home directory
ENV HOMEBREW_PREFIX="/home/gitpod/.linuxbrew"

# Install Homebrew
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

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
