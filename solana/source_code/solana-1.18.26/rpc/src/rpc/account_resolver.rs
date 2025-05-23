use {
    solana_runtime::bank::Bank,
    solana_sdk::{account::AccountSharedData, pubkey::Pubkey},
    std::collections::HashMap,
};


// 优先从 overwrite_accounts 获取数据，否则从 bank 里面获取数据
pub(crate) fn get_account_from_overwrites_or_bank(
    pubkey: &Pubkey,
    bank: &Bank,
    overwrite_accounts: Option<&HashMap<Pubkey, AccountSharedData>>,
) -> Option<AccountSharedData> {
    overwrite_accounts
        .and_then(|accounts| accounts.get(pubkey).cloned())
        .or_else(|| bank.get_account(pubkey))
}
