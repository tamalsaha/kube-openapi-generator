# coding: utf-8

"""
    stash-server

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: v0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.models.com_github_appscode_stash_apis_stash_v1alpha1_backend import ComGithubAppscodeStashApisStashV1alpha1Backend  # noqa: F401,E501
from swagger_client.models.com_github_appscode_stash_apis_stash_v1alpha1_local_spec import ComGithubAppscodeStashApisStashV1alpha1LocalSpec  # noqa: F401,E501
from swagger_client.models.com_github_appscode_stash_apis_stash_v1alpha1_local_typed_reference import ComGithubAppscodeStashApisStashV1alpha1LocalTypedReference  # noqa: F401,E501
from swagger_client.models.io_k8s_api_core_v1_local_object_reference import IoK8sApiCoreV1LocalObjectReference  # noqa: F401,E501


class ComGithubAppscodeStashApisStashV1alpha1RecoverySpec(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'backend': 'ComGithubAppscodeStashApisStashV1alpha1Backend',
        'image_pull_secrets': 'list[IoK8sApiCoreV1LocalObjectReference]',
        'node_name': 'str',
        'paths': 'list[str]',
        'pod_ordinal': 'str',
        'recovered_volumes': 'list[ComGithubAppscodeStashApisStashV1alpha1LocalSpec]',
        'workload': 'ComGithubAppscodeStashApisStashV1alpha1LocalTypedReference'
    }

    attribute_map = {
        'backend': 'backend',
        'image_pull_secrets': 'imagePullSecrets',
        'node_name': 'nodeName',
        'paths': 'paths',
        'pod_ordinal': 'podOrdinal',
        'recovered_volumes': 'recoveredVolumes',
        'workload': 'workload'
    }

    def __init__(self, backend=None, image_pull_secrets=None, node_name=None, paths=None, pod_ordinal=None, recovered_volumes=None, workload=None):  # noqa: E501
        """ComGithubAppscodeStashApisStashV1alpha1RecoverySpec - a model defined in Swagger"""  # noqa: E501

        self._backend = None
        self._image_pull_secrets = None
        self._node_name = None
        self._paths = None
        self._pod_ordinal = None
        self._recovered_volumes = None
        self._workload = None
        self.discriminator = None

        if backend is not None:
            self.backend = backend
        if image_pull_secrets is not None:
            self.image_pull_secrets = image_pull_secrets
        if node_name is not None:
            self.node_name = node_name
        if paths is not None:
            self.paths = paths
        if pod_ordinal is not None:
            self.pod_ordinal = pod_ordinal
        if recovered_volumes is not None:
            self.recovered_volumes = recovered_volumes
        if workload is not None:
            self.workload = workload

    @property
    def backend(self):
        """Gets the backend of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The backend of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: ComGithubAppscodeStashApisStashV1alpha1Backend
        """
        return self._backend

    @backend.setter
    def backend(self, backend):
        """Sets the backend of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param backend: The backend of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: ComGithubAppscodeStashApisStashV1alpha1Backend
        """

        self._backend = backend

    @property
    def image_pull_secrets(self):
        """Gets the image_pull_secrets of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The image_pull_secrets of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: list[IoK8sApiCoreV1LocalObjectReference]
        """
        return self._image_pull_secrets

    @image_pull_secrets.setter
    def image_pull_secrets(self, image_pull_secrets):
        """Sets the image_pull_secrets of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param image_pull_secrets: The image_pull_secrets of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: list[IoK8sApiCoreV1LocalObjectReference]
        """

        self._image_pull_secrets = image_pull_secrets

    @property
    def node_name(self):
        """Gets the node_name of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The node_name of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: str
        """
        return self._node_name

    @node_name.setter
    def node_name(self, node_name):
        """Sets the node_name of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param node_name: The node_name of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: str
        """

        self._node_name = node_name

    @property
    def paths(self):
        """Gets the paths of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The paths of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: list[str]
        """
        return self._paths

    @paths.setter
    def paths(self, paths):
        """Sets the paths of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param paths: The paths of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: list[str]
        """

        self._paths = paths

    @property
    def pod_ordinal(self):
        """Gets the pod_ordinal of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The pod_ordinal of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: str
        """
        return self._pod_ordinal

    @pod_ordinal.setter
    def pod_ordinal(self, pod_ordinal):
        """Sets the pod_ordinal of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param pod_ordinal: The pod_ordinal of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: str
        """

        self._pod_ordinal = pod_ordinal

    @property
    def recovered_volumes(self):
        """Gets the recovered_volumes of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The recovered_volumes of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: list[ComGithubAppscodeStashApisStashV1alpha1LocalSpec]
        """
        return self._recovered_volumes

    @recovered_volumes.setter
    def recovered_volumes(self, recovered_volumes):
        """Sets the recovered_volumes of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param recovered_volumes: The recovered_volumes of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: list[ComGithubAppscodeStashApisStashV1alpha1LocalSpec]
        """

        self._recovered_volumes = recovered_volumes

    @property
    def workload(self):
        """Gets the workload of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501


        :return: The workload of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :rtype: ComGithubAppscodeStashApisStashV1alpha1LocalTypedReference
        """
        return self._workload

    @workload.setter
    def workload(self, workload):
        """Sets the workload of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.


        :param workload: The workload of this ComGithubAppscodeStashApisStashV1alpha1RecoverySpec.  # noqa: E501
        :type: ComGithubAppscodeStashApisStashV1alpha1LocalTypedReference
        """

        self._workload = workload

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ComGithubAppscodeStashApisStashV1alpha1RecoverySpec):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
