import logging
from flask import Flask

LOG_FORMAT = "%()s | %(name)s | %(levelname)s | %(message)s"


def create_app(name, static_folder, config):
    logging.basicConfig(format=LOG_FORMAT, level=config.LOG_LEVEL)
    log = logging.getLogger(__name__)

    app = Flask(name, static_folder=static_folder)
    app.config.from_object(config)

    if (config.ENV == BackendMode.DEVELOPMENT.value
            or config.ENV == BackendMode.DEVELOPMENT_FULL.value):
        log.warn("RUNNING IN DEVELOPMENT MODE")

    # Register the blueprints

    return app
