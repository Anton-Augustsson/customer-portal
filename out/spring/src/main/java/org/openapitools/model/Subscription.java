package org.openapitools.model;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import org.openapitools.model.Camera;
import org.openapitools.jackson.nullable.JsonNullable;
import java.time.OffsetDateTime;
import javax.validation.Valid;
import javax.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import javax.annotation.Generated;

/**
 * Subscription
 */

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-02-18T17:28:49.134660478Z[Etc/UTC]")
public class Subscription {

  private Long id;

  private Long subscriberId;

  /**
   * Gets or Sets status
   */
  public enum StatusEnum {
    ACTIVE("active"),
    
    INACTIVE("inactive"),
    
    SUSPENDED("suspended"),
    
    PROCESSING("processing");

    private String value;

    StatusEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static StatusEnum fromValue(String value) {
      for (StatusEnum b : StatusEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  private StatusEnum status;

  private Camera camera;

  public Subscription id(Long id) {
    this.id = id;
    return this;
  }

  /**
   * Get id
   * @return id
  */
  
  @Schema(name = "id", example = "10", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("id")
  public Long getId() {
    return id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public Subscription subscriberId(Long subscriberId) {
    this.subscriberId = subscriberId;
    return this;
  }

  /**
   * Get subscriberId
   * @return subscriberId
  */
  
  @Schema(name = "subscriberId", example = "311109", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("subscriberId")
  public Long getSubscriberId() {
    return subscriberId;
  }

  public void setSubscriberId(Long subscriberId) {
    this.subscriberId = subscriberId;
  }

  public Subscription status(StatusEnum status) {
    this.status = status;
    return this;
  }

  /**
   * Get status
   * @return status
  */
  
  @Schema(name = "status", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("status")
  public StatusEnum getStatus() {
    return status;
  }

  public void setStatus(StatusEnum status) {
    this.status = status;
  }

  public Subscription camera(Camera camera) {
    this.camera = camera;
    return this;
  }

  /**
   * Get camera
   * @return camera
  */
  @Valid 
  @Schema(name = "camera", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("camera")
  public Camera getCamera() {
    return camera;
  }

  public void setCamera(Camera camera) {
    this.camera = camera;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Subscription subscription = (Subscription) o;
    return Objects.equals(this.id, subscription.id) &&
        Objects.equals(this.subscriberId, subscription.subscriberId) &&
        Objects.equals(this.status, subscription.status) &&
        Objects.equals(this.camera, subscription.camera);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, subscriberId, status, camera);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Subscription {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    subscriberId: ").append(toIndentedString(subscriberId)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    camera: ").append(toIndentedString(camera)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

