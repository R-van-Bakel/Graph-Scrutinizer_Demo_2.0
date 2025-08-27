FROM ubuntu:22.04

# Install wget
RUN apt-get update && \
    apt-get install -y wget

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

EXPOSE 8888

# Download Python dependencies
# RUN pip install --no-cache-dir -r ts2g2_demo/ts2g2/ -r multi-summaries_demo/multi-summaries/
RUN pip install --no-cache-dir -r ts2g2_demo/ts2g2/requirements.txt
RUN pip install --no-cache-dir matplotlib scikit-learn scipy

CMD ["bash"]
