package com.mon.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotNull;

@Getter
@Setter
@Builder
public class AccountPostDTO {

    @Email
    @NotNull
    @JsonProperty("email")
    private String email;

    @NotNull
    @JsonProperty("name")
    private String name;

    @NotNull
    @JsonProperty("date_of_birth")
    private String dateOfBirth;

    @NotNull
    @JsonProperty("username")
    private String username;

}
