/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech) (7.4.0-SNAPSHOT).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package org.openapitools.api;

import org.openapitools.model.Subscriber;
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
public interface SubscribersApi {

    default Optional<NativeWebRequest> getRequest() {
        return Optional.empty();
    }

    /**
     * POST /subscribers
     * Create a single subscribers
     *
     * @param subscriber Update an existent pet in the store (required)
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "addSubscriber",
        description = "Create a single subscribers",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response")
        }
    )
    @RequestMapping(
        method = RequestMethod.POST,
        value = "/subscribers",
        consumes = { "application/json" }
    )
    
    default ResponseEntity<Void> addSubscriber(
        @Parameter(name = "Subscriber", description = "Update an existent pet in the store", required = true) @Valid @RequestBody Subscriber subscriber
    ) {
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }


    /**
     * GET /subscribers
     * Returns a list of all subscribers
     *
     * @return Successful response (status code 200)
     */
    @Operation(
        operationId = "getSubscribers",
        description = "Returns a list of all subscribers",
        tags = { "subscriber" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Successful response", content = {
                @Content(mediaType = "application/json", array = @ArraySchema(schema = @Schema(implementation = Subscriber.class)))
            })
        }
    )
    @RequestMapping(
        method = RequestMethod.GET,
        value = "/subscribers",
        produces = { "application/json" }
    )
    
    default ResponseEntity<List<Subscriber>> getSubscribers(
        
    ) {
        getRequest().ifPresent(request -> {
            for (MediaType mediaType: MediaType.parseMediaTypes(request.getHeader("Accept"))) {
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "[ { \"name\" : \"bob smith\", \"id\" : 311109, \"username\" : \"bob123\" }, { \"name\" : \"bob smith\", \"id\" : 311109, \"username\" : \"bob123\" } ]";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
            }
        });
        return new ResponseEntity<>(HttpStatus.NOT_IMPLEMENTED);

    }

}
