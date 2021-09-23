SHELL=bash

install-deps:
	pushd ../../base/backend && \
	pip install -e . && \
	popd
	pip install -r requirements.txt

run:
	APP_PREFIX=/producer \
	gunicorn -w 3 --bind 0.0.0.0:50001 --access-logfile - entrypoint:app

run-dev:
	UI_FLAVOR=default \
	BACKEND_MODE=dev \
	APP_PREFIX=/ \
	APP_SECURE_COOKIES=False \
	gunicorn -w 3 --bind 0.0.0.0:50001 --access-logfile - entrypoint:app
