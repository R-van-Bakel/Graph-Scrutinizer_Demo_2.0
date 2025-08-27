# Graph Scrutinizer Demo
## Introduction
This repository contains three demos to run for three Graph-Scrutinizer components: *Go-Network*, *Multi-Summary* and *TS2G2*. Each demo will involve running some code an answering questions in a Google form. Each folder will contain its own `README.md` file with further instructions on running the respective demo.

## Requirements
- Git
    - Git is a distributed version control system that offer various functions. We will need it to download some other code of the demos later.
    - [Installation Guide](https://github.com/git-guides/install-git)
- Docker
    - Docker is a tool that offers several functions for creating and sharing a controled environments for running programs in. We will be using Docker so you won't have to manually install any further dependencies.
    - [Installation Guide](https://github.com/git-guides/install-git)

## Setting Up
Firstly, install the dependencies mentioned above. Now all you have to do is run the command below to set up the Docker image for the demos. This may take some time, so we recommend to run it now, so you can read the rest of the instructions during the installation process.

```
make
```

That's it! This will run the script in our `Makefile`, which will download the TS2G2 and multi-summary libraries. After that it will set up a Docker image and install all the required dependecies. This Docker image contains an Ubuntu 22.04 installation from which you can run the programs in each of the upcomming demos.

## Working with the Docker Container
In order to run programs via our Docker setup, you will first have to start a Docker container from the image we you have just built. You can do this by running the following command:

```
make run
```

This will start a Docker container, along with a terminal to interact with it. In this terminal you can run any of the normal command, such as `cd` to change between directories, but also some dependency specific ones, such as `go`, `python` and `g++`. The container is set up in such a way that you can edit any of the files in the demo directory and the Docker image will automatically update. This means that you can freely edit any code in the demos in you own favorite IDE, while running all the programs via the container. Each demo will give you further instructions as to what to do with the Docker container.

Finally, once you are done with all the demos, you can run the following command to clean up. It will delete the cloned repositories and it will clean up the Docker images and cache. If you just delete this repository, you might still have some dangling Docker images and cached data left (which can be **gigabytes** of data).

```
make clean
```

After running this command, you can delete this repository without leaving hidden data around. Thank you for your time! Happy testing!

## Demos
Please go over each of the following demos. You can find detailed instructions in the `README.md` files of each respective demo. Each demo also has a corresponding Google form, in which we ask for some feedback relating to the respective demo.

### Go-Network
This demo is contained in `go-network_demo/`. For this demo you will import and use the Go-Network. The Go-Network is a library that offers graph generation and sampling algortihms. This demo uses the Go programming language.

### Multi-Summaries
This demo is contained in `multi-summaries_demo/`. For this demo, you will use the Multi-Summaries tool to create (condensed) multi-summaries from graphs and analyze them. This demo uses some C++ code in the background, but if comes with a Python interface.

### TS2G2
This demo is contained in `ts2g2_demo/`. For this demo you will use the TS2G2 tool to turn time series into graphs and to turn graphs back into time series again. This demo uses Python.