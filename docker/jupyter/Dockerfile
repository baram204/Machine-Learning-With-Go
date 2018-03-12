FROM yunabe/lgo:latest

WORKDIR /mlgo_execercise

# for password
RUN jupyter notebook --generate-config

# for extensions
RUN pip install jupyter_contrib_nbextensions
RUN cd /home/gopher/.local/bin && ./jupyter-contrib-nbextension install --user
RUN cd /home/gopher/.local/bin && ./jupyter-nbextensions_configurator enable --user

#for ML
RUN go get github.com/kniren/gota/dataframe

#for DB
RUN go get github.com/nanobox-io/golang-scribble
RUN go get github.com/lib/pq
RUN go get github.com/go-xorm/xorm
RUN go get github.com/go-xorm/cmd/xorm
RUN go get -u github.com/basgys/goxml2json
RUN lgo install

# To use JupyterLab, replace "notebook" with "lab".
CMD ["jupyter", "notebook", "--ip=0.0.0.0"]
