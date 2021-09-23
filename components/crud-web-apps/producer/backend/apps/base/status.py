import datetime as dt
from zoracloud.crud_backend import api, status

EVENT_TYPE_WARNING = "Warning"
STOP_ANNOTATION = "kubeflow-resource-stopped"


def process_status(producer):
    """
    Return status and reason. Status may be [running|waiting|warning|error]
    """
    # Check if the producer is stopped
    readyReplicas = producer.get("status", {}).get("readyReplicas", 0)
    metadata = producer.get("metadata", {})
    annotations = metadata.get("annotations", {})

    if STOP_ANNOTATION in annotations:
        if readyReplicas == 0:
            return status.create_status(
                status.STATUS_PHASE.STOPPED,
                "No Pods are currently running for this producer Server.",
            )
        else:
            return status.create_status(
                status.STATUS_PHASE.TERMINATING, "producer Server is stopping."
            )

    # If the producer is being deleted, the status will be waiting
    if "deletionTimestamp" in metadata:
        return status.create_status(
            status.STATUS_PHASE.TERMINATING, "Deleting this producer server"
        )

    # Check the status
    state = producer.get("status", {}).get("containerState", "")

    # Use conditions on the Jupyter producer (i.e., s) to determine overall
    # status. If no container state is available, we try to extract information
    # about why the producer is not starting from the producer's events
    # (see find_error_event)
    if readyReplicas == 1:
        return status.create_status(status.STATUS_PHASE.READY, "Running")

    if "waiting" in state:
        return status.create_status(
            status.STATUS_PHASE.WAITING, state["waiting"]["reason"]
        )

    # Provide the user with detailed information (if any) about
    # why the producer is not starting
    producer_events = get_producer_events(producer)
    status_val, reason = status.STATUS_PHASE.WAITING, "Scheduling the Pod"
    status_event, reason_event = find_error_event(producer_events)
    if status_event is not None:
        status_val, reason = status_event, reason_event

    return status.create_status(status_val, reason)


def get_producer_events(producer):
    name = producer["metadata"]["name"]
    namespace = producer["metadata"]["namespace"]

    nb_creation_time = dt.datetime.strptime(
        producer["metadata"]["creationTimestamp"], "%Y-%m-%dT%H:%M:%SZ"
    )

    nb_events = api.list_producer_events(name, namespace).items
    # User can delete and then create a nb server with the same name
    # Make sure previous events are not taken into account
    nb_events = filter(
        lambda e: event_timestamp(e) >= nb_creation_time, nb_events,
    )

    return nb_events


def find_error_event(producer_events):
    """
    Returns status and reason from the latest event that surfaces the cause
    of why the resource could not be created. For a producer, it can be due to:
          EVENT_TYPE      EVENT_REASON      DESCRIPTION
          Warning         FailedCreate      pods "x" is forbidden: error
            looking up service account ... (originated in statefulset)
          Warning         FailedScheduling  0/1 nodes are available: 1
            Insufficient cpu (originated in pod)
    """
    for e in sorted(producer_events, key=event_timestamp, reverse=True):
        if e.type == EVENT_TYPE_WARNING:
            return status.STATUS_PHASE.WAITING, e.message

    return None, None


def event_timestamp(event):
    return event.metadata.creation_timestamp.replace(tzinfo=None)
