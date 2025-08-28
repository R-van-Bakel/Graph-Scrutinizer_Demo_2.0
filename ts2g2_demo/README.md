# TS2G2
## Introduction
For this demo you will use the TS2G2 tool to turn time series into graphs and to turn graphs back into time series. This demo uses Python and Jupyter notebooks.

# Google Form
We have a [Google form](https://docs.google.com/forms/d/e/1FAIpQLSfvQgYfAaGZT3zr9S74O1VlaZg1Lwj1bh5Rq-Wxj5AX8T4mCA/viewform?usp=header) in which you get to answer some questions, relating to the usabilty of the tool and ease of the demo. Please read through the respective section of the form briefly, before continuing with the instructions.

# Instructions
For this demo you will have to go over some Jupyter notebooks. We have set this up so that you can host Jupyter notebook from the Docker container, while editing in your own browser. To get started, run the following instructions:

1. If you haven't done so already, open any terminal, navigate to the parent directory (`Graph-Scrutinizer_Demo_2.0/`) and start the Docker container with the following command:

    ```
    make
    ```

2. In the terminal running the Docker container, navigate to the current directory (`Graph-Scrutinizer_Demo_2.0/ts2g2_demo`).

3. In the terminal, run the following command:

    ```
    jupyter notebook --ip=0.0.0.0 --no-browser --allow-root
    ```

    This will run Jupyter notebook on the Docker container.

4. The previous command should have printed a link you can copy and paste into your browser to access the Jupyter notebook session. Please do so.

5. In the Jupyter notebook app, navigate to `ts2g2/tutorials/`. This should contain several Jupyter notebooks (`.ipynb` files).

6. Please go through at least three of the notebooks. We recommend to start with `review-demo-amazon-stocks.ipynb`. It might take some time to run the first code block, because Jupyter is loading the kernel in the background.

7. After you are done, you can close the Jupyer browser tabs. Then in the Docker terminal hit `CTRL+C`. Jupyter will ask if you want to exit. Just type `y` and press enter. This should end the Jupyter process runing in the background. You are now free to reuse the Docker terminal for other demos.

7. Finally fill in and submit the Google form provided above.