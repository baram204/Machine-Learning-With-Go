FROM yunabe/lgo:latest

WORKDIR /mlgo_execercise

# for password
RUN jupyter notebook --generate-config

# for extensions
RUN pip install jupyter_contrib_nbextensions
RUN cd /home/gopher/.local/bin && ./jupyter-contrib-nbextension install --user
RUN cd /home/gopher/.local/bin && ./jupyter-nbextensions_configurator enable --user

#for DB
RUN go get github.com/nanobox-io/golang-scribble
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/go-xorm/xorm
RUN go get github.com/go-xorm/cmd/xorm
RUN lgo install

# To use JupyterLab, replace "notebook" with "lab".
CMD ["jupyter", "notebook", "--ip=0.0.0.0"]
