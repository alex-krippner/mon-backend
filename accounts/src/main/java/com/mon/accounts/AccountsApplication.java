package com.mon.accounts;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDate;
import java.util.*;

@SpringBootApplication
public class AccountsApplication {

    public static void main(String[] args) {
        SpringApplication.run(AccountsApplication.class, args);
    }

}

class Account {

    private final String id;
    private String name;
    private LocalDate dateOfBirth;
    private String username;
    private String email;

    public Account(String name, String dateOfBirth, String username, String email) {

        this.id = UUID.randomUUID().toString();
        this.name = name;
        this.dateOfBirth = LocalDate.parse(dateOfBirth);
        this.email = email;
        this.username = username;
    }

    public String getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public LocalDate getDateOfBirth() {
        return dateOfBirth;
    }

    public void setDateOfBirth(String dateOfBirth) {
        this.dateOfBirth = LocalDate.parse(dateOfBirth);
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}

@RestController
@RequestMapping("/accounts")
class RestApiAccountsController {
    private final List<Account> accounts = new ArrayList<>();


    public RestApiAccountsController() {
        accounts.addAll(List.of(
                new Account("John", "1988-09-29", "jonny", "jon@mail.com"),
                new Account("Jane", "1955-03-22", "jane", "jane@mail.com"),
                new Account("Bob", "1998-01-11", "bo", "bo@mail.com")
        ));
    }

    @GetMapping
    Iterable<Account> getAccounts() {
        return accounts;
    }

    @GetMapping("/{id}")
    Optional<Account> getAccountById(@PathVariable String id) {
        for (Account a : accounts) {
            if (a.getId().equals(id)) {
                return Optional.of(a);
            }
        }

        return Optional.empty();
    }

    @PostMapping
    Account postAccount(@RequestBody Account account) {
        accounts.add(account);
        return account;
    }

    @PutMapping("/{id}")
    ResponseEntity<Account> putAccount(@PathVariable String id, @RequestBody Account account) {
        int accountIndex = -1;

        for (Account a : accounts) {
            if (a.getId().equals(id)) {
                accountIndex = accounts.indexOf(a);
                accounts.set(accountIndex, account);
            }
        }

        return (accountIndex == -1) ? new ResponseEntity<>(postAccount(account), HttpStatus.CREATED) : new ResponseEntity<>(account, HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    void deleteAccount(@PathVariable String id) {
        accounts.removeIf(a -> a.getId().equals(id));
    }

}
