package com.mon.service;

import com.mon.domain.Account;
import com.mon.repository.AccountRepository;
import lombok.NonNull;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
@RequiredArgsConstructor(onConstructor = @__(@Autowired))
public class AccountService {

    private final AccountRepository accountRepository;

    public Iterable<Account> getAll() {
        return accountRepository.findAll();
    }


    public Optional<Account> getById(@NonNull Long id) {
        return accountRepository.findById(id);
    }


    public Account create(@NonNull Account account) {
        return accountRepository.save(account);
    }


    public Account update(@NonNull Account account) {
        return accountRepository.save(account);
    }


    public void delete(@NonNull Long id) {
        accountRepository.deleteById(id);
    }


    public boolean exists(@NonNull Long id) {
        return accountRepository.existsById(id);
    }


}
