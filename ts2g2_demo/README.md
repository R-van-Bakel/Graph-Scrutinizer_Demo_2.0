# TS2G2
## Introduction
For this demo you will use the TS2G2 tool to turn time series into graphs and to turn graphs back into time series. This demo uses Python and Jupyter notebooks.

# Google Form
We have a [Google form](https://docs.google.com/forms/d/e/1FAIpQLSfvQgYfAaGZT3zr9S74O1VlaZg1Lwj1bh5Rq-Wxj5AX8T4mCA/viewform?usp=header) in which you get to answer some questions, relating to the usability of the tool and ease of the demo. Please read through the respective section of the form briefly, before continuing with the instructions.

# Instructions
For this demo you will have to go over some Jupyter notebooks. We have set this up so that you can host Jupyter notebook from the Docker container, while editing in your own browser. To get started, run the following instructions:

1. If you haven't started the Docker container already, open any terminal, navigate to the parent directory (`Graph-Scrutinizer_Demo_2.0/`) and start the Docker container with the following command:

    ```
    make
    ```

2. In the terminal running the Docker container, navigate to the current directory (`Graph-Scrutinizer_Demo_2.0/ts2g2_demo`). If you did this correctly, then the `ls` command should show this `README.md` file and the `ts2g2` directory.

3. In the terminal, run the following command:

    ```
    jupyter notebook --ip=0.0.0.0 --no-browser --allow-root
    ```

    This will run Jupyter notebook on the Docker container.

4. The previous command should have printed a link you can copy and paste into your browser to access the Jupyter notebook session. Please do so.

5. In the Jupyter notebook app, navigate to `ts2g2/tutorials/`. This should contain several Jupyter notebooks (`.ipynb` files).

6. Please go through at least three of the notebooks. We provide more details for running these notebooks in the [section below](#running-the-notebooks). It might take some time to run the first code block, because Jupyter might still be loading the kernel in the background.

7. After you are done, you can close the Jupyter browser tabs. Then in the Docker terminal hit `CTRL+C`. Jupyter will ask if you want to exit. Just type `y` and press `enter`. This should end the Jupyter process running in the background. You are now free to reuse the Docker terminal for other demos.

7. Finally fill in and submit the Google form provided above.

## Running the Notebooks
### Overview
Please start with the "overview" notebook. This notebook will give a high-level overview about the TS2G2 library.

Next we provide six notebooks on time series to graph conversion methods: `visibility-strategy`, `ordinal-partition`, `quantiles`, `proximity-networks`, `sliding-window-strategy`, `timeseries-correlation`. Please run these notebooks in any order.

We also provide one notebook for graph to time series conversion: `g2ts-strategies`. Please go over that notebook now.


Currently, the `embeddings` notebook does not fully run, but please have a look at the notebook nonetheless. There should be some pre-generated outputs that you could have a look at. 


The `compact` notebook offers a final, brief, demonstration of all the previously discussed methods. Please go over this notebook and run all the code within.

### Interactive Demo
For the final part of the TS2G2 tutorial, please start by going over the `amazon-stocks` demo. It gives some examples of turning time series to graphs and back again, on Amazon stock data. After that, go over the `beta-testing` notebook. This final notebook looks very similar to the previous one, but it contains some further instructions for you to play around with the tool. The comments flagged with `TODO` contain theses instructions.