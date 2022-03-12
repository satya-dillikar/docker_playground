# set base image (host OS)
FROM python:3.9

RUN pip install --upgrade pip

RUN adduser --disabled-password -q worker
USER worker

# set the working directory in the container
WORKDIR /code

# add `/app/node_modules/.bin` to $PATH
ENV PATH /home/worker/.local/bin:$PATH

COPY --chown=worker:worker . /code

# install dependencies
COPY --chown=worker:worker requirements.txt requirements.txt
RUN pip install --user -r requirements.txt

LABEL maintainer="Satya Dillikar <satya.dillikar@gmail.com>" \
      version="1.0.0"

ENTRYPOINT ["python"]
# command to run on container start
CMD [ "./server/main.py" ]
