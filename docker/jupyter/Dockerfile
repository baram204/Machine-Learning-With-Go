FROM yunabe/lgo:latest

WORKDIR /mlgo_execercise

RUN jupyter notebook --generate-config
RUN pip install jupyter_contrib_nbextensions
RUN cd /home/gopher/.local/bin && ./jupyter-contrib-nbextension install --user
RUN cd /home/gopher/.local/bin && ./jupyter-nbextensions_configurator enable --user

# To use JupyterLab, replace "notebook" with "lab".
CMD ["jupyter", "notebook", "--ip=0.0.0.0"]
