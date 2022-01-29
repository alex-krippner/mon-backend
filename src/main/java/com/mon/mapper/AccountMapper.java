package com.mon.mapper;

import com.mon.domain.Account;
import com.mon.model.AccountDTO;
import com.mon.model.AccountPostDTO;
import org.mapstruct.Mapper;

@Mapper
public interface AccountMapper {

    AccountDTO accountToAccountDTO(Account account);

    Account accountPostDTOToAccount(AccountPostDTO accountPostDTO);

    Account accountDTOToAccount(AccountDTO account);

}
