TS2G2_REPO = https://github.com/graph-massivizer/ts2g2.git
TS2G2_FOLDER = ts2g2_demo/ts2g2/
SUMMARIZER_REPO = https://github.com/R-van-Bakel/Multi-Summaries.git
SUMMARIZER_FOLDER = multi-summaries_demo/multi-summaries/
PWD = $(shell pwd)
IMAGE = scrutinizer-demo-image

.PHONY: all clone build clean run

all: clone build

clone:
	@if test -d $(TS2G2_FOLDER); then \
		echo "TS2G2 is already cloned. Skipping cloning."; \
	else \
		git clone --depth=1 $(TS2G2_REPO) $(TS2G2_FOLDER); \
	fi
	@if test -d $(SUMMARIZER_FOLDER); then \
		echo "Multi-summaries is already cloned. Skipping cloning."; \
	else \
		git clone --depth=1 $(SUMMARIZER_REPO) $(SUMMARIZER_FOLDER); \
	fi

build:
	docker buildx build -t $(IMAGE) . --load

clean:
	@if test -d $(TS2G2_FOLDER); then \
		echo Removing local TS2G2 repository; \
		rm -rf $(TS2G2_FOLDER); \
	else \
		echo "Local TS2G2 repository not found. Skipping removing."; \
	fi
	@if test -d $(SUMMARIZER_FOLDER); then \
		echo Removing local multi-summaries repository; \
		rm -rf $(SUMMARIZER_FOLDER); \
	else \
		echo "Local multi-summaries repository not found. Skipping removing."; \
	fi
	@if [ -n "$$(docker images -q $(IMAGE))" ]; then \
		echo Removing docker image: $(IMAGE); \
		docker rmi -f $(IMAGE); \
	else \
		echo Docker image: $(IMAGE) not found. Skipping removing.; \
	fi
	docker image prune -f
	docker buildx prune -f

run:
	docker run -it --rm -p 8888:8888 -v $(PWD):/app $(IMAGE)