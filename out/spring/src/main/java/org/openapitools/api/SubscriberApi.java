/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech) (7.4.0-SNAPSHOT).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package org.openapitools.api;

import org.openapitools.model.Subscriber;
import org.openapitools.model.Subscription;
import io.swagger.v3.oas.annotations.ExternalDocumentation;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.Parameters;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import io.swagger.v3.oas.annotations.enums.ParameterIn;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.context.request.NativeWebRequest;
import org.springframework.web.multipart.MultipartFile;

import javax.validation.Valid;
import javax.validation.constraints.*;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Generated;

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-02-18T17:28:49.134660478Z[Etc/UTC]")
@Validated
@Tag(name = "subscriber", description = "Managing subscribers")
public interface SubscriberApi {

    default Optional<NativeWebRequest> getRequest() {
        return Optional.empty();
    }

    /**
     * DELETE /subscriber/{subscriberId}
     * Delete a single subscribers
     *
     * @param subscriberId ID of subscriber to delete (required)
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "deleteSubscriber",
        description = "Delete a single subscribers",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response")
        }
    )
    @RequestMapping(
        method = RequestMethod.DELETE,
        value = "/subscriber/{subscriberId}"
    )
    
    default ResponseEntity<Void> deleteSubscriber(
        @Parameter(name = "subscriberId", description = "ID of subscriber to delete", required = true, in = ParameterIn.PATH) @PathVariable("subscriberId") Long subscriberId
    ) {
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }


    /**
     * GET /subscriber/{subscriberId}
     * Get a single subscribers
     *
     * @param subscriberId ID of subscriber to get data from (required)
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "getSubscriberById",
        description = "Get a single subscribers",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response", content = {
                @Content(mediaType = "application/json", schema = @Schema(implementation = Subscriber.class))
            })
        }
    )
    @RequestMapping(
        method = RequestMethod.GET,
        value = "/subscriber/{subscriberId}",
        produces = { "application/json" }
    )
    
    default ResponseEntity<Subscriber> getSubscriberById(
        @Parameter(name = "subscriberId", description = "ID of subscriber to get data from", required = true, in = ParameterIn.PATH) @PathVariable("subscriberId") Long subscriberId
    ) {
        getRequest().ifPresent(request -> {
            for (MediaType mediaType: MediaType.parseMediaTypes(request.getHeader("Accept"))) {
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "{ \"name\" : \"bob smith\", \"id\" : 311109, \"username\" : \"bob123\" }";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
            }
        });
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }


    /**
     * GET /subscriber/{subscriberId}/subscription
     * Returns a list of all of the subscription of a subscriber
     *
     * @param subscriberId ID of subscriber to get the subscription from (required)
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "getSubscribersSubscriptions",
        description = "Returns a list of all of the subscription of a subscriber",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response", content = {
                @Content(mediaType = "application/json", array = @ArraySchema(schema = @Schema(implementation = Subscription.class)))
            })
        }
    )
    @RequestMapping(
        method = RequestMethod.GET,
        value = "/subscriber/{subscriberId}/subscription",
        produces = { "application/json" }
    )
    
    default ResponseEntity<List<Subscription>> getSubscribersSubscriptions(
        @Parameter(name = "subscriberId", description = "ID of subscriber to get the subscription from", required = true, in = ParameterIn.PATH) @PathVariable("subscriberId") Long subscriberId
    ) {
        getRequest().ifPresent(request -> {
            for (MediaType mediaType: MediaType.parseMediaTypes(request.getHeader("Accept"))) {
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "[ { \"subscriberId\" : 311109, \"id\" : 10, \"camera\" : { \"name\" : \"camera1\", \"location\" : { \"latitude\" : 41.4, \"longitude\" : 2.1 }, \"id\" : 11 }, \"status\" : \"active\" }, { \"subscriberId\" : 311109, \"id\" : 10, \"camera\" : { \"name\" : \"camera1\", \"location\" : { \"latitude\" : 41.4, \"longitude\" : 2.1 }, \"id\" : 11 }, \"status\" : \"active\" } ]";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
            }
        });
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }


    /**
     * PUT /subscriber/{subscriberId}
     * Update a single subscribers
     *
     * @param subscriberId ID of subscriber to update (required)
     * @param subscriber Update subscriber information (required)
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "updateSubscriber",
        description = "Update a single subscribers",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response")
        }
    )
    @RequestMapping(
        method = RequestMethod.PUT,
        value = "/subscriber/{subscriberId}",
        consumes = { "application/json" }
    )
    
    default ResponseEntity<Void> updateSubscriber(
        @Parameter(name = "subscriberId", description = "ID of subscriber to update", required = true, in = ParameterIn.PATH) @PathVariable("subscriberId") Long subscriberId,
        @Parameter(name = "Subscriber", description = "Update subscriber information", required = true) @Valid @RequestBody Subscriber subscriber
    ) {
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }

}
