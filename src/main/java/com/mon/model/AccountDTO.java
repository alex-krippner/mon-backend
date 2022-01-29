package com.mon.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

import javax.validation.constraints.NotNull;
import java.time.LocalDate;

@Getter
@Setter
@Builder
public class AccountDTO {

    @NotNull
    @JsonProperty("id")
    private Long id;
    @JsonProperty
    private String name;
    @JsonProperty("date_of_birth")
    private LocalDate dateOfBirth;
    @JsonProperty
    private String username;
    @JsonProperty
    private String email;
}
