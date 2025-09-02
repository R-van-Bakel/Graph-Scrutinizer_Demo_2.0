# Multi-Summary
For this demo, you will use the Multi-Summary tool to create (condensed) multi-summaries from graphs and analyze them. This demo uses some C++ code, along with a Python interface.

# Google Form
We have a [Google form](https://docs.google.com/forms/d/e/1FAIpQLSfvQgYfAaGZT3zr9S74O1VlaZg1Lwj1bh5Rq-Wxj5AX8T4mCA/viewform?usp=header) in which you get to answer some questions, relating to the usability of the tool and ease of the demo. Please read through the respective section of the form briefly, before continuing with the instructions.

# Instructions
1. If you haven't started the Docker container already, open any terminal, navigate to the parent directory (`Graph-Scrutinizer_Demo_2.0/`) and start the Docker container with the following command:

    ```
    make
    ```

2. In the terminal running the Docker container, navigate to the current directory (`Graph-Scrutinizer_Demo_2.0/multi-summaries_demo`). If you did this correctly, then the `ls` command should show several files, including `multi-summaries_demo.ipynb`.

3. In the terminal, run the following command:

    ```
    jupyter notebook --ip=0.0.0.0 --no-browser --allow-root
    ```

    This will run Jupyter notebook on the Docker container.

4. The previous command should have printed a link you can copy and paste into your browser to access the Jupyter notebook session. Please do so.

6. Open the `multi-summaries_demo.ipynb` notebook and follow its instructions. It might take some time to run the first code block, because Jupyter is loading the kernel in the background.

7. After you are done, you can close the Jupyter browser tabs. Then in the Docker terminal hit `CTRL+C`. Jupyter will ask if you want to exit. Just type `y` and press enter. This should end the Jupyter process running in the background. You are now free to reuse the Docker terminal for other demos.

7. Finally fill in and submit the Google form provided above.