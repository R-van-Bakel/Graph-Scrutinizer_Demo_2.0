FROM ubuntu:22.04

# Install wget
RUN apt-get update && \
    apt-get install -y wget

# Install git
RUN apt-get update && \
    apt-get install -y git

# Install time
RUN apt update && \
    apt install time

# Install Python
RUN apt-get update && \
    apt-get install -y python3.10 python3.10-venv python3-pip

# Install Go
RUN wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz && \
    rm go1.25.0.linux-amd64.tar.gz
# Expose the `go` command
ENV PATH="/usr/local/go/bin:${PATH}"

# Install G++13 for C++
RUN apt-get update && \
    apt-get install -y g++-13

# Clean up package index files
RUN rm -rf /var/lib/apt/lists/*

# Clone the local directory into the container
WORKDIR /app
COPY . /app

# Download Go requirements
COPY go-network_demo/ ./
RUN go mod download

# Setup a virtual Python enviroment
RUN python3 -m venv /venv
ENV PATH="/venv/bin:$PATH"

# Expose a port to access Jupyter notebook later
EXPOSE 8888

# Download Python dependencies
RUN pip install --no-cache-dir -r ts2g2_demo/ts2g2/requirements.txt
RUN pip install --no-cache-dir -r ts2g2_demo/ts2g2/tutorials/requirements.txt
RUN pip install --no-cache-dir ts2g2_demo/ts2g2/ multi-summaries_demo/multi-summaries/code/python/
RUN pip install --no-cache-dir sktime gensim

# Add ts2g2 to the Python path (otherwise some demos won't work)
ENV PYTHONPATH="/app/ts2g2_demo/ts2g2"

CMD ["bash"]
