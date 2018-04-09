/*
 * stash-server
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1alpha1
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import io.swagger.client.model.IoK8sApimachineryPkgApisMetaV1APIGroup;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * APIGroupList is a list of APIGroup, to allow clients to discover the API at /apis.
 */
@ApiModel(description = "APIGroupList is a list of APIGroup, to allow clients to discover the API at /apis.")
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2018-04-09T00:59:16.508-07:00")
public class IoK8sApimachineryPkgApisMetaV1APIGroupList {
  @SerializedName("apiVersion")
  private String apiVersion = null;

  @SerializedName("groups")
  private List<IoK8sApimachineryPkgApisMetaV1APIGroup> groups = new ArrayList<IoK8sApimachineryPkgApisMetaV1APIGroup>();

  @SerializedName("kind")
  private String kind = null;

  public IoK8sApimachineryPkgApisMetaV1APIGroupList apiVersion(String apiVersion) {
    this.apiVersion = apiVersion;
    return this;
  }

   /**
   * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
   * @return apiVersion
  **/
  @ApiModelProperty(value = "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources")
  public String getApiVersion() {
    return apiVersion;
  }

  public void setApiVersion(String apiVersion) {
    this.apiVersion = apiVersion;
  }

  public IoK8sApimachineryPkgApisMetaV1APIGroupList groups(List<IoK8sApimachineryPkgApisMetaV1APIGroup> groups) {
    this.groups = groups;
    return this;
  }

  public IoK8sApimachineryPkgApisMetaV1APIGroupList addGroupsItem(IoK8sApimachineryPkgApisMetaV1APIGroup groupsItem) {
    this.groups.add(groupsItem);
    return this;
  }

   /**
   * groups is a list of APIGroup.
   * @return groups
  **/
  @ApiModelProperty(required = true, value = "groups is a list of APIGroup.")
  public List<IoK8sApimachineryPkgApisMetaV1APIGroup> getGroups() {
    return groups;
  }

  public void setGroups(List<IoK8sApimachineryPkgApisMetaV1APIGroup> groups) {
    this.groups = groups;
  }

  public IoK8sApimachineryPkgApisMetaV1APIGroupList kind(String kind) {
    this.kind = kind;
    return this;
  }

   /**
   * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
   * @return kind
  **/
  @ApiModelProperty(value = "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds")
  public String getKind() {
    return kind;
  }

  public void setKind(String kind) {
    this.kind = kind;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    IoK8sApimachineryPkgApisMetaV1APIGroupList ioK8sApimachineryPkgApisMetaV1APIGroupList = (IoK8sApimachineryPkgApisMetaV1APIGroupList) o;
    return Objects.equals(this.apiVersion, ioK8sApimachineryPkgApisMetaV1APIGroupList.apiVersion) &&
        Objects.equals(this.groups, ioK8sApimachineryPkgApisMetaV1APIGroupList.groups) &&
        Objects.equals(this.kind, ioK8sApimachineryPkgApisMetaV1APIGroupList.kind);
  }

  @Override
  public int hashCode() {
    return Objects.hash(apiVersion, groups, kind);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class IoK8sApimachineryPkgApisMetaV1APIGroupList {\n");
    
    sb.append("    apiVersion: ").append(toIndentedString(apiVersion)).append("\n");
    sb.append("    groups: ").append(toIndentedString(groups)).append("\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
