#![allow(clippy::result_large_err)]
use anchor_lang::prelude::*;
use instructions::*;
pub mod instructions;

declare_id!("FvQXmfpyLX2BJFJFtFXNK2XHuBNqsv2hpYyUMHz7r8Kt");

#[program]
pub mod pda_rent_payer {
    use super::*;

    pub fn init_rent_vault(ctx: Context<InitRentVault>, fund_lamports: u64) -> Result<()> {
        init_rent_vault::init_rent_vault(ctx, fund_lamports)
    }

    pub fn create_new_account(ctx: Context<CreateNewAccount>) -> Result<()> {
        create_new_account::create_new_account(ctx)
    }
}
