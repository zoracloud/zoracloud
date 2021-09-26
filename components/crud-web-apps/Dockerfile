# --- Build the backend kubeflow-wheel ---
FROM python:3.7-slim-buster AS backend-kubeflow-wheel
WORKDIR /src

COPY ./base/backend/ .
RUN python3 setup.py bdist_wheel

# --- Build the frontend kubeflow library ---
FROM node:12-buster-slim as frontend-kubeflow-lib

WORKDIR /src

ENV NG_CLI_ANALYTICS "ci"
COPY ./base/frontend/kubeflow-common-lib/package*.json ./
RUN npm ci

COPY ./base/frontend/kubeflow-common-lib/ .
RUN npm run build

# --- Build the frontend ---
FROM node:12-buster-slim as frontend

WORKDIR /src

COPY ./dataproducer/frontend/package*.json ./
COPY ./dataproducer/frontend/tsconfig*.json ./
COPY ./dataproducer/frontend/angular.json ./
COPY ./dataproducer/frontend/src ./src


ENV NG_CLI_ANALYTICS "ci"
RUN npm ci
COPY --from=frontend-kubeflow-lib /src/dist/kubeflow/ ./node_modules/kubeflow/

RUN npm run build -- --output-path=./dist/default --configuration=production
RUN npm run build -- --output-path=./dist/rok --configuration=rok-prod

# Web App
FROM python:3.7-slim-buster

WORKDIR /package
COPY --from=backend-kubeflow-wheel /src .
RUN pip3 install .

WORKDIR /src
COPY ./dataproducer/backend/requirements.txt .
RUN pip3 install -r requirements.txt

COPY ./dataproducer/backend/apps/ ./apps
COPY ./dataproducer/backend/entrypoint.py .

COPY --from=frontend /src/dist/default/ /src/apps/default/static/

ENTRYPOINT ["/bin/bash","-c","gunicorn -w 3 --bind 0.0.0.0:5000 --access-logfile - entrypoint:app"]

