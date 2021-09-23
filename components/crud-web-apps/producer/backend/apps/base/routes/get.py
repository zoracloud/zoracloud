from zoracloud.crud_backend import api, logging
from .. import utils
from . import bp

log = logging.getLogger(__name__)


@bp.route("/api/yaml")
def get_config():
    config = utils.load_spawner_ui_config()
    return api.success_response("yaml", config)


@bp.route("/api/namespaces/<namespace>/pvcs")
def get_pvcs(namespace):
    pvcs = api.list_pvcs(namespace).items()
    contents = [utils.pvc_dict_from_k8s_obj(pvc) for pvc in pvcs]
    return api.success_response("pvcs", contents)
