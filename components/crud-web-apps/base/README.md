## Backend

The backend will be exposing a base backend which will be taking care of:
* Serving the Single Page Application
* Adding liveness/readiness probes
* Authentication based on the `kubeflow-userid` header
* Authorization using SubjectAccessReviews
* Uniform logging
* Exceptions handling
* Common helper functions for dates, yaml file parsing etc.

### Supported ENV Vars

The following is a list of ENV var that can be set for any web app that is using this base app.
This is list is incomplete, we will be adding more variables in the future.
| ENV Var | Description |
| - | - |
| CSRF_SAMESITE | Controls the [SameSite value](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#SameSite) of the CSRF cookie |

### How to use

In order to use this code during the development process one could use the `-e` [flag](https://pip.pypa.io/en/stable/reference/pip_install/#install-editable) with `pip install`. For example:

```bash
# I'm currently in /components/crud-web-apps/volumes/backend
cd ../../common/backend && pip install -e .
```

This will install all the dependencies of the package and you will now be able to include code from `kubeflow.kubeflow.crud_backend` in you current Python environment.

In order to build a Docker image and use this code you coud build a wheel and then install it:
```dockerfile
### Docker
FROM python:3.7 AS backend-kubeflow-wheel

WORKDIR /src
COPY ./components/crud-web-apps/common/backend .

RUN python3 setup.py bdist_wheel

...
# Web App
FROM python:3.7

WORKDIR /package
COPY --from=backend-kubeflow-wheel /src .
RUN pip3 install .
...
```