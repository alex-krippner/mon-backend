package com.mon.controller;

import com.mon.domain.Account;
import com.mon.mapper.AccountMapper;
import com.mon.model.AccountDTO;
import com.mon.model.AccountPostDTO;
import com.mon.service.AccountService;
import lombok.NonNull;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.ArrayList;
import java.util.Optional;

@RestController
@RequestMapping("/accounts")
@RequiredArgsConstructor(onConstructor = @__(@Autowired))
class AccountController {
    private final AccountService accountService;
    private final AccountMapper accountMapper;

    /**
     * Get all accounts
     *
     * @return all accounts
     */

    @GetMapping
    public ResponseEntity<ArrayList<AccountDTO>> getAccounts() {

        ArrayList<AccountDTO> accounts = new ArrayList();
        accountService.getAll().forEach(a -> accounts.add(accountMapper.accountToAccountDTO(a)));
        return ResponseEntity.ok(accounts);
    }

    /**
     * Find an account by its id
     *
     * @param id
     * @return found account
     */

    @GetMapping("/{id}")
    public ResponseEntity<AccountDTO> getAccountById(@PathVariable Long id) {
//        Optional<Account> account = accountService.getById(id);
//        return ResponseEntity.ok(accountMapper.accountToAccountGetDTO(account));

        return accountService.getById(id)
                .map(accountMapper::accountToAccountDTO)
                .map(ResponseEntity::ok).orElse(ResponseEntity.notFound().build());
    }

    /**
     * Create a new account
     *
     * @param accountPostDTO
     * @return http response
     */

    @PostMapping
    public ResponseEntity<?> postAccount(@RequestBody AccountPostDTO accountPostDTO) {
        Account account = accountMapper.accountPostDTOToAccount(accountPostDTO);
        accountService.create(account);

        URI uri = UriComponentsBuilder.fromUriString("/accounts").pathSegment(String.valueOf(account.getId())).build().toUri();

        return ResponseEntity.created(uri).build();
    }

    /**
     * Updates an entire account
     *
     * @param accountDTO
     * @return updated account
     */
    @PutMapping
    public ResponseEntity<AccountDTO> putAccount(@RequestBody AccountDTO accountDTO) {

        if (accountService.exists(accountDTO.getId())) {
            Account account = accountMapper.accountDTOToAccount(accountDTO);
            Account updatedAccount = accountService.update(account);

            return ResponseEntity.ok(accountMapper.accountToAccountDTO(updatedAccount));
        } else {
            return ResponseEntity.notFound().build();
        }
    }

    /**
     * Update single entity fields
     *
     * @param accountDTO
     * @return updated account
     */

    @PatchMapping("/{id}")
    public ResponseEntity<AccountDTO> patchAccount(@PathVariable Long id, @RequestBody AccountDTO accountDTO) {

        try {
            Account account = getPatchedAccount(id, accountDTO);
            Account updatedAccount = accountService.update(account);

            return ResponseEntity.ok(accountMapper.accountToAccountDTO(updatedAccount));
        } catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }

    private Account getPatchedAccount(@NonNull Long id, @NonNull AccountDTO accountDTO) throws Exception {
        final Account account = accountService.getById(id).orElseThrow(Exception::new);

        Optional.ofNullable(accountDTO.getEmail()).ifPresent(account::setEmail);
        Optional.ofNullable(accountDTO.getDateOfBirth()).ifPresent(account::setDateOfBirth);
        Optional.ofNullable(accountDTO.getName()).ifPresent(account::setName);
        Optional.ofNullable(accountDTO.getUsername()).ifPresent(account::setUsername);

        return account;
    }

    /**
     * Deletes an account
     *
     * @param id
     */

    @DeleteMapping("/{id}")
    void deleteAccount(@PathVariable Long id) {
        accountService.delete(id);
    }

}
