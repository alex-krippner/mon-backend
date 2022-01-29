package com.mon.domain;

import com.mon.repository.AccountRepository;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.time.LocalDate;

@Component
@AllArgsConstructor
public class DataLoader {
    private final AccountRepository repository;

    @PostConstruct
    private void loadData() {
        repository.deleteAll();

        repository.save(
                new Account(81L, "John", LocalDate.parse("1988-09-29"), "jonny", "jon@mail.com"));
    }
}
