FROM gitpod/workspace-full 

WORKDIR /usr/app

# Create the gitpod user. UID must be 33333.
# RUN useradd -l -u 33333 -md /home/gitpod -s /bin/bash -p gitpod gitpod

# Install required packages
RUN apt-get update && \
    apt-get install -y fish

# Switch to the gitpod user
# USER gitpod
