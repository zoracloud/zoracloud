import os

from kubernetes import client
from werkzeug import exceptions

from zoracloud.crud_backend import helpers, logging

from . import status

log = logging.getLogger(__name__)

FILE_ABS_PATH = os.path.abspath(os.path.dirname(__file__))

NOTEBOOK_TEMPLATE_YAML = os.path.join(
    FILE_ABS_PATH, "yaml/notebook_template.yaml"
)

# The production configuration is mounted on the app's pod via a configmap
DEV_CONFIG = os.path.join(FILE_ABS_PATH, "yaml/spawner_ui_config.yaml")
CONFIGS = [
    "/etc/config/spawner_ui_config.yaml",
    DEV_CONFIG,
]
