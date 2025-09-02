# Graph Scrutinizer Demo
## Introduction
This repository contains three demos to run for three Graph-Scrutinizer components: *Go-Network*, *Multi-Summary* and *TS2G2*. Each demo will involve running some code an answering questions in a Google form. Each folder will contain its own `README.md` file with further instructions on running the respective demo.

## Requirements
We try to minimize the amount of dependencies you need to manually install. The following three tools should help with that, but they themselves still need to be installed manually:

- Git
    - Git is a distributed version control system that offer various functions. We will need it to download some other code of the demos later. You can run `git --version` in a terminal to see if you have it correctly installed.
    - [Installation Guide](https://github.com/git-guides/install-git)
- Docker
    - Docker is a tool that offers several functions for creating and sharing a controlled environments for running programs in. We will be using Docker so you won't have to manually install any further dependencies. You can run `docker --version` in a terminal to see if you have it correctly installed. Make sure the Docker engine is running, either via command line interface or via the dedicated Docker Desktop gui. Starting the Docker engine can be slightly different on different operating systems, but once it is started, the `docker info` command should give you info on the Docker system (otherwise it will give an error).
    - [Installation Guide](https://github.com/git-guides/install-git)
- Make (GNU)
    - Make is a build automation tool that we use to make it easier to install all the further dependencies. You can run `make --version` in a terminal to see if you have it correctly installed.
    - Some operating systems come with this tool out of the box, but you might have to install it yourself. The installation is very operating system dependent, so please find a dedicated installation guide online for your own operating system if need be.

## Setting Up
Firstly, install the dependencies mentioned above. Now all you have to do is run the command below, in the current directory (`Graph-Scrutinizer_Demo_2.0/`), to set up the Docker image for the demos. This may take some time, so we recommend to run it now, so you can read the rest of the instructions during the installation process.

```
make
```

That's it! This will run the script in our `Makefile`, which will download the **TS2G2** and **multi-summary** libraries (the **Go-Network** does not need to be cloned). After that it will set up a Docker image and install all the required dependencies. This Docker image contains an Ubuntu 22.04 installation from which you can run the programs in each of the upcoming demos. If you are unfamiliar with Docker, then just know that a Docker image is like a "blueprint" for creating a Docker container (see the section below).

## Working with the Docker Container
In order to run programs via our Docker setup, you will first have to start a Docker container from the image you have just built. This container is effectively a controlled virtual environment with its own operating system (Ubuntu 22.04 in our case) and pre-installed dependencies (Go, Python, C++ and more in our case).


You can start a Docker container from our image, by running the following command:

```
make run
```

This will start a Docker container, along with a (virtual) terminal to interact with it. In this terminal you can run any of the normal command, such as `cd` to change between directories, but also some dependency specific ones, such as `go`, `python` and `g++`. The terminal should start out in a directory, called `/app`. This corresponds to the `Graph-Scrutinizer_Demo_2.0/` directory on your own machine. If you run `ls` in this terminal, you should see all the the demo directories (along with some other files). The container is set up in such a way that you can edit any of the files in the demo directories on your own machine and the Docker image will automatically update. This means that you can freely play around with any code in the demos in you own favorite IDE, while actually running all the code via the container. Each demo will give you further instructions as to what to do with the Docker container.

## Demos
Please go over each of the following demos. You can find detailed instructions in the `README.md` files of each respective demo. We also ask you to fill in a [Google form](https://docs.google.com/forms/d/e/1FAIpQLSfvQgYfAaGZT3zr9S74O1VlaZg1Lwj1bh5Rq-Wxj5AX8T4mCA/viewform?usp=header) containing questions for each demo.

### Go-Network
This demo is contained in `go-network_demo/`. For this demo you will import and use the Go-Network. The Go-Network is a library that offers graph generation and sampling algorithms. This demo uses the Go programming language.

### Multi-Summaries
This demo is contained in `multi-summaries_demo/`. For this demo, you will use the Multi-Summaries tool to create (condensed) multi-summaries from graphs and analyze them. This demo mainly uses C++ code in the background, but it comes with a Python interface, so you will not have to manually work with any C++ code.

### TS2G2
This demo is contained in `ts2g2_demo/`. For this demo you will use the TS2G2 tool via a selection of Jupyter notebooks. This demo uses Python.

## Cleaning Up
Finally, once you are done with all the demos and you have filled in the Google form, you can run the following command to clean up. It will delete the cloned repositories and it will clean up the Docker images and cache. If you were to just delete this repository, you might still have some dangling Docker images and cached data left (which can be **gigabytes** of data).

```
make clean
```

After running this command, you can delete this repository without leaving hidden data around. Thank you for your time! Happy testing!